package main

import (
	"context"
	"crypto/ecdsa"
	"errors"
	randomnessgeneratorAPI "example1/agent/contractapi"
	"example1/agent/usecases"
	"example1/agent/usecases/service"
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
	"strings"
	"time"
)

// Config holds the configuration parameters
type Config struct {
	NodeURL            string
	Mnemonic           string
	HexContractAddress string
	DerivationPath     string
	GasLimit           uint64
	GasFeeCap          *big.Int
	GasTipCap          *big.Int
	HTTPServerAddr     string
}

// Deploying contracts with the account: 0x84ac82e5Ae41685D76021b909Db4f8E7C4bE279E
// RandomnessGenerator deployed at: 0x833F2FE4BFF66e712aA2C676ce8AdF08fd65B028

type RNGAgent struct {
	EthClient       *ethclient.Client
	ContractAPI     *randomnessgeneratorAPI.Contract
	ContractAddress common.Address
	DB              *leveldb.DB
	Auth            *bind.TransactOpts
	Generator       usecases.RandomnessGenerator
}

func NewRNGAgent(cfg Config) (*RNGAgent, error) {
	client, err := ethclient.Dial(cfg.NodeURL)
	if err != nil {
		return nil, fmt.Errorf("can't connect to evm: %v", err)
	}
	log.Printf("connected to Ethereum node at %s", cfg.NodeURL)

	db, err := leveldb.OpenFile("./db", nil)
	if err != nil {
		return nil, fmt.Errorf("can't connect to db: %v", err)
	}

	contractAddress := common.HexToAddress(cfg.HexContractAddress)
	contract, err := randomnessgeneratorAPI.NewContract(contractAddress, client)
	if err != nil {
		return nil, fmt.Errorf("can't use rng smart-contract API: %v", err)
	}

	exists, err := contractExists(client, contractAddress)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, fmt.Errorf("contract does not exist at address: %s", contractAddress.Hex())
	}

	g, err := service.NewRandomnessGenerator()
	if err != nil {
		return nil, fmt.Errorf("can't connect to randomness generator service: %v", err)
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
	ID *big.Int
}

func parseRandomnessRequestedEvent(data []byte) (*RandomnessRequestedEvent, error) {
	parsedABI, err := abi.JSON(strings.NewReader(`[{"anonymous":false,"inputs":[{"indexed":false,"name":"id","type":"uint256"}],"name":"RandomnessRequestedEvent","type":"event"}]`))
	if err != nil {
		return nil, fmt.Errorf("failed to parse ABI: %w", err)
	}

	var event RandomnessRequestedEvent
	err = parsedABI.UnpackIntoInterface(&event, "RandomnessRequestedEvent", data)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack event data: %w", err)
	}

	if event.ID == nil {
		return nil, fmt.Errorf("randomness ID is nil")
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

			var processedEvents []*big.Int
			for _, event := range events {
				r, err := parseRandomnessRequestedEvent(event.Data)
				if err != nil {
					log.Printf("error parsing randomness requested event: %v", err)
					continue
				}

				randID := r.ID
				randomness, err := a.Generator.GenerateUInt256()
				if err != nil {
					log.Printf("can't generate u256 random: %v", err)
					continue
				}

				_, err = a.DB.Get(randID.Bytes(), nil)
				if err == nil {
					log.Printf("randomness already put in db for ID: %s", randID.String())
					continue
				} else if !errors.Is(err, leveldb.ErrNotFound) {
					log.Printf("error while getting randomness key in db: %v", err)
					continue
				}

				err = a.DB.Put(randID.Bytes(), randomness.Bytes(), nil)
				if err != nil {
					log.Printf("error putting [key;value] = [%s;%d] into DB: %v", randID.String(), randomness, err)
					continue
				}

				log.Printf("[%s:%s]", randID.String(), randomness.String())

				auth := createTransactOpts(a.EthClient, config)
				tx, err := a.ContractAPI.PostRandomness(auth, randID, randomness)
				if err != nil {
					log.Printf("error PostRandomness tx: %v", err)
					continue
				}
				err = waitForTransaction(a.EthClient, tx)
				if err != nil {
					log.Println(a.Auth.From.String())
					log.Printf("PostRandomness tx failed: %v", err)
					continue
				}

				processedEvents = append(processedEvents, event.EventId)
			}

			if len(processedEvents) > 0 {
				auth := createTransactOpts(a.EthClient, config)
				tx, err := a.ContractAPI.EraseEvents(auth, processedEvents, RandomnessRequested)
				if err != nil {
					log.Printf("error sending EraseEvents tx: %v", err)
					continue
				}
				err = waitForTransaction(a.EthClient, tx)
				if err != nil {
					log.Printf("EraseEvents tx failed: %v", err)
				}
			}

			time.Sleep(time.Second)
		}
	}
}

