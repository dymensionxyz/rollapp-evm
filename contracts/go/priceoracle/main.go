package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
)

// Config holds the configuration parameters
type Config struct {
	NodeURL           string
	Mnemonic          string
	ContractAddress   string
	BaseTokenAddress  string
	QuoteTokenAddress string
	DerivationPath    string
	GasLimit          uint64
	GasFeeCap         *big.Int
	GasTipCap         *big.Int
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Initialize configuration
	config := Config{
		NodeURL:           "http://127.0.0.1:8545", // Local Hardhat node
		Mnemonic:          "depend version wrestle document episode celery nuclear main penalty hundred trap scale candy donate search glory build valve round athlete become beauty indicate hamster",
		ContractAddress:   "0x676E400d0200Ac8f3903A3CDC7cc3feaF21004d0",
		BaseTokenAddress:  "0xde0b295669a9fd93d5f28d9ec85e40f4cb697bae",
		QuoteTokenAddress: "0x0000000000000000000000000000000000000000",
		DerivationPath:    "m/44'/60'/0'/0/0",
		GasLimit:          60000,
		GasFeeCap:         big.NewInt(30000000000), // 30 Gwei
		GasTipCap:         big.NewInt(2000000000),  // 2 Gwei
	}

	// Step 1: Derive the private key from the mnemonic
	privateKey, fromAddress := derivePrivateKey(config)
	fmt.Printf("Derived Address: %s\n", fromAddress.Hex())

	// Step 2: Connect to the Ethereum node
	client := connectToEthereumNode(config.NodeURL)
	defer client.Close()

	// Step 3: Check account balance
	checkAccountBalance(client, fromAddress)

	// Step 4: Create a signed transactor
	auth := createTransactor(client, privateKey, config)

	// Step 5: Instantiate the contract
	priceOracle := instantiateContract(config.ContractAddress, client)

	// Step 6: Verify ownership
	verifyOwnership(priceOracle, fromAddress)

	// Step 7: Prepare parameters for updatePrice
	priceWithProof := preparePriceWithProof()

	// Step 8: Call updatePrice
	tx := callUpdatePrice(client, priceOracle, auth, config, priceWithProof)

	// Step 9: Wait for the transaction to be mined
	waitForTransaction(client, tx)
}

// derivePrivateKey derives the ECDSA private key and returns the key along with the address
func derivePrivateKey(config Config) (*ecdsa.PrivateKey, common.Address) {
	seed := bip39.NewSeed(config.Mnemonic, "") // No passphrase
	masterKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		log.Fatalf("Error creating master key: %v", err)
	}

	childKey, err := deriveKey(masterKey, config.DerivationPath)
	if err != nil {
		log.Fatalf("Error deriving key: %v", err)
	}

	privateKeyECDSA, err := crypto.ToECDSA(childKey.Key)
	if err != nil {
		log.Fatalf("Error converting to ECDSA: %v", err)
	}

	publicKey := privateKeyECDSA.Public().(*ecdsa.PublicKey)
	fromAddress := crypto.PubkeyToAddress(*publicKey)

	return privateKeyECDSA, fromAddress
}

// connectToEthereumNode establishes a connection to the Ethereum node
func connectToEthereumNode(nodeURL string) *ethclient.Client {
	client, err := ethclient.Dial(nodeURL)
	if err != nil {
		log.Fatalf("Error connecting to Ethereum node: %v", err)
	}
	fmt.Println("Connected to Ethereum node at", nodeURL)
	return client
}

// checkAccountBalance retrieves and displays the account balance
func checkAccountBalance(client *ethclient.Client, address common.Address) {
	balance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Fatalf("Failed to retrieve account balance: %v", err)
	}
	fmt.Printf("Account balance: %s wei\n", balance.String())
}

// createTransactor creates a signed transactor for sending transactions
func createTransactor(client *ethclient.Client, privateKey *ecdsa.PrivateKey, config Config) *bind.TransactOpts {
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatalf("Error getting Chain ID: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatalf("Error creating signed transactor: %v", err)
	}

	auth.GasLimit = config.GasLimit
	auth.GasFeeCap = config.GasFeeCap
	auth.GasTipCap = config.GasTipCap

	return auth
}

// instantiateContract initializes the contract instance
func instantiateContract(contractAddressHex string, client *ethclient.Client) *Contract {
	contractAddress := common.HexToAddress(contractAddressHex)
	priceOracle, err := NewContract(contractAddress, client)
	if err != nil {
		log.Fatalf("Error instantiating the contract: %v", err)
	}
	fmt.Println("PriceOracle contract instantiated at", contractAddress.Hex())
	return priceOracle
}

