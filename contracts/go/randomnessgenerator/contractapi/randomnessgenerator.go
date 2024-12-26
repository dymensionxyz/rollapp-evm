// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractapi

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

// RandomnessGeneratorMetaData contains all meta data concerning the RandomnessGenerator contract.
var RandomnessGeneratorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_writer\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"randomnessId\",\"type\":\"uint256\"}],\"name\":\"EventNewRandomnessRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"randomnessId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"randomnessValue\",\"type\":\"uint256\"}],\"name\":\"EventRandomnessProvided\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getRandomness\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"randomness\",\"type\":\"uint256\"}],\"name\":\"postRandomness\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"randomnessId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"randomnessJobs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestRandomness\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"writer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// RandomnessGeneratorABI is the input ABI used to generate the binding from.
// Deprecated: Use RandomnessGeneratorMetaData.ABI instead.
var RandomnessGeneratorABI = RandomnessGeneratorMetaData.ABI

// RandomnessGenerator is an auto generated Go binding around an Ethereum contract.
type RandomnessGenerator struct {
	RandomnessGeneratorCaller     // Read-only binding to the contract
	RandomnessGeneratorTransactor // Write-only binding to the contract
	RandomnessGeneratorFilterer   // Log filterer for contract events
}

// RandomnessGeneratorCaller is an auto generated read-only Go binding around an Ethereum contract.
type RandomnessGeneratorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RandomnessGeneratorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RandomnessGeneratorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RandomnessGeneratorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RandomnessGeneratorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RandomnessGeneratorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RandomnessGeneratorSession struct {
	Contract     *RandomnessGenerator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// RandomnessGeneratorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RandomnessGeneratorCallerSession struct {
	Contract *RandomnessGeneratorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// RandomnessGeneratorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RandomnessGeneratorTransactorSession struct {
	Contract     *RandomnessGeneratorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// RandomnessGeneratorRaw is an auto generated low-level Go binding around an Ethereum contract.
type RandomnessGeneratorRaw struct {
	Contract *RandomnessGenerator // Generic contract binding to access the raw methods on
}

// RandomnessGeneratorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RandomnessGeneratorCallerRaw struct {
	Contract *RandomnessGeneratorCaller // Generic read-only contract binding to access the raw methods on
}

// RandomnessGeneratorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RandomnessGeneratorTransactorRaw struct {
	Contract *RandomnessGeneratorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRandomnessGenerator creates a new instance of RandomnessGenerator, bound to a specific deployed contract.
func NewRandomnessGenerator(address common.Address, backend bind.ContractBackend) (*RandomnessGenerator, error) {
	contract, err := bindRandomnessGenerator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RandomnessGenerator{RandomnessGeneratorCaller: RandomnessGeneratorCaller{contract: contract}, RandomnessGeneratorTransactor: RandomnessGeneratorTransactor{contract: contract}, RandomnessGeneratorFilterer: RandomnessGeneratorFilterer{contract: contract}}, nil
}

// NewRandomnessGeneratorCaller creates a new read-only instance of RandomnessGenerator, bound to a specific deployed contract.
func NewRandomnessGeneratorCaller(address common.Address, caller bind.ContractCaller) (*RandomnessGeneratorCaller, error) {
	contract, err := bindRandomnessGenerator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RandomnessGeneratorCaller{contract: contract}, nil
}

// NewRandomnessGeneratorTransactor creates a new write-only instance of RandomnessGenerator, bound to a specific deployed contract.
func NewRandomnessGeneratorTransactor(address common.Address, transactor bind.ContractTransactor) (*RandomnessGeneratorTransactor, error) {
	contract, err := bindRandomnessGenerator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RandomnessGeneratorTransactor{contract: contract}, nil
}

// NewRandomnessGeneratorFilterer creates a new log filterer instance of RandomnessGenerator, bound to a specific deployed contract.
func NewRandomnessGeneratorFilterer(address common.Address, filterer bind.ContractFilterer) (*RandomnessGeneratorFilterer, error) {
	contract, err := bindRandomnessGenerator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RandomnessGeneratorFilterer{contract: contract}, nil
}

// bindRandomnessGenerator binds a generic wrapper to an already deployed contract.
func bindRandomnessGenerator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RandomnessGeneratorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RandomnessGenerator *RandomnessGeneratorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RandomnessGenerator.Contract.RandomnessGeneratorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RandomnessGenerator *RandomnessGeneratorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RandomnessGenerator.Contract.RandomnessGeneratorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RandomnessGenerator *RandomnessGeneratorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RandomnessGenerator.Contract.RandomnessGeneratorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RandomnessGenerator *RandomnessGeneratorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RandomnessGenerator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RandomnessGenerator *RandomnessGeneratorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RandomnessGenerator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RandomnessGenerator *RandomnessGeneratorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RandomnessGenerator.Contract.contract.Transact(opts, method, params...)
}

