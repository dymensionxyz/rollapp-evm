package main

import (
	"bytes"
	"context"
	"encoding/binary"
	"encoding/gob"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
	"golang.org/x/sync/errgroup"
	"log"
	"math/big"
	"net/http"
	randomnessgeneratorAPI "randomnessgenerator/agent/contractapi"
	"randomnessgenerator/agent/service"
	"strings"
	"time"
)

type Config struct {
	NodeURL              string
	Mnemonic             string
	HexContractAddress   string
	DerivationPath       string
	GasLimit             uint64
	GasFeeCap            *big.Int
	GasTipCap            *big.Int
	HTTPServerAddr       string
	RandomnessServiceURL string
}

type RNGAgent struct {
	EthClient       *ethclient.Client
	ContractAPI     *randomnessgeneratorAPI.Contract
	ContractAddress common.Address
	DB              *leveldb.DB
	Auth            *bind.TransactOpts
	Generator       *service.RandomnessGenerator
}

func NewRNGAgent(cfg Config) (*RNGAgent, error) {
	client, err := ethclient.Dial(cfg.NodeURL)
	if err != nil {
		return nil, fmt.Errorf("can't connect to evm: %w", err)
	}
	log.Printf("connected to Ethereum node at %s", cfg.NodeURL)

	db, err := leveldb.OpenFile("./db", nil)
	if err != nil {
		return nil, fmt.Errorf("can't connect to db: %w", err)
	}

	contractAddress := common.HexToAddress(cfg.HexContractAddress)
	contract, err := randomnessgeneratorAPI.NewContract(contractAddress, client)
	if err != nil {
		return nil, fmt.Errorf("can't use rng smart-contract API: %w", err)
	}

	exists, err := contractExists(client, contractAddress)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, fmt.Errorf("contract does not exist at address: %s", contractAddress.Hex())
	}

	g, err := service.NewRandomnessGenerator(cfg.RandomnessServiceURL)
	if err != nil {
		return nil, fmt.Errorf("can't connect to randomness generator service: %w", err)
	}

	auth := createTransactOpts(client, cfg)
	return &RNGAgent{
		EthClient:       client,
		ContractAddress: contractAddress,
		DB:              db,
		ContractAPI:     contract,
		Auth:            auth,
		Generator:       g,
	}, nil
}

const (
	RandomnessRequested = 0
)

type RandomnessRequestedEvent struct {
	ID uint64
}

func parseRandomnessRequestedEvent(data []byte) (*RandomnessRequestedEvent, error) {
	parsedABI, err := abi.JSON(strings.NewReader(`[{"anonymous":false,"inputs":[{"indexed":false,"name":"id","type":"uint64"}],"name":"RandomnessRequestedEvent","type":"event"}]`))
	if err != nil {
		return nil, fmt.Errorf("failed to parse ABI: %w", err)
	}

	var event RandomnessRequestedEvent
	err = parsedABI.UnpackIntoInterface(&event, "RandomnessRequestedEvent", data)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack event data: %w", err)
	}

	return &event, nil
}

