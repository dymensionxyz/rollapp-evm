package cpc

import (
	"github.com/dymensionxyz/rollapp-evm/x/cpc/abi"
	"github.com/ethereum/go-ethereum/common"
	corevm "github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"
	evmkeeper "github.com/evmos/evmos/v12/x/evm/keeper"
	"math/big"
)

// contract

var _ evmkeeper.CustomPrecompiledContractI = &priceFeedCustomPrecompiledContract{}

// priceFeedCustomPrecompiledContract
type priceFeedCustomPrecompiledContract struct {
	keeper    *evmkeeper.Keeper // FIXME: replace with something that can be used to get the price
	executors []evmkeeper.ExtendedCustomPrecompiledContractMethodExecutorI
}

// NewPreFeedCustomPrecompiledContract creates a new price-feed custom precompiled contract.
func NewPreFeedCustomPrecompiledContract(keeper *evmkeeper.Keeper) evmkeeper.CustomPrecompiledContractI {
	contract := &priceFeedCustomPrecompiledContract{
		keeper: keeper,
	}

	contract.executors = []evmkeeper.ExtendedCustomPrecompiledContractMethodExecutorI{
		&priceFeedCustomPrecompiledContractRoGetPrice{contract},
	}

	return contract
}

func (m priceFeedCustomPrecompiledContract) GetName() string {
	return abi.PriceFeedCpcInfo.Name
}

func (m priceFeedCustomPrecompiledContract) GetAddress() common.Address {
	return CpcPriceFeedFixedAddress
}

func (m priceFeedCustomPrecompiledContract) GetMethodExecutors() []evmkeeper.ExtendedCustomPrecompiledContractMethodExecutorI {
	return m.executors
}

// getPrice(string)

var _ evmkeeper.ExtendedCustomPrecompiledContractMethodExecutorI = &priceFeedCustomPrecompiledContractRoGetPrice{}

type priceFeedCustomPrecompiledContractRoGetPrice struct {
	contract *priceFeedCustomPrecompiledContract
}

func (e priceFeedCustomPrecompiledContractRoGetPrice) Execute(_ corevm.ContractRef, _ common.Address, input []byte, _ evmkeeper.CpcExecutorEnv) ([]byte, error) {
	ips, err := abi.PriceFeedCpcInfo.UnpackMethodInput("getPrice", input)
	if err != nil {
		return nil, err
	}

	denom := ips[0].(string)
	var price *big.Int

	{ // FIXME: implement it correctly
		// generate random price
		price = new(big.Int).SetBytes(crypto.Keccak256([]byte(denom)))
	}

	return abi.PriceFeedCpcInfo.PackMethodOutput("getPrice", price, true)
}

func (e priceFeedCustomPrecompiledContractRoGetPrice) Method4BytesSignatures() []byte {
	return []byte{0x52, 0x4f, 0x38, 0x89}
}

func (e priceFeedCustomPrecompiledContractRoGetPrice) RequireGas() uint64 {
	return 500_000
}

func (e priceFeedCustomPrecompiledContractRoGetPrice) ReadOnly() bool {
	return true
}
