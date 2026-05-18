// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package baccarat

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

// BaccaratMetaData contains all meta data concerning the Baccarat contract.
var BaccaratMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"token\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"balance\",\"type\":\"int256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"token\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"balance\",\"type\":\"int256\"}],\"name\":\"DepositPrize\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"token\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"balance\",\"type\":\"int256\"}],\"name\":\"Settle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"token\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"balance\",\"type\":\"int256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"token\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"balance\",\"type\":\"int256\"}],\"name\":\"WithdrawPrize\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"token\",\"type\":\"uint8\"}],\"name\":\"bet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"token\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"token\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"depositPrizePool\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"token\",\"type\":\"uint8\"}],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"token\",\"type\":\"uint8\"}],\"name\":\"getPrizePool\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"roundId\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"ID\",\"type\":\"uint64\"}],\"name\":\"setRoundId\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"SettleAmount\",\"type\":\"int256\"},{\"internalType\":\"uint8\",\"name\":\"token\",\"type\":\"uint8\"}],\"name\":\"settle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"token\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"token\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawPrizePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawPrizePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// BaccaratABI is the input ABI used to generate the binding from.
// Deprecated: Use BaccaratMetaData.ABI instead.
var BaccaratABI = BaccaratMetaData.ABI

// Baccarat is an auto generated Go binding around an Ethereum contract.
type Baccarat struct {
	BaccaratCaller     // Read-only binding to the contract
	BaccaratTransactor // Write-only binding to the contract
	BaccaratFilterer   // Log filterer for contract events
}

