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

// EventManagerEvent is an auto generated low-level Go binding around an user-defined struct.
type EventManagerEvent struct {
	EventId   *big.Int
	EventType uint16
	Data      []byte
}

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_writer\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"eventIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint16\",\"name\":\"eventType\",\"type\":\"uint16\"}],\"name\":\"eraseEvents\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getRandomness\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"eventType\",\"type\":\"uint16\"}],\"name\":\"pollEvents\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"eventId\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"eventType\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structEventManager.Event[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"randomness\",\"type\":\"uint256\"}],\"name\":\"postRandomness\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"randomnessId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"randomnessJobs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestRandomness\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"writer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f5ffd5b506040516115d03803806115d0833981810160405281019061003191906100e9565b61280081816001819055508060025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050505f60038190555050610114565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6100b88261008f565b9050919050565b6100c8816100ae565b81146100d2575f5ffd5b50565b5f815190506100e3816100bf565b92915050565b5f602082840312156100fe576100fd61008b565b5b5f61010b848285016100d5565b91505092915050565b6114af806101215f395ff3fe608060405234801561000f575f5ffd5b5060043610610086575f3560e01c8063a13e993f11610059578063a13e993f14610124578063b20e730b14610142578063cae62d3e1461015e578063f8413b071461018e57610086565b8063453a2abc1461008a578063453f4f62146100a85780635612cd17146100d85780638a54f929146100f4575b5f5ffd5b6100926101ac565b60405161009f919061089a565b60405180910390f35b6100c260048036038101906100bd91906108f7565b6101d1565b6040516100cf9190610931565b60405180910390f35b6100f260048036038101906100ed9190610ad1565b610232565b005b61010e600480360381019061010991906108f7565b610462565b60405161011b9190610931565b60405180910390f35b61012c610477565b6040516101399190610931565b60405180910390f35b61015c60048036038101906101579190610b2b565b61047d565b005b61017860048036038101906101739190610b69565b610579565b6040516101859190610d1a565b60405180910390f35b6101966106a1565b6040516101a39190610931565b60405180910390f35b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f5f60045f8481526020019081526020015f205490505f8103610229576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161022090610d94565b60405180910390fd5b80915050919050565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146102c1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102b890610dfc565b60405180910390fd5b5f5f5f8361ffff1681526020019081526020015f2090505f5f90505b818054905081101561045c575f5f90505f5f90505b85518110156104415785818151811061030e5761030d610e1a565b5b602002602001015184848154811061032957610328610e1a565b5b905f5260205f2090600302015f01540361043457836001858054905061034f9190610e74565b815481106103605761035f610e1a565b5b905f5260205f20906003020184848154811061037f5761037e610e1a565b5b905f5260205f2090600302015f820154815f0155600182015f9054906101000a900461ffff16816001015f6101000a81548161ffff021916908361ffff160217905550600282018160020190816103d691906110b9565b50905050838054806103eb576103ea61119e565b5b600190038181905f5260205f2090600302015f5f82015f9055600182015f6101000a81549061ffff0219169055600282015f6104279190610803565b5050905560019150610441565b80806001019150506102f2565b5080610456578180610452906111cb565b9250505b506102dd565b50505050565b6004602052805f5260405f205f915090505481565b60035481565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461050c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105039061125c565b60405180910390fd5b5f60045f8481526020019081526020015f20541461055f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610556906112c4565b60405180910390fd5b8060045f8481526020019081526020015f20819055505050565b60605f5f8361ffff1681526020019081526020015f20805480602002602001604051908101604052809291908181526020015f905b82821015610696578382905f5260205f2090600302016040518060600160405290815f8201548152602001600182015f9054906101000a900461ffff1661ffff1661ffff16815260200160028201805461060790610ed4565b80601f016020809104026020016040519081016040528092919081815260200182805461063390610ed4565b801561067e5780601f106106555761010080835404028352916020019161067e565b820191905f5260205f20905b81548152906001019060200180831161066157829003601f168201915b505050505081525050815260200190600101906105ae565b505050509050919050565b5f600160035f8282546106b491906112e2565b925050819055505f6003546040516020016106cf9190610931565b60405160208183030381529060405290506106ff6003545f60018111156106f9576106f8611315565b5b83610708565b60035491505090565b6001545f5f8461ffff1681526020019081526020015f208054905010610763576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161075a9061138c565b60405180910390fd5b5f5f8361ffff1681526020019081526020015f2060405180606001604052808581526020018461ffff16815260200183815250908060018154018082558091505060019003905f5260205f2090600302015f909190919091505f820151815f01556020820151816001015f6101000a81548161ffff021916908361ffff16021790555060408201518160020190816107fb91906113aa565b505050505050565b50805461080f90610ed4565b5f825580601f10610820575061083d565b601f0160209004905f5260205f209081019061083c9190610840565b5b50565b5b80821115610857575f815f905550600101610841565b5090565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6108848261085b565b9050919050565b6108948161087a565b82525050565b5f6020820190506108ad5f83018461088b565b92915050565b5f604051905090565b5f5ffd5b5f5ffd5b5f819050919050565b6108d6816108c4565b81146108e0575f5ffd5b50565b5f813590506108f1816108cd565b92915050565b5f6020828403121561090c5761090b6108bc565b5b5f610919848285016108e3565b91505092915050565b61092b816108c4565b82525050565b5f6020820190506109445f830184610922565b92915050565b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6109948261094e565b810181811067ffffffffffffffff821117156109b3576109b261095e565b5b80604052505050565b5f6109c56108b3565b90506109d1828261098b565b919050565b5f67ffffffffffffffff8211156109f0576109ef61095e565b5b602082029050602081019050919050565b5f5ffd5b5f610a17610a12846109d6565b6109bc565b90508083825260208201905060208402830185811115610a3a57610a39610a01565b5b835b81811015610a635780610a4f88826108e3565b845260208401935050602081019050610a3c565b5050509392505050565b5f82601f830112610a8157610a8061094a565b5b8135610a91848260208601610a05565b91505092915050565b5f61ffff82169050919050565b610ab081610a9a565b8114610aba575f5ffd5b50565b5f81359050610acb81610aa7565b92915050565b5f5f60408385031215610ae757610ae66108bc565b5b5f83013567ffffffffffffffff811115610b0457610b036108c0565b5b610b1085828601610a6d565b9250506020610b2185828601610abd565b9150509250929050565b5f5f60408385031215610b4157610b406108bc565b5b5f610b4e858286016108e3565b9250506020610b5f858286016108e3565b9150509250929050565b5f60208284031215610b7e57610b7d6108bc565b5b5f610b8b84828501610abd565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b610bc6816108c4565b82525050565b610bd581610a9a565b82525050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f610c0d82610bdb565b610c178185610be5565b9350610c27818560208601610bf5565b610c308161094e565b840191505092915050565b5f606083015f830151610c505f860182610bbd565b506020830151610c636020860182610bcc565b5060408301518482036040860152610c7b8282610c03565b9150508091505092915050565b5f610c938383610c3b565b905092915050565b5f602082019050919050565b5f610cb182610b94565b610cbb8185610b9e565b935083602082028501610ccd85610bae565b805f5b85811015610d085784840389528151610ce98582610c88565b9450610cf483610c9b565b925060208a01995050600181019050610cd0565b50829750879550505050505092915050565b5f6020820190508181035f830152610d328184610ca7565b905092915050565b5f82825260208201905092915050565b7f52616e646f6d6e657373206e6f7420706f7374656400000000000000000000005f82015250565b5f610d7e601583610d3a565b9150610d8982610d4a565b602082019050919050565b5f6020820190508181035f830152610dab81610d72565b9050919050565b7f4f6e6c79207772697465722063616e206572617365206576656e7473000000005f82015250565b5f610de6601c83610d3a565b9150610df182610db2565b602082019050919050565b5f6020820190508181035f830152610e1381610dda565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610e7e826108c4565b9150610e89836108c4565b9250828203905081811115610ea157610ea0610e47565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680610eeb57607f821691505b602082108103610efe57610efd610ea7565b5b50919050565b5f81549050610f1281610ed4565b9050919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302610f757fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610f3a565b610f7f8683610f3a565b95508019841693508086168417925050509392505050565b5f819050919050565b5f610fba610fb5610fb0846108c4565b610f97565b6108c4565b9050919050565b5f819050919050565b610fd383610fa0565b610fe7610fdf82610fc1565b848454610f46565b825550505050565b5f5f905090565b610ffe610fef565b611009818484610fca565b505050565b5b8181101561102c576110215f82610ff6565b60018101905061100f565b5050565b601f8211156110715761104281610f19565b61104b84610f2b565b8101602085101561105a578190505b61106e61106685610f2b565b83018261100e565b50505b505050565b5f82821c905092915050565b5f6110915f1984600802611076565b1980831691505092915050565b5f6110a98383611082565b9150826002028217905092915050565b8181036110c757505061119c565b6110d082610f04565b67ffffffffffffffff8111156110e9576110e861095e565b5b6110f38254610ed4565b6110fe828285611030565b5f601f83116001811461112b575f8415611119578287015490505b611123858261109e565b865550611195565b601f19841661113987610f19565b965061114486610f19565b5f5b8281101561116b57848901548255600182019150600185019450602081019050611146565b868310156111885784890154611184601f891682611082565b8355505b6001600288020188555050505b5050505050505b565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffd5b5f6111d5826108c4565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361120757611206610e47565b5b600182019050919050565b7f4f6e6c79207772697465722063616e20706f73742072616e646f6d6e657373005f82015250565b5f611246601f83610d3a565b915061125182611212565b602082019050919050565b5f6020820190508181035f8301526112738161123a565b9050919050565b7f52616e646f6d6e65737320616c726561647920706f73746564000000000000005f82015250565b5f6112ae601983610d3a565b91506112b98261127a565b602082019050919050565b5f6020820190508181035f8301526112db816112a2565b9050919050565b5f6112ec826108c4565b91506112f7836108c4565b925082820190508082111561130f5761130e610e47565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b7f4576656e74206275666665722069732066756c6c0000000000000000000000005f82015250565b5f611376601483610d3a565b915061138182611342565b602082019050919050565b5f6020820190508181035f8301526113a38161136a565b9050919050565b6113b382610bdb565b67ffffffffffffffff8111156113cc576113cb61095e565b5b6113d68254610ed4565b6113e1828285611030565b5f60209050601f831160018114611412575f8415611400578287015190505b61140a858261109e565b865550611471565b601f19841661142086610f19565b5f5b8281101561144757848901518255600182019150602085019450602081019050611422565b868310156114645784890151611460601f891682611082565b8355505b6001600288020188555050505b50505050505056fea2646970667358221220a4ed611c4b44b9a4f17de8d03a2293744384064afd1150f1c61972e4d714435a64736f6c634300081c0033",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// ContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractMetaData.Bin instead.
var ContractBin = ContractMetaData.Bin

