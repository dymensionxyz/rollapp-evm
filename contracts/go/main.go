package main

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
)

func main() {
	// Enhanced Logging
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Configuration
	nodeURL := "http://127.0.0.1:8545"                                                                                                                                // URL of your local Hardhat node
	mnemonic := "penalty useful movie rookie toilet album abuse rude sing size meadow noodle wise pen castle trust proud chalk loud era universe can reflect clarify" // Your Hardhat mnemonic
	contractAddressHex := "0x08ae2aBDCa46Aa7aF765369Df47395EfA62ba3F9"                                                                                                // Deployed contract address
	baseAddressHex := "0xYourBaseTokenAddress"                                                                                                                        // Replace with the actual token/base address
	quoteAddressHex := "0xYourQuoteTokenAddress"                                                                                                                      // Replace with the actual token/quote address

	// 1. Derive the private key from the mnemonic
	seed := bip39.NewSeed(mnemonic, "") // Passphrase not used
	masterKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		log.Fatalf("Error creating master key: %v", err)
	}

	// Derived path for the first account according to BIP44 (m/44'/60'/0'/0/0)
	derivationPath := "m/44'/60'/0'/0/0"
	childKey, err := deriveKey(masterKey, derivationPath)
	if err != nil {
		log.Fatalf("Error deriving key: %v", err)
	}

	privateKeyECDSA, err := crypto.ToECDSA(childKey.Key)
	if err != nil {
		log.Fatalf("Error converting to ECDSA: %v", err)
	}

	fmt.Printf("Clave pública derivada: %x\n", privateKeyECDSA.Public())

	// 2. Connect to the Ethereum node
	client, err := ethclient.Dial(nodeURL)
	if err != nil {
		log.Fatalf("Error connecting to Ethereum node: %v", err)
	}
	defer client.Close()
	fmt.Println("Connected to Ethereum node at", nodeURL)

	// 3. Check account balance
	publicKey := privateKeyECDSA.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatalf("Error converting public key")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Printf("Using account: %s\n", fromAddress.Hex())

	balance, err := client.BalanceAt(context.Background(), fromAddress, nil)
	if err != nil {
		log.Fatalf("Failed to retrieve account balance: %v", err)
	}
	fmt.Printf("Account balance: %s wei\n", balance.String())

	// 4. Create a signed transactor
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatalf("Error getting Chain ID: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKeyECDSA, chainID)
	if err != nil {
		log.Fatalf("Error creating signed transactor: %v", err)
	}

	// Optional: Set gas limit and prices
	auth.GasLimit = uint64(500000)          // Increased gas limit
	auth.GasPrice = big.NewInt(30000000000) // 30 Gwei, adjust as needed

	// 5. Instantiate the contract
	contractAddress := common.HexToAddress(contractAddressHex)
	priceOracle, err := NewContracts(contractAddress, client)
	if err != nil {
		log.Fatalf("Error instantiating the contract: %v", err)
	}
	fmt.Println("PriceOracle contract instantiated at", contractAddress.Hex())

	// 6. Query the contract owner
	ownerAddress, err := priceOracle.Owner(&bind.CallOpts{
		Context: context.Background(),
	})
	if err != nil {
		log.Fatalf("Error querying contract owner: %v", err)
	}
	fmt.Printf("Contract Owner: %s\n", ownerAddress.Hex())

	// 7. Verify if the current account is the owner
	if ownerAddress.Hex() != fromAddress.Hex() {
		log.Fatalf("Error: The current account (%s) is not the owner of the contract (%s).", fromAddress.Hex(), ownerAddress.Hex())
	}
	fmt.Println("Verification successful: You are the owner of the contract.")

	// 8. Prepare parameters for updatePrice
	baseAddress := common.HexToAddress(baseAddressHex)
	quoteAddress := common.HexToAddress(quoteAddressHex)
	price := big.NewInt(1000) // For example, 1000 (adjust according to contract precision)

	// Generate a valid Merkle proof
	merkleProof, err := generateMerkleProof(1000)
	if err != nil {
		log.Fatalf("Failed to generate Merkle proof: %v", err)
	}

	// Populate PriceProof with actual data
	priceProof := PriceOraclePriceProof{
		CreationHeight:     big.NewInt(123456),     // Replace with actual block height
		CreationTimeUnixMs: big.NewInt(1617181920), // Replace with actual timestamp in ms
		Height:             big.NewInt(12345678),   // Replace with current block height
		Revision:           big.NewInt(1),          // Replace with the correct revision number
		MerkleProof:        merkleProof,            // Use the generated proof
	}

	// Prepare the PriceWithProof structure
	priceWithProof := PriceOraclePriceWithProof{
		Price: price,
		Proof: priceProof,
	}

	fmt.Printf("Auth From Address: %s\n", auth.From.Hex())
	if auth.From != fromAddress {
		log.Fatalf("auth.From (%s) no coincide con fromAddress (%s)", auth.From.Hex(), fromAddress.Hex())
	} else {
		fmt.Println("auth.From coincide con fromAddress")
	}
	fmt.Printf("Chain ID: %s\n", chainID.String())
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatalf("Error obteniendo el nonce: %v", err)
	}
	fmt.Printf("Nonce: %d\n", nonce)
	auth.Nonce = big.NewInt(int64(nonce))
	fmt.Printf("Base Address: %s\n", baseAddress.Hex())
	fmt.Printf("Quote Address: %s\n", quoteAddress.Hex())
	fmt.Printf("Price: %s\n", price.String())
	fmt.Printf("PriceProof: %+v\n", priceProof)
	fmt.Printf("PriceWithProof: %+v\n", priceWithProof)

	// 9. Call updatePrice
	tx, err := priceOracle.UpdatePrice(auth, baseAddress, quoteAddress, priceWithProof)
	if err != nil {
		log.Fatalf("Error calling updatePrice: %v", err)
	}

	fmt.Printf("Transaction sent: %s\n", tx.Hash().Hex())

	// 10. Wait for the transaction to be mined
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatalf("Error waiting for transaction confirmation: %v", err)
	}

	if receipt.Status == 1 {
		fmt.Println("Price updated successfully")
	} else {
		fmt.Println("Transaction failed")
		fmt.Printf("Transaction Receipt: %+v\n", receipt)

		// Attempt to extract the revert reason
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
		fmt.Sscanf(p, "%d", &index)
		var err error
		if hardened {
			key, err = key.NewChildKey(uint32(index) + bip32.FirstHardenedChild)
		} else {
			key, err = key.NewChildKey(uint32(index))
		}
		if err != nil {
			return nil, err
		}
	}
	return key, nil
}

