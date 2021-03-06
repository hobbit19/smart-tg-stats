// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// SmartTgStatsABI is the input ABI used to generate the binding from.
const SmartTgStatsABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"status\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"CEOWithdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxRequestID\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"statusRenew\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Requests\",\"outputs\":[{\"name\":\"channel\",\"type\":\"string\"},{\"name\":\"postID\",\"type\":\"uint32\"},{\"name\":\"bid\",\"type\":\"uint256\"},{\"name\":\"willUpdate\",\"type\":\"bool\"},{\"name\":\"timeCreated\",\"type\":\"uint256\"},{\"name\":\"timeFirst\",\"type\":\"uint256\"},{\"name\":\"timeLast\",\"type\":\"uint256\"},{\"name\":\"lastSubscribers\",\"type\":\"uint32\"},{\"name\":\"lastPostViews\",\"type\":\"uint32\"},{\"name\":\"postTime\",\"type\":\"uint256\"},{\"name\":\"error\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_requestID\",\"type\":\"uint256\"},{\"name\":\"_willUpdate\",\"type\":\"bool\"},{\"name\":\"_subscribers\",\"type\":\"uint32\"},{\"name\":\"_postViews\",\"type\":\"uint32\"},{\"name\":\"_postTime\",\"type\":\"uint256\"},{\"name\":\"_error\",\"type\":\"uint8\"}],\"name\":\"CEOAddResponse\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_channel\",\"type\":\"string\"},{\"name\":\"_postID\",\"type\":\"uint32\"}],\"name\":\"AddRequest\",\"outputs\":[{\"name\":\"requestID\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_isOn\",\"type\":\"bool\"}],\"name\":\"CEORenewStatus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_isOn\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"NewRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"NewResponse\",\"type\":\"event\"}]"