// DeployContract deploys a new Ethereum contract, binding an instance of Contract to it.
func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend, _writer common.Address) (common.Address, *types.Transaction, *Contract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractBin), backend, _writer)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

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

// PollEvents is a free data retrieval call binding the contract method 0xcae62d3e.
//
// Solidity: function pollEvents(uint16 eventType) view returns((uint256,uint16,bytes)[])
func (_Contract *ContractCaller) PollEvents(opts *bind.CallOpts, eventType uint16) ([]EventManagerEvent, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "pollEvents", eventType)

	if err != nil {
		return *new([]EventManagerEvent), err
	}

	out0 := *abi.ConvertType(out[0], new([]EventManagerEvent)).(*[]EventManagerEvent)

	return out0, err

}

// PollEvents is a free data retrieval call binding the contract method 0xcae62d3e.
//
// Solidity: function pollEvents(uint16 eventType) view returns((uint256,uint16,bytes)[])
func (_Contract *ContractSession) PollEvents(eventType uint16) ([]EventManagerEvent, error) {
	return _Contract.Contract.PollEvents(&_Contract.CallOpts, eventType)
}

// PollEvents is a free data retrieval call binding the contract method 0xcae62d3e.
//
// Solidity: function pollEvents(uint16 eventType) view returns((uint256,uint16,bytes)[])
func (_Contract *ContractCallerSession) PollEvents(eventType uint16) ([]EventManagerEvent, error) {
	return _Contract.Contract.PollEvents(&_Contract.CallOpts, eventType)
}

