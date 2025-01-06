// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// EventManagerEvent is an auto generated low-level Go binding around an user-defined struct.
type EventManagerEvent struct {
	EventId   uint64
	EventType uint16
	Data      []byte
}

// RandomnessGeneratorUnprocessedRandomness is an auto generated low-level Go binding around an user-defined struct.
type RandomnessGeneratorUnprocessedRandomness struct {
	RandomnessId uint64
}

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_writer\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"eventType\",\"type\":\"uint16\"}],\"name\":\"getEvents\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"eventId\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"eventType\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structEventManager.Event[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getRandomness\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUnprocessedRandomness\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"randomnessId\",\"type\":\"uint64\"}],\"internalType\":\"structRandomnessGenerator.UnprocessedRandomness[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"randomness\",\"type\":\"uint256\"}],\"name\":\"postRandomness\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"randomnessId\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"randomnessJobs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestRandomness\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"writer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

// GetEvents is a free data retrieval call binding the contract method 0x22ac1500.
//
// Solidity: function getEvents(uint16 eventType) view returns((uint64,uint16,bytes)[])
func (_Contract *ContractCaller) GetEvents(opts *bind.CallOpts, eventType uint16) ([]EventManagerEvent, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getEvents", eventType)

	if err != nil {
		return *new([]EventManagerEvent), err
	}

	out0 := *abi.ConvertType(out[0], new([]EventManagerEvent)).(*[]EventManagerEvent)

	return out0, err

}

// GetEvents is a free data retrieval call binding the contract method 0x22ac1500.
//
// Solidity: function getEvents(uint16 eventType) view returns((uint64,uint16,bytes)[])
func (_Contract *ContractSession) GetEvents(eventType uint16) ([]EventManagerEvent, error) {
	return _Contract.Contract.GetEvents(&_Contract.CallOpts, eventType)
}

// GetEvents is a free data retrieval call binding the contract method 0x22ac1500.
//
// Solidity: function getEvents(uint16 eventType) view returns((uint64,uint16,bytes)[])
func (_Contract *ContractCallerSession) GetEvents(eventType uint16) ([]EventManagerEvent, error) {
	return _Contract.Contract.GetEvents(&_Contract.CallOpts, eventType)
}

