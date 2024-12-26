package main

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"example1/contractapi"
	"fmt"
	"net/http"
	"time"

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
	"strings"
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
	ContractAPI     *contractapi.RandomnessGenerator
	ContractAddress common.Address
	DB              *leveldb.DB
	Auth            *bind.TransactOpts
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
	contract, err := contractapi.NewRandomnessGenerator(contractAddress, client)
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

	auth := createTransactOpts(client, cfg)
	return &RNGAgent{
		EthClient:       client,
		ContractAddress: contractAddress,
		DB:              db,
		ContractAPI:     contract,
		Auth:            auth,
	}, nil
}

type EventNewRandomnessRequest struct {
	RandomnessID *big.Int
}

func generateRandomUint256() (*big.Int, error) {
	maxInt := new(big.Int).Lsh(big.NewInt(1), 256)
	maxInt.Sub(maxInt, big.NewInt(1))

	randomNumber, err := rand.Int(rand.Reader, maxInt)
	if err != nil {
		return nil, fmt.Errorf("failed to generate random uint256: %v", err)
	}

	return randomNumber, nil
}

func (a *RNGAgent) ListenForSmartContractEvents() error {
	contractABI, err := abi.JSON(strings.NewReader(contractapi.RandomnessGeneratorMetaData.ABI))
	if err != nil {
		return fmt.Errorf("error parsing contract ABI: %v", err)
	}

	eventNewRandomnessRequestID := contractABI.Events["EventNewRandomnessRequest"].ID

	query := ethereum.FilterQuery{
		Addresses: []common.Address{a.ContractAddress},
		Topics:    [][]common.Hash{{eventNewRandomnessRequestID}},
	}

	logs := make(chan types.Log)
	sub, err := a.EthClient.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		return fmt.Errorf("error subscribing to logs: %v", err)
	}

	//sub.Unsubscribe()

	println("KAL SOBAKI")

	for {
		select {

		case err := <-sub.Err():
			log.Printf("error with subscription: %v", err)

			dialRetries := 3
			for err != nil && dialRetries > 0 {
				log.Println("dialRetries: ", dialRetries)
				if dialRetries != 3 {
					time.Sleep(1000 * time.Millisecond)
				}
				a.EthClient, err = ethclient.Dial("ws://127.0.0.1:8546")
				dialRetries -= 1
			}

			if err == nil {
				println("OK")
			}

			return err
		case l := <-logs:
			var event EventNewRandomnessRequest
			err := contractABI.UnpackIntoInterface(&event, "EventNewRandomnessRequest", l.Data)
			if err != nil {
				log.Printf("error unpacking log: %v", err)
				continue
			}

			randID := event.RandomnessID
			fmt.Printf("tx[%s], Randomness requested: ID: %s\n", l.TxHash.Hex(), randID.String())

			randomness, err := generateRandomUint256()
			if err != nil {
				log.Printf("can't generate u256 random: %v", err)
				continue
			}

			err = a.DB.Put(randID.Bytes(), randomness.Bytes(), nil)
			if err != nil {
				log.Printf("error putting [key;value] = [%s;%d] into DB: %v", randID.String(), randomness, err)
				continue
			}

			log.Printf("[%s:%s]", randID.String(), randomness.String())

			_, err = a.ContractAPI.PostRandomness(a.Auth, randID, randomness)
			if err != nil {
				log.Printf("%v", err)
			}
		case <-time.After(60 * time.Second):
			log.Println("timeout 1min retry dial")
			// https://github.com/ethereum/go-ethereum/blob/245f3146c26698193c4b479e7bc5825b058c444a/rpc/subscription.go#L243
			sub.Unsubscribe()
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
		if err == leveldb.ErrNotFound {
			http.Error(w, "Randomness not found", http.StatusNotFound)
		} else {
			http.Error(w, fmt.Sprintf("Error retrieving randomness: %v", err), http.StatusInternalServerError)
		}
		return
	}

	randomness := new(big.Int).SetBytes(randomnessBytes)
	_, _ = fmt.Fprintf(w, randomness.String())
}

func (a *RNGAgent) StartHTTPServer(address string) error {
	http.HandleFunc("/randomness", a.handleGetRandomness)

	log.Printf("Starting HTTP server at %s", address)
	err := http.ListenAndServe(address, nil)
	return fmt.Errorf("failed to start HTTP server: %v", err)
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	config := Config{
		NodeURL:            "ws://127.0.0.1:8546", // Local Hardhat node
		Mnemonic:           "depend version wrestle document episode celery nuclear main penalty hundred trap scale candy donate search glory build valve round athlete become beauty indicate hamster",
		HexContractAddress: "0x676E400d0200Ac8f3903A3CDC7cc3feaF21004d0", // Replace with the correct contract address
		DerivationPath:     "m/44'/60'/0'/0/0",
		GasLimit:           60000,
		GasFeeCap:          big.NewInt(30000000000), // 30 Gwei
		GasTipCap:          big.NewInt(2000000000),  // 2 Gwei
		HTTPServerAddr:     ":8080",
	}

	ctx := context.Background()
	g, ctx := errgroup.WithContext(ctx)

	agent, err := NewRNGAgent(config)
	if err != nil {
		log.Fatalf("error while creating rng agent: %v", err)
	}

	g.Go(agent.ListenForSmartContractEvents)
	g.Go(func() error { return agent.StartHTTPServer(config.HTTPServerAddr) })

	if err := g.Wait(); err != nil {
		log.Fatalf("failed: %v", err)
	}
}

func contractExists(client *ethclient.Client, address common.Address) (bool, error) {
	code, err := client.CodeAt(context.Background(), address, nil) // Проверка кода на контракте
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