// SmartTgStatsBin is the compiled bytecode used for deploying new contracts.
const SmartTgStatsBin = `6080604052604051602080610de48339810180604052602081101561002357600080fd5b8101908080519060200190929190505050600073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461008e57600080fd5b336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600060146101000a81548160ff0219169083151502179055504260018190555050610ce5806100ff6000396000f3fe60806040526004361061007b5760003560e01c806340ac99931161004e57806340ac999314610160578063b3b4090c14610290578063c15f19fc1461030e578063dcfd0e11146103ed5761007b565b8063200d2ed21461008057806321213cac146100af57806325093fff1461010a5780633a72507814610135575b600080fd5b34801561008c57600080fd5b5061009561042a565b604051808215151515815260200191505060405180910390f35b3480156100bb57600080fd5b50610108600480360360408110156100d257600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061043d565b005b34801561011657600080fd5b5061011f61053c565b6040518082815260200191505060405180910390f35b34801561014157600080fd5b5061014a610542565b6040518082815260200191505060405180910390f35b34801561016c57600080fd5b506101996004803603602081101561018357600080fd5b8101908080359060200190929190505050610548565b60405180806020018c63ffffffff1663ffffffff1681526020018b81526020018a1515151581526020018981526020018881526020018781526020018663ffffffff1663ffffffff1681526020018563ffffffff1663ffffffff16815260200184815260200183600381111561020b57fe5b60ff16815260200182810382528d818151815260200191508051906020019080838360005b8381101561024b578082015181840152602081019050610230565b50505050905090810190601f1680156102785780820380516001836020036101000a031916815260200191505b509c5050505050505050505050505060405180910390f35b34801561029c57600080fd5b5061030c600480360360c08110156102b357600080fd5b8101908080359060200190929190803515159060200190929190803563ffffffff169060200190929190803563ffffffff16906020019092919080359060200190929190803560ff169060200190929190505050610684565b005b6103d76004803603604081101561032457600080fd5b810190808035906020019064010000000081111561034157600080fd5b82018360208201111561035357600080fd5b8035906020019184600183028401116401000000008311171561037557600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290803563ffffffff1690602001909291905050506108b0565b6040518082815260200191505060405180910390f35b3480156103f957600080fd5b506104286004803603602081101561041057600080fd5b81019080803515159060200190929190505050610b3c565b005b600060149054906101000a900460ff1681565b600073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141561049857600080fd5b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146104f157600080fd5b8173ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f19350505050158015610537573d6000803e3d6000fd5b505050565b60035481565b60015481565b6002602052806000526040600020600091509050806000018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156105f45780601f106105c9576101008083540402835291602001916105f4565b820191906000526020600020905b8154815290600101906020018083116105d757829003601f168201915b5050505050908060010160009054906101000a900463ffffffff16908060020154908060030160009054906101000a900460ff16908060040154908060050154908060060154908060070160009054906101000a900463ffffffff16908060070160049054906101000a900463ffffffff16908060080154908060090160009054906101000a900460ff1690508b565b600073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614156106df57600080fd5b3373ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461073857600080fd5b60008611801561074a57506003548611155b61075357600080fd5b600115156002600088815260200190815260200160002060030160009054906101000a900460ff1615151461078757600080fd5b846002600088815260200190815260200160002060030160006101000a81548160ff021916908315150217905550836002600088815260200190815260200160002060070160006101000a81548163ffffffff021916908363ffffffff160217905550826002600088815260200190815260200160002060070160046101000a81548163ffffffff021916908363ffffffff160217905550816002600088815260200190815260200160002060080181905550806002600088815260200190815260200160002060090160006101000a81548160ff0219169083600381111561086c57fe5b02179055507f53b55b3d6349be0707243ec0a2bc86efff3ab477b9aabf8a02d4e3dfdc450b7c866040518082815260200191505060405180910390a1505050505050565b60008073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141561090c57600080fd5b60011515600060149054906101000a900460ff1615151461092c57600080fd5b60006001544203101561093e57600080fd5b6212750060015442031061095157600080fd5b600160036000828254019250508190555060035490506040518061016001604052808481526020018363ffffffff1681526020013481526020016001151581526020014281526020016000815260200160008152602001600063ffffffff168152602001600063ffffffff16815260200160008152602001600060038111156109d657fe5b815250600260008381526020019081526020016000206000820151816000019080519060200190610a08929190610c14565b5060208201518160010160006101000a81548163ffffffff021916908363ffffffff1602179055506040820151816002015560608201518160030160006101000a81548160ff0219169083151502179055506080820151816004015560a0820151816005015560c0820151816006015560e08201518160070160006101000a81548163ffffffff021916908363ffffffff1602179055506101008201518160070160046101000a81548163ffffffff021916908363ffffffff16021790555061012082015181600801556101408201518160090160006101000a81548160ff02191690836003811115610af757fe5b02179055509050507fe5e7dd91b3ed7fb84c335f117423d1b9bbbed2d76c57e81520dd681b9ede9885816040518082815260200191505060405180910390a192915050565b600073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415610b9757600080fd5b3373ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610bf057600080fd5b80600060146101000a81548160ff0219169083151502179055504260018190555050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610c5557805160ff1916838001178555610c83565b82800160010185558215610c83579182015b82811115610c82578251825591602001919060010190610c67565b5b509050610c909190610c94565b5090565b610cb691905b80821115610cb2576000816000905550600101610c9a565b5090565b9056fea165627a7a72305820d1d88af0b1b2567d38164d92d022c5a5063e648be5dd2dbfdd0932afa223ff9b0029`

// DeploySmartTgStats deploys a new Ethereum contract, binding an instance of SmartTgStats to it.
func DeploySmartTgStats(auth *bind.TransactOpts, backend bind.ContractBackend, _isOn bool) (common.Address, *types.Transaction, *SmartTgStats, error) {
	parsed, err := abi.JSON(strings.NewReader(SmartTgStatsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SmartTgStatsBin), backend, _isOn)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SmartTgStats{SmartTgStatsCaller: SmartTgStatsCaller{contract: contract}, SmartTgStatsTransactor: SmartTgStatsTransactor{contract: contract}, SmartTgStatsFilterer: SmartTgStatsFilterer{contract: contract}}, nil
}

