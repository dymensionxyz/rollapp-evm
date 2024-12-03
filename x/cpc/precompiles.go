package cpc

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
)

const (
	cpcAddrNoncePriceFeed byte = iota + 1
)

// CpcPriceFeedFixedAddress is the address of the price-feed custom precompiled contract.
var CpcPriceFeedFixedAddress common.Address

func init() {
	generatedCpcAddresses := make(map[common.Address]struct{})

	// generateCpcAddress generates a custom precompiled contract address based on the contract address nonce.
	generateCpcAddress := func(contractAddrNonce byte) common.Address {
		if contractAddrNonce == 0 {
			panic("contract address nonce cannot be zero")
		}
		bz := make([]byte, 20)
		bz[0] = 0xCC
		bz[1] = contractAddrNonce
		bz[2] = 0x01 // avoid collision with custom precompiled contracts
		bz[19] = contractAddrNonce

		addr := common.BytesToAddress(bz)
		if _, ok := generatedCpcAddresses[addr]; ok {
			panic(fmt.Sprintf("generated address %s already exists", addr.Hex()))
		}
		generatedCpcAddresses[addr] = struct{}{}

		return addr
	}

	CpcPriceFeedFixedAddress = generateCpcAddress(cpcAddrNoncePriceFeed)
}
