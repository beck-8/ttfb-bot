// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package registry

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
	ServiceProvider common.Address
	Payee           common.Address
	Name            string
	Description     string
	IsActive        bool
}

// Struct2 is an auto generated low-level Go binding around an user-defined struct.
type Struct2 struct {
	ProviderId              *big.Int
	ProviderInfo            Struct0
	Product                 Struct1
	ProductCapabilityValues [][]byte
}

// Struct1 is an auto generated low-level Go binding around an user-defined struct.
type Struct1 struct {
	ProductType    uint8
	CapabilityKeys []string
	IsActive       bool
}

// ServiceProviderRegistryMetaData contains all meta data concerning the ServiceProviderRegistry contract.
var ServiceProviderRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"getProviderCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getProviderWithProduct\",\"inputs\":[{\"name\":\"providerId\",\"type\":\"uint256\"},{\"name\":\"productType\",\"type\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"components\":[{\"name\":\"providerId\",\"type\":\"uint256\"},{\"name\":\"providerInfo\",\"type\":\"tuple\",\"components\":[{\"name\":\"serviceProvider\",\"type\":\"address\"},{\"name\":\"payee\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"description\",\"type\":\"string\"},{\"name\":\"isActive\",\"type\":\"bool\"}]},{\"name\":\"product\",\"type\":\"tuple\",\"components\":[{\"name\":\"productType\",\"type\":\"uint8\"},{\"name\":\"capabilityKeys\",\"type\":\"string[]\"},{\"name\":\"isActive\",\"type\":\"bool\"}]},{\"name\":\"productCapabilityValues\",\"type\":\"bytes[]\"}]}],\"stateMutability\":\"view\"}]",
}

// ServiceProviderRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use ServiceProviderRegistryMetaData.ABI instead.
var ServiceProviderRegistryABI = ServiceProviderRegistryMetaData.ABI

// ServiceProviderRegistry is an auto generated Go binding around an Ethereum contract.
type ServiceProviderRegistry struct {
	ServiceProviderRegistryCaller     // Read-only binding to the contract
	ServiceProviderRegistryTransactor // Write-only binding to the contract
	ServiceProviderRegistryFilterer   // Log filterer for contract events
}

// ServiceProviderRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ServiceProviderRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ServiceProviderRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ServiceProviderRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ServiceProviderRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ServiceProviderRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ServiceProviderRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ServiceProviderRegistrySession struct {
	Contract     *ServiceProviderRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ServiceProviderRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ServiceProviderRegistryCallerSession struct {
	Contract *ServiceProviderRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// ServiceProviderRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ServiceProviderRegistryTransactorSession struct {
	Contract     *ServiceProviderRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// ServiceProviderRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ServiceProviderRegistryRaw struct {
	Contract *ServiceProviderRegistry // Generic contract binding to access the raw methods on
}

// ServiceProviderRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ServiceProviderRegistryCallerRaw struct {
	Contract *ServiceProviderRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// ServiceProviderRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ServiceProviderRegistryTransactorRaw struct {
	Contract *ServiceProviderRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewServiceProviderRegistry creates a new instance of ServiceProviderRegistry, bound to a specific deployed contract.
func NewServiceProviderRegistry(address common.Address, backend bind.ContractBackend) (*ServiceProviderRegistry, error) {
	contract, err := bindServiceProviderRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ServiceProviderRegistry{ServiceProviderRegistryCaller: ServiceProviderRegistryCaller{contract: contract}, ServiceProviderRegistryTransactor: ServiceProviderRegistryTransactor{contract: contract}, ServiceProviderRegistryFilterer: ServiceProviderRegistryFilterer{contract: contract}}, nil
}

// NewServiceProviderRegistryCaller creates a new read-only instance of ServiceProviderRegistry, bound to a specific deployed contract.
func NewServiceProviderRegistryCaller(address common.Address, caller bind.ContractCaller) (*ServiceProviderRegistryCaller, error) {
	contract, err := bindServiceProviderRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ServiceProviderRegistryCaller{contract: contract}, nil
}

// NewServiceProviderRegistryTransactor creates a new write-only instance of ServiceProviderRegistry, bound to a specific deployed contract.
func NewServiceProviderRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*ServiceProviderRegistryTransactor, error) {
	contract, err := bindServiceProviderRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ServiceProviderRegistryTransactor{contract: contract}, nil
}

