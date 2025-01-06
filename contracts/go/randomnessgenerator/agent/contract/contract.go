package contract

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log/slog"
	"randomnessgenerator/agent/config"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
)

type RNGClient struct {
	logger *slog.Logger

	config      config.ContractConfig
	ethClient   *ethclient.Client
	contractAPI *Contract
	txAuth      *bind.TransactOpts
}

func NewRNGClient(ctx context.Context, logger *slog.Logger, config config.ContractConfig) (*RNGClient, error) {
	client, err := ethclient.Dial(config.NodeURL)
	if err != nil {
		return nil, fmt.Errorf("eth client dial: %w", err)
	}

	contractAddress := common.HexToAddress(config.ContractAddress)

	contract, err := NewContract(contractAddress, client)
	if err != nil {
		return nil, fmt.Errorf("can't use rng smart-contract API: %w", err)
	}

	exists, err := contractExists(ctx, client, contractAddress)
	if err != nil {
		return nil, fmt.Errorf("contract exists: %w", err)
	}
	if !exists {
		return nil, fmt.Errorf("contract does not exist at address: %s", contractAddress.Hex())
	}

	priKey, _, err := derivePrivateKey(config.Mnemonic, config.DerivationPath)
	if err != nil {
		return nil, fmt.Errorf("derive private key: %w", err)
	}

	auth, err := createTransactor(ctx, client, priKey, config)
	if err != nil {
		return nil, fmt.Errorf("create transactor: %w", err)
	}

	return &RNGClient{
		logger:      logger,
		config:      config,
		ethClient:   client,
		contractAPI: contract,
		txAuth:      auth,
	}, nil
}

func contractExists(ctx context.Context, client *ethclient.Client, address common.Address) (bool, error) {
	code, err := client.CodeAt(ctx, address, nil)
	if err != nil {
		return false, fmt.Errorf("code at: %w", err)
	}
	return len(code) > 0, nil
}

// createTransactor creates a signed transactor for sending transactions
func createTransactor(ctx context.Context, client *ethclient.Client, privateKey *ecdsa.PrivateKey, config config.ContractConfig) (*bind.TransactOpts, error) {
	chainID, err := client.NetworkID(ctx)
	if err != nil {
		return nil, fmt.Errorf("get network ID: %w", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, fmt.Errorf("new keyed transactor with chain id %s: %w", chainID, err)
	}

	auth.GasLimit = config.GasLimit
	auth.GasFeeCap = config.GasFeeCap
	auth.GasTipCap = config.GasTipCap

	return auth, nil
}

// derivePrivateKey derives the ECDSA private key and returns the key along with the address
func derivePrivateKey(mnemonic, derivationPath string) (*ecdsa.PrivateKey, common.Address, error) {
	seed := bip39.NewSeed(mnemonic, "") // No passphrase
	masterKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		return nil, common.Address{}, fmt.Errorf("new master key: %w", err)
	}

	childKey, err := deriveKey(masterKey, derivationPath)
	if err != nil {
		return nil, common.Address{}, fmt.Errorf("derive key from master key and derivatoin path: %w", err)
	}

	privateKeyECDSA, err := crypto.ToECDSA(childKey.Key)
	if err != nil {
		return nil, common.Address{}, fmt.Errorf("bytes to ECDSA: %w", err)
	}

	publicKey := privateKeyECDSA.Public().(*ecdsa.PublicKey)
	fromAddress := crypto.PubkeyToAddress(*publicKey)

	return privateKeyECDSA, fromAddress, nil
}

// deriveKey derives a private key using a BIP32 derivation path
func deriveKey(master *bip32.Key, path string) (*bip32.Key, error) {
	parts := strings.Split(path, "/")
	key := master
	for _, p := range parts[1:] { // Skip the 'm'
		hardened := false
		if strings.HasSuffix(p, "'") {
			hardened = true
			p = strings.TrimSuffix(p, "'")
		}
		index := 0
		_, err := fmt.Sscanf(p, "%d", &index)
		if err != nil {
			return nil, fmt.Errorf("invalid path element %s: %v", p, err)
		}

		if hardened {
			key, err = key.NewChildKey(uint32(index) + bip32.FirstHardenedChild)
		} else {
			key, err = key.NewChildKey(uint32(index))
		}
		if err != nil {
			return nil, fmt.Errorf("failed to derive key at index %d: %v", index, err)
		}
	}
	return key, nil
}
