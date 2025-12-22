// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package warmstorage

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

// WarmStorageServiceMetaData contains all meta data concerning the WarmStorageService contract.
var WarmStorageServiceMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"serviceProviderRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"viewContractAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pdpVerifierAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\"}]",
}

// WarmStorageServiceABI is the input ABI used to generate the binding from.
// Deprecated: Use WarmStorageServiceMetaData.ABI instead.
var WarmStorageServiceABI = WarmStorageServiceMetaData.ABI

// WarmStorageService is an auto generated Go binding around an Ethereum contract.
type WarmStorageService struct {
	WarmStorageServiceCaller     // Read-only binding to the contract
	WarmStorageServiceTransactor // Write-only binding to the contract
	WarmStorageServiceFilterer   // Log filterer for contract events
}

// WarmStorageServiceCaller is an auto generated read-only Go binding around an Ethereum contract.
type WarmStorageServiceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WarmStorageServiceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WarmStorageServiceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WarmStorageServiceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WarmStorageServiceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WarmStorageServiceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WarmStorageServiceSession struct {
	Contract     *WarmStorageService // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// WarmStorageServiceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WarmStorageServiceCallerSession struct {
	Contract *WarmStorageServiceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// WarmStorageServiceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WarmStorageServiceTransactorSession struct {
	Contract     *WarmStorageServiceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// WarmStorageServiceRaw is an auto generated low-level Go binding around an Ethereum contract.
type WarmStorageServiceRaw struct {
	Contract *WarmStorageService // Generic contract binding to access the raw methods on
}

// WarmStorageServiceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WarmStorageServiceCallerRaw struct {
	Contract *WarmStorageServiceCaller // Generic read-only contract binding to access the raw methods on
}

// WarmStorageServiceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WarmStorageServiceTransactorRaw struct {
	Contract *WarmStorageServiceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWarmStorageService creates a new instance of WarmStorageService, bound to a specific deployed contract.
func NewWarmStorageService(address common.Address, backend bind.ContractBackend) (*WarmStorageService, error) {
	contract, err := bindWarmStorageService(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WarmStorageService{WarmStorageServiceCaller: WarmStorageServiceCaller{contract: contract}, WarmStorageServiceTransactor: WarmStorageServiceTransactor{contract: contract}, WarmStorageServiceFilterer: WarmStorageServiceFilterer{contract: contract}}, nil
}

// NewWarmStorageServiceCaller creates a new read-only instance of WarmStorageService, bound to a specific deployed contract.
func NewWarmStorageServiceCaller(address common.Address, caller bind.ContractCaller) (*WarmStorageServiceCaller, error) {
	contract, err := bindWarmStorageService(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WarmStorageServiceCaller{contract: contract}, nil
}

// NewWarmStorageServiceTransactor creates a new write-only instance of WarmStorageService, bound to a specific deployed contract.
func NewWarmStorageServiceTransactor(address common.Address, transactor bind.ContractTransactor) (*WarmStorageServiceTransactor, error) {
	contract, err := bindWarmStorageService(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WarmStorageServiceTransactor{contract: contract}, nil
}

// NewWarmStorageServiceFilterer creates a new log filterer instance of WarmStorageService, bound to a specific deployed contract.
func NewWarmStorageServiceFilterer(address common.Address, filterer bind.ContractFilterer) (*WarmStorageServiceFilterer, error) {
	contract, err := bindWarmStorageService(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WarmStorageServiceFilterer{contract: contract}, nil
}

// bindWarmStorageService binds a generic wrapper to an already deployed contract.
func bindWarmStorageService(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WarmStorageServiceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WarmStorageService *WarmStorageServiceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WarmStorageService.Contract.WarmStorageServiceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WarmStorageService *WarmStorageServiceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WarmStorageService.Contract.WarmStorageServiceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WarmStorageService *WarmStorageServiceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WarmStorageService.Contract.WarmStorageServiceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WarmStorageService *WarmStorageServiceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WarmStorageService.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WarmStorageService *WarmStorageServiceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WarmStorageService.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WarmStorageService *WarmStorageServiceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WarmStorageService.Contract.contract.Transact(opts, method, params...)
}

// PdpVerifierAddress is a free data retrieval call binding the contract method 0xde4b6b71.
//
// Solidity: function pdpVerifierAddress() view returns(address)
func (_WarmStorageService *WarmStorageServiceCaller) PdpVerifierAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WarmStorageService.contract.Call(opts, &out, "pdpVerifierAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PdpVerifierAddress is a free data retrieval call binding the contract method 0xde4b6b71.
//
// Solidity: function pdpVerifierAddress() view returns(address)
func (_WarmStorageService *WarmStorageServiceSession) PdpVerifierAddress() (common.Address, error) {
	return _WarmStorageService.Contract.PdpVerifierAddress(&_WarmStorageService.CallOpts)
}

// PdpVerifierAddress is a free data retrieval call binding the contract method 0xde4b6b71.
//
// Solidity: function pdpVerifierAddress() view returns(address)
func (_WarmStorageService *WarmStorageServiceCallerSession) PdpVerifierAddress() (common.Address, error) {
	return _WarmStorageService.Contract.PdpVerifierAddress(&_WarmStorageService.CallOpts)
}

// ServiceProviderRegistry is a free data retrieval call binding the contract method 0x05f892ec.
//
// Solidity: function serviceProviderRegistry() view returns(address)
func (_WarmStorageService *WarmStorageServiceCaller) ServiceProviderRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WarmStorageService.contract.Call(opts, &out, "serviceProviderRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ServiceProviderRegistry is a free data retrieval call binding the contract method 0x05f892ec.
//
// Solidity: function serviceProviderRegistry() view returns(address)
func (_WarmStorageService *WarmStorageServiceSession) ServiceProviderRegistry() (common.Address, error) {
	return _WarmStorageService.Contract.ServiceProviderRegistry(&_WarmStorageService.CallOpts)
}

// ServiceProviderRegistry is a free data retrieval call binding the contract method 0x05f892ec.
//
// Solidity: function serviceProviderRegistry() view returns(address)
func (_WarmStorageService *WarmStorageServiceCallerSession) ServiceProviderRegistry() (common.Address, error) {
	return _WarmStorageService.Contract.ServiceProviderRegistry(&_WarmStorageService.CallOpts)
}

// ViewContractAddress is a free data retrieval call binding the contract method 0x7a9ebc15.
//
// Solidity: function viewContractAddress() view returns(address)
func (_WarmStorageService *WarmStorageServiceCaller) ViewContractAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WarmStorageService.contract.Call(opts, &out, "viewContractAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ViewContractAddress is a free data retrieval call binding the contract method 0x7a9ebc15.
//
// Solidity: function viewContractAddress() view returns(address)
func (_WarmStorageService *WarmStorageServiceSession) ViewContractAddress() (common.Address, error) {
	return _WarmStorageService.Contract.ViewContractAddress(&_WarmStorageService.CallOpts)
}

// ViewContractAddress is a free data retrieval call binding the contract method 0x7a9ebc15.
//
// Solidity: function viewContractAddress() view returns(address)
func (_WarmStorageService *WarmStorageServiceCallerSession) ViewContractAddress() (common.Address, error) {
	return _WarmStorageService.Contract.ViewContractAddress(&_WarmStorageService.CallOpts)
}