// NewServiceProviderRegistryFilterer creates a new log filterer instance of ServiceProviderRegistry, bound to a specific deployed contract.
func NewServiceProviderRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*ServiceProviderRegistryFilterer, error) {
	contract, err := bindServiceProviderRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ServiceProviderRegistryFilterer{contract: contract}, nil
}

// bindServiceProviderRegistry binds a generic wrapper to an already deployed contract.
func bindServiceProviderRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ServiceProviderRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ServiceProviderRegistry *ServiceProviderRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ServiceProviderRegistry.Contract.ServiceProviderRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ServiceProviderRegistry *ServiceProviderRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ServiceProviderRegistry.Contract.ServiceProviderRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ServiceProviderRegistry *ServiceProviderRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ServiceProviderRegistry.Contract.ServiceProviderRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ServiceProviderRegistry *ServiceProviderRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ServiceProviderRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ServiceProviderRegistry *ServiceProviderRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ServiceProviderRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ServiceProviderRegistry *ServiceProviderRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ServiceProviderRegistry.Contract.contract.Transact(opts, method, params...)
}

// GetProviderCount is a free data retrieval call binding the contract method 0x46ce4175.
//
// Solidity: function getProviderCount() view returns(uint256)
func (_ServiceProviderRegistry *ServiceProviderRegistryCaller) GetProviderCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ServiceProviderRegistry.contract.Call(opts, &out, "getProviderCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProviderCount is a free data retrieval call binding the contract method 0x46ce4175.
//
// Solidity: function getProviderCount() view returns(uint256)
func (_ServiceProviderRegistry *ServiceProviderRegistrySession) GetProviderCount() (*big.Int, error) {
	return _ServiceProviderRegistry.Contract.GetProviderCount(&_ServiceProviderRegistry.CallOpts)
}

// GetProviderCount is a free data retrieval call binding the contract method 0x46ce4175.
//
// Solidity: function getProviderCount() view returns(uint256)
func (_ServiceProviderRegistry *ServiceProviderRegistryCallerSession) GetProviderCount() (*big.Int, error) {
	return _ServiceProviderRegistry.Contract.GetProviderCount(&_ServiceProviderRegistry.CallOpts)
}

// GetProviderWithProduct is a free data retrieval call binding the contract method 0xadd33358.
//
// Solidity: function getProviderWithProduct(uint256 providerId, uint8 productType) view returns((uint256,(address,address,string,string,bool),(uint8,string[],bool),bytes[]))
func (_ServiceProviderRegistry *ServiceProviderRegistryCaller) GetProviderWithProduct(opts *bind.CallOpts, providerId *big.Int, productType uint8) (Struct2, error) {
	var out []interface{}
	err := _ServiceProviderRegistry.contract.Call(opts, &out, "getProviderWithProduct", providerId, productType)

	if err != nil {
		return *new(Struct2), err
	}

	out0 := *abi.ConvertType(out[0], new(Struct2)).(*Struct2)

	return out0, err

}

// GetProviderWithProduct is a free data retrieval call binding the contract method 0xadd33358.
//
// Solidity: function getProviderWithProduct(uint256 providerId, uint8 productType) view returns((uint256,(address,address,string,string,bool),(uint8,string[],bool),bytes[]))
func (_ServiceProviderRegistry *ServiceProviderRegistrySession) GetProviderWithProduct(providerId *big.Int, productType uint8) (Struct2, error) {
	return _ServiceProviderRegistry.Contract.GetProviderWithProduct(&_ServiceProviderRegistry.CallOpts, providerId, productType)
}

// GetProviderWithProduct is a free data retrieval call binding the contract method 0xadd33358.
//
// Solidity: function getProviderWithProduct(uint256 providerId, uint8 productType) view returns((uint256,(address,address,string,string,bool),(uint8,string[],bool),bytes[]))
func (_ServiceProviderRegistry *ServiceProviderRegistryCallerSession) GetProviderWithProduct(providerId *big.Int, productType uint8) (Struct2, error) {
	return _ServiceProviderRegistry.Contract.GetProviderWithProduct(&_ServiceProviderRegistry.CallOpts, providerId, productType)
}
