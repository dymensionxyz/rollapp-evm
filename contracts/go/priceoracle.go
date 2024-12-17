// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// PriceOracleAssetInfo is an auto generated low-level Go binding around an user-defined struct.
type PriceOracleAssetInfo struct {
	LocalNetworkName      common.Address
	OracleNetworkName     string
	LocalNetworkPrecision *big.Int
}

// PriceOracleGetPriceResponse is an auto generated low-level Go binding around an user-defined struct.
type PriceOracleGetPriceResponse struct {
	Price     *big.Int
	IsInverse bool
}

// PriceOraclePriceProof is an auto generated low-level Go binding around an user-defined struct.
type PriceOraclePriceProof struct {
	CreationHeight     *big.Int
	CreationTimeUnixMs *big.Int
	Height             *big.Int
	Revision           *big.Int
	MerkleProof        []byte
}

// PriceOraclePriceWithProof is an auto generated low-level Go binding around an user-defined struct.
type PriceOraclePriceWithProof struct {
	Price *big.Int
	Proof PriceOraclePriceProof
}

// ContractsMetaData contains all meta data concerning the Contracts contract.
var ContractsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_expirationOffset\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"localNetworkName\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"oracleNetworkName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"localNetworkPrecision\",\"type\":\"uint256\"}],\"internalType\":\"structPriceOracle.AssetInfo[]\",\"name\":\"_assetInfos\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"boundThreshold\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"initializer\",\"type\":\"address\"}],\"name\":\"OracleInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"base\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"quote\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"PriceUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"SCALE_FACTOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"boundThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"expirationOffset\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"base\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quote\",\"type\":\"address\"}],\"name\":\"getPrice\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"is_inverse\",\"type\":\"bool\"}],\"internalType\":\"structPriceOracle.GetPriceResponse\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"localNetworkToOracleNetworkDenoms\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"precisionMapping\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"base\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quote\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"priceInBound\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"prices_cache\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"base\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quote\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"creationHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"creationTimeUnixMs\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revision\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"merkleProof\",\"type\":\"bytes\"}],\"internalType\":\"structPriceOracle.PriceProof\",\"name\":\"proof\",\"type\":\"tuple\"}],\"internalType\":\"structPriceOracle.PriceWithProof\",\"name\":\"priceWithProof\",\"type\":\"tuple\"}],\"name\":\"updatePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ContractsABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractsMetaData.ABI instead.
var ContractsABI = ContractsMetaData.ABI

// Contracts is an auto generated Go binding around an Ethereum contract.
type Contracts struct {
	ContractsCaller     // Read-only binding to the contract
	ContractsTransactor // Write-only binding to the contract
	ContractsFilterer   // Log filterer for contract events
}

// ContractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractsSession struct {
	Contract     *Contracts        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractsCallerSession struct {
	Contract *ContractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ContractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractsTransactorSession struct {
	Contract     *ContractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ContractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractsRaw struct {
	Contract *Contracts // Generic contract binding to access the raw methods on
}

// ContractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractsCallerRaw struct {
	Contract *ContractsCaller // Generic read-only contract binding to access the raw methods on
}

// ContractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractsTransactorRaw struct {
	Contract *ContractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContracts creates a new instance of Contracts, bound to a specific deployed contract.
func NewContracts(address common.Address, backend bind.ContractBackend) (*Contracts, error) {
	contract, err := bindContracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// NewContractsCaller creates a new read-only instance of Contracts, bound to a specific deployed contract.
func NewContractsCaller(address common.Address, caller bind.ContractCaller) (*ContractsCaller, error) {
	contract, err := bindContracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsCaller{contract: contract}, nil
}

// NewContractsTransactor creates a new write-only instance of Contracts, bound to a specific deployed contract.
func NewContractsTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractsTransactor, error) {
	contract, err := bindContracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsTransactor{contract: contract}, nil
}

// NewContractsFilterer creates a new log filterer instance of Contracts, bound to a specific deployed contract.
func NewContractsFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractsFilterer, error) {
	contract, err := bindContracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractsFilterer{contract: contract}, nil
}

// bindContracts binds a generic wrapper to an already deployed contract.
func bindContracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.ContractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transact(opts, method, params...)
}