// GetRandomness is a free data retrieval call binding the contract method 0x453f4f62.
//
// Solidity: function getRandomness(uint256 id) view returns(uint256)
func (_Contract *ContractCaller) GetRandomness(opts *bind.CallOpts, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getRandomness", id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRandomness is a free data retrieval call binding the contract method 0x453f4f62.
//
// Solidity: function getRandomness(uint256 id) view returns(uint256)
func (_Contract *ContractSession) GetRandomness(id *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetRandomness(&_Contract.CallOpts, id)
}

// GetRandomness is a free data retrieval call binding the contract method 0x453f4f62.
//
// Solidity: function getRandomness(uint256 id) view returns(uint256)
func (_Contract *ContractCallerSession) GetRandomness(id *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetRandomness(&_Contract.CallOpts, id)
}

// GetUnprocessedRandomness is a free data retrieval call binding the contract method 0x98cee726.
//
// Solidity: function getUnprocessedRandomness() view returns((uint64)[])
func (_Contract *ContractCaller) GetUnprocessedRandomness(opts *bind.CallOpts) ([]RandomnessGeneratorUnprocessedRandomness, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getUnprocessedRandomness")

	if err != nil {
		return *new([]RandomnessGeneratorUnprocessedRandomness), err
	}

	out0 := *abi.ConvertType(out[0], new([]RandomnessGeneratorUnprocessedRandomness)).(*[]RandomnessGeneratorUnprocessedRandomness)

	return out0, err

}

// GetUnprocessedRandomness is a free data retrieval call binding the contract method 0x98cee726.
//
// Solidity: function getUnprocessedRandomness() view returns((uint64)[])
func (_Contract *ContractSession) GetUnprocessedRandomness() ([]RandomnessGeneratorUnprocessedRandomness, error) {
	return _Contract.Contract.GetUnprocessedRandomness(&_Contract.CallOpts)
}

// GetUnprocessedRandomness is a free data retrieval call binding the contract method 0x98cee726.
//
// Solidity: function getUnprocessedRandomness() view returns((uint64)[])
func (_Contract *ContractCallerSession) GetUnprocessedRandomness() ([]RandomnessGeneratorUnprocessedRandomness, error) {
	return _Contract.Contract.GetUnprocessedRandomness(&_Contract.CallOpts)
}

// RandomnessId is a free data retrieval call binding the contract method 0xa13e993f.
//
// Solidity: function randomnessId() view returns(uint64)
func (_Contract *ContractCaller) RandomnessId(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "randomnessId")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// RandomnessId is a free data retrieval call binding the contract method 0xa13e993f.
//
// Solidity: function randomnessId() view returns(uint64)
func (_Contract *ContractSession) RandomnessId() (uint64, error) {
	return _Contract.Contract.RandomnessId(&_Contract.CallOpts)
}

// RandomnessId is a free data retrieval call binding the contract method 0xa13e993f.
//
// Solidity: function randomnessId() view returns(uint64)
func (_Contract *ContractCallerSession) RandomnessId() (uint64, error) {
	return _Contract.Contract.RandomnessId(&_Contract.CallOpts)
}

// RandomnessJobs is a free data retrieval call binding the contract method 0x8a54f929.
//
// Solidity: function randomnessJobs(uint256 ) view returns(uint256)
func (_Contract *ContractCaller) RandomnessJobs(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "randomnessJobs", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RandomnessJobs is a free data retrieval call binding the contract method 0x8a54f929.
//
// Solidity: function randomnessJobs(uint256 ) view returns(uint256)
func (_Contract *ContractSession) RandomnessJobs(arg0 *big.Int) (*big.Int, error) {
	return _Contract.Contract.RandomnessJobs(&_Contract.CallOpts, arg0)
}

// RandomnessJobs is a free data retrieval call binding the contract method 0x8a54f929.
//
// Solidity: function randomnessJobs(uint256 ) view returns(uint256)
func (_Contract *ContractCallerSession) RandomnessJobs(arg0 *big.Int) (*big.Int, error) {
	return _Contract.Contract.RandomnessJobs(&_Contract.CallOpts, arg0)
}

// Writer is a free data retrieval call binding the contract method 0x453a2abc.
//
// Solidity: function writer() view returns(address)
func (_Contract *ContractCaller) Writer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "writer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Writer is a free data retrieval call binding the contract method 0x453a2abc.
//
// Solidity: function writer() view returns(address)
func (_Contract *ContractSession) Writer() (common.Address, error) {
	return _Contract.Contract.Writer(&_Contract.CallOpts)
}

// Writer is a free data retrieval call binding the contract method 0x453a2abc.
//
// Solidity: function writer() view returns(address)
func (_Contract *ContractCallerSession) Writer() (common.Address, error) {
	return _Contract.Contract.Writer(&_Contract.CallOpts)
}

// PostRandomness is a paid mutator transaction binding the contract method 0x1207e18d.
//
// Solidity: function postRandomness(uint64 id, uint256 randomness) returns()
func (_Contract *ContractTransactor) PostRandomness(opts *bind.TransactOpts, id uint64, randomness *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "postRandomness", id, randomness)
}

// PostRandomness is a paid mutator transaction binding the contract method 0x1207e18d.
//
// Solidity: function postRandomness(uint64 id, uint256 randomness) returns()
func (_Contract *ContractSession) PostRandomness(id uint64, randomness *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.PostRandomness(&_Contract.TransactOpts, id, randomness)
}

// PostRandomness is a paid mutator transaction binding the contract method 0x1207e18d.
//
// Solidity: function postRandomness(uint64 id, uint256 randomness) returns()
func (_Contract *ContractTransactorSession) PostRandomness(id uint64, randomness *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.PostRandomness(&_Contract.TransactOpts, id, randomness)
}

// RequestRandomness is a paid mutator transaction binding the contract method 0xf8413b07.
//
// Solidity: function requestRandomness() returns(uint256)
func (_Contract *ContractTransactor) RequestRandomness(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "requestRandomness")
}

// RequestRandomness is a paid mutator transaction binding the contract method 0xf8413b07.
//
// Solidity: function requestRandomness() returns(uint256)
func (_Contract *ContractSession) RequestRandomness() (*types.Transaction, error) {
	return _Contract.Contract.RequestRandomness(&_Contract.TransactOpts)
}

// RequestRandomness is a paid mutator transaction binding the contract method 0xf8413b07.
//
// Solidity: function requestRandomness() returns(uint256)
func (_Contract *ContractTransactorSession) RequestRandomness() (*types.Transaction, error) {
	return _Contract.Contract.RequestRandomness(&_Contract.TransactOpts)
}