// getRevertReason attempts to extract the revert reason from a failed transaction
func getRevertReason(client *ethclient.Client, txHash common.Hash) (string, error) {
	// Get the transaction
	tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		return "", fmt.Errorf("failed to get transaction: %v", err)
	}
	if isPending {
		return "", fmt.Errorf("transaction is still pending")
	}

	// Get the receipt
	receipt, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		return "", fmt.Errorf("failed to get receipt: %v", err)
	}

	// Check if status is failed
	if receipt.Status != 0 {
		return "", fmt.Errorf("transaction did not fail")
	}

	// Get the revert reason by calling eth_call with the same data
	msg := ethereum.CallMsg{
		To:   tx.To(),
		Data: tx.Data(),
		// Include gas, gas price, value if necessary
	}

	// Perform a call
	res, err := client.CallContract(context.Background(), msg, receipt.BlockNumber)
	if err != nil {
		return "", fmt.Errorf("failed to call contract: %v", err)
	}

	if len(res) < 4 {
		return "No revert reason", nil
	}

	// The revert reason is ABI encoded: first 4 bytes are the function selector for Error(string)
	// The rest is the ABI-encoded string
	methodID := res[:4]
	// The method ID for Error(string) is 0x08c379a0
	if bytes.Equal(methodID, []byte{0x08, 0xc3, 0x79, 0xa0}) { // Error(string)
		// Decode the string
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

	return "Could not decode revert reason", nil
}

// generateMerkleProof generates a valid Merkle proof for the given price
// You need to implement this based on your Merkle tree structure
func generateMerkleProof(price int64) ([]byte, error) {
	// Placeholder implementation
	// TODO: Implement Merkle proof generation based on your application's requirements

	// Example: Return a dummy Merkle proof
	return []byte{0x12, 0x34, 0x56, 0x78}, nil
}