// SCALEFACTOR is a free data retrieval call binding the contract method 0xce4b5bbe.
//
// Solidity: function SCALE_FACTOR() view returns(uint256)
func (_Contracts *ContractsCaller) SCALEFACTOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "SCALE_FACTOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SCALEFACTOR is a free data retrieval call binding the contract method 0xce4b5bbe.
//
// Solidity: function SCALE_FACTOR() view returns(uint256)
func (_Contracts *ContractsSession) SCALEFACTOR() (*big.Int, error) {
	return _Contracts.Contract.SCALEFACTOR(&_Contracts.CallOpts)
}

// SCALEFACTOR is a free data retrieval call binding the contract method 0xce4b5bbe.
//
// Solidity: function SCALE_FACTOR() view returns(uint256)
func (_Contracts *ContractsCallerSession) SCALEFACTOR() (*big.Int, error) {
	return _Contracts.Contract.SCALEFACTOR(&_Contracts.CallOpts)
}

// BoundThreshold is a free data retrieval call binding the contract method 0x5fca4148.
//
// Solidity: function boundThreshold() view returns(uint256)
func (_Contracts *ContractsCaller) BoundThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "boundThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BoundThreshold is a free data retrieval call binding the contract method 0x5fca4148.
//
// Solidity: function boundThreshold() view returns(uint256)
func (_Contracts *ContractsSession) BoundThreshold() (*big.Int, error) {
	return _Contracts.Contract.BoundThreshold(&_Contracts.CallOpts)
}

// BoundThreshold is a free data retrieval call binding the contract method 0x5fca4148.
//
// Solidity: function boundThreshold() view returns(uint256)
func (_Contracts *ContractsCallerSession) BoundThreshold() (*big.Int, error) {
	return _Contracts.Contract.BoundThreshold(&_Contracts.CallOpts)
}

// ExpirationOffset is a free data retrieval call binding the contract method 0xa12ae8e4.
//
// Solidity: function expirationOffset() view returns(uint256)
func (_Contracts *ContractsCaller) ExpirationOffset(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "expirationOffset")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExpirationOffset is a free data retrieval call binding the contract method 0xa12ae8e4.
//
// Solidity: function expirationOffset() view returns(uint256)
func (_Contracts *ContractsSession) ExpirationOffset() (*big.Int, error) {
	return _Contracts.Contract.ExpirationOffset(&_Contracts.CallOpts)
}

// ExpirationOffset is a free data retrieval call binding the contract method 0xa12ae8e4.
//
// Solidity: function expirationOffset() view returns(uint256)
func (_Contracts *ContractsCallerSession) ExpirationOffset() (*big.Int, error) {
	return _Contracts.Contract.ExpirationOffset(&_Contracts.CallOpts)
}

// GetPrice is a free data retrieval call binding the contract method 0xac41865a.
//
// Solidity: function getPrice(address base, address quote) view returns((uint256,bool))
func (_Contracts *ContractsCaller) GetPrice(opts *bind.CallOpts, base common.Address, quote common.Address) (PriceOracleGetPriceResponse, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getPrice", base, quote)

	if err != nil {
		return *new(PriceOracleGetPriceResponse), err
	}

	out0 := *abi.ConvertType(out[0], new(PriceOracleGetPriceResponse)).(*PriceOracleGetPriceResponse)

	return out0, err

}

// GetPrice is a free data retrieval call binding the contract method 0xac41865a.
//
// Solidity: function getPrice(address base, address quote) view returns((uint256,bool))
func (_Contracts *ContractsSession) GetPrice(base common.Address, quote common.Address) (PriceOracleGetPriceResponse, error) {
	return _Contracts.Contract.GetPrice(&_Contracts.CallOpts, base, quote)
}