// SmartTgStats is an auto generated Go binding around an Ethereum contract.
type SmartTgStats struct {
	SmartTgStatsCaller     // Read-only binding to the contract
	SmartTgStatsTransactor // Write-only binding to the contract
	SmartTgStatsFilterer   // Log filterer for contract events
}

// SmartTgStatsCaller is an auto generated read-only Go binding around an Ethereum contract.
type SmartTgStatsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmartTgStatsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SmartTgStatsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmartTgStatsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SmartTgStatsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmartTgStatsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SmartTgStatsSession struct {
	Contract     *SmartTgStats     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SmartTgStatsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SmartTgStatsCallerSession struct {
	Contract *SmartTgStatsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// SmartTgStatsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SmartTgStatsTransactorSession struct {
	Contract     *SmartTgStatsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// SmartTgStatsRaw is an auto generated low-level Go binding around an Ethereum contract.
type SmartTgStatsRaw struct {
	Contract *SmartTgStats // Generic contract binding to access the raw methods on
}

// SmartTgStatsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SmartTgStatsCallerRaw struct {
	Contract *SmartTgStatsCaller // Generic read-only contract binding to access the raw methods on
}

// SmartTgStatsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SmartTgStatsTransactorRaw struct {
	Contract *SmartTgStatsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSmartTgStats creates a new instance of SmartTgStats, bound to a specific deployed contract.
func NewSmartTgStats(address common.Address, backend bind.ContractBackend) (*SmartTgStats, error) {
	contract, err := bindSmartTgStats(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SmartTgStats{SmartTgStatsCaller: SmartTgStatsCaller{contract: contract}, SmartTgStatsTransactor: SmartTgStatsTransactor{contract: contract}, SmartTgStatsFilterer: SmartTgStatsFilterer{contract: contract}}, nil
}

// NewSmartTgStatsCaller creates a new read-only instance of SmartTgStats, bound to a specific deployed contract.
func NewSmartTgStatsCaller(address common.Address, caller bind.ContractCaller) (*SmartTgStatsCaller, error) {
	contract, err := bindSmartTgStats(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SmartTgStatsCaller{contract: contract}, nil
}

// NewSmartTgStatsTransactor creates a new write-only instance of SmartTgStats, bound to a specific deployed contract.
func NewSmartTgStatsTransactor(address common.Address, transactor bind.ContractTransactor) (*SmartTgStatsTransactor, error) {
	contract, err := bindSmartTgStats(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SmartTgStatsTransactor{contract: contract}, nil
}

// NewSmartTgStatsFilterer creates a new log filterer instance of SmartTgStats, bound to a specific deployed contract.
func NewSmartTgStatsFilterer(address common.Address, filterer bind.ContractFilterer) (*SmartTgStatsFilterer, error) {
	contract, err := bindSmartTgStats(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SmartTgStatsFilterer{contract: contract}, nil
}

// bindSmartTgStats binds a generic wrapper to an already deployed contract.
func bindSmartTgStats(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SmartTgStatsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SmartTgStats *SmartTgStatsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SmartTgStats.Contract.SmartTgStatsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SmartTgStats *SmartTgStatsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SmartTgStats.Contract.SmartTgStatsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SmartTgStats *SmartTgStatsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SmartTgStats.Contract.SmartTgStatsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SmartTgStats *SmartTgStatsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SmartTgStats.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SmartTgStats *SmartTgStatsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SmartTgStats.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SmartTgStats *SmartTgStatsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SmartTgStats.Contract.contract.Transact(opts, method, params...)
}

// Requests is a free data retrieval call binding the contract method 0x40ac9993.
//
// Solidity: function Requests(uint256 ) constant returns(string channel, uint32 postID, uint256 bid, bool willUpdate, uint256 timeCreated, uint256 timeFirst, uint256 timeLast, uint32 lastSubscribers, uint32 lastPostViews, uint256 postTime, uint8 error)
func (_SmartTgStats *SmartTgStatsCaller) Requests(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Channel         string
	PostID          uint32
	Bid             *big.Int
	WillUpdate      bool
	TimeCreated     *big.Int
	TimeFirst       *big.Int
	TimeLast        *big.Int
	LastSubscribers uint32
	LastPostViews   uint32
	PostTime        *big.Int
	Error           uint8
}, error) {
	ret := new(struct {
		Channel         string
		PostID          uint32
		Bid             *big.Int
		WillUpdate      bool
		TimeCreated     *big.Int
		TimeFirst       *big.Int
		TimeLast        *big.Int
		LastSubscribers uint32
		LastPostViews   uint32
		PostTime        *big.Int
		Error           uint8
	})
	out := ret
	err := _SmartTgStats.contract.Call(opts, out, "Requests", arg0)
	return *ret, err
}

// Requests is a free data retrieval call binding the contract method 0x40ac9993.
//
// Solidity: function Requests(uint256 ) constant returns(string channel, uint32 postID, uint256 bid, bool willUpdate, uint256 timeCreated, uint256 timeFirst, uint256 timeLast, uint32 lastSubscribers, uint32 lastPostViews, uint256 postTime, uint8 error)
func (_SmartTgStats *SmartTgStatsSession) Requests(arg0 *big.Int) (struct {
	Channel         string
	PostID          uint32
	Bid             *big.Int
	WillUpdate      bool
	TimeCreated     *big.Int
	TimeFirst       *big.Int
	TimeLast        *big.Int
	LastSubscribers uint32
	LastPostViews   uint32
	PostTime        *big.Int
	Error           uint8
}, error) {
	return _SmartTgStats.Contract.Requests(&_SmartTgStats.CallOpts, arg0)
}

// Requests is a free data retrieval call binding the contract method 0x40ac9993.
//
// Solidity: function Requests(uint256 ) constant returns(string channel, uint32 postID, uint256 bid, bool willUpdate, uint256 timeCreated, uint256 timeFirst, uint256 timeLast, uint32 lastSubscribers, uint32 lastPostViews, uint256 postTime, uint8 error)
func (_SmartTgStats *SmartTgStatsCallerSession) Requests(arg0 *big.Int) (struct {
	Channel         string
	PostID          uint32
	Bid             *big.Int
	WillUpdate      bool
	TimeCreated     *big.Int
	TimeFirst       *big.Int
	TimeLast        *big.Int
	LastSubscribers uint32
	LastPostViews   uint32
	PostTime        *big.Int
	Error           uint8
}, error) {
	return _SmartTgStats.Contract.Requests(&_SmartTgStats.CallOpts, arg0)
}

// MaxRequestID is a free data retrieval call binding the contract method 0x25093fff.
//
// Solidity: function maxRequestID() constant returns(uint256)
func (_SmartTgStats *SmartTgStatsCaller) MaxRequestID(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SmartTgStats.contract.Call(opts, out, "maxRequestID")
	return *ret0, err
}

// MaxRequestID is a free data retrieval call binding the contract method 0x25093fff.
//
// Solidity: function maxRequestID() constant returns(uint256)
func (_SmartTgStats *SmartTgStatsSession) MaxRequestID() (*big.Int, error) {
	return _SmartTgStats.Contract.MaxRequestID(&_SmartTgStats.CallOpts)
}

// MaxRequestID is a free data retrieval call binding the contract method 0x25093fff.
//
// Solidity: function maxRequestID() constant returns(uint256)
func (_SmartTgStats *SmartTgStatsCallerSession) MaxRequestID() (*big.Int, error) {
	return _SmartTgStats.Contract.MaxRequestID(&_SmartTgStats.CallOpts)
}

// Status is a free data retrieval call binding the contract method 0x200d2ed2.
//
// Solidity: function status() constant returns(bool)
func (_SmartTgStats *SmartTgStatsCaller) Status(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SmartTgStats.contract.Call(opts, out, "status")
	return *ret0, err
}

// Status is a free data retrieval call binding the contract method 0x200d2ed2.
//
// Solidity: function status() constant returns(bool)
func (_SmartTgStats *SmartTgStatsSession) Status() (bool, error) {
	return _SmartTgStats.Contract.Status(&_SmartTgStats.CallOpts)
}

// Status is a free data retrieval call binding the contract method 0x200d2ed2.
//
// Solidity: function status() constant returns(bool)
func (_SmartTgStats *SmartTgStatsCallerSession) Status() (bool, error) {
	return _SmartTgStats.Contract.Status(&_SmartTgStats.CallOpts)
}

// StatusRenew is a free data retrieval call binding the contract method 0x3a725078.
//
// Solidity: function statusRenew() constant returns(uint256)
func (_SmartTgStats *SmartTgStatsCaller) StatusRenew(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SmartTgStats.contract.Call(opts, out, "statusRenew")
	return *ret0, err
}

// StatusRenew is a free data retrieval call binding the contract method 0x3a725078.
//
// Solidity: function statusRenew() constant returns(uint256)
func (_SmartTgStats *SmartTgStatsSession) StatusRenew() (*big.Int, error) {
	return _SmartTgStats.Contract.StatusRenew(&_SmartTgStats.CallOpts)
}

// StatusRenew is a free data retrieval call binding the contract method 0x3a725078.
//
// Solidity: function statusRenew() constant returns(uint256)
func (_SmartTgStats *SmartTgStatsCallerSession) StatusRenew() (*big.Int, error) {
	return _SmartTgStats.Contract.StatusRenew(&_SmartTgStats.CallOpts)
}

// AddRequest is a paid mutator transaction binding the contract method 0xc15f19fc.
//
// Solidity: function AddRequest(string _channel, uint32 _postID) returns(uint256 requestID)
func (_SmartTgStats *SmartTgStatsTransactor) AddRequest(opts *bind.TransactOpts, _channel string, _postID uint32) (*types.Transaction, error) {
	return _SmartTgStats.contract.Transact(opts, "AddRequest", _channel, _postID)
}

// AddRequest is a paid mutator transaction binding the contract method 0xc15f19fc.
//
// Solidity: function AddRequest(string _channel, uint32 _postID) returns(uint256 requestID)
func (_SmartTgStats *SmartTgStatsSession) AddRequest(_channel string, _postID uint32) (*types.Transaction, error) {
	return _SmartTgStats.Contract.AddRequest(&_SmartTgStats.TransactOpts, _channel, _postID)
}

// AddRequest is a paid mutator transaction binding the contract method 0xc15f19fc.
//
// Solidity: function AddRequest(string _channel, uint32 _postID) returns(uint256 requestID)
func (_SmartTgStats *SmartTgStatsTransactorSession) AddRequest(_channel string, _postID uint32) (*types.Transaction, error) {
	return _SmartTgStats.Contract.AddRequest(&_SmartTgStats.TransactOpts, _channel, _postID)
}

// CEOAddResponse is a paid mutator transaction binding the contract method 0xb3b4090c.
//
// Solidity: function CEOAddResponse(uint256 _requestID, bool _willUpdate, uint32 _subscribers, uint32 _postViews, uint256 _postTime, uint8 _error) returns()
func (_SmartTgStats *SmartTgStatsTransactor) CEOAddResponse(opts *bind.TransactOpts, _requestID *big.Int, _willUpdate bool, _subscribers uint32, _postViews uint32, _postTime *big.Int, _error uint8) (*types.Transaction, error) {
	return _SmartTgStats.contract.Transact(opts, "CEOAddResponse", _requestID, _willUpdate, _subscribers, _postViews, _postTime, _error)
}

// CEOAddResponse is a paid mutator transaction binding the contract method 0xb3b4090c.
//
// Solidity: function CEOAddResponse(uint256 _requestID, bool _willUpdate, uint32 _subscribers, uint32 _postViews, uint256 _postTime, uint8 _error) returns()
func (_SmartTgStats *SmartTgStatsSession) CEOAddResponse(_requestID *big.Int, _willUpdate bool, _subscribers uint32, _postViews uint32, _postTime *big.Int, _error uint8) (*types.Transaction, error) {
	return _SmartTgStats.Contract.CEOAddResponse(&_SmartTgStats.TransactOpts, _requestID, _willUpdate, _subscribers, _postViews, _postTime, _error)
}

// CEOAddResponse is a paid mutator transaction binding the contract method 0xb3b4090c.
//
// Solidity: function CEOAddResponse(uint256 _requestID, bool _willUpdate, uint32 _subscribers, uint32 _postViews, uint256 _postTime, uint8 _error) returns()
func (_SmartTgStats *SmartTgStatsTransactorSession) CEOAddResponse(_requestID *big.Int, _willUpdate bool, _subscribers uint32, _postViews uint32, _postTime *big.Int, _error uint8) (*types.Transaction, error) {
	return _SmartTgStats.Contract.CEOAddResponse(&_SmartTgStats.TransactOpts, _requestID, _willUpdate, _subscribers, _postViews, _postTime, _error)
}

// CEORenewStatus is a paid mutator transaction binding the contract method 0xdcfd0e11.
//
// Solidity: function CEORenewStatus(bool _isOn) returns()
func (_SmartTgStats *SmartTgStatsTransactor) CEORenewStatus(opts *bind.TransactOpts, _isOn bool) (*types.Transaction, error) {
	return _SmartTgStats.contract.Transact(opts, "CEORenewStatus", _isOn)
}

// CEORenewStatus is a paid mutator transaction binding the contract method 0xdcfd0e11.
//
// Solidity: function CEORenewStatus(bool _isOn) returns()
func (_SmartTgStats *SmartTgStatsSession) CEORenewStatus(_isOn bool) (*types.Transaction, error) {
	return _SmartTgStats.Contract.CEORenewStatus(&_SmartTgStats.TransactOpts, _isOn)
}

// CEORenewStatus is a paid mutator transaction binding the contract method 0xdcfd0e11.
//
// Solidity: function CEORenewStatus(bool _isOn) returns()
func (_SmartTgStats *SmartTgStatsTransactorSession) CEORenewStatus(_isOn bool) (*types.Transaction, error) {
	return _SmartTgStats.Contract.CEORenewStatus(&_SmartTgStats.TransactOpts, _isOn)
}

// CEOWithdraw is a paid mutator transaction binding the contract method 0x21213cac.
//
// Solidity: function CEOWithdraw(address _to, uint256 _amount) returns()
func (_SmartTgStats *SmartTgStatsTransactor) CEOWithdraw(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SmartTgStats.contract.Transact(opts, "CEOWithdraw", _to, _amount)
}

// CEOWithdraw is a paid mutator transaction binding the contract method 0x21213cac.
//
// Solidity: function CEOWithdraw(address _to, uint256 _amount) returns()
func (_SmartTgStats *SmartTgStatsSession) CEOWithdraw(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SmartTgStats.Contract.CEOWithdraw(&_SmartTgStats.TransactOpts, _to, _amount)
}

// CEOWithdraw is a paid mutator transaction binding the contract method 0x21213cac.
//
// Solidity: function CEOWithdraw(address _to, uint256 _amount) returns()
func (_SmartTgStats *SmartTgStatsTransactorSession) CEOWithdraw(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SmartTgStats.Contract.CEOWithdraw(&_SmartTgStats.TransactOpts, _to, _amount)
}

// SmartTgStatsNewRequestIterator is returned from FilterNewRequest and is used to iterate over the raw logs and unpacked data for NewRequest events raised by the SmartTgStats contract.
type SmartTgStatsNewRequestIterator struct {
	Event *SmartTgStatsNewRequest // Event containing the contract specifics and raw log

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
func (it *SmartTgStatsNewRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartTgStatsNewRequest)
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
		it.Event = new(SmartTgStatsNewRequest)
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
func (it *SmartTgStatsNewRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartTgStatsNewRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartTgStatsNewRequest represents a NewRequest event raised by the SmartTgStats contract.
type SmartTgStatsNewRequest struct {
	*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterNewRequest is a free log retrieval operation binding the contract event 0xe5e7dd91b3ed7fb84c335f117423d1b9bbbed2d76c57e81520dd681b9ede9885.
//
// Solidity: event NewRequest(uint256 )
func (_SmartTgStats *SmartTgStatsFilterer) FilterNewRequest(opts *bind.FilterOpts) (*SmartTgStatsNewRequestIterator, error) {

	logs, sub, err := _SmartTgStats.contract.FilterLogs(opts, "NewRequest")
	if err != nil {
		return nil, err
	}
	return &SmartTgStatsNewRequestIterator{contract: _SmartTgStats.contract, event: "NewRequest", logs: logs, sub: sub}, nil
}

// WatchNewRequest is a free log subscription operation binding the contract event 0xe5e7dd91b3ed7fb84c335f117423d1b9bbbed2d76c57e81520dd681b9ede9885.
//
// Solidity: event NewRequest(uint256 )
func (_SmartTgStats *SmartTgStatsFilterer) WatchNewRequest(opts *bind.WatchOpts, sink chan<- *SmartTgStatsNewRequest) (event.Subscription, error) {

	logs, sub, err := _SmartTgStats.contract.WatchLogs(opts, "NewRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartTgStatsNewRequest)
				if err := _SmartTgStats.contract.UnpackLog(event, "NewRequest", log); err != nil {
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

// SmartTgStatsNewResponseIterator is returned from FilterNewResponse and is used to iterate over the raw logs and unpacked data for NewResponse events raised by the SmartTgStats contract.
type SmartTgStatsNewResponseIterator struct {
	Event *SmartTgStatsNewResponse // Event containing the contract specifics and raw log

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
func (it *SmartTgStatsNewResponseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartTgStatsNewResponse)
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
		it.Event = new(SmartTgStatsNewResponse)
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
func (it *SmartTgStatsNewResponseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartTgStatsNewResponseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartTgStatsNewResponse represents a NewResponse event raised by the SmartTgStats contract.
type SmartTgStatsNewResponse struct {
	*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterNewResponse is a free log retrieval operation binding the contract event 0x53b55b3d6349be0707243ec0a2bc86efff3ab477b9aabf8a02d4e3dfdc450b7c.
//
// Solidity: event NewResponse(uint256 )
func (_SmartTgStats *SmartTgStatsFilterer) FilterNewResponse(opts *bind.FilterOpts) (*SmartTgStatsNewResponseIterator, error) {

	logs, sub, err := _SmartTgStats.contract.FilterLogs(opts, "NewResponse")
	if err != nil {
		return nil, err
	}
	return &SmartTgStatsNewResponseIterator{contract: _SmartTgStats.contract, event: "NewResponse", logs: logs, sub: sub}, nil
}

// WatchNewResponse is a free log subscription operation binding the contract event 0x53b55b3d6349be0707243ec0a2bc86efff3ab477b9aabf8a02d4e3dfdc450b7c.
//
// Solidity: event NewResponse(uint256 )
func (_SmartTgStats *SmartTgStatsFilterer) WatchNewResponse(opts *bind.WatchOpts, sink chan<- *SmartTgStatsNewResponse) (event.Subscription, error) {

	logs, sub, err := _SmartTgStats.contract.WatchLogs(opts, "NewResponse")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartTgStatsNewResponse)
				if err := _SmartTgStats.contract.UnpackLog(event, "NewResponse", log); err != nil {
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
