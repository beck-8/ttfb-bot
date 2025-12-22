// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package view

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

// Struct0 is an auto generated low-level Go binding around an user-defined struct.
type Struct0 struct {
	PdpRailId       *big.Int
	CacheMissRailId *big.Int
	CdnRailId       *big.Int
	Payer           common.Address
	Payee           common.Address
	ServiceProvider common.Address
	CommissionBps   *big.Int
	ClientDataSetId *big.Int
	PdpEndEpoch     *big.Int
	ProviderId      *big.Int
	DataSetId       *big.Int
}

// WarmStorageViewMetaData contains all meta data concerning the WarmStorageView contract.
var WarmStorageViewMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"getDataSet\",\"inputs\":[{\"name\":\"dataSetId\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"info\",\"type\":\"tuple\",\"components\":[{\"name\":\"pdpRailId\",\"type\":\"uint256\"},{\"name\":\"cacheMissRailId\",\"type\":\"uint256\"},{\"name\":\"cdnRailId\",\"type\":\"uint256\"},{\"name\":\"payer\",\"type\":\"address\"},{\"name\":\"payee\",\"type\":\"address\"},{\"name\":\"serviceProvider\",\"type\":\"address\"},{\"name\":\"commissionBps\",\"type\":\"uint256\"},{\"name\":\"clientDataSetId\",\"type\":\"uint256\"},{\"name\":\"pdpEndEpoch\",\"type\":\"uint256\"},{\"name\":\"providerId\",\"type\":\"uint256\"},{\"name\":\"dataSetId\",\"type\":\"uint256\"}]}],\"stateMutability\":\"view\"}]",
}

// WarmStorageViewABI is the input ABI used to generate the binding from.
// Deprecated: Use WarmStorageViewMetaData.ABI instead.
var WarmStorageViewABI = WarmStorageViewMetaData.ABI

// WarmStorageView is an auto generated Go binding around an Ethereum contract.
type WarmStorageView struct {
	WarmStorageViewCaller     // Read-only binding to the contract
	WarmStorageViewTransactor // Write-only binding to the contract
	WarmStorageViewFilterer   // Log filterer for contract events
}

// WarmStorageViewCaller is an auto generated read-only Go binding around an Ethereum contract.
type WarmStorageViewCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WarmStorageViewTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WarmStorageViewTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WarmStorageViewFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WarmStorageViewFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WarmStorageViewSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WarmStorageViewSession struct {
	Contract     *WarmStorageView  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WarmStorageViewCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WarmStorageViewCallerSession struct {
	Contract *WarmStorageViewCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// WarmStorageViewTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WarmStorageViewTransactorSession struct {
	Contract     *WarmStorageViewTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// WarmStorageViewRaw is an auto generated low-level Go binding around an Ethereum contract.
type WarmStorageViewRaw struct {
	Contract *WarmStorageView // Generic contract binding to access the raw methods on
}

// WarmStorageViewCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WarmStorageViewCallerRaw struct {
	Contract *WarmStorageViewCaller // Generic read-only contract binding to access the raw methods on
}

// WarmStorageViewTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WarmStorageViewTransactorRaw struct {
	Contract *WarmStorageViewTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWarmStorageView creates a new instance of WarmStorageView, bound to a specific deployed contract.
func NewWarmStorageView(address common.Address, backend bind.ContractBackend) (*WarmStorageView, error) {
	contract, err := bindWarmStorageView(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WarmStorageView{WarmStorageViewCaller: WarmStorageViewCaller{contract: contract}, WarmStorageViewTransactor: WarmStorageViewTransactor{contract: contract}, WarmStorageViewFilterer: WarmStorageViewFilterer{contract: contract}}, nil
}

// NewWarmStorageViewCaller creates a new read-only instance of WarmStorageView, bound to a specific deployed contract.
func NewWarmStorageViewCaller(address common.Address, caller bind.ContractCaller) (*WarmStorageViewCaller, error) {
	contract, err := bindWarmStorageView(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WarmStorageViewCaller{contract: contract}, nil
}

// NewWarmStorageViewTransactor creates a new write-only instance of WarmStorageView, bound to a specific deployed contract.
func NewWarmStorageViewTransactor(address common.Address, transactor bind.ContractTransactor) (*WarmStorageViewTransactor, error) {
	contract, err := bindWarmStorageView(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WarmStorageViewTransactor{contract: contract}, nil
}

// NewWarmStorageViewFilterer creates a new log filterer instance of WarmStorageView, bound to a specific deployed contract.
func NewWarmStorageViewFilterer(address common.Address, filterer bind.ContractFilterer) (*WarmStorageViewFilterer, error) {
	contract, err := bindWarmStorageView(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WarmStorageViewFilterer{contract: contract}, nil
}

// bindWarmStorageView binds a generic wrapper to an already deployed contract.
func bindWarmStorageView(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WarmStorageViewMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WarmStorageView *WarmStorageViewRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WarmStorageView.Contract.WarmStorageViewCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WarmStorageView *WarmStorageViewRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WarmStorageView.Contract.WarmStorageViewTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WarmStorageView *WarmStorageViewRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WarmStorageView.Contract.WarmStorageViewTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WarmStorageView *WarmStorageViewCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WarmStorageView.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WarmStorageView *WarmStorageViewTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WarmStorageView.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WarmStorageView *WarmStorageViewTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WarmStorageView.Contract.contract.Transact(opts, method, params...)
}

// GetDataSet is a free data retrieval call binding the contract method 0xbdaac056.
//
// Solidity: function getDataSet(uint256 dataSetId) view returns((uint256,uint256,uint256,address,address,address,uint256,uint256,uint256,uint256,uint256) info)
func (_WarmStorageView *WarmStorageViewCaller) GetDataSet(opts *bind.CallOpts, dataSetId *big.Int) (Struct0, error) {
	var out []interface{}
	err := _WarmStorageView.contract.Call(opts, &out, "getDataSet", dataSetId)

	if err != nil {
		return *new(Struct0), err
	}

	out0 := *abi.ConvertType(out[0], new(Struct0)).(*Struct0)

	return out0, err

}

// GetDataSet is a free data retrieval call binding the contract method 0xbdaac056.
//
// Solidity: function getDataSet(uint256 dataSetId) view returns((uint256,uint256,uint256,address,address,address,uint256,uint256,uint256,uint256,uint256) info)
func (_WarmStorageView *WarmStorageViewSession) GetDataSet(dataSetId *big.Int) (Struct0, error) {
	return _WarmStorageView.Contract.GetDataSet(&_WarmStorageView.CallOpts, dataSetId)
}

// GetDataSet is a free data retrieval call binding the contract method 0xbdaac056.
//
// Solidity: function getDataSet(uint256 dataSetId) view returns((uint256,uint256,uint256,address,address,address,uint256,uint256,uint256,uint256,uint256) info)
func (_WarmStorageView *WarmStorageViewCallerSession) GetDataSet(dataSetId *big.Int) (Struct0, error) {
	return _WarmStorageView.Contract.GetDataSet(&_WarmStorageView.CallOpts, dataSetId)
}