// GetPrice is a free data retrieval call binding the contract method 0xac41865a.
//
// Solidity: function getPrice(address base, address quote) view returns((uint256,bool))
func (_Contracts *ContractsCallerSession) GetPrice(base common.Address, quote common.Address) (PriceOracleGetPriceResponse, error) {
	return _Contracts.Contract.GetPrice(&_Contracts.CallOpts, base, quote)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Contracts *ContractsCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "initialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Contracts *ContractsSession) Initialized() (bool, error) {
	return _Contracts.Contract.Initialized(&_Contracts.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Contracts *ContractsCallerSession) Initialized() (bool, error) {
	return _Contracts.Contract.Initialized(&_Contracts.CallOpts)
}

// LocalNetworkToOracleNetworkDenoms is a free data retrieval call binding the contract method 0x5282fe92.
//
// Solidity: function localNetworkToOracleNetworkDenoms(address ) view returns(string)
func (_Contracts *ContractsCaller) LocalNetworkToOracleNetworkDenoms(opts *bind.CallOpts, arg0 common.Address) (string, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "localNetworkToOracleNetworkDenoms", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// LocalNetworkToOracleNetworkDenoms is a free data retrieval call binding the contract method 0x5282fe92.
//
// Solidity: function localNetworkToOracleNetworkDenoms(address ) view returns(string)
func (_Contracts *ContractsSession) LocalNetworkToOracleNetworkDenoms(arg0 common.Address) (string, error) {
	return _Contracts.Contract.LocalNetworkToOracleNetworkDenoms(&_Contracts.CallOpts, arg0)
}

// LocalNetworkToOracleNetworkDenoms is a free data retrieval call binding the contract method 0x5282fe92.
//
// Solidity: function localNetworkToOracleNetworkDenoms(address ) view returns(string)
func (_Contracts *ContractsCallerSession) LocalNetworkToOracleNetworkDenoms(arg0 common.Address) (string, error) {
	return _Contracts.Contract.LocalNetworkToOracleNetworkDenoms(&_Contracts.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contracts *ContractsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contracts *ContractsSession) Owner() (common.Address, error) {
	return _Contracts.Contract.Owner(&_Contracts.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contracts *ContractsCallerSession) Owner() (common.Address, error) {
	return _Contracts.Contract.Owner(&_Contracts.CallOpts)
}

// PrecisionMapping is a free data retrieval call binding the contract method 0x84be6db4.
//
// Solidity: function precisionMapping(address ) view returns(uint256)
func (_Contracts *ContractsCaller) PrecisionMapping(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "precisionMapping", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PrecisionMapping is a free data retrieval call binding the contract method 0x84be6db4.
//
// Solidity: function precisionMapping(address ) view returns(uint256)
func (_Contracts *ContractsSession) PrecisionMapping(arg0 common.Address) (*big.Int, error) {
	return _Contracts.Contract.PrecisionMapping(&_Contracts.CallOpts, arg0)
}

// PrecisionMapping is a free data retrieval call binding the contract method 0x84be6db4.
//
// Solidity: function precisionMapping(address ) view returns(uint256)
func (_Contracts *ContractsCallerSession) PrecisionMapping(arg0 common.Address) (*big.Int, error) {
	return _Contracts.Contract.PrecisionMapping(&_Contracts.CallOpts, arg0)
}

// PriceInBound is a free data retrieval call binding the contract method 0x27b5a917.
//
// Solidity: function priceInBound(address base, address quote, uint256 price) view returns(bool)
func (_Contracts *ContractsCaller) PriceInBound(opts *bind.CallOpts, base common.Address, quote common.Address, price *big.Int) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "priceInBound", base, quote, price)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PriceInBound is a free data retrieval call binding the contract method 0x27b5a917.
//
// Solidity: function priceInBound(address base, address quote, uint256 price) view returns(bool)
func (_Contracts *ContractsSession) PriceInBound(base common.Address, quote common.Address, price *big.Int) (bool, error) {
	return _Contracts.Contract.PriceInBound(&_Contracts.CallOpts, base, quote, price)
}

// PriceInBound is a free data retrieval call binding the contract method 0x27b5a917.
//
// Solidity: function priceInBound(address base, address quote, uint256 price) view returns(bool)
func (_Contracts *ContractsCallerSession) PriceInBound(base common.Address, quote common.Address, price *big.Int) (bool, error) {
	return _Contracts.Contract.PriceInBound(&_Contracts.CallOpts, base, quote, price)
}

// PricesCache is a free data retrieval call binding the contract method 0x5e0ce1f6.
//
// Solidity: function prices_cache(address , address ) view returns(uint256 price, uint256 expiration, bool exists)
func (_Contracts *ContractsCaller) PricesCache(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (struct {
	Price      *big.Int
	Expiration *big.Int
	Exists     bool
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "prices_cache", arg0, arg1)

	outstruct := new(struct {
		Price      *big.Int
		Expiration *big.Int
		Exists     bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Price = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Expiration = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Exists = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// PricesCache is a free data retrieval call binding the contract method 0x5e0ce1f6.
//
// Solidity: function prices_cache(address , address ) view returns(uint256 price, uint256 expiration, bool exists)
func (_Contracts *ContractsSession) PricesCache(arg0 common.Address, arg1 common.Address) (struct {
	Price      *big.Int
	Expiration *big.Int
	Exists     bool
}, error) {
	return _Contracts.Contract.PricesCache(&_Contracts.CallOpts, arg0, arg1)
}

// PricesCache is a free data retrieval call binding the contract method 0x5e0ce1f6.
//
// Solidity: function prices_cache(address , address ) view returns(uint256 price, uint256 expiration, bool exists)
func (_Contracts *ContractsCallerSession) PricesCache(arg0 common.Address, arg1 common.Address) (struct {
	Price      *big.Int
	Expiration *big.Int
	Exists     bool
}, error) {
	return _Contracts.Contract.PricesCache(&_Contracts.CallOpts, arg0, arg1)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Contracts *ContractsTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Contracts *ContractsSession) Initialize() (*types.Transaction, error) {
	return _Contracts.Contract.Initialize(&_Contracts.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Contracts *ContractsTransactorSession) Initialize() (*types.Transaction, error) {
	return _Contracts.Contract.Initialize(&_Contracts.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contracts *ContractsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contracts *ContractsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.TransferOwnership(&_Contracts.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contracts *ContractsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.TransferOwnership(&_Contracts.TransactOpts, newOwner)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0x42630dee.
//
// Solidity: function updatePrice(address base, address quote, (uint256,(uint256,uint256,uint256,uint256,bytes)) priceWithProof) returns()
func (_Contracts *ContractsTransactor) UpdatePrice(opts *bind.TransactOpts, base common.Address, quote common.Address, priceWithProof PriceOraclePriceWithProof) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "updatePrice", base, quote, priceWithProof)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0x42630dee.
//
// Solidity: function updatePrice(address base, address quote, (uint256,(uint256,uint256,uint256,uint256,bytes)) priceWithProof) returns()
func (_Contracts *ContractsSession) UpdatePrice(base common.Address, quote common.Address, priceWithProof PriceOraclePriceWithProof) (*types.Transaction, error) {
	return _Contracts.Contract.UpdatePrice(&_Contracts.TransactOpts, base, quote, priceWithProof)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0x42630dee.
//
// Solidity: function updatePrice(address base, address quote, (uint256,(uint256,uint256,uint256,uint256,bytes)) priceWithProof) returns()
func (_Contracts *ContractsTransactorSession) UpdatePrice(base common.Address, quote common.Address, priceWithProof PriceOraclePriceWithProof) (*types.Transaction, error) {
	return _Contracts.Contract.UpdatePrice(&_Contracts.TransactOpts, base, quote, priceWithProof)
}

// ContractsOracleInitializedIterator is returned from FilterOracleInitialized and is used to iterate over the raw logs and unpacked data for OracleInitialized events raised by the Contracts contract.
type ContractsOracleInitializedIterator struct {
	Event *ContractsOracleInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsOracleInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsOracleInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsOracleInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsOracleInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsOracleInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsOracleInitialized represents a OracleInitialized event raised by the Contracts contract.
type ContractsOracleInitialized struct {
	Initializer common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOracleInitialized is a free log retrieval operation binding the contract event 0xdfd7c5122793c91a1300dd5e12580dd7d076e1ee1a4509e0fd0900e9d808c16e.
//
// Solidity: event OracleInitialized(address indexed initializer)
func (_Contracts *ContractsFilterer) FilterOracleInitialized(opts *bind.FilterOpts, initializer []common.Address) (*ContractsOracleInitializedIterator, error) {

	var initializerRule []interface{}
	for _, initializerItem := range initializer {
		initializerRule = append(initializerRule, initializerItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "OracleInitialized", initializerRule)
	if err != nil {
		return nil, err
	}
	return &ContractsOracleInitializedIterator{contract: _Contracts.contract, event: "OracleInitialized", logs: logs, sub: sub}, nil
}

// WatchOracleInitialized is a free log subscription operation binding the contract event 0xdfd7c5122793c91a1300dd5e12580dd7d076e1ee1a4509e0fd0900e9d808c16e.
//
// Solidity: event OracleInitialized(address indexed initializer)
func (_Contracts *ContractsFilterer) WatchOracleInitialized(opts *bind.WatchOpts, sink chan<- *ContractsOracleInitialized, initializer []common.Address) (event.Subscription, error) {

	var initializerRule []interface{}
	for _, initializerItem := range initializer {
		initializerRule = append(initializerRule, initializerItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "OracleInitialized", initializerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsOracleInitialized)
				if err := _Contracts.contract.UnpackLog(event, "OracleInitialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOracleInitialized is a log parse operation binding the contract event 0xdfd7c5122793c91a1300dd5e12580dd7d076e1ee1a4509e0fd0900e9d808c16e.
//
// Solidity: event OracleInitialized(address indexed initializer)
func (_Contracts *ContractsFilterer) ParseOracleInitialized(log types.Log) (*ContractsOracleInitialized, error) {
	event := new(ContractsOracleInitialized)
	if err := _Contracts.contract.UnpackLog(event, "OracleInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Contracts contract.
type ContractsOwnershipTransferredIterator struct {
	Event *ContractsOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsOwnershipTransferred represents a OwnershipTransferred event raised by the Contracts contract.
type ContractsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contracts *ContractsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractsOwnershipTransferredIterator{contract: _Contracts.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contracts *ContractsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsOwnershipTransferred)
				if err := _Contracts.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contracts *ContractsFilterer) ParseOwnershipTransferred(log types.Log) (*ContractsOwnershipTransferred, error) {
	event := new(ContractsOwnershipTransferred)
	if err := _Contracts.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsPriceUpdatedIterator is returned from FilterPriceUpdated and is used to iterate over the raw logs and unpacked data for PriceUpdated events raised by the Contracts contract.
type ContractsPriceUpdatedIterator struct {
	Event *ContractsPriceUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsPriceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsPriceUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsPriceUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsPriceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsPriceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsPriceUpdated represents a PriceUpdated event raised by the Contracts contract.
type ContractsPriceUpdated struct {
	Base  common.Address
	Quote common.Address
	Price *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPriceUpdated is a free log retrieval operation binding the contract event 0xb71c154260e8508e211e2ace194becba2c6d7e727c3ed292fe4787458969cd10.
//
// Solidity: event PriceUpdated(address indexed base, address indexed quote, uint256 price)
func (_Contracts *ContractsFilterer) FilterPriceUpdated(opts *bind.FilterOpts, base []common.Address, quote []common.Address) (*ContractsPriceUpdatedIterator, error) {

	var baseRule []interface{}
	for _, baseItem := range base {
		baseRule = append(baseRule, baseItem)
	}
	var quoteRule []interface{}
	for _, quoteItem := range quote {
		quoteRule = append(quoteRule, quoteItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "PriceUpdated", baseRule, quoteRule)
	if err != nil {
		return nil, err
	}
	return &ContractsPriceUpdatedIterator{contract: _Contracts.contract, event: "PriceUpdated", logs: logs, sub: sub}, nil
}

// WatchPriceUpdated is a free log subscription operation binding the contract event 0xb71c154260e8508e211e2ace194becba2c6d7e727c3ed292fe4787458969cd10.
//
// Solidity: event PriceUpdated(address indexed base, address indexed quote, uint256 price)
func (_Contracts *ContractsFilterer) WatchPriceUpdated(opts *bind.WatchOpts, sink chan<- *ContractsPriceUpdated, base []common.Address, quote []common.Address) (event.Subscription, error) {

	var baseRule []interface{}
	for _, baseItem := range base {
		baseRule = append(baseRule, baseItem)
	}
	var quoteRule []interface{}
	for _, quoteItem := range quote {
		quoteRule = append(quoteRule, quoteItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "PriceUpdated", baseRule, quoteRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsPriceUpdated)
				if err := _Contracts.contract.UnpackLog(event, "PriceUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePriceUpdated is a log parse operation binding the contract event 0xb71c154260e8508e211e2ace194becba2c6d7e727c3ed292fe4787458969cd10.
//
// Solidity: event PriceUpdated(address indexed base, address indexed quote, uint256 price)
func (_Contracts *ContractsFilterer) ParsePriceUpdated(log types.Log) (*ContractsPriceUpdated, error) {
	event := new(ContractsPriceUpdated)
	if err := _Contracts.contract.UnpackLog(event, "PriceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
