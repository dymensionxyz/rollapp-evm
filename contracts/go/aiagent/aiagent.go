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

// AIOracleMetaData contains all meta data concerning the AIOracle contract.
var AIOracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_authorizedWriter\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddWhitelisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"promptId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"answer\",\"type\":\"string\"}],\"name\":\"AnswerSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldAIAgent\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAIAgent\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"promptId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"prompt\",\"type\":\"string\"}],\"name\":\"PromptSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"RemoveWhitelisted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addWhitelistAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"aiAgent\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"answers\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getAnswer\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isWhitelistedPrompter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestPromptId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removeWhitelistAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"answer\",\"type\":\"string\"}],\"name\":\"submitAnswer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"prompt\",\"type\":\"string\"}],\"name\":\"submitPrompt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAIAgent\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AIOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use AIOracleMetaData.ABI instead.
var AIOracleABI = AIOracleMetaData.ABI

// AIOracle is an auto generated Go binding around an Ethereum contract.
type AIOracle struct {
	AIOracleCaller     // Read-only binding to the contract
	AIOracleTransactor // Write-only binding to the contract
	AIOracleFilterer   // Log filterer for contract events
}

// AIOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type AIOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AIOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AIOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AIOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AIOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AIOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AIOracleSession struct {
	Contract     *AIOracle         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AIOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AIOracleCallerSession struct {
	Contract *AIOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// AIOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AIOracleTransactorSession struct {
	Contract     *AIOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// AIOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type AIOracleRaw struct {
	Contract *AIOracle // Generic contract binding to access the raw methods on
}

// AIOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AIOracleCallerRaw struct {
	Contract *AIOracleCaller // Generic read-only contract binding to access the raw methods on
}

// AIOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AIOracleTransactorRaw struct {
	Contract *AIOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAIOracle creates a new instance of AIOracle, bound to a specific deployed contract.
func NewAIOracle(address common.Address, backend bind.ContractBackend) (*AIOracle, error) {
	contract, err := bindAIOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AIOracle{AIOracleCaller: AIOracleCaller{contract: contract}, AIOracleTransactor: AIOracleTransactor{contract: contract}, AIOracleFilterer: AIOracleFilterer{contract: contract}}, nil
}

// NewAIOracleCaller creates a new read-only instance of AIOracle, bound to a specific deployed contract.
func NewAIOracleCaller(address common.Address, caller bind.ContractCaller) (*AIOracleCaller, error) {
	contract, err := bindAIOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AIOracleCaller{contract: contract}, nil
}

// NewAIOracleTransactor creates a new write-only instance of AIOracle, bound to a specific deployed contract.
func NewAIOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*AIOracleTransactor, error) {
	contract, err := bindAIOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AIOracleTransactor{contract: contract}, nil
}

// NewAIOracleFilterer creates a new log filterer instance of AIOracle, bound to a specific deployed contract.
func NewAIOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*AIOracleFilterer, error) {
	contract, err := bindAIOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AIOracleFilterer{contract: contract}, nil
}

// bindAIOracle binds a generic wrapper to an already deployed contract.
func bindAIOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AIOracleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AIOracle *AIOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AIOracle.Contract.AIOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AIOracle *AIOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AIOracle.Contract.AIOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AIOracle *AIOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AIOracle.Contract.AIOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AIOracle *AIOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AIOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AIOracle *AIOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AIOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AIOracle *AIOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AIOracle.Contract.contract.Transact(opts, method, params...)
}