func (a *RNGAgent) ListenForSmartContractEvents(ctx context.Context, config Config) error {
	for {
		select {
		case <-ctx.Done():
			log.Println("context canceled, exiting event loop")
			return ctx.Err()
		default:
			events, err := a.ContractAPI.PollEvents(&bind.CallOpts{Context: ctx}, RandomnessRequested)
			if err != nil {
				log.Printf("error polling events from contract: %v", err)
				continue
			}

			for _, event := range events {
				r, err := parseRandomnessRequestedEvent(event.Data)
				if err != nil {
					log.Printf("error parsing randomness requested event: %v", err)
					continue
				}

				randID := r.ID
				randomnessResp, err := a.Generator.GenerateUInt256()
				if err != nil {
					log.Printf("can't generate u256 random: %v", err)
					continue
				}

				var buf bytes.Buffer
				enc := gob.NewEncoder(&buf)
				err = enc.Encode(randomnessResp)
				if err != nil {
					log.Printf("randomness response serialization err: %v", err)
				}

				key := make([]byte, 8)
				binary.BigEndian.PutUint64(key, randID)
				err = a.DB.Put(key, buf.Bytes(), nil)
				if err != nil {
					log.Printf("error putting [key;value] = [%s;%d] into DB: %v", randID, randomnessResp.Randomness, err)
					continue
				}

				log.Printf("[%d:%s]", randID, randomnessResp.Randomness.String())

				auth := createTransactOpts(a.EthClient, config)
				tx, err := a.ContractAPI.PostRandomness(auth, randID, randomnessResp.Randomness)
				if err != nil {
					log.Printf("error PostRandomness tx: %v", err)
					continue
				}
				err = waitForTransaction(ctx, a.EthClient, tx)
				if err != nil {
					log.Println(a.Auth.From.String())
					log.Printf("PostRandomness tx failed: %v", err)
					continue
				}
			}

			time.Sleep(time.Second)
		}
	}
}

func (a *RNGAgent) handleGetRandomness(w http.ResponseWriter, r *http.Request) {
	ids, ok := r.URL.Query()["id"]
	if !ok || len(ids[0]) < 1 {
		http.Error(w, "missing ID parameter", http.StatusBadRequest)
		return
	}
	id := ids[0]

	idBytes, success := new(big.Int).SetString(id, 10)
	if !success {
		http.Error(w, "invalid ID format", http.StatusBadRequest)
		return
	}

	randomnessBytes, err := a.DB.Get(idBytes.Bytes(), nil)
	if err != nil {
		if errors.Is(err, leveldb.ErrNotFound) {
			http.Error(w, "randomness not found", http.StatusNotFound)
		} else {
			http.Error(w, fmt.Sprintf("error retrieving randomness: %v", err), http.StatusInternalServerError)
		}
		return
	}

	var randomnessResp service.RandomnessResponse
	buf := bytes.NewBuffer(randomnessBytes)
	dec := gob.NewDecoder(buf)
	err = dec.Decode(&randomnessResp)
	if err != nil {
		http.Error(w, fmt.Sprintf("error decoding randomness: %v", err), http.StatusInternalServerError)
		return
	}

	randomnessStr := randomnessResp.Randomness.String()

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(struct {
		RequestID  string `json:"requestID"`
		Randomness string `json:"randomness"`
	}{
		RequestID:  randomnessResp.RequestID,
		Randomness: randomnessStr,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("error encoding response to JSON: %v", err), http.StatusInternalServerError)
		return
	}
}

func (a *RNGAgent) StartHTTPServer(ctx context.Context, address string) error {
	server := &http.Server{
		Addr: address,
	}

	http.HandleFunc("/randomness", a.handleGetRandomness)

	go func() {
		<-ctx.Done()
		_ = server.Shutdown(context.Background())
	}()

	log.Printf("Starting HTTP server at %s", address)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		return err
	}
	return nil
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	config := Config{
		NodeURL:              "http://127.0.0.1:8545",
		Mnemonic:             "depend version wrestle document episode celery nuclear main penalty hundred trap scale candy donate search glory build valve round athlete become beauty indicate hamster",
		HexContractAddress:   "0x676E400d0200Ac8f3903A3CDC7cc3feaF21004d0",
		DerivationPath:       "m/44'/60'/0'/0/0",
		GasLimit:             60000000,
		GasFeeCap:            big.NewInt(30000000000), // 30 Gwei
		GasTipCap:            big.NewInt(2000000000),  // 2 Gwei
		HTTPServerAddr:       ":8080",
		RandomnessServiceURL: "http://127.0.0.1:8081/generate",
	}

	ctx := context.Background()
	g, ctx := errgroup.WithContext(ctx)

	agent, err := NewRNGAgent(config)
	if err != nil {
		log.Fatalf("error while creating rng agent: %v", err)
	}

	g.Go(func() error { return agent.ListenForSmartContractEvents(ctx, config) })
	g.Go(func() error { return agent.StartHTTPServer(ctx, config.HTTPServerAddr) })

	if err := g.Wait(); err != nil {
		log.Fatalf("failed: %v", err)
	}
}

