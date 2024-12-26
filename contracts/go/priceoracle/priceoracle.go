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

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_expirationOffset\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"localNetworkName\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"oracleNetworkName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"localNetworkPrecision\",\"type\":\"uint256\"}],\"internalType\":\"structPriceOracle.AssetInfo[]\",\"name\":\"_assetInfos\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"boundThreshold\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"initializer\",\"type\":\"address\"}],\"name\":\"OracleInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"base\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"quote\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"PriceUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"SCALE_FACTOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"boundThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"expirationOffset\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"base\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quote\",\"type\":\"address\"}],\"name\":\"getPrice\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"is_inverse\",\"type\":\"bool\"}],\"internalType\":\"structPriceOracle.GetPriceResponse\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"localNetworkToOracleNetworkDenoms\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"precisionMapping\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"base\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quote\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"priceInBound\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"prices_cache\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"base\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quote\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"creationHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"creationTimeUnixMs\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revision\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"merkleProof\",\"type\":\"bytes\"}],\"internalType\":\"structPriceOracle.PriceProof\",\"name\":\"proof\",\"type\":\"tuple\"}],\"internalType\":\"structPriceOracle.PriceWithProof\",\"name\":\"priceWithProof\",\"type\":\"tuple\"}],\"name\":\"updatePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// SCALEFACTOR is a free data retrieval call binding the contract method 0xce4b5bbe.
//
// Solidity: function SCALE_FACTOR() view returns(uint256)
func (_Contract *ContractCaller) SCALEFACTOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "SCALE_FACTOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SCALEFACTOR is a free data retrieval call binding the contract method 0xce4b5bbe.
//
// Solidity: function SCALE_FACTOR() view returns(uint256)
func (_Contract *ContractSession) SCALEFACTOR() (*big.Int, error) {
	return _Contract.Contract.SCALEFACTOR(&_Contract.CallOpts)
}

// SCALEFACTOR is a free data retrieval call binding the contract method 0xce4b5bbe.
//
// Solidity: function SCALE_FACTOR() view returns(uint256)
func (_Contract *ContractCallerSession) SCALEFACTOR() (*big.Int, error) {
	return _Contract.Contract.SCALEFACTOR(&_Contract.CallOpts)
}

// BoundThreshold is a free data retrieval call binding the contract method 0x5fca4148.
//
// Solidity: function boundThreshold() view returns(uint256)
func (_Contract *ContractCaller) BoundThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "boundThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BoundThreshold is a free data retrieval call binding the contract method 0x5fca4148.
//
// Solidity: function boundThreshold() view returns(uint256)
func (_Contract *ContractSession) BoundThreshold() (*big.Int, error) {
	return _Contract.Contract.BoundThreshold(&_Contract.CallOpts)
}

// BoundThreshold is a free data retrieval call binding the contract method 0x5fca4148.
//
// Solidity: function boundThreshold() view returns(uint256)
func (_Contract *ContractCallerSession) BoundThreshold() (*big.Int, error) {
	return _Contract.Contract.BoundThreshold(&_Contract.CallOpts)
}

// ExpirationOffset is a free data retrieval call binding the contract method 0xa12ae8e4.
//
// Solidity: function expirationOffset() view returns(uint256)
func (_Contract *ContractCaller) ExpirationOffset(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "expirationOffset")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExpirationOffset is a free data retrieval call binding the contract method 0xa12ae8e4.
//
// Solidity: function expirationOffset() view returns(uint256)
func (_Contract *ContractSession) ExpirationOffset() (*big.Int, error) {
	return _Contract.Contract.ExpirationOffset(&_Contract.CallOpts)
}

// ExpirationOffset is a free data retrieval call binding the contract method 0xa12ae8e4.
//
// Solidity: function expirationOffset() view returns(uint256)
func (_Contract *ContractCallerSession) ExpirationOffset() (*big.Int, error) {
	return _Contract.Contract.ExpirationOffset(&_Contract.CallOpts)
}

// GetPrice is a free data retrieval call binding the contract method 0xac41865a.
//
// Solidity: function getPrice(address base, address quote) view returns((uint256,bool))
func (_Contract *ContractCaller) GetPrice(opts *bind.CallOpts, base common.Address, quote common.Address) (PriceOracleGetPriceResponse, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getPrice", base, quote)

	if err != nil {
		return *new(PriceOracleGetPriceResponse), err
	}

	out0 := *abi.ConvertType(out[0], new(PriceOracleGetPriceResponse)).(*PriceOracleGetPriceResponse)

	return out0, err

}