// AiAgent is a free data retrieval call binding the contract method 0xa6d5b732.
//
// Solidity: function aiAgent() view returns(address)
func (_AIOracle *AIOracleCaller) AiAgent(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AIOracle.contract.Call(opts, &out, "aiAgent")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AiAgent is a free data retrieval call binding the contract method 0xa6d5b732.
//
// Solidity: function aiAgent() view returns(address)
func (_AIOracle *AIOracleSession) AiAgent() (common.Address, error) {
	return _AIOracle.Contract.AiAgent(&_AIOracle.CallOpts)
}

// AiAgent is a free data retrieval call binding the contract method 0xa6d5b732.
//
// Solidity: function aiAgent() view returns(address)
func (_AIOracle *AIOracleCallerSession) AiAgent() (common.Address, error) {
	return _AIOracle.Contract.AiAgent(&_AIOracle.CallOpts)
}

// Answers is a free data retrieval call binding the contract method 0x17599cc5.
//
// Solidity: function answers(uint256 ) view returns(string)
func (_AIOracle *AIOracleCaller) Answers(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _AIOracle.contract.Call(opts, &out, "answers", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Answers is a free data retrieval call binding the contract method 0x17599cc5.
//
// Solidity: function answers(uint256 ) view returns(string)
func (_AIOracle *AIOracleSession) Answers(arg0 *big.Int) (string, error) {
	return _AIOracle.Contract.Answers(&_AIOracle.CallOpts, arg0)
}

// Answers is a free data retrieval call binding the contract method 0x17599cc5.
//
// Solidity: function answers(uint256 ) view returns(string)
func (_AIOracle *AIOracleCallerSession) Answers(arg0 *big.Int) (string, error) {
	return _AIOracle.Contract.Answers(&_AIOracle.CallOpts, arg0)
}

// GetAnswer is a free data retrieval call binding the contract method 0xb5ab58dc.
//
// Solidity: function getAnswer(uint256 id) view returns(string)
func (_AIOracle *AIOracleCaller) GetAnswer(opts *bind.CallOpts, id *big.Int) (string, error) {
	var out []interface{}
	err := _AIOracle.contract.Call(opts, &out, "getAnswer", id)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetAnswer is a free data retrieval call binding the contract method 0xb5ab58dc.
//
// Solidity: function getAnswer(uint256 id) view returns(string)
func (_AIOracle *AIOracleSession) GetAnswer(id *big.Int) (string, error) {
	return _AIOracle.Contract.GetAnswer(&_AIOracle.CallOpts, id)
}

// GetAnswer is a free data retrieval call binding the contract method 0xb5ab58dc.
//
// Solidity: function getAnswer(uint256 id) view returns(string)
func (_AIOracle *AIOracleCallerSession) GetAnswer(id *big.Int) (string, error) {
	return _AIOracle.Contract.GetAnswer(&_AIOracle.CallOpts, id)
}

// IsWhitelistedPrompter is a free data retrieval call binding the contract method 0x01ce61be.
//
// Solidity: function isWhitelistedPrompter(address account) view returns(bool)
func (_AIOracle *AIOracleCaller) IsWhitelistedPrompter(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _AIOracle.contract.Call(opts, &out, "isWhitelistedPrompter", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWhitelistedPrompter is a free data retrieval call binding the contract method 0x01ce61be.
//
// Solidity: function isWhitelistedPrompter(address account) view returns(bool)
func (_AIOracle *AIOracleSession) IsWhitelistedPrompter(account common.Address) (bool, error) {
	return _AIOracle.Contract.IsWhitelistedPrompter(&_AIOracle.CallOpts, account)
}

// IsWhitelistedPrompter is a free data retrieval call binding the contract method 0x01ce61be.
//
// Solidity: function isWhitelistedPrompter(address account) view returns(bool)
func (_AIOracle *AIOracleCallerSession) IsWhitelistedPrompter(account common.Address) (bool, error) {
	return _AIOracle.Contract.IsWhitelistedPrompter(&_AIOracle.CallOpts, account)
}

// LatestPromptId is a free data retrieval call binding the contract method 0x6f2a2816.
//
// Solidity: function latestPromptId() view returns(uint256)
func (_AIOracle *AIOracleCaller) LatestPromptId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AIOracle.contract.Call(opts, &out, "latestPromptId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestPromptId is a free data retrieval call binding the contract method 0x6f2a2816.
//
// Solidity: function latestPromptId() view returns(uint256)
func (_AIOracle *AIOracleSession) LatestPromptId() (*big.Int, error) {
	return _AIOracle.Contract.LatestPromptId(&_AIOracle.CallOpts)
}

// LatestPromptId is a free data retrieval call binding the contract method 0x6f2a2816.
//
// Solidity: function latestPromptId() view returns(uint256)
func (_AIOracle *AIOracleCallerSession) LatestPromptId() (*big.Int, error) {
	return _AIOracle.Contract.LatestPromptId(&_AIOracle.CallOpts)
}

// AddWhitelistAddress is a paid mutator transaction binding the contract method 0x94a7ef15.
//
// Solidity: function addWhitelistAddress(address account) returns()
func (_AIOracle *AIOracleTransactor) AddWhitelistAddress(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _AIOracle.contract.Transact(opts, "addWhitelistAddress", account)
}

// AddWhitelistAddress is a paid mutator transaction binding the contract method 0x94a7ef15.
//
// Solidity: function addWhitelistAddress(address account) returns()
func (_AIOracle *AIOracleSession) AddWhitelistAddress(account common.Address) (*types.Transaction, error) {
	return _AIOracle.Contract.AddWhitelistAddress(&_AIOracle.TransactOpts, account)
}

// AddWhitelistAddress is a paid mutator transaction binding the contract method 0x94a7ef15.
//
// Solidity: function addWhitelistAddress(address account) returns()
func (_AIOracle *AIOracleTransactorSession) AddWhitelistAddress(account common.Address) (*types.Transaction, error) {
	return _AIOracle.Contract.AddWhitelistAddress(&_AIOracle.TransactOpts, account)
}

// RemoveWhitelistAddress is a paid mutator transaction binding the contract method 0xb7ecbaae.
//
// Solidity: function removeWhitelistAddress(address account) returns()
func (_AIOracle *AIOracleTransactor) RemoveWhitelistAddress(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _AIOracle.contract.Transact(opts, "removeWhitelistAddress", account)
}

// RemoveWhitelistAddress is a paid mutator transaction binding the contract method 0xb7ecbaae.
//
// Solidity: function removeWhitelistAddress(address account) returns()
func (_AIOracle *AIOracleSession) RemoveWhitelistAddress(account common.Address) (*types.Transaction, error) {
	return _AIOracle.Contract.RemoveWhitelistAddress(&_AIOracle.TransactOpts, account)
}

// RemoveWhitelistAddress is a paid mutator transaction binding the contract method 0xb7ecbaae.
//
// Solidity: function removeWhitelistAddress(address account) returns()
func (_AIOracle *AIOracleTransactorSession) RemoveWhitelistAddress(account common.Address) (*types.Transaction, error) {
	return _AIOracle.Contract.RemoveWhitelistAddress(&_AIOracle.TransactOpts, account)
}

// SubmitAnswer is a paid mutator transaction binding the contract method 0x01fa3bec.
//
// Solidity: function submitAnswer(uint256 id, string answer) returns()
func (_AIOracle *AIOracleTransactor) SubmitAnswer(opts *bind.TransactOpts, id *big.Int, answer string) (*types.Transaction, error) {
	return _AIOracle.contract.Transact(opts, "submitAnswer", id, answer)
}

// SubmitAnswer is a paid mutator transaction binding the contract method 0x01fa3bec.
//
// Solidity: function submitAnswer(uint256 id, string answer) returns()
func (_AIOracle *AIOracleSession) SubmitAnswer(id *big.Int, answer string) (*types.Transaction, error) {
	return _AIOracle.Contract.SubmitAnswer(&_AIOracle.TransactOpts, id, answer)
}

// SubmitAnswer is a paid mutator transaction binding the contract method 0x01fa3bec.
//
// Solidity: function submitAnswer(uint256 id, string answer) returns()
func (_AIOracle *AIOracleTransactorSession) SubmitAnswer(id *big.Int, answer string) (*types.Transaction, error) {
	return _AIOracle.Contract.SubmitAnswer(&_AIOracle.TransactOpts, id, answer)
}

// SubmitPrompt is a paid mutator transaction binding the contract method 0x28b43144.
//
// Solidity: function submitPrompt(string prompt) returns(uint256)
func (_AIOracle *AIOracleTransactor) SubmitPrompt(opts *bind.TransactOpts, prompt string) (*types.Transaction, error) {
	return _AIOracle.contract.Transact(opts, "submitPrompt", prompt)
}

// SubmitPrompt is a paid mutator transaction binding the contract method 0x28b43144.
//
// Solidity: function submitPrompt(string prompt) returns(uint256)
func (_AIOracle *AIOracleSession) SubmitPrompt(prompt string) (*types.Transaction, error) {
	return _AIOracle.Contract.SubmitPrompt(&_AIOracle.TransactOpts, prompt)
}

// SubmitPrompt is a paid mutator transaction binding the contract method 0x28b43144.
//
// Solidity: function submitPrompt(string prompt) returns(uint256)
func (_AIOracle *AIOracleTransactorSession) SubmitPrompt(prompt string) (*types.Transaction, error) {
	return _AIOracle.Contract.SubmitPrompt(&_AIOracle.TransactOpts, prompt)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newAIAgent) returns()
func (_AIOracle *AIOracleTransactor) TransferOwnership(opts *bind.TransactOpts, newAIAgent common.Address) (*types.Transaction, error) {
	return _AIOracle.contract.Transact(opts, "transferOwnership", newAIAgent)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newAIAgent) returns()
func (_AIOracle *AIOracleSession) TransferOwnership(newAIAgent common.Address) (*types.Transaction, error) {
	return _AIOracle.Contract.TransferOwnership(&_AIOracle.TransactOpts, newAIAgent)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newAIAgent) returns()
func (_AIOracle *AIOracleTransactorSession) TransferOwnership(newAIAgent common.Address) (*types.Transaction, error) {
	return _AIOracle.Contract.TransferOwnership(&_AIOracle.TransactOpts, newAIAgent)
}

// AIOracleAddWhitelistedIterator is returned from FilterAddWhitelisted and is used to iterate over the raw logs and unpacked data for AddWhitelisted events raised by the AIOracle contract.
type AIOracleAddWhitelistedIterator struct {
	Event *AIOracleAddWhitelisted // Event containing the contract specifics and raw log

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
func (it *AIOracleAddWhitelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AIOracleAddWhitelisted)
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
		it.Event = new(AIOracleAddWhitelisted)
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
func (it *AIOracleAddWhitelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AIOracleAddWhitelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AIOracleAddWhitelisted represents a AddWhitelisted event raised by the AIOracle contract.
type AIOracleAddWhitelisted struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAddWhitelisted is a free log retrieval operation binding the contract event 0xf3e0a9bccfdae73de3642e074bd7547b27b8788b6b7db7e51b25d86ea5ca8767.
//
// Solidity: event AddWhitelisted(address indexed account)
func (_AIOracle *AIOracleFilterer) FilterAddWhitelisted(opts *bind.FilterOpts, account []common.Address) (*AIOracleAddWhitelistedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AIOracle.contract.FilterLogs(opts, "AddWhitelisted", accountRule)
	if err != nil {
		return nil, err
	}
	return &AIOracleAddWhitelistedIterator{contract: _AIOracle.contract, event: "AddWhitelisted", logs: logs, sub: sub}, nil
}

// WatchAddWhitelisted is a free log subscription operation binding the contract event 0xf3e0a9bccfdae73de3642e074bd7547b27b8788b6b7db7e51b25d86ea5ca8767.
//
// Solidity: event AddWhitelisted(address indexed account)
func (_AIOracle *AIOracleFilterer) WatchAddWhitelisted(opts *bind.WatchOpts, sink chan<- *AIOracleAddWhitelisted, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AIOracle.contract.WatchLogs(opts, "AddWhitelisted", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AIOracleAddWhitelisted)
				if err := _AIOracle.contract.UnpackLog(event, "AddWhitelisted", log); err != nil {
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

// ParseAddWhitelisted is a log parse operation binding the contract event 0xf3e0a9bccfdae73de3642e074bd7547b27b8788b6b7db7e51b25d86ea5ca8767.
//
// Solidity: event AddWhitelisted(address indexed account)
func (_AIOracle *AIOracleFilterer) ParseAddWhitelisted(log types.Log) (*AIOracleAddWhitelisted, error) {
	event := new(AIOracleAddWhitelisted)
	if err := _AIOracle.contract.UnpackLog(event, "AddWhitelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AIOracleAnswerSubmittedIterator is returned from FilterAnswerSubmitted and is used to iterate over the raw logs and unpacked data for AnswerSubmitted events raised by the AIOracle contract.
type AIOracleAnswerSubmittedIterator struct {
	Event *AIOracleAnswerSubmitted // Event containing the contract specifics and raw log

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
func (it *AIOracleAnswerSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AIOracleAnswerSubmitted)
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
		it.Event = new(AIOracleAnswerSubmitted)
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
func (it *AIOracleAnswerSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AIOracleAnswerSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AIOracleAnswerSubmitted represents a AnswerSubmitted event raised by the AIOracle contract.
type AIOracleAnswerSubmitted struct {
	PromptId *big.Int
	Answer   string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAnswerSubmitted is a free log retrieval operation binding the contract event 0x04cf4660632bddc9871c3099c79fc6e85c889dbee3192ea0cfa1fde0419dfd7d.
//
// Solidity: event AnswerSubmitted(uint256 promptId, string answer)
func (_AIOracle *AIOracleFilterer) FilterAnswerSubmitted(opts *bind.FilterOpts) (*AIOracleAnswerSubmittedIterator, error) {

	logs, sub, err := _AIOracle.contract.FilterLogs(opts, "AnswerSubmitted")
	if err != nil {
		return nil, err
	}
	return &AIOracleAnswerSubmittedIterator{contract: _AIOracle.contract, event: "AnswerSubmitted", logs: logs, sub: sub}, nil
}

// WatchAnswerSubmitted is a free log subscription operation binding the contract event 0x04cf4660632bddc9871c3099c79fc6e85c889dbee3192ea0cfa1fde0419dfd7d.
//
// Solidity: event AnswerSubmitted(uint256 promptId, string answer)
func (_AIOracle *AIOracleFilterer) WatchAnswerSubmitted(opts *bind.WatchOpts, sink chan<- *AIOracleAnswerSubmitted) (event.Subscription, error) {

	logs, sub, err := _AIOracle.contract.WatchLogs(opts, "AnswerSubmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AIOracleAnswerSubmitted)
				if err := _AIOracle.contract.UnpackLog(event, "AnswerSubmitted", log); err != nil {
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

// ParseAnswerSubmitted is a log parse operation binding the contract event 0x04cf4660632bddc9871c3099c79fc6e85c889dbee3192ea0cfa1fde0419dfd7d.
//
// Solidity: event AnswerSubmitted(uint256 promptId, string answer)
func (_AIOracle *AIOracleFilterer) ParseAnswerSubmitted(log types.Log) (*AIOracleAnswerSubmitted, error) {
	event := new(AIOracleAnswerSubmitted)
	if err := _AIOracle.contract.UnpackLog(event, "AnswerSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AIOracleOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AIOracle contract.
type AIOracleOwnershipTransferredIterator struct {
	Event *AIOracleOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AIOracleOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AIOracleOwnershipTransferred)
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
		it.Event = new(AIOracleOwnershipTransferred)
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
func (it *AIOracleOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AIOracleOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AIOracleOwnershipTransferred represents a OwnershipTransferred event raised by the AIOracle contract.
type AIOracleOwnershipTransferred struct {
	OldAIAgent common.Address
	NewAIAgent common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldAIAgent, address indexed newAIAgent)
func (_AIOracle *AIOracleFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, oldAIAgent []common.Address, newAIAgent []common.Address) (*AIOracleOwnershipTransferredIterator, error) {

	var oldAIAgentRule []interface{}
	for _, oldAIAgentItem := range oldAIAgent {
		oldAIAgentRule = append(oldAIAgentRule, oldAIAgentItem)
	}
	var newAIAgentRule []interface{}
	for _, newAIAgentItem := range newAIAgent {
		newAIAgentRule = append(newAIAgentRule, newAIAgentItem)
	}

	logs, sub, err := _AIOracle.contract.FilterLogs(opts, "OwnershipTransferred", oldAIAgentRule, newAIAgentRule)
	if err != nil {
		return nil, err
	}
	return &AIOracleOwnershipTransferredIterator{contract: _AIOracle.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldAIAgent, address indexed newAIAgent)
func (_AIOracle *AIOracleFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AIOracleOwnershipTransferred, oldAIAgent []common.Address, newAIAgent []common.Address) (event.Subscription, error) {

	var oldAIAgentRule []interface{}
	for _, oldAIAgentItem := range oldAIAgent {
		oldAIAgentRule = append(oldAIAgentRule, oldAIAgentItem)
	}
	var newAIAgentRule []interface{}
	for _, newAIAgentItem := range newAIAgent {
		newAIAgentRule = append(newAIAgentRule, newAIAgentItem)
	}

	logs, sub, err := _AIOracle.contract.WatchLogs(opts, "OwnershipTransferred", oldAIAgentRule, newAIAgentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AIOracleOwnershipTransferred)
				if err := _AIOracle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
// Solidity: event OwnershipTransferred(address indexed oldAIAgent, address indexed newAIAgent)
func (_AIOracle *AIOracleFilterer) ParseOwnershipTransferred(log types.Log) (*AIOracleOwnershipTransferred, error) {
	event := new(AIOracleOwnershipTransferred)
	if err := _AIOracle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AIOraclePromptSubmittedIterator is returned from FilterPromptSubmitted and is used to iterate over the raw logs and unpacked data for PromptSubmitted events raised by the AIOracle contract.
type AIOraclePromptSubmittedIterator struct {
	Event *AIOraclePromptSubmitted // Event containing the contract specifics and raw log

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
func (it *AIOraclePromptSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AIOraclePromptSubmitted)
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
		it.Event = new(AIOraclePromptSubmitted)
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
func (it *AIOraclePromptSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AIOraclePromptSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AIOraclePromptSubmitted represents a PromptSubmitted event raised by the AIOracle contract.
type AIOraclePromptSubmitted struct {
	PromptId *big.Int
	Prompt   string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPromptSubmitted is a free log retrieval operation binding the contract event 0xf53cb06c91632882a222e576c4100718a072c6648442e3cb944b44358b5ab839.
//
// Solidity: event PromptSubmitted(uint256 promptId, string prompt)
func (_AIOracle *AIOracleFilterer) FilterPromptSubmitted(opts *bind.FilterOpts) (*AIOraclePromptSubmittedIterator, error) {

	logs, sub, err := _AIOracle.contract.FilterLogs(opts, "PromptSubmitted")
	if err != nil {
		return nil, err
	}
	return &AIOraclePromptSubmittedIterator{contract: _AIOracle.contract, event: "PromptSubmitted", logs: logs, sub: sub}, nil
}

// WatchPromptSubmitted is a free log subscription operation binding the contract event 0xf53cb06c91632882a222e576c4100718a072c6648442e3cb944b44358b5ab839.
//
// Solidity: event PromptSubmitted(uint256 promptId, string prompt)
func (_AIOracle *AIOracleFilterer) WatchPromptSubmitted(opts *bind.WatchOpts, sink chan<- *AIOraclePromptSubmitted) (event.Subscription, error) {

	logs, sub, err := _AIOracle.contract.WatchLogs(opts, "PromptSubmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AIOraclePromptSubmitted)
				if err := _AIOracle.contract.UnpackLog(event, "PromptSubmitted", log); err != nil {
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

// ParsePromptSubmitted is a log parse operation binding the contract event 0xf53cb06c91632882a222e576c4100718a072c6648442e3cb944b44358b5ab839.
//
// Solidity: event PromptSubmitted(uint256 promptId, string prompt)
func (_AIOracle *AIOracleFilterer) ParsePromptSubmitted(log types.Log) (*AIOraclePromptSubmitted, error) {
	event := new(AIOraclePromptSubmitted)
	if err := _AIOracle.contract.UnpackLog(event, "PromptSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AIOracleRemoveWhitelistedIterator is returned from FilterRemoveWhitelisted and is used to iterate over the raw logs and unpacked data for RemoveWhitelisted events raised by the AIOracle contract.
type AIOracleRemoveWhitelistedIterator struct {
	Event *AIOracleRemoveWhitelisted // Event containing the contract specifics and raw log

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
func (it *AIOracleRemoveWhitelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AIOracleRemoveWhitelisted)
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
		it.Event = new(AIOracleRemoveWhitelisted)
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
func (it *AIOracleRemoveWhitelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AIOracleRemoveWhitelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AIOracleRemoveWhitelisted represents a RemoveWhitelisted event raised by the AIOracle contract.
type AIOracleRemoveWhitelisted struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRemoveWhitelisted is a free log retrieval operation binding the contract event 0x12891dbc60d241c27b09600bf192c7e0ce5128d76790bb872a2a4649de301583.
//
// Solidity: event RemoveWhitelisted(address indexed account)
func (_AIOracle *AIOracleFilterer) FilterRemoveWhitelisted(opts *bind.FilterOpts, account []common.Address) (*AIOracleRemoveWhitelistedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AIOracle.contract.FilterLogs(opts, "RemoveWhitelisted", accountRule)
	if err != nil {
		return nil, err
	}
	return &AIOracleRemoveWhitelistedIterator{contract: _AIOracle.contract, event: "RemoveWhitelisted", logs: logs, sub: sub}, nil
}

// WatchRemoveWhitelisted is a free log subscription operation binding the contract event 0x12891dbc60d241c27b09600bf192c7e0ce5128d76790bb872a2a4649de301583.
//
// Solidity: event RemoveWhitelisted(address indexed account)
func (_AIOracle *AIOracleFilterer) WatchRemoveWhitelisted(opts *bind.WatchOpts, sink chan<- *AIOracleRemoveWhitelisted, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AIOracle.contract.WatchLogs(opts, "RemoveWhitelisted", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AIOracleRemoveWhitelisted)
				if err := _AIOracle.contract.UnpackLog(event, "RemoveWhitelisted", log); err != nil {
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

// ParseRemoveWhitelisted is a log parse operation binding the contract event 0x12891dbc60d241c27b09600bf192c7e0ce5128d76790bb872a2a4649de301583.
//
// Solidity: event RemoveWhitelisted(address indexed account)
func (_AIOracle *AIOracleFilterer) ParseRemoveWhitelisted(log types.Log) (*AIOracleRemoveWhitelisted, error) {
	event := new(AIOracleRemoveWhitelisted)
	if err := _AIOracle.contract.UnpackLog(event, "RemoveWhitelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