// GetRandomness is a free data retrieval call binding the contract method 0x453f4f62.
//
// Solidity: function getRandomness(uint256 id) view returns(uint256)
func (_RandomnessGenerator *RandomnessGeneratorCaller) GetRandomness(opts *bind.CallOpts, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RandomnessGenerator.contract.Call(opts, &out, "getRandomness", id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRandomness is a free data retrieval call binding the contract method 0x453f4f62.
//
// Solidity: function getRandomness(uint256 id) view returns(uint256)
func (_RandomnessGenerator *RandomnessGeneratorSession) GetRandomness(id *big.Int) (*big.Int, error) {
	return _RandomnessGenerator.Contract.GetRandomness(&_RandomnessGenerator.CallOpts, id)
}

// GetRandomness is a free data retrieval call binding the contract method 0x453f4f62.
//
// Solidity: function getRandomness(uint256 id) view returns(uint256)
func (_RandomnessGenerator *RandomnessGeneratorCallerSession) GetRandomness(id *big.Int) (*big.Int, error) {
	return _RandomnessGenerator.Contract.GetRandomness(&_RandomnessGenerator.CallOpts, id)
}

// RandomnessId is a free data retrieval call binding the contract method 0xa13e993f.
//
// Solidity: function randomnessId() view returns(uint256)
func (_RandomnessGenerator *RandomnessGeneratorCaller) RandomnessId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RandomnessGenerator.contract.Call(opts, &out, "randomnessId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RandomnessId is a free data retrieval call binding the contract method 0xa13e993f.
//
// Solidity: function randomnessId() view returns(uint256)
func (_RandomnessGenerator *RandomnessGeneratorSession) RandomnessId() (*big.Int, error) {
	return _RandomnessGenerator.Contract.RandomnessId(&_RandomnessGenerator.CallOpts)
}

// RandomnessId is a free data retrieval call binding the contract method 0xa13e993f.
//
// Solidity: function randomnessId() view returns(uint256)
func (_RandomnessGenerator *RandomnessGeneratorCallerSession) RandomnessId() (*big.Int, error) {
	return _RandomnessGenerator.Contract.RandomnessId(&_RandomnessGenerator.CallOpts)
}

// RandomnessJobs is a free data retrieval call binding the contract method 0x8a54f929.
//
// Solidity: function randomnessJobs(uint256 ) view returns(uint256)
func (_RandomnessGenerator *RandomnessGeneratorCaller) RandomnessJobs(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RandomnessGenerator.contract.Call(opts, &out, "randomnessJobs", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RandomnessJobs is a free data retrieval call binding the contract method 0x8a54f929.
//
// Solidity: function randomnessJobs(uint256 ) view returns(uint256)
func (_RandomnessGenerator *RandomnessGeneratorSession) RandomnessJobs(arg0 *big.Int) (*big.Int, error) {
	return _RandomnessGenerator.Contract.RandomnessJobs(&_RandomnessGenerator.CallOpts, arg0)
}

// RandomnessJobs is a free data retrieval call binding the contract method 0x8a54f929.
//
// Solidity: function randomnessJobs(uint256 ) view returns(uint256)
func (_RandomnessGenerator *RandomnessGeneratorCallerSession) RandomnessJobs(arg0 *big.Int) (*big.Int, error) {
	return _RandomnessGenerator.Contract.RandomnessJobs(&_RandomnessGenerator.CallOpts, arg0)
}

// Writer is a free data retrieval call binding the contract method 0x453a2abc.
//
// Solidity: function writer() view returns(address)
func (_RandomnessGenerator *RandomnessGeneratorCaller) Writer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RandomnessGenerator.contract.Call(opts, &out, "writer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Writer is a free data retrieval call binding the contract method 0x453a2abc.
//
// Solidity: function writer() view returns(address)
func (_RandomnessGenerator *RandomnessGeneratorSession) Writer() (common.Address, error) {
	return _RandomnessGenerator.Contract.Writer(&_RandomnessGenerator.CallOpts)
}

// Writer is a free data retrieval call binding the contract method 0x453a2abc.
//
// Solidity: function writer() view returns(address)
func (_RandomnessGenerator *RandomnessGeneratorCallerSession) Writer() (common.Address, error) {
	return _RandomnessGenerator.Contract.Writer(&_RandomnessGenerator.CallOpts)
}

// PostRandomness is a paid mutator transaction binding the contract method 0xb20e730b.
//
// Solidity: function postRandomness(uint256 id, uint256 randomness) returns()
func (_RandomnessGenerator *RandomnessGeneratorTransactor) PostRandomness(opts *bind.TransactOpts, id *big.Int, randomness *big.Int) (*types.Transaction, error) {
	return _RandomnessGenerator.contract.Transact(opts, "postRandomness", id, randomness)
}

// PostRandomness is a paid mutator transaction binding the contract method 0xb20e730b.
//
// Solidity: function postRandomness(uint256 id, uint256 randomness) returns()
func (_RandomnessGenerator *RandomnessGeneratorSession) PostRandomness(id *big.Int, randomness *big.Int) (*types.Transaction, error) {
	return _RandomnessGenerator.Contract.PostRandomness(&_RandomnessGenerator.TransactOpts, id, randomness)
}

// PostRandomness is a paid mutator transaction binding the contract method 0xb20e730b.
//
// Solidity: function postRandomness(uint256 id, uint256 randomness) returns()
func (_RandomnessGenerator *RandomnessGeneratorTransactorSession) PostRandomness(id *big.Int, randomness *big.Int) (*types.Transaction, error) {
	return _RandomnessGenerator.Contract.PostRandomness(&_RandomnessGenerator.TransactOpts, id, randomness)
}

// RequestRandomness is a paid mutator transaction binding the contract method 0xf8413b07.
//
// Solidity: function requestRandomness() returns(uint256)
func (_RandomnessGenerator *RandomnessGeneratorTransactor) RequestRandomness(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RandomnessGenerator.contract.Transact(opts, "requestRandomness")
}

// RequestRandomness is a paid mutator transaction binding the contract method 0xf8413b07.
//
// Solidity: function requestRandomness() returns(uint256)
func (_RandomnessGenerator *RandomnessGeneratorSession) RequestRandomness() (*types.Transaction, error) {
	return _RandomnessGenerator.Contract.RequestRandomness(&_RandomnessGenerator.TransactOpts)
}

// RequestRandomness is a paid mutator transaction binding the contract method 0xf8413b07.
//
// Solidity: function requestRandomness() returns(uint256)
func (_RandomnessGenerator *RandomnessGeneratorTransactorSession) RequestRandomness() (*types.Transaction, error) {
	return _RandomnessGenerator.Contract.RequestRandomness(&_RandomnessGenerator.TransactOpts)
}

// RandomnessGeneratorEventNewRandomnessRequestIterator is returned from FilterEventNewRandomnessRequest and is used to iterate over the raw logs and unpacked data for EventNewRandomnessRequest events raised by the RandomnessGenerator contract.
type RandomnessGeneratorEventNewRandomnessRequestIterator struct {
	Event *RandomnessGeneratorEventNewRandomnessRequest // Event containing the contract specifics and raw log

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
func (it *RandomnessGeneratorEventNewRandomnessRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RandomnessGeneratorEventNewRandomnessRequest)
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
		it.Event = new(RandomnessGeneratorEventNewRandomnessRequest)
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
func (it *RandomnessGeneratorEventNewRandomnessRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RandomnessGeneratorEventNewRandomnessRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RandomnessGeneratorEventNewRandomnessRequest represents a EventNewRandomnessRequest event raised by the RandomnessGenerator contract.
type RandomnessGeneratorEventNewRandomnessRequest struct {
	RandomnessId *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterEventNewRandomnessRequest is a free log retrieval operation binding the contract event 0x41cd70c4c50ab29a462813018b5f492d489144743dcf020302faadf882815ffe.
//
// Solidity: event EventNewRandomnessRequest(uint256 randomnessId)
func (_RandomnessGenerator *RandomnessGeneratorFilterer) FilterEventNewRandomnessRequest(opts *bind.FilterOpts) (*RandomnessGeneratorEventNewRandomnessRequestIterator, error) {

	logs, sub, err := _RandomnessGenerator.contract.FilterLogs(opts, "EventNewRandomnessRequest")
	if err != nil {
		return nil, err
	}
	return &RandomnessGeneratorEventNewRandomnessRequestIterator{contract: _RandomnessGenerator.contract, event: "EventNewRandomnessRequest", logs: logs, sub: sub}, nil
}

// WatchEventNewRandomnessRequest is a free log subscription operation binding the contract event 0x41cd70c4c50ab29a462813018b5f492d489144743dcf020302faadf882815ffe.
//
// Solidity: event EventNewRandomnessRequest(uint256 randomnessId)
func (_RandomnessGenerator *RandomnessGeneratorFilterer) WatchEventNewRandomnessRequest(opts *bind.WatchOpts, sink chan<- *RandomnessGeneratorEventNewRandomnessRequest) (event.Subscription, error) {

	logs, sub, err := _RandomnessGenerator.contract.WatchLogs(opts, "EventNewRandomnessRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RandomnessGeneratorEventNewRandomnessRequest)
				if err := _RandomnessGenerator.contract.UnpackLog(event, "EventNewRandomnessRequest", log); err != nil {
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

// ParseEventNewRandomnessRequest is a log parse operation binding the contract event 0x41cd70c4c50ab29a462813018b5f492d489144743dcf020302faadf882815ffe.
//
// Solidity: event EventNewRandomnessRequest(uint256 randomnessId)
func (_RandomnessGenerator *RandomnessGeneratorFilterer) ParseEventNewRandomnessRequest(log types.Log) (*RandomnessGeneratorEventNewRandomnessRequest, error) {
	event := new(RandomnessGeneratorEventNewRandomnessRequest)
	if err := _RandomnessGenerator.contract.UnpackLog(event, "EventNewRandomnessRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RandomnessGeneratorEventRandomnessProvidedIterator is returned from FilterEventRandomnessProvided and is used to iterate over the raw logs and unpacked data for EventRandomnessProvided events raised by the RandomnessGenerator contract.
type RandomnessGeneratorEventRandomnessProvidedIterator struct {
	Event *RandomnessGeneratorEventRandomnessProvided // Event containing the contract specifics and raw log

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
func (it *RandomnessGeneratorEventRandomnessProvidedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RandomnessGeneratorEventRandomnessProvided)
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
		it.Event = new(RandomnessGeneratorEventRandomnessProvided)
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
func (it *RandomnessGeneratorEventRandomnessProvidedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RandomnessGeneratorEventRandomnessProvidedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RandomnessGeneratorEventRandomnessProvided represents a EventRandomnessProvided event raised by the RandomnessGenerator contract.
type RandomnessGeneratorEventRandomnessProvided struct {
	RandomnessId    *big.Int
	RandomnessValue *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterEventRandomnessProvided is a free log retrieval operation binding the contract event 0x99cd0ec730b51933c4220821b1fd2269322b52b1eb83b391e42e4541a1eaf488.
//
// Solidity: event EventRandomnessProvided(uint256 randomnessId, uint256 randomnessValue)
func (_RandomnessGenerator *RandomnessGeneratorFilterer) FilterEventRandomnessProvided(opts *bind.FilterOpts) (*RandomnessGeneratorEventRandomnessProvidedIterator, error) {

	logs, sub, err := _RandomnessGenerator.contract.FilterLogs(opts, "EventRandomnessProvided")
	if err != nil {
		return nil, err
	}
	return &RandomnessGeneratorEventRandomnessProvidedIterator{contract: _RandomnessGenerator.contract, event: "EventRandomnessProvided", logs: logs, sub: sub}, nil
}

// WatchEventRandomnessProvided is a free log subscription operation binding the contract event 0x99cd0ec730b51933c4220821b1fd2269322b52b1eb83b391e42e4541a1eaf488.
//
// Solidity: event EventRandomnessProvided(uint256 randomnessId, uint256 randomnessValue)
func (_RandomnessGenerator *RandomnessGeneratorFilterer) WatchEventRandomnessProvided(opts *bind.WatchOpts, sink chan<- *RandomnessGeneratorEventRandomnessProvided) (event.Subscription, error) {

	logs, sub, err := _RandomnessGenerator.contract.WatchLogs(opts, "EventRandomnessProvided")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RandomnessGeneratorEventRandomnessProvided)
				if err := _RandomnessGenerator.contract.UnpackLog(event, "EventRandomnessProvided", log); err != nil {
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

// ParseEventRandomnessProvided is a log parse operation binding the contract event 0x99cd0ec730b51933c4220821b1fd2269322b52b1eb83b391e42e4541a1eaf488.
//
// Solidity: event EventRandomnessProvided(uint256 randomnessId, uint256 randomnessValue)
func (_RandomnessGenerator *RandomnessGeneratorFilterer) ParseEventRandomnessProvided(log types.Log) (*RandomnessGeneratorEventRandomnessProvided, error) {
	event := new(RandomnessGeneratorEventRandomnessProvided)
	if err := _RandomnessGenerator.contract.UnpackLog(event, "EventRandomnessProvided", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