// RandomnessId is a free data retrieval call binding the contract method 0xa13e993f.
//
// Solidity: function randomnessId() view returns(uint256)
func (_Contract *ContractCaller) RandomnessId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "randomnessId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RandomnessId is a free data retrieval call binding the contract method 0xa13e993f.
//
// Solidity: function randomnessId() view returns(uint256)
func (_Contract *ContractSession) RandomnessId() (*big.Int, error) {
	return _Contract.Contract.RandomnessId(&_Contract.CallOpts)
}

// RandomnessId is a free data retrieval call binding the contract method 0xa13e993f.
//
// Solidity: function randomnessId() view returns(uint256)
func (_Contract *ContractCallerSession) RandomnessId() (*big.Int, error) {
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

// EraseEvents is a paid mutator transaction binding the contract method 0x5612cd17.
//
// Solidity: function eraseEvents(uint256[] eventIds, uint16 eventType) returns()
func (_Contract *ContractTransactor) EraseEvents(opts *bind.TransactOpts, eventIds []*big.Int, eventType uint16) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "eraseEvents", eventIds, eventType)
}

// EraseEvents is a paid mutator transaction binding the contract method 0x5612cd17.
//
// Solidity: function eraseEvents(uint256[] eventIds, uint16 eventType) returns()
func (_Contract *ContractSession) EraseEvents(eventIds []*big.Int, eventType uint16) (*types.Transaction, error) {
	return _Contract.Contract.EraseEvents(&_Contract.TransactOpts, eventIds, eventType)
}