func waitForTransaction(ctx context.Context, client *ethclient.Client, tx *types.Transaction) error {
	receipt, err := bind.WaitMined(ctx, client, tx)
	if err != nil {
		return fmt.Errorf("error waiting for transaction confirmation: %w", err)
	}

	if receipt.Status == 1 {
		return nil
	}

	revertReason, err := getRevertReason(client, tx.Hash())
	if err != nil {
		return fmt.Errorf("error getting revert reason: %w", err)
	}
	return fmt.Errorf("tx reverted: hash: %s, reason: %s", tx.Hash().String(), revertReason)
}

func getRevertReason(client *ethclient.Client, txHash common.Hash) (string, error) {
	receipt, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		return "", fmt.Errorf("failed to get receipt: %w", err)
	}

	if receipt.Status != 0 {
		return "", fmt.Errorf("transaction did not fail")
	}

	tx, _, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		return "", fmt.Errorf("failed to get transaction: %w", err)
	}

	msg := ethereum.CallMsg{
		To:   tx.To(),
		Data: tx.Data(),
	}

	res, err := client.CallContract(context.Background(), msg, receipt.BlockNumber)
	if err != nil {
		return "", fmt.Errorf("failed to call contract: %w", err)
	}

	if len(res) < 4 {
		return "No revert reason", nil
	}

	const errorMethodID = "0x08c379a0"
	if fmt.Sprintf("0x%x", res[:4]) != errorMethodID {
		return "Could not decode revert reason", nil
	}

	abiError, err := abi.JSON(strings.NewReader(`[{"inputs":[{"internalType":"string","name":"reason","type":"string"}],"name":"Error","type":"function"}]`))
	if err != nil {
		return "", fmt.Errorf("failed to parse ABI: %w", err)
	}

	var errorMsg string
	err = abiError.UnpackIntoInterface(&errorMsg, "Error", res[4:])
	if err != nil {
		return "", fmt.Errorf("failed to unpack revert reason: %w", err)
	}

	return errorMsg, nil
}

func contractExists(client *ethclient.Client, address common.Address) (bool, error) {
	code, err := client.CodeAt(context.Background(), address, nil)
	if err != nil {
		return false, fmt.Errorf("failed to check contract existence: %w", err)
	}
	return len(code) > 0, nil
}

func createTransactOpts(client *ethclient.Client, config Config) *bind.TransactOpts {
	seed := bip39.NewSeed(config.Mnemonic, "")
	masterKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		log.Fatalf("error creating master key: %v", err)
	}

	parts := strings.Split(config.DerivationPath, "/")
	key := masterKey
	for _, p := range parts[1:] {
		hardened := false
		if strings.HasSuffix(p, "'") {
			hardened = true
			p = strings.TrimSuffix(p, "'")
		}
		index := 0
		_, err := fmt.Sscanf(p, "%d", &index)
		if err != nil {
			log.Fatalf("invalid path element %s: %v", p, err)
		}

		if hardened {
			key, err = key.NewChildKey(uint32(index) + bip32.FirstHardenedChild)
		} else {
			key, err = key.NewChildKey(uint32(index))
		}
		if err != nil {
			log.Fatalf("failed to derive key at index %d: %v", index, err)
		}
	}

	privateKeyECDSA, err := crypto.ToECDSA(key.Key)
	if err != nil {
		log.Fatalf("error converting to ECDSA: %v", err)
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatalf("error getting Chain ID: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKeyECDSA, chainID)
	if err != nil {
		log.Fatalf("error creating signed transactor: %v", err)
	}

	auth.GasLimit = config.GasLimit
	auth.GasFeeCap = config.GasFeeCap
	auth.GasTipCap = config.GasTipCap

	return auth
}