// verifyOwnership checks if the derived address is the contract owner
func verifyOwnership(priceOracle *Contract, fromAddress common.Address) {
	ownerAddress, err := priceOracle.Owner(&bind.CallOpts{Context: context.Background()})
	if err != nil {
		log.Fatalf("Error querying contract owner: %v", err)
	}

	if !addressesEqual(ownerAddress, fromAddress) {
		log.Fatalf("Error: The current account (%s) is not the owner of the contract (%s).", fromAddress.Hex(), ownerAddress.Hex())
	}
	fmt.Println("Verification successful: You are the owner of the contract.")
}

// preparePriceWithProof constructs the PriceWithProof structure
func preparePriceWithProof() PriceOraclePriceWithProof {
	// Generate a valid Merkle proof
	merkleProof, err := generateMerkleProof(1000)
	if err != nil {
		log.Fatalf("Failed to generate Merkle proof: %v", err)
	}

	// Populate PriceProof with actual data
	priceProof := PriceOraclePriceProof{
		CreationHeight:     big.NewInt(123456),                      // Replace with actual block height
		CreationTimeUnixMs: big.NewInt(time.Now().UnixNano() / 1e6), // Current timestamp in ms
		Height:             big.NewInt(12345678),                    // Replace with current block height
		Revision:           big.NewInt(1),                           // Current revision number
		MerkleProof:        merkleProof,                             // Generated proof
	}

	// Prepare the PriceWithProof structure
	priceWithProof := PriceOraclePriceWithProof{
		Price: big.NewInt(1000),
		Proof: priceProof,
	}

	return priceWithProof
}

// callUpdatePrice sends the updatePrice transaction to the contract
func callUpdatePrice(client *ethclient.Client, priceOracle *Contract, auth *bind.TransactOpts, config Config, priceWithProof PriceOraclePriceWithProof) *types.Transaction {
	baseAddress := common.HexToAddress(config.BaseTokenAddress)
	quoteAddress := common.HexToAddress(config.QuoteTokenAddress)

	nonce, err := client.PendingNonceAt(context.Background(), auth.From)
	if err != nil {
		log.Fatalf("Error getting nonce: %v", err)
	}
	auth.Nonce = big.NewInt(int64(nonce))

	tx, err := priceOracle.UpdatePrice(auth, baseAddress, quoteAddress, priceWithProof)
	if err != nil {
		log.Fatalf("Error calling updatePrice: %v", err)
	}

	fmt.Printf("Transaction sent: %s\n", tx.Hash().Hex())
	return tx
}

// waitForTransaction waits for the transaction to be mined and checks its status
func waitForTransaction(client *ethclient.Client, tx *types.Transaction) {
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatalf("Error waiting for transaction confirmation: %v", err)
	}

	if receipt.Status == 1 {
		fmt.Println("Price updated successfully")
	} else {
		fmt.Println("Transaction failed")
		fmt.Printf("Transaction Receipt: %+v\n", receipt)

		revertReason, err := getRevertReason(client, tx.Hash())
		if err != nil {
			log.Printf("Failed to get revert reason: %v", err)
		} else {
			fmt.Printf("Revert Reason: %s\n", revertReason)
		}
	}
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

// getRevertReason attempts to extract the revert reason from a failed transaction
func getRevertReason(client *ethclient.Client, txHash common.Hash) (string, error) {
	// Get the transaction receipt
	receipt, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		return "", fmt.Errorf("failed to get receipt: %v", err)
	}

	// Check if transaction failed
	if receipt.Status != 0 {
		return "", fmt.Errorf("transaction did not fail")
	}

	// Get the transaction data
	tx, _, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		return "", fmt.Errorf("failed to get transaction: %v", err)
	}

	// Prepare the call message
	msg := ethereum.CallMsg{
		To:   tx.To(),
		Data: tx.Data(),
	}

	// Perform a call to get the revert reason
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

	// Decode the revert reason
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

// generateMerkleProof generates a valid Merkle proof for the given price
// Implement this based on your Merkle tree structure
func generateMerkleProof(price int64) ([]byte, error) {
	// Placeholder implementation
	// TODO: Implement Merkle proof generation based on your application's requirements

	// Example: Return a dummy Merkle proof
	return []byte{0x12, 0x34, 0x56, 0x78}, nil
}

// addressesEqual compares two Ethereum addresses case-insensitively
func addressesEqual(a, b common.Address) bool {
	return strings.ToLower(a.Hex()) == strings.ToLower(b.Hex())
}