// BaccaratCaller is an auto generated read-only Go binding around an Ethereum contract.
type BaccaratCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaccaratTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BaccaratTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaccaratFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BaccaratFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaccaratSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BaccaratSession struct {
	Contract     *Baccarat         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BaccaratCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BaccaratCallerSession struct {
	Contract *BaccaratCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BaccaratTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BaccaratTransactorSession struct {
	Contract     *BaccaratTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BaccaratRaw is an auto generated low-level Go binding around an Ethereum contract.
type BaccaratRaw struct {
	Contract *Baccarat // Generic contract binding to access the raw methods on
}

// BaccaratCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BaccaratCallerRaw struct {
	Contract *BaccaratCaller // Generic read-only contract binding to access the raw methods on
}

// BaccaratTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BaccaratTransactorRaw struct {
	Contract *BaccaratTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBaccarat creates a new instance of Baccarat, bound to a specific deployed contract.
func NewBaccarat(address common.Address, backend bind.ContractBackend) (*Baccarat, error) {
	contract, err := bindBaccarat(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Baccarat{BaccaratCaller: BaccaratCaller{contract: contract}, BaccaratTransactor: BaccaratTransactor{contract: contract}, BaccaratFilterer: BaccaratFilterer{contract: contract}}, nil
}

// NewBaccaratCaller creates a new read-only instance of Baccarat, bound to a specific deployed contract.
func NewBaccaratCaller(address common.Address, caller bind.ContractCaller) (*BaccaratCaller, error) {
	contract, err := bindBaccarat(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BaccaratCaller{contract: contract}, nil
}

// NewBaccaratTransactor creates a new write-only instance of Baccarat, bound to a specific deployed contract.
func NewBaccaratTransactor(address common.Address, transactor bind.ContractTransactor) (*BaccaratTransactor, error) {
	contract, err := bindBaccarat(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BaccaratTransactor{contract: contract}, nil
}

// NewBaccaratFilterer creates a new log filterer instance of Baccarat, bound to a specific deployed contract.
func NewBaccaratFilterer(address common.Address, filterer bind.ContractFilterer) (*BaccaratFilterer, error) {
	contract, err := bindBaccarat(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BaccaratFilterer{contract: contract}, nil
}

// bindBaccarat binds a generic wrapper to an already deployed contract.
func bindBaccarat(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BaccaratMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Baccarat *BaccaratRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Baccarat.Contract.BaccaratCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Baccarat *BaccaratRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Baccarat.Contract.BaccaratTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Baccarat *BaccaratRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Baccarat.Contract.BaccaratTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Baccarat *BaccaratCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Baccarat.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Baccarat *BaccaratTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Baccarat.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Baccarat *BaccaratTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Baccarat.Contract.contract.Transact(opts, method, params...)
}

// GetBalance is a free data retrieval call binding the contract method 0xe3a12480.
//
// Solidity: function getBalance(uint8 token) view returns(int256)
func (_Baccarat *BaccaratCaller) GetBalance(opts *bind.CallOpts, token uint8) (*big.Int, error) {
	var out []interface{}
	err := _Baccarat.contract.Call(opts, &out, "getBalance", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0xe3a12480.
//
// Solidity: function getBalance(uint8 token) view returns(int256)
func (_Baccarat *BaccaratSession) GetBalance(token uint8) (*big.Int, error) {
	return _Baccarat.Contract.GetBalance(&_Baccarat.CallOpts, token)
}

// GetBalance is a free data retrieval call binding the contract method 0xe3a12480.
//
// Solidity: function getBalance(uint8 token) view returns(int256)
func (_Baccarat *BaccaratCallerSession) GetBalance(token uint8) (*big.Int, error) {
	return _Baccarat.Contract.GetBalance(&_Baccarat.CallOpts, token)
}

// GetPrizePool is a free data retrieval call binding the contract method 0xfe09ac4f.
//
// Solidity: function getPrizePool(uint8 token) view returns(int256)
func (_Baccarat *BaccaratCaller) GetPrizePool(opts *bind.CallOpts, token uint8) (*big.Int, error) {
	var out []interface{}
	err := _Baccarat.contract.Call(opts, &out, "getPrizePool", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrizePool is a free data retrieval call binding the contract method 0xfe09ac4f.
//
// Solidity: function getPrizePool(uint8 token) view returns(int256)
func (_Baccarat *BaccaratSession) GetPrizePool(token uint8) (*big.Int, error) {
	return _Baccarat.Contract.GetPrizePool(&_Baccarat.CallOpts, token)
}

// GetPrizePool is a free data retrieval call binding the contract method 0xfe09ac4f.
//
// Solidity: function getPrizePool(uint8 token) view returns(int256)
func (_Baccarat *BaccaratCallerSession) GetPrizePool(token uint8) (*big.Int, error) {
	return _Baccarat.Contract.GetPrizePool(&_Baccarat.CallOpts, token)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Baccarat *BaccaratCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Baccarat.contract.Call(opts, &out, "isOwner")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Baccarat *BaccaratSession) IsOwner() (bool, error) {
	return _Baccarat.Contract.IsOwner(&_Baccarat.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Baccarat *BaccaratCallerSession) IsOwner() (bool, error) {
	return _Baccarat.Contract.IsOwner(&_Baccarat.CallOpts)
}

// RoundId is a free data retrieval call binding the contract method 0x8cd221c9.
//
// Solidity: function roundId() view returns(uint64)
func (_Baccarat *BaccaratCaller) RoundId(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Baccarat.contract.Call(opts, &out, "roundId")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// RoundId is a free data retrieval call binding the contract method 0x8cd221c9.
//
// Solidity: function roundId() view returns(uint64)
func (_Baccarat *BaccaratSession) RoundId() (uint64, error) {
	return _Baccarat.Contract.RoundId(&_Baccarat.CallOpts)
}

// RoundId is a free data retrieval call binding the contract method 0x8cd221c9.
//
// Solidity: function roundId() view returns(uint64)
func (_Baccarat *BaccaratCallerSession) RoundId() (uint64, error) {
	return _Baccarat.Contract.RoundId(&_Baccarat.CallOpts)
}

// Bet is a paid mutator transaction binding the contract method 0xd0821b0e.
//
// Solidity: function bet(uint8 token) returns()
func (_Baccarat *BaccaratTransactor) Bet(opts *bind.TransactOpts, token uint8) (*types.Transaction, error) {
	return _Baccarat.contract.Transact(opts, "bet", token)
}

// Bet is a paid mutator transaction binding the contract method 0xd0821b0e.
//
// Solidity: function bet(uint8 token) returns()
func (_Baccarat *BaccaratSession) Bet(token uint8) (*types.Transaction, error) {
	return _Baccarat.Contract.Bet(&_Baccarat.TransactOpts, token)
}

// Bet is a paid mutator transaction binding the contract method 0xd0821b0e.
//
// Solidity: function bet(uint8 token) returns()
func (_Baccarat *BaccaratTransactorSession) Bet(token uint8) (*types.Transaction, error) {
	return _Baccarat.Contract.Bet(&_Baccarat.TransactOpts, token)
}

// Deposit is a paid mutator transaction binding the contract method 0xf4d4c9d7.
//
// Solidity: function deposit(uint8 token, uint256 amount) payable returns()
func (_Baccarat *BaccaratTransactor) Deposit(opts *bind.TransactOpts, token uint8, amount *big.Int) (*types.Transaction, error) {
	return _Baccarat.contract.Transact(opts, "deposit", token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xf4d4c9d7.
//
// Solidity: function deposit(uint8 token, uint256 amount) payable returns()
func (_Baccarat *BaccaratSession) Deposit(token uint8, amount *big.Int) (*types.Transaction, error) {
	return _Baccarat.Contract.Deposit(&_Baccarat.TransactOpts, token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xf4d4c9d7.
//
// Solidity: function deposit(uint8 token, uint256 amount) payable returns()
func (_Baccarat *BaccaratTransactorSession) Deposit(token uint8, amount *big.Int) (*types.Transaction, error) {
	return _Baccarat.Contract.Deposit(&_Baccarat.TransactOpts, token, amount)
}

// DepositPrizePool is a paid mutator transaction binding the contract method 0xfe6c1826.
//
// Solidity: function depositPrizePool(uint8 token, uint256 amount) payable returns()
func (_Baccarat *BaccaratTransactor) DepositPrizePool(opts *bind.TransactOpts, token uint8, amount *big.Int) (*types.Transaction, error) {
	return _Baccarat.contract.Transact(opts, "depositPrizePool", token, amount)
}

// DepositPrizePool is a paid mutator transaction binding the contract method 0xfe6c1826.
//
// Solidity: function depositPrizePool(uint8 token, uint256 amount) payable returns()
func (_Baccarat *BaccaratSession) DepositPrizePool(token uint8, amount *big.Int) (*types.Transaction, error) {
	return _Baccarat.Contract.DepositPrizePool(&_Baccarat.TransactOpts, token, amount)
}

// DepositPrizePool is a paid mutator transaction binding the contract method 0xfe6c1826.
//
// Solidity: function depositPrizePool(uint8 token, uint256 amount) payable returns()
func (_Baccarat *BaccaratTransactorSession) DepositPrizePool(token uint8, amount *big.Int) (*types.Transaction, error) {
	return _Baccarat.Contract.DepositPrizePool(&_Baccarat.TransactOpts, token, amount)
}

// SetRoundId is a paid mutator transaction binding the contract method 0x1d3c8b6e.
//
// Solidity: function setRoundId(uint64 ID) returns()
func (_Baccarat *BaccaratTransactor) SetRoundId(opts *bind.TransactOpts, ID uint64) (*types.Transaction, error) {
	return _Baccarat.contract.Transact(opts, "setRoundId", ID)
}

// SetRoundId is a paid mutator transaction binding the contract method 0x1d3c8b6e.
//
// Solidity: function setRoundId(uint64 ID) returns()
func (_Baccarat *BaccaratSession) SetRoundId(ID uint64) (*types.Transaction, error) {
	return _Baccarat.Contract.SetRoundId(&_Baccarat.TransactOpts, ID)
}

// SetRoundId is a paid mutator transaction binding the contract method 0x1d3c8b6e.
//
// Solidity: function setRoundId(uint64 ID) returns()
func (_Baccarat *BaccaratTransactorSession) SetRoundId(ID uint64) (*types.Transaction, error) {
	return _Baccarat.Contract.SetRoundId(&_Baccarat.TransactOpts, ID)
}

// Settle is a paid mutator transaction binding the contract method 0x4fea5bff.
//
// Solidity: function settle(address player, int256 SettleAmount, uint8 token) returns()
func (_Baccarat *BaccaratTransactor) Settle(opts *bind.TransactOpts, player common.Address, SettleAmount *big.Int, token uint8) (*types.Transaction, error) {
	return _Baccarat.contract.Transact(opts, "settle", player, SettleAmount, token)
}

// Settle is a paid mutator transaction binding the contract method 0x4fea5bff.
//
// Solidity: function settle(address player, int256 SettleAmount, uint8 token) returns()
func (_Baccarat *BaccaratSession) Settle(player common.Address, SettleAmount *big.Int, token uint8) (*types.Transaction, error) {
	return _Baccarat.Contract.Settle(&_Baccarat.TransactOpts, player, SettleAmount, token)
}

// Settle is a paid mutator transaction binding the contract method 0x4fea5bff.
//
// Solidity: function settle(address player, int256 SettleAmount, uint8 token) returns()
func (_Baccarat *BaccaratTransactorSession) Settle(player common.Address, SettleAmount *big.Int, token uint8) (*types.Transaction, error) {
	return _Baccarat.Contract.Settle(&_Baccarat.TransactOpts, player, SettleAmount, token)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Baccarat *BaccaratTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Baccarat.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Baccarat *BaccaratSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Baccarat.Contract.TransferOwnership(&_Baccarat.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Baccarat *BaccaratTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Baccarat.Contract.TransferOwnership(&_Baccarat.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3f489914.
//
// Solidity: function withdraw(uint8 token, uint256 amount) payable returns()
func (_Baccarat *BaccaratTransactor) Withdraw(opts *bind.TransactOpts, token uint8, amount *big.Int) (*types.Transaction, error) {
	return _Baccarat.contract.Transact(opts, "withdraw", token, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3f489914.
//
// Solidity: function withdraw(uint8 token, uint256 amount) payable returns()
func (_Baccarat *BaccaratSession) Withdraw(token uint8, amount *big.Int) (*types.Transaction, error) {
	return _Baccarat.Contract.Withdraw(&_Baccarat.TransactOpts, token, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3f489914.
//
// Solidity: function withdraw(uint8 token, uint256 amount) payable returns()
func (_Baccarat *BaccaratTransactorSession) Withdraw(token uint8, amount *big.Int) (*types.Transaction, error) {
	return _Baccarat.Contract.Withdraw(&_Baccarat.TransactOpts, token, amount)
}

// WithdrawPrizePool is a paid mutator transaction binding the contract method 0x02f97038.
//
// Solidity: function withdrawPrizePool(uint8 token, uint256 amount) returns()
func (_Baccarat *BaccaratTransactor) WithdrawPrizePool(opts *bind.TransactOpts, token uint8, amount *big.Int) (*types.Transaction, error) {
	return _Baccarat.contract.Transact(opts, "withdrawPrizePool", token, amount)
}

// WithdrawPrizePool is a paid mutator transaction binding the contract method 0x02f97038.
//
// Solidity: function withdrawPrizePool(uint8 token, uint256 amount) returns()
func (_Baccarat *BaccaratSession) WithdrawPrizePool(token uint8, amount *big.Int) (*types.Transaction, error) {
	return _Baccarat.Contract.WithdrawPrizePool(&_Baccarat.TransactOpts, token, amount)
}

// WithdrawPrizePool is a paid mutator transaction binding the contract method 0x02f97038.
//
// Solidity: function withdrawPrizePool(uint8 token, uint256 amount) returns()
func (_Baccarat *BaccaratTransactorSession) WithdrawPrizePool(token uint8, amount *big.Int) (*types.Transaction, error) {
	return _Baccarat.Contract.WithdrawPrizePool(&_Baccarat.TransactOpts, token, amount)
}

// WithdrawPrizePool0 is a paid mutator transaction binding the contract method 0x05e1b115.
//
// Solidity: function withdrawPrizePool(uint256 amount) returns()
func (_Baccarat *BaccaratTransactor) WithdrawPrizePool0(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Baccarat.contract.Transact(opts, "withdrawPrizePool0", amount)
}

// WithdrawPrizePool0 is a paid mutator transaction binding the contract method 0x05e1b115.
//
// Solidity: function withdrawPrizePool(uint256 amount) returns()
func (_Baccarat *BaccaratSession) WithdrawPrizePool0(amount *big.Int) (*types.Transaction, error) {
	return _Baccarat.Contract.WithdrawPrizePool0(&_Baccarat.TransactOpts, amount)
}

// WithdrawPrizePool0 is a paid mutator transaction binding the contract method 0x05e1b115.
//
// Solidity: function withdrawPrizePool(uint256 amount) returns()
func (_Baccarat *BaccaratTransactorSession) WithdrawPrizePool0(amount *big.Int) (*types.Transaction, error) {
	return _Baccarat.Contract.WithdrawPrizePool0(&_Baccarat.TransactOpts, amount)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Baccarat *BaccaratTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Baccarat.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Baccarat *BaccaratSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Baccarat.Contract.Fallback(&_Baccarat.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Baccarat *BaccaratTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Baccarat.Contract.Fallback(&_Baccarat.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Baccarat *BaccaratTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Baccarat.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Baccarat *BaccaratSession) Receive() (*types.Transaction, error) {
	return _Baccarat.Contract.Receive(&_Baccarat.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Baccarat *BaccaratTransactorSession) Receive() (*types.Transaction, error) {
	return _Baccarat.Contract.Receive(&_Baccarat.TransactOpts)
}

// BaccaratDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Baccarat contract.
type BaccaratDepositIterator struct {
	Event *BaccaratDeposit // Event containing the contract specifics and raw log

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
func (it *BaccaratDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaccaratDeposit)
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
		it.Event = new(BaccaratDeposit)
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
func (it *BaccaratDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaccaratDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaccaratDeposit represents a Deposit event raised by the Baccarat contract.
type BaccaratDeposit struct {
	Player  common.Address
	Token   uint8
	Amount  *big.Int
	Balance *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x8d08000e640da164220ee64e90ba424b5a3d8ade3b569979e3bbb27011df9c4b.
//
// Solidity: event Deposit(address indexed player, uint8 indexed token, uint256 amount, int256 balance)
func (_Baccarat *BaccaratFilterer) FilterDeposit(opts *bind.FilterOpts, player []common.Address, token []uint8) (*BaccaratDepositIterator, error) {

	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Baccarat.contract.FilterLogs(opts, "Deposit", playerRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &BaccaratDepositIterator{contract: _Baccarat.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x8d08000e640da164220ee64e90ba424b5a3d8ade3b569979e3bbb27011df9c4b.
//
// Solidity: event Deposit(address indexed player, uint8 indexed token, uint256 amount, int256 balance)
func (_Baccarat *BaccaratFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *BaccaratDeposit, player []common.Address, token []uint8) (event.Subscription, error) {

	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Baccarat.contract.WatchLogs(opts, "Deposit", playerRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaccaratDeposit)
				if err := _Baccarat.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x8d08000e640da164220ee64e90ba424b5a3d8ade3b569979e3bbb27011df9c4b.
//
// Solidity: event Deposit(address indexed player, uint8 indexed token, uint256 amount, int256 balance)
func (_Baccarat *BaccaratFilterer) ParseDeposit(log types.Log) (*BaccaratDeposit, error) {
	event := new(BaccaratDeposit)
	if err := _Baccarat.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaccaratDepositPrizeIterator is returned from FilterDepositPrize and is used to iterate over the raw logs and unpacked data for DepositPrize events raised by the Baccarat contract.
type BaccaratDepositPrizeIterator struct {
	Event *BaccaratDepositPrize // Event containing the contract specifics and raw log

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
func (it *BaccaratDepositPrizeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaccaratDepositPrize)
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
		it.Event = new(BaccaratDepositPrize)
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
func (it *BaccaratDepositPrizeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaccaratDepositPrizeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaccaratDepositPrize represents a DepositPrize event raised by the Baccarat contract.
type BaccaratDepositPrize struct {
	Player  common.Address
	Token   uint8
	Amount  *big.Int
	Balance *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDepositPrize is a free log retrieval operation binding the contract event 0x674005fe57ef38bc71125a69bebece64697aa1391449437cb6adcccb3029d315.
//
// Solidity: event DepositPrize(address indexed player, uint8 indexed token, int256 amount, int256 balance)
func (_Baccarat *BaccaratFilterer) FilterDepositPrize(opts *bind.FilterOpts, player []common.Address, token []uint8) (*BaccaratDepositPrizeIterator, error) {

	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Baccarat.contract.FilterLogs(opts, "DepositPrize", playerRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &BaccaratDepositPrizeIterator{contract: _Baccarat.contract, event: "DepositPrize", logs: logs, sub: sub}, nil
}

// WatchDepositPrize is a free log subscription operation binding the contract event 0x674005fe57ef38bc71125a69bebece64697aa1391449437cb6adcccb3029d315.
//
// Solidity: event DepositPrize(address indexed player, uint8 indexed token, int256 amount, int256 balance)
func (_Baccarat *BaccaratFilterer) WatchDepositPrize(opts *bind.WatchOpts, sink chan<- *BaccaratDepositPrize, player []common.Address, token []uint8) (event.Subscription, error) {

	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Baccarat.contract.WatchLogs(opts, "DepositPrize", playerRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaccaratDepositPrize)
				if err := _Baccarat.contract.UnpackLog(event, "DepositPrize", log); err != nil {
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

// ParseDepositPrize is a log parse operation binding the contract event 0x674005fe57ef38bc71125a69bebece64697aa1391449437cb6adcccb3029d315.
//
// Solidity: event DepositPrize(address indexed player, uint8 indexed token, int256 amount, int256 balance)
func (_Baccarat *BaccaratFilterer) ParseDepositPrize(log types.Log) (*BaccaratDepositPrize, error) {
	event := new(BaccaratDepositPrize)
	if err := _Baccarat.contract.UnpackLog(event, "DepositPrize", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaccaratSettleIterator is returned from FilterSettle and is used to iterate over the raw logs and unpacked data for Settle events raised by the Baccarat contract.
type BaccaratSettleIterator struct {
	Event *BaccaratSettle // Event containing the contract specifics and raw log

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
func (it *BaccaratSettleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaccaratSettle)
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
		it.Event = new(BaccaratSettle)
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
func (it *BaccaratSettleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaccaratSettleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaccaratSettle represents a Settle event raised by the Baccarat contract.
type BaccaratSettle struct {
	Player  common.Address
	Token   uint8
	Amount  *big.Int
	Balance *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSettle is a free log retrieval operation binding the contract event 0xeef32898a33873f49b102087ddbb024589a505b133d0c8f0fb8a09836d28946b.
//
// Solidity: event Settle(address indexed player, uint8 indexed token, int256 amount, int256 balance)
func (_Baccarat *BaccaratFilterer) FilterSettle(opts *bind.FilterOpts, player []common.Address, token []uint8) (*BaccaratSettleIterator, error) {

	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Baccarat.contract.FilterLogs(opts, "Settle", playerRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &BaccaratSettleIterator{contract: _Baccarat.contract, event: "Settle", logs: logs, sub: sub}, nil
}

// WatchSettle is a free log subscription operation binding the contract event 0xeef32898a33873f49b102087ddbb024589a505b133d0c8f0fb8a09836d28946b.
//
// Solidity: event Settle(address indexed player, uint8 indexed token, int256 amount, int256 balance)
func (_Baccarat *BaccaratFilterer) WatchSettle(opts *bind.WatchOpts, sink chan<- *BaccaratSettle, player []common.Address, token []uint8) (event.Subscription, error) {

	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Baccarat.contract.WatchLogs(opts, "Settle", playerRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaccaratSettle)
				if err := _Baccarat.contract.UnpackLog(event, "Settle", log); err != nil {
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

// ParseSettle is a log parse operation binding the contract event 0xeef32898a33873f49b102087ddbb024589a505b133d0c8f0fb8a09836d28946b.
//
// Solidity: event Settle(address indexed player, uint8 indexed token, int256 amount, int256 balance)
func (_Baccarat *BaccaratFilterer) ParseSettle(log types.Log) (*BaccaratSettle, error) {
	event := new(BaccaratSettle)
	if err := _Baccarat.contract.UnpackLog(event, "Settle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaccaratWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Baccarat contract.
type BaccaratWithdrawIterator struct {
	Event *BaccaratWithdraw // Event containing the contract specifics and raw log

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
func (it *BaccaratWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaccaratWithdraw)
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
		it.Event = new(BaccaratWithdraw)
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
func (it *BaccaratWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaccaratWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaccaratWithdraw represents a Withdraw event raised by the Baccarat contract.
type BaccaratWithdraw struct {
	Player  common.Address
	Token   uint8
	Amount  *big.Int
	Balance *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x4e3faed108cc0bdcf97f4fa492bb3b4853f2cb3e107d07d9f5a6155395e6e8fa.
//
// Solidity: event Withdraw(address indexed player, uint8 indexed token, uint256 amount, int256 balance)
func (_Baccarat *BaccaratFilterer) FilterWithdraw(opts *bind.FilterOpts, player []common.Address, token []uint8) (*BaccaratWithdrawIterator, error) {

	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Baccarat.contract.FilterLogs(opts, "Withdraw", playerRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &BaccaratWithdrawIterator{contract: _Baccarat.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x4e3faed108cc0bdcf97f4fa492bb3b4853f2cb3e107d07d9f5a6155395e6e8fa.
//
// Solidity: event Withdraw(address indexed player, uint8 indexed token, uint256 amount, int256 balance)
func (_Baccarat *BaccaratFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *BaccaratWithdraw, player []common.Address, token []uint8) (event.Subscription, error) {

	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Baccarat.contract.WatchLogs(opts, "Withdraw", playerRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaccaratWithdraw)
				if err := _Baccarat.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x4e3faed108cc0bdcf97f4fa492bb3b4853f2cb3e107d07d9f5a6155395e6e8fa.
//
// Solidity: event Withdraw(address indexed player, uint8 indexed token, uint256 amount, int256 balance)
func (_Baccarat *BaccaratFilterer) ParseWithdraw(log types.Log) (*BaccaratWithdraw, error) {
	event := new(BaccaratWithdraw)
	if err := _Baccarat.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaccaratWithdrawPrizeIterator is returned from FilterWithdrawPrize and is used to iterate over the raw logs and unpacked data for WithdrawPrize events raised by the Baccarat contract.
type BaccaratWithdrawPrizeIterator struct {
	Event *BaccaratWithdrawPrize // Event containing the contract specifics and raw log

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
func (it *BaccaratWithdrawPrizeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaccaratWithdrawPrize)
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
		it.Event = new(BaccaratWithdrawPrize)
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
func (it *BaccaratWithdrawPrizeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaccaratWithdrawPrizeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaccaratWithdrawPrize represents a WithdrawPrize event raised by the Baccarat contract.
type BaccaratWithdrawPrize struct {
	Player  common.Address
	Token   uint8
	Amount  *big.Int
	Balance *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawPrize is a free log retrieval operation binding the contract event 0x1877f475a1659d79f43a1e30e8513520783424feea7b0ac2c31bf54c8bef8cdf.
//
// Solidity: event WithdrawPrize(address indexed player, uint8 indexed token, int256 amount, int256 balance)
func (_Baccarat *BaccaratFilterer) FilterWithdrawPrize(opts *bind.FilterOpts, player []common.Address, token []uint8) (*BaccaratWithdrawPrizeIterator, error) {

	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Baccarat.contract.FilterLogs(opts, "WithdrawPrize", playerRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &BaccaratWithdrawPrizeIterator{contract: _Baccarat.contract, event: "WithdrawPrize", logs: logs, sub: sub}, nil
}

// WatchWithdrawPrize is a free log subscription operation binding the contract event 0x1877f475a1659d79f43a1e30e8513520783424feea7b0ac2c31bf54c8bef8cdf.
//
// Solidity: event WithdrawPrize(address indexed player, uint8 indexed token, int256 amount, int256 balance)
func (_Baccarat *BaccaratFilterer) WatchWithdrawPrize(opts *bind.WatchOpts, sink chan<- *BaccaratWithdrawPrize, player []common.Address, token []uint8) (event.Subscription, error) {

	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Baccarat.contract.WatchLogs(opts, "WithdrawPrize", playerRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaccaratWithdrawPrize)
				if err := _Baccarat.contract.UnpackLog(event, "WithdrawPrize", log); err != nil {
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

// ParseWithdrawPrize is a log parse operation binding the contract event 0x1877f475a1659d79f43a1e30e8513520783424feea7b0ac2c31bf54c8bef8cdf.
//
// Solidity: event WithdrawPrize(address indexed player, uint8 indexed token, int256 amount, int256 balance)
func (_Baccarat *BaccaratFilterer) ParseWithdrawPrize(log types.Log) (*BaccaratWithdrawPrize, error) {
	event := new(BaccaratWithdrawPrize)
	if err := _Baccarat.contract.UnpackLog(event, "WithdrawPrize", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
