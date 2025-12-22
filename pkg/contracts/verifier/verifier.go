// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package verifier

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
	Data []byte
}

// PDPVerifierMetaData contains all meta data concerning the PDPVerifier contract.
var PDPVerifierMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"getNextDataSetId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNextPieceId\",\"inputs\":[{\"name\":\"setId\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPieceCid\",\"inputs\":[{\"name\":\"setId\",\"type\":\"uint256\"},{\"name\":\"pieceId\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"components\":[{\"name\":\"data\",\"type\":\"bytes\"}]}],\"stateMutability\":\"view\"}]",
}

// PDPVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use PDPVerifierMetaData.ABI instead.
var PDPVerifierABI = PDPVerifierMetaData.ABI

// PDPVerifier is an auto generated Go binding around an Ethereum contract.
type PDPVerifier struct {
	PDPVerifierCaller     // Read-only binding to the contract
	PDPVerifierTransactor // Write-only binding to the contract
	PDPVerifierFilterer   // Log filterer for contract events
}

// PDPVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type PDPVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PDPVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PDPVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PDPVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PDPVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PDPVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PDPVerifierSession struct {
	Contract     *PDPVerifier      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PDPVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PDPVerifierCallerSession struct {
	Contract *PDPVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// PDPVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PDPVerifierTransactorSession struct {
	Contract     *PDPVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// PDPVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type PDPVerifierRaw struct {
	Contract *PDPVerifier // Generic contract binding to access the raw methods on
}

// PDPVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PDPVerifierCallerRaw struct {
	Contract *PDPVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// PDPVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PDPVerifierTransactorRaw struct {
	Contract *PDPVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPDPVerifier creates a new instance of PDPVerifier, bound to a specific deployed contract.
func NewPDPVerifier(address common.Address, backend bind.ContractBackend) (*PDPVerifier, error) {
	contract, err := bindPDPVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PDPVerifier{PDPVerifierCaller: PDPVerifierCaller{contract: contract}, PDPVerifierTransactor: PDPVerifierTransactor{contract: contract}, PDPVerifierFilterer: PDPVerifierFilterer{contract: contract}}, nil
}

// NewPDPVerifierCaller creates a new read-only instance of PDPVerifier, bound to a specific deployed contract.
func NewPDPVerifierCaller(address common.Address, caller bind.ContractCaller) (*PDPVerifierCaller, error) {
	contract, err := bindPDPVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PDPVerifierCaller{contract: contract}, nil
}

// NewPDPVerifierTransactor creates a new write-only instance of PDPVerifier, bound to a specific deployed contract.
func NewPDPVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*PDPVerifierTransactor, error) {
	contract, err := bindPDPVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PDPVerifierTransactor{contract: contract}, nil
}

// NewPDPVerifierFilterer creates a new log filterer instance of PDPVerifier, bound to a specific deployed contract.
func NewPDPVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*PDPVerifierFilterer, error) {
	contract, err := bindPDPVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PDPVerifierFilterer{contract: contract}, nil
}

// bindPDPVerifier binds a generic wrapper to an already deployed contract.
func bindPDPVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PDPVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PDPVerifier *PDPVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PDPVerifier.Contract.PDPVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PDPVerifier *PDPVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PDPVerifier.Contract.PDPVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PDPVerifier *PDPVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PDPVerifier.Contract.PDPVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PDPVerifier *PDPVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PDPVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PDPVerifier *PDPVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PDPVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PDPVerifier *PDPVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PDPVerifier.Contract.contract.Transact(opts, method, params...)
}

// GetNextDataSetId is a free data retrieval call binding the contract method 0x442cded3.
//
// Solidity: function getNextDataSetId() view returns(uint64)
func (_PDPVerifier *PDPVerifierCaller) GetNextDataSetId(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _PDPVerifier.contract.Call(opts, &out, "getNextDataSetId")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetNextDataSetId is a free data retrieval call binding the contract method 0x442cded3.
//
// Solidity: function getNextDataSetId() view returns(uint64)
func (_PDPVerifier *PDPVerifierSession) GetNextDataSetId() (uint64, error) {
	return _PDPVerifier.Contract.GetNextDataSetId(&_PDPVerifier.CallOpts)
}

// GetNextDataSetId is a free data retrieval call binding the contract method 0x442cded3.
//
// Solidity: function getNextDataSetId() view returns(uint64)
func (_PDPVerifier *PDPVerifierCallerSession) GetNextDataSetId() (uint64, error) {
	return _PDPVerifier.Contract.GetNextDataSetId(&_PDPVerifier.CallOpts)
}

// GetNextPieceId is a free data retrieval call binding the contract method 0x1c5ae80f.
//
// Solidity: function getNextPieceId(uint256 setId) view returns(uint256)
func (_PDPVerifier *PDPVerifierCaller) GetNextPieceId(opts *bind.CallOpts, setId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PDPVerifier.contract.Call(opts, &out, "getNextPieceId", setId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextPieceId is a free data retrieval call binding the contract method 0x1c5ae80f.
//
// Solidity: function getNextPieceId(uint256 setId) view returns(uint256)
func (_PDPVerifier *PDPVerifierSession) GetNextPieceId(setId *big.Int) (*big.Int, error) {
	return _PDPVerifier.Contract.GetNextPieceId(&_PDPVerifier.CallOpts, setId)
}

// GetNextPieceId is a free data retrieval call binding the contract method 0x1c5ae80f.
//
// Solidity: function getNextPieceId(uint256 setId) view returns(uint256)
func (_PDPVerifier *PDPVerifierCallerSession) GetNextPieceId(setId *big.Int) (*big.Int, error) {
	return _PDPVerifier.Contract.GetNextPieceId(&_PDPVerifier.CallOpts, setId)
}

// GetPieceCid is a free data retrieval call binding the contract method 0x25bbbedf.
//
// Solidity: function getPieceCid(uint256 setId, uint256 pieceId) view returns((bytes))
func (_PDPVerifier *PDPVerifierCaller) GetPieceCid(opts *bind.CallOpts, setId *big.Int, pieceId *big.Int) (Struct0, error) {
	var out []interface{}
	err := _PDPVerifier.contract.Call(opts, &out, "getPieceCid", setId, pieceId)

	if err != nil {
		return *new(Struct0), err
	}

	out0 := *abi.ConvertType(out[0], new(Struct0)).(*Struct0)

	return out0, err

}

// GetPieceCid is a free data retrieval call binding the contract method 0x25bbbedf.
//
// Solidity: function getPieceCid(uint256 setId, uint256 pieceId) view returns((bytes))
func (_PDPVerifier *PDPVerifierSession) GetPieceCid(setId *big.Int, pieceId *big.Int) (Struct0, error) {
	return _PDPVerifier.Contract.GetPieceCid(&_PDPVerifier.CallOpts, setId, pieceId)
}

// GetPieceCid is a free data retrieval call binding the contract method 0x25bbbedf.
//
// Solidity: function getPieceCid(uint256 setId, uint256 pieceId) view returns((bytes))
func (_PDPVerifier *PDPVerifierCallerSession) GetPieceCid(setId *big.Int, pieceId *big.Int) (Struct0, error) {
	return _PDPVerifier.Contract.GetPieceCid(&_PDPVerifier.CallOpts, setId, pieceId)
}