// GetPrice is a free data retrieval call binding the contract method 0xac41865a.
//
// Solidity: function getPrice(address base, address quote) view returns((uint256,bool))
func (_Contract *ContractSession) GetPrice(base common.Address, quote common.Address) (PriceOracleGetPriceResponse, error) {
	return _Contract.Contract.GetPrice(&_Contract.CallOpts, base, quote)
}

// GetPrice is a free data retrieval call binding the contract method 0xac41865a.
//
// Solidity: function getPrice(address base, address quote) view returns((uint256,bool))
func (_Contract *ContractCallerSession) GetPrice(base common.Address, quote common.Address) (PriceOracleGetPriceResponse, error) {
	return _Contract.Contract.GetPrice(&_Contract.CallOpts, base, quote)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Contract *ContractCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "initialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Contract *ContractSession) Initialized() (bool, error) {
	return _Contract.Contract.Initialized(&_Contract.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Contract *ContractCallerSession) Initialized() (bool, error) {
	return _Contract.Contract.Initialized(&_Contract.CallOpts)
}

// LocalNetworkToOracleNetworkDenoms is a free data retrieval call binding the contract method 0x5282fe92.
//
// Solidity: function localNetworkToOracleNetworkDenoms(address ) view returns(string)
func (_Contract *ContractCaller) LocalNetworkToOracleNetworkDenoms(opts *bind.CallOpts, arg0 common.Address) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "localNetworkToOracleNetworkDenoms", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// LocalNetworkToOracleNetworkDenoms is a free data retrieval call binding the contract method 0x5282fe92.
//
// Solidity: function localNetworkToOracleNetworkDenoms(address ) view returns(string)
func (_Contract *ContractSession) LocalNetworkToOracleNetworkDenoms(arg0 common.Address) (string, error) {
	return _Contract.Contract.LocalNetworkToOracleNetworkDenoms(&_Contract.CallOpts, arg0)
}

// LocalNetworkToOracleNetworkDenoms is a free data retrieval call binding the contract method 0x5282fe92.
//
// Solidity: function localNetworkToOracleNetworkDenoms(address ) view returns(string)
func (_Contract *ContractCallerSession) LocalNetworkToOracleNetworkDenoms(arg0 common.Address) (string, error) {
	return _Contract.Contract.LocalNetworkToOracleNetworkDenoms(&_Contract.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCallerSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// PrecisionMapping is a free data retrieval call binding the contract method 0x84be6db4.
//
// Solidity: function precisionMapping(address ) view returns(uint256)
func (_Contract *ContractCaller) PrecisionMapping(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "precisionMapping", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PrecisionMapping is a free data retrieval call binding the contract method 0x84be6db4.
//
// Solidity: function precisionMapping(address ) view returns(uint256)
func (_Contract *ContractSession) PrecisionMapping(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.PrecisionMapping(&_Contract.CallOpts, arg0)
}

// PrecisionMapping is a free data retrieval call binding the contract method 0x84be6db4.
//
// Solidity: function precisionMapping(address ) view returns(uint256)
func (_Contract *ContractCallerSession) PrecisionMapping(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.PrecisionMapping(&_Contract.CallOpts, arg0)
}

// PriceInBound is a free data retrieval call binding the contract method 0x27b5a917.
//
// Solidity: function priceInBound(address base, address quote, uint256 price) view returns(bool)
func (_Contract *ContractCaller) PriceInBound(opts *bind.CallOpts, base common.Address, quote common.Address, price *big.Int) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "priceInBound", base, quote, price)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PriceInBound is a free data retrieval call binding the contract method 0x27b5a917.
//
// Solidity: function priceInBound(address base, address quote, uint256 price) view returns(bool)
func (_Contract *ContractSession) PriceInBound(base common.Address, quote common.Address, price *big.Int) (bool, error) {
	return _Contract.Contract.PriceInBound(&_Contract.CallOpts, base, quote, price)
}

// PriceInBound is a free data retrieval call binding the contract method 0x27b5a917.
//
// Solidity: function priceInBound(address base, address quote, uint256 price) view returns(bool)
func (_Contract *ContractCallerSession) PriceInBound(base common.Address, quote common.Address, price *big.Int) (bool, error) {
	return _Contract.Contract.PriceInBound(&_Contract.CallOpts, base, quote, price)
}

// PricesCache is a free data retrieval call binding the contract method 0x5e0ce1f6.
//
// Solidity: function prices_cache(address , address ) view returns(uint256 price, uint256 expiration, bool exists)
func (_Contract *ContractCaller) PricesCache(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (struct {
	Price      *big.Int
	Expiration *big.Int
	Exists     bool
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "prices_cache", arg0, arg1)

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
func (_Contract *ContractSession) PricesCache(arg0 common.Address, arg1 common.Address) (struct {
	Price      *big.Int
	Expiration *big.Int
	Exists     bool
}, error) {
	return _Contract.Contract.PricesCache(&_Contract.CallOpts, arg0, arg1)
}

// PricesCache is a free data retrieval call binding the contract method 0x5e0ce1f6.
//
// Solidity: function prices_cache(address , address ) view returns(uint256 price, uint256 expiration, bool exists)
func (_Contract *ContractCallerSession) PricesCache(arg0 common.Address, arg1 common.Address) (struct {
	Price      *big.Int
	Expiration *big.Int
	Exists     bool
}, error) {
	return _Contract.Contract.PricesCache(&_Contract.CallOpts, arg0, arg1)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Contract *ContractTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Contract *ContractSession) Initialize() (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Contract *ContractTransactorSession) Initialize() (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0x42630dee.
//
// Solidity: function updatePrice(address base, address quote, (uint256,(uint256,uint256,uint256,uint256,bytes)) priceWithProof) returns()
func (_Contract *ContractTransactor) UpdatePrice(opts *bind.TransactOpts, base common.Address, quote common.Address, priceWithProof PriceOraclePriceWithProof) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updatePrice", base, quote, priceWithProof)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0x42630dee.
//
// Solidity: function updatePrice(address base, address quote, (uint256,(uint256,uint256,uint256,uint256,bytes)) priceWithProof) returns()
func (_Contract *ContractSession) UpdatePrice(base common.Address, quote common.Address, priceWithProof PriceOraclePriceWithProof) (*types.Transaction, error) {
	return _Contract.Contract.UpdatePrice(&_Contract.TransactOpts, base, quote, priceWithProof)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0x42630dee.
//
// Solidity: function updatePrice(address base, address quote, (uint256,(uint256,uint256,uint256,uint256,bytes)) priceWithProof) returns()
func (_Contract *ContractTransactorSession) UpdatePrice(base common.Address, quote common.Address, priceWithProof PriceOraclePriceWithProof) (*types.Transaction, error) {
	return _Contract.Contract.UpdatePrice(&_Contract.TransactOpts, base, quote, priceWithProof)
}

// ContractOracleInitializedIterator is returned from FilterOracleInitialized and is used to iterate over the raw logs and unpacked data for OracleInitialized events raised by the Contract contract.
type ContractOracleInitializedIterator struct {
	Event *ContractOracleInitialized // Event containing the contract specifics and raw log

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
func (it *ContractOracleInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOracleInitialized)
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
		it.Event = new(ContractOracleInitialized)
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
func (it *ContractOracleInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOracleInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOracleInitialized represents a OracleInitialized event raised by the Contract contract.
type ContractOracleInitialized struct {
	Initializer common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOracleInitialized is a free log retrieval operation binding the contract event 0xdfd7c5122793c91a1300dd5e12580dd7d076e1ee1a4509e0fd0900e9d808c16e.
//
// Solidity: event OracleInitialized(address indexed initializer)
func (_Contract *ContractFilterer) FilterOracleInitialized(opts *bind.FilterOpts, initializer []common.Address) (*ContractOracleInitializedIterator, error) {

	var initializerRule []interface{}
	for _, initializerItem := range initializer {
		initializerRule = append(initializerRule, initializerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "OracleInitialized", initializerRule)
	if err != nil {
		return nil, err
	}
	return &ContractOracleInitializedIterator{contract: _Contract.contract, event: "OracleInitialized", logs: logs, sub: sub}, nil
}

// WatchOracleInitialized is a free log subscription operation binding the contract event 0xdfd7c5122793c91a1300dd5e12580dd7d076e1ee1a4509e0fd0900e9d808c16e.
//
// Solidity: event OracleInitialized(address indexed initializer)
func (_Contract *ContractFilterer) WatchOracleInitialized(opts *bind.WatchOpts, sink chan<- *ContractOracleInitialized, initializer []common.Address) (event.Subscription, error) {

	var initializerRule []interface{}
	for _, initializerItem := range initializer {
		initializerRule = append(initializerRule, initializerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "OracleInitialized", initializerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOracleInitialized)
				if err := _Contract.contract.UnpackLog(event, "OracleInitialized", log); err != nil {
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
func (_Contract *ContractFilterer) ParseOracleInitialized(log types.Log) (*ContractOracleInitialized, error) {
	event := new(ContractOracleInitialized)
	if err := _Contract.contract.UnpackLog(event, "OracleInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Contract contract.
type ContractOwnershipTransferredIterator struct {
	Event *ContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOwnershipTransferred)
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
		it.Event = new(ContractOwnershipTransferred)
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
func (it *ContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOwnershipTransferred represents a OwnershipTransferred event raised by the Contract contract.
type ContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractOwnershipTransferredIterator{contract: _Contract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOwnershipTransferred)
				if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Contract *ContractFilterer) ParseOwnershipTransferred(log types.Log) (*ContractOwnershipTransferred, error) {
	event := new(ContractOwnershipTransferred)
	if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractPriceUpdatedIterator is returned from FilterPriceUpdated and is used to iterate over the raw logs and unpacked data for PriceUpdated events raised by the Contract contract.
type ContractPriceUpdatedIterator struct {
	Event *ContractPriceUpdated // Event containing the contract specifics and raw log

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
func (it *ContractPriceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractPriceUpdated)
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
		it.Event = new(ContractPriceUpdated)
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
func (it *ContractPriceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractPriceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractPriceUpdated represents a PriceUpdated event raised by the Contract contract.
type ContractPriceUpdated struct {
	Base  common.Address
	Quote common.Address
	Price *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPriceUpdated is a free log retrieval operation binding the contract event 0xb71c154260e8508e211e2ace194becba2c6d7e727c3ed292fe4787458969cd10.
//
// Solidity: event PriceUpdated(address indexed base, address indexed quote, uint256 price)
func (_Contract *ContractFilterer) FilterPriceUpdated(opts *bind.FilterOpts, base []common.Address, quote []common.Address) (*ContractPriceUpdatedIterator, error) {

	var baseRule []interface{}
	for _, baseItem := range base {
		baseRule = append(baseRule, baseItem)
	}
	var quoteRule []interface{}
	for _, quoteItem := range quote {
		quoteRule = append(quoteRule, quoteItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "PriceUpdated", baseRule, quoteRule)
	if err != nil {
		return nil, err
	}
	return &ContractPriceUpdatedIterator{contract: _Contract.contract, event: "PriceUpdated", logs: logs, sub: sub}, nil
}

// WatchPriceUpdated is a free log subscription operation binding the contract event 0xb71c154260e8508e211e2ace194becba2c6d7e727c3ed292fe4787458969cd10.
//
// Solidity: event PriceUpdated(address indexed base, address indexed quote, uint256 price)
func (_Contract *ContractFilterer) WatchPriceUpdated(opts *bind.WatchOpts, sink chan<- *ContractPriceUpdated, base []common.Address, quote []common.Address) (event.Subscription, error) {

	var baseRule []interface{}
	for _, baseItem := range base {
		baseRule = append(baseRule, baseItem)
	}
	var quoteRule []interface{}
	for _, quoteItem := range quote {
		quoteRule = append(quoteRule, quoteItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "PriceUpdated", baseRule, quoteRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractPriceUpdated)
				if err := _Contract.contract.UnpackLog(event, "PriceUpdated", log); err != nil {
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
func (_Contract *ContractFilterer) ParsePriceUpdated(log types.Log) (*ContractPriceUpdated, error) {
	event := new(ContractPriceUpdated)
	if err := _Contract.contract.UnpackLog(event, "PriceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