func (a *RNGAgent) handleGetRandomness(w http.ResponseWriter, r *http.Request) {
	ids, ok := r.URL.Query()["id"]
	if !ok || len(ids[0]) < 1 {
		http.Error(w, "Missing ID parameter", http.StatusBadRequest)
		return
	}
	id := ids[0]

	idBytes, success := new(big.Int).SetString(id, 10)
	if !success {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	randomnessBytes, err := a.DB.Get(idBytes.Bytes(), nil)
	if err != nil {
		if errors.Is(err, leveldb.ErrNotFound) {
			http.Error(w, "Randomness not found", http.StatusNotFound)
		} else {
			http.Error(w, fmt.Sprintf("Error retrieving randomness: %v", err), http.StatusInternalServerError)
		}
		return
	}

	randomness := new(big.Int).SetBytes(randomnessBytes)
	_, _ = fmt.Fprintf(w, randomness.String())
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
		NodeURL:            "http://127.0.0.1:8545", // Local Hardhat node
		Mnemonic:           "depend version wrestle document episode celery nuclear main penalty hundred trap scale candy donate search glory build valve round athlete become beauty indicate hamster",
		HexContractAddress: "0x676E400d0200Ac8f3903A3CDC7cc3feaF21004d0", // Replace with the correct contract address
		DerivationPath:     "m/44'/60'/0'/0/0",
		GasLimit:           1e7,
		GasFeeCap:          big.NewInt(3e15),             // 30 Gwei
		GasTipCap:          big.NewInt(2000000000000000), // 2 Gwei
		HTTPServerAddr:     ":8080",
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

func waitForTransaction(client *ethclient.Client, tx *types.Transaction) error {
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		return fmt.Errorf("error waiting for transaction confirmation: %v", err)
	}

	if receipt.Status == 1 {
		return nil
	}

	revertReason, err := getRevertReason(client, tx.Hash())
	if err != nil {
		return err
	}
	return fmt.Errorf("tx[%s] failed, revert reason: %s", tx.Hash().String(), revertReason)
}

func getRevertReason(client *ethclient.Client, txHash common.Hash) (string, error) {
	receipt, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		return "", fmt.Errorf("failed to get receipt: %v", err)
	}

	if receipt.Status != 0 {
		return "", fmt.Errorf("transaction did not fail")
	}

	tx, _, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		return "", fmt.Errorf("failed to get transaction: %v", err)
	}

	msg := ethereum.CallMsg{
		To:   tx.To(),
		Data: tx.Data(),
	}

	res, err := client.CallContract(context.Background(), msg, receipt.BlockNumber)
	if err != nil {
		return "", fmt.Errorf("failed to call contract: %v", err)
	}

	if len(res) < 4 {
		return "No revert reason", nil
	}

	// The revert reason is ABI encoded: first 4 bytes are the function selector for Error(string)
	const errorMethodID = "0x08c379a0"
	if fmt.Sprintf("0x%x", res[:4]) != errorMethodID {
		return "Could not decode revert reason", nil
	}

	abiError, err := abi.JSON(strings.NewReader(`[{"inputs":[{"internalType":"string","name":"reason","type":"string"}],"name":"Error","type":"function"}]`))
	if err != nil {
		return "", fmt.Errorf("failed to parse ABI: %v", err)
	}

	var errorMsg string
	err = abiError.UnpackIntoInterface(&errorMsg, "Error", res[4:])
	if err != nil {
		return "", fmt.Errorf("failed to unpack revert reason: %v", err)
	}

	return errorMsg, nil
}

func contractExists(client *ethclient.Client, address common.Address) (bool, error) {
	code, err := client.CodeAt(context.Background(), address, nil)
	if err != nil {
		return false, fmt.Errorf("failed to check contract existence: %v", err)
	}
	return len(code) > 0, nil
}

func createTransactOpts(client *ethclient.Client, config Config) *bind.TransactOpts {
	seed := bip39.NewSeed(config.Mnemonic, "")
	masterKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		log.Fatalf("Error creating master key: %v", err)
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
			log.Fatalf("Invalid path element %s: %v", p, err)
		}

		if hardened {
			key, err = key.NewChildKey(uint32(index) + bip32.FirstHardenedChild)
		} else {
			key, err = key.NewChildKey(uint32(index))
		}
		if err != nil {
			log.Fatalf("Failed to derive key at index %d: %v", index, err)
		}
	}

	privateKeyECDSA, err := crypto.ToECDSA(key.Key)
	if err != nil {
		log.Fatalf("Error converting to ECDSA: %v", err)
	}

	publicKey := privateKeyECDSA.Public().(*ecdsa.PublicKey)
	fromAddress := crypto.PubkeyToAddress(*publicKey)
	fmt.Printf("Derived Address: %s\n", fromAddress.Hex())

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatalf("Error getting Chain ID: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKeyECDSA, chainID)
	if err != nil {
		log.Fatalf("Error creating signed transactor: %v", err)
	}

	auth.GasLimit = config.GasLimit
	auth.GasFeeCap = config.GasFeeCap
	auth.GasTipCap = config.GasTipCap

	return auth
}