// EraseEvents is a paid mutator transaction binding the contract method 0x5612cd17.
//
// Solidity: function eraseEvents(uint256[] eventIds, uint16 eventType) returns()
func (_Contract *ContractTransactorSession) EraseEvents(eventIds []*big.Int, eventType uint16) (*types.Transaction, error) {
	return _Contract.Contract.EraseEvents(&_Contract.TransactOpts, eventIds, eventType)
}

// PostRandomness is a paid mutator transaction binding the contract method 0xb20e730b.
//
// Solidity: function postRandomness(uint256 id, uint256 randomness) returns()
func (_Contract *ContractTransactor) PostRandomness(opts *bind.TransactOpts, id *big.Int, randomness *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "postRandomness", id, randomness)
}

// PostRandomness is a paid mutator transaction binding the contract method 0xb20e730b.
//
// Solidity: function postRandomness(uint256 id, uint256 randomness) returns()
func (_Contract *ContractSession) PostRandomness(id *big.Int, randomness *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.PostRandomness(&_Contract.TransactOpts, id, randomness)
}

// PostRandomness is a paid mutator transaction binding the contract method 0xb20e730b.
//
// Solidity: function postRandomness(uint256 id, uint256 randomness) returns()
func (_Contract *ContractTransactorSession) PostRandomness(id *big.Int, randomness *big.Int) (*types.Transaction, error) {
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
