// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package web3

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

// DealBatch is an auto generated low-level Go binding around an user-defined struct.
type DealBatch struct {
	BatchId        *big.Int
	CommPcEndIndex *big.Int
	Commpa         [32]byte
	DealId         []byte
	ProviderId     []byte
	MarketActorId  []byte
}

// DealRequest is an auto generated low-level Go binding around an user-defined struct.
type DealRequest struct {
	PieceCid  []byte
	Label     string
	FetchLink string
}

// RequestId is an auto generated low-level Go binding around an user-defined struct.
type RequestId struct {
	RequestId [32]byte
	Valid     bool
}

// Web3MetaData contains all meta data concerning the Web3 contract.
var Web3MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"batchId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"pieceCidLeft\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"pieceCidRight\",\"type\":\"bytes32\"}],\"name\":\"BatchIdLeafNodes\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"pieceCid\",\"type\":\"bytes\"}],\"name\":\"DealProposalCreate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"received\",\"type\":\"string\"}],\"name\":\"ReceivedDataCap\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"aggregator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"batches\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"batch_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"commPcEndIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"commpa\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"deal_id\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"provider_id\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"market_actor_id\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"batchesLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_marketActorId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_dealId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_providerId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"commPcStartIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"commPcEndIndex\",\"type\":\"uint256\"}],\"name\":\"complete\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batchId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_left\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_right\",\"type\":\"bytes32\"}],\"name\":\"computeNode\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"dealRequestIdx\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"dealRequests\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"piece_cid\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"fetchLink\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dealsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_cid\",\"type\":\"bytes\"}],\"name\":\"getAllDeals\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"batch_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"commPcEndIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"commpa\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"deal_id\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"provider_id\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"market_actor_id\",\"type\":\"bytes\"}],\"internalType\":\"structDealBatch\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"piece_cid\",\"type\":\"bytes32\"}],\"name\":\"getCIDFromPieceCid\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getDealByIndex\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"piece_cid\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"fetchLink\",\"type\":\"string\"}],\"internalType\":\"structDealRequest\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_cid\",\"type\":\"bytes\"}],\"name\":\"getPieceCidfromCid\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"cid\",\"type\":\"bytes\"}],\"name\":\"getProposalIdSet\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"}],\"internalType\":\"structRequestId\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"pieceRequests\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"pieceStatus\",\"outputs\":[{\"internalType\":\"enumDealClient.Status\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"pieceToBatchId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newAggregator\",\"type\":\"address\"}],\"name\":\"setAggregator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_cid\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"fetchLink\",\"type\":\"string\"}],\"name\":\"submit\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// Web3ABI is the input ABI used to generate the binding from.
// Deprecated: Use Web3MetaData.ABI instead.
var Web3ABI = Web3MetaData.ABI

// Web3 is an auto generated Go binding around an Ethereum contract.
type Web3 struct {
	Web3Caller     // Read-only binding to the contract
	Web3Transactor // Write-only binding to the contract
	Web3Filterer   // Log filterer for contract events
}

// Web3Caller is an auto generated read-only Go binding around an Ethereum contract.
type Web3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Web3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Web3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Web3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Web3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Web3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Web3Session struct {
	Contract     *Web3             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Web3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Web3CallerSession struct {
	Contract *Web3Caller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// Web3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Web3TransactorSession struct {
	Contract     *Web3Transactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Web3Raw is an auto generated low-level Go binding around an Ethereum contract.
type Web3Raw struct {
	Contract *Web3 // Generic contract binding to access the raw methods on
}

// Web3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Web3CallerRaw struct {
	Contract *Web3Caller // Generic read-only contract binding to access the raw methods on
}

// Web3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Web3TransactorRaw struct {
	Contract *Web3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewWeb3 creates a new instance of Web3, bound to a specific deployed contract.
func NewWeb3(address common.Address, backend bind.ContractBackend) (*Web3, error) {
	contract, err := bindWeb3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Web3{Web3Caller: Web3Caller{contract: contract}, Web3Transactor: Web3Transactor{contract: contract}, Web3Filterer: Web3Filterer{contract: contract}}, nil
}

// NewWeb3Caller creates a new read-only instance of Web3, bound to a specific deployed contract.
func NewWeb3Caller(address common.Address, caller bind.ContractCaller) (*Web3Caller, error) {
	contract, err := bindWeb3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Web3Caller{contract: contract}, nil
}

// NewWeb3Transactor creates a new write-only instance of Web3, bound to a specific deployed contract.
func NewWeb3Transactor(address common.Address, transactor bind.ContractTransactor) (*Web3Transactor, error) {
	contract, err := bindWeb3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Web3Transactor{contract: contract}, nil
}

// NewWeb3Filterer creates a new log filterer instance of Web3, bound to a specific deployed contract.
func NewWeb3Filterer(address common.Address, filterer bind.ContractFilterer) (*Web3Filterer, error) {
	contract, err := bindWeb3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Web3Filterer{contract: contract}, nil
}

// bindWeb3 binds a generic wrapper to an already deployed contract.
func bindWeb3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Web3MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Web3 *Web3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Web3.Contract.Web3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Web3 *Web3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Web3.Contract.Web3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Web3 *Web3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Web3.Contract.Web3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Web3 *Web3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Web3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Web3 *Web3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Web3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Web3 *Web3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Web3.Contract.contract.Transact(opts, method, params...)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_Web3 *Web3Caller) Aggregator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Web3.contract.Call(opts, &out, "aggregator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_Web3 *Web3Session) Aggregator() (common.Address, error) {
	return _Web3.Contract.Aggregator(&_Web3.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_Web3 *Web3CallerSession) Aggregator() (common.Address, error) {
	return _Web3.Contract.Aggregator(&_Web3.CallOpts)
}

// Batches is a free data retrieval call binding the contract method 0xb32c4d8d.
//
// Solidity: function batches(uint256 ) view returns(uint256 batch_id, uint256 commPcEndIndex, bytes32 commpa, bytes deal_id, bytes provider_id, bytes market_actor_id)
func (_Web3 *Web3Caller) Batches(opts *bind.CallOpts, arg0 *big.Int) (struct {
	BatchId        *big.Int
	CommPcEndIndex *big.Int
	Commpa         [32]byte
	DealId         []byte
	ProviderId     []byte
	MarketActorId  []byte
}, error) {
	var out []interface{}
	err := _Web3.contract.Call(opts, &out, "batches", arg0)

	outstruct := new(struct {
		BatchId        *big.Int
		CommPcEndIndex *big.Int
		Commpa         [32]byte
		DealId         []byte
		ProviderId     []byte
		MarketActorId  []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BatchId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.CommPcEndIndex = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Commpa = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	outstruct.DealId = *abi.ConvertType(out[3], new([]byte)).(*[]byte)
	outstruct.ProviderId = *abi.ConvertType(out[4], new([]byte)).(*[]byte)
	outstruct.MarketActorId = *abi.ConvertType(out[5], new([]byte)).(*[]byte)

	return *outstruct, err

}

// Batches is a free data retrieval call binding the contract method 0xb32c4d8d.
//
// Solidity: function batches(uint256 ) view returns(uint256 batch_id, uint256 commPcEndIndex, bytes32 commpa, bytes deal_id, bytes provider_id, bytes market_actor_id)
func (_Web3 *Web3Session) Batches(arg0 *big.Int) (struct {
	BatchId        *big.Int
	CommPcEndIndex *big.Int
	Commpa         [32]byte
	DealId         []byte
	ProviderId     []byte
	MarketActorId  []byte
}, error) {
	return _Web3.Contract.Batches(&_Web3.CallOpts, arg0)
}

// Batches is a free data retrieval call binding the contract method 0xb32c4d8d.
//
// Solidity: function batches(uint256 ) view returns(uint256 batch_id, uint256 commPcEndIndex, bytes32 commpa, bytes deal_id, bytes provider_id, bytes market_actor_id)
func (_Web3 *Web3CallerSession) Batches(arg0 *big.Int) (struct {
	BatchId        *big.Int
	CommPcEndIndex *big.Int
	Commpa         [32]byte
	DealId         []byte
	ProviderId     []byte
	MarketActorId  []byte
}, error) {
	return _Web3.Contract.Batches(&_Web3.CallOpts, arg0)
}

// BatchesLength is a free data retrieval call binding the contract method 0x590ba313.
//
// Solidity: function batchesLength() view returns(uint256)
func (_Web3 *Web3Caller) BatchesLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Web3.contract.Call(opts, &out, "batchesLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BatchesLength is a free data retrieval call binding the contract method 0x590ba313.
//
// Solidity: function batchesLength() view returns(uint256)
func (_Web3 *Web3Session) BatchesLength() (*big.Int, error) {
	return _Web3.Contract.BatchesLength(&_Web3.CallOpts)
}

// BatchesLength is a free data retrieval call binding the contract method 0x590ba313.
//
// Solidity: function batchesLength() view returns(uint256)
func (_Web3 *Web3CallerSession) BatchesLength() (*big.Int, error) {
	return _Web3.Contract.BatchesLength(&_Web3.CallOpts)
}

// ComputeNode is a free data retrieval call binding the contract method 0x1b93f2cb.
//
// Solidity: function computeNode(bytes32 _left, bytes32 _right) pure returns(bytes32)
func (_Web3 *Web3Caller) ComputeNode(opts *bind.CallOpts, _left [32]byte, _right [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Web3.contract.Call(opts, &out, "computeNode", _left, _right)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ComputeNode is a free data retrieval call binding the contract method 0x1b93f2cb.
//
// Solidity: function computeNode(bytes32 _left, bytes32 _right) pure returns(bytes32)
func (_Web3 *Web3Session) ComputeNode(_left [32]byte, _right [32]byte) ([32]byte, error) {
	return _Web3.Contract.ComputeNode(&_Web3.CallOpts, _left, _right)
}

// ComputeNode is a free data retrieval call binding the contract method 0x1b93f2cb.
//
// Solidity: function computeNode(bytes32 _left, bytes32 _right) pure returns(bytes32)
func (_Web3 *Web3CallerSession) ComputeNode(_left [32]byte, _right [32]byte) ([32]byte, error) {
	return _Web3.Contract.ComputeNode(&_Web3.CallOpts, _left, _right)
}

// DealRequestIdx is a free data retrieval call binding the contract method 0x1adf7eb1.
//
// Solidity: function dealRequestIdx(bytes32 ) view returns(uint256 idx, bool valid)
func (_Web3 *Web3Caller) DealRequestIdx(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Idx   *big.Int
	Valid bool
}, error) {
	var out []interface{}
	err := _Web3.contract.Call(opts, &out, "dealRequestIdx", arg0)

	outstruct := new(struct {
		Idx   *big.Int
		Valid bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Idx = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Valid = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// DealRequestIdx is a free data retrieval call binding the contract method 0x1adf7eb1.
//
// Solidity: function dealRequestIdx(bytes32 ) view returns(uint256 idx, bool valid)
func (_Web3 *Web3Session) DealRequestIdx(arg0 [32]byte) (struct {
	Idx   *big.Int
	Valid bool
}, error) {
	return _Web3.Contract.DealRequestIdx(&_Web3.CallOpts, arg0)
}

// DealRequestIdx is a free data retrieval call binding the contract method 0x1adf7eb1.
//
// Solidity: function dealRequestIdx(bytes32 ) view returns(uint256 idx, bool valid)
func (_Web3 *Web3CallerSession) DealRequestIdx(arg0 [32]byte) (struct {
	Idx   *big.Int
	Valid bool
}, error) {
	return _Web3.Contract.DealRequestIdx(&_Web3.CallOpts, arg0)
}

// DealRequests is a free data retrieval call binding the contract method 0xbad79de9.
//
// Solidity: function dealRequests(uint256 ) view returns(bytes piece_cid, string label, string fetchLink)
func (_Web3 *Web3Caller) DealRequests(opts *bind.CallOpts, arg0 *big.Int) (struct {
	PieceCid  []byte
	Label     string
	FetchLink string
}, error) {
	var out []interface{}
	err := _Web3.contract.Call(opts, &out, "dealRequests", arg0)

	outstruct := new(struct {
		PieceCid  []byte
		Label     string
		FetchLink string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PieceCid = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.Label = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.FetchLink = *abi.ConvertType(out[2], new(string)).(*string)

	return *outstruct, err

}

// DealRequests is a free data retrieval call binding the contract method 0xbad79de9.
//
// Solidity: function dealRequests(uint256 ) view returns(bytes piece_cid, string label, string fetchLink)
func (_Web3 *Web3Session) DealRequests(arg0 *big.Int) (struct {
	PieceCid  []byte
	Label     string
	FetchLink string
}, error) {
	return _Web3.Contract.DealRequests(&_Web3.CallOpts, arg0)
}

// DealRequests is a free data retrieval call binding the contract method 0xbad79de9.
//
// Solidity: function dealRequests(uint256 ) view returns(bytes piece_cid, string label, string fetchLink)
func (_Web3 *Web3CallerSession) DealRequests(arg0 *big.Int) (struct {
	PieceCid  []byte
	Label     string
	FetchLink string
}, error) {
	return _Web3.Contract.DealRequests(&_Web3.CallOpts, arg0)
}

// DealsLength is a free data retrieval call binding the contract method 0x297ab486.
//
// Solidity: function dealsLength() view returns(uint256)
func (_Web3 *Web3Caller) DealsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Web3.contract.Call(opts, &out, "dealsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DealsLength is a free data retrieval call binding the contract method 0x297ab486.
//
// Solidity: function dealsLength() view returns(uint256)
func (_Web3 *Web3Session) DealsLength() (*big.Int, error) {
	return _Web3.Contract.DealsLength(&_Web3.CallOpts)
}

// DealsLength is a free data retrieval call binding the contract method 0x297ab486.
//
// Solidity: function dealsLength() view returns(uint256)
func (_Web3 *Web3CallerSession) DealsLength() (*big.Int, error) {
	return _Web3.Contract.DealsLength(&_Web3.CallOpts)
}

// GetAllDeals is a free data retrieval call binding the contract method 0x496a9d21.
//
// Solidity: function getAllDeals(bytes _cid) view returns((uint256,uint256,bytes32,bytes,bytes,bytes))
func (_Web3 *Web3Caller) GetAllDeals(opts *bind.CallOpts, _cid []byte) (DealBatch, error) {
	var out []interface{}
	err := _Web3.contract.Call(opts, &out, "getAllDeals", _cid)

	if err != nil {
		return *new(DealBatch), err
	}

	out0 := *abi.ConvertType(out[0], new(DealBatch)).(*DealBatch)

	return out0, err

}

// GetAllDeals is a free data retrieval call binding the contract method 0x496a9d21.
//
// Solidity: function getAllDeals(bytes _cid) view returns((uint256,uint256,bytes32,bytes,bytes,bytes))
func (_Web3 *Web3Session) GetAllDeals(_cid []byte) (DealBatch, error) {
	return _Web3.Contract.GetAllDeals(&_Web3.CallOpts, _cid)
}

// GetAllDeals is a free data retrieval call binding the contract method 0x496a9d21.
//
// Solidity: function getAllDeals(bytes _cid) view returns((uint256,uint256,bytes32,bytes,bytes,bytes))
func (_Web3 *Web3CallerSession) GetAllDeals(_cid []byte) (DealBatch, error) {
	return _Web3.Contract.GetAllDeals(&_Web3.CallOpts, _cid)
}

// GetCIDFromPieceCid is a free data retrieval call binding the contract method 0x12a8ad56.
//
// Solidity: function getCIDFromPieceCid(bytes32 piece_cid) pure returns(bytes)
func (_Web3 *Web3Caller) GetCIDFromPieceCid(opts *bind.CallOpts, piece_cid [32]byte) ([]byte, error) {
	var out []interface{}
	err := _Web3.contract.Call(opts, &out, "getCIDFromPieceCid", piece_cid)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetCIDFromPieceCid is a free data retrieval call binding the contract method 0x12a8ad56.
//
// Solidity: function getCIDFromPieceCid(bytes32 piece_cid) pure returns(bytes)
func (_Web3 *Web3Session) GetCIDFromPieceCid(piece_cid [32]byte) ([]byte, error) {
	return _Web3.Contract.GetCIDFromPieceCid(&_Web3.CallOpts, piece_cid)
}

// GetCIDFromPieceCid is a free data retrieval call binding the contract method 0x12a8ad56.
//
// Solidity: function getCIDFromPieceCid(bytes32 piece_cid) pure returns(bytes)
func (_Web3 *Web3CallerSession) GetCIDFromPieceCid(piece_cid [32]byte) ([]byte, error) {
	return _Web3.Contract.GetCIDFromPieceCid(&_Web3.CallOpts, piece_cid)
}

// GetDealByIndex is a free data retrieval call binding the contract method 0x96925ae6.
//
// Solidity: function getDealByIndex(uint256 index) view returns((bytes,string,string))
func (_Web3 *Web3Caller) GetDealByIndex(opts *bind.CallOpts, index *big.Int) (DealRequest, error) {
	var out []interface{}
	err := _Web3.contract.Call(opts, &out, "getDealByIndex", index)

	if err != nil {
		return *new(DealRequest), err
	}

	out0 := *abi.ConvertType(out[0], new(DealRequest)).(*DealRequest)

	return out0, err

}

// GetDealByIndex is a free data retrieval call binding the contract method 0x96925ae6.
//
// Solidity: function getDealByIndex(uint256 index) view returns((bytes,string,string))
func (_Web3 *Web3Session) GetDealByIndex(index *big.Int) (DealRequest, error) {
	return _Web3.Contract.GetDealByIndex(&_Web3.CallOpts, index)
}

// GetDealByIndex is a free data retrieval call binding the contract method 0x96925ae6.
//
// Solidity: function getDealByIndex(uint256 index) view returns((bytes,string,string))
func (_Web3 *Web3CallerSession) GetDealByIndex(index *big.Int) (DealRequest, error) {
	return _Web3.Contract.GetDealByIndex(&_Web3.CallOpts, index)
}

// GetPieceCidfromCid is a free data retrieval call binding the contract method 0xaa746068.
//
// Solidity: function getPieceCidfromCid(bytes _cid) pure returns(bytes32)
func (_Web3 *Web3Caller) GetPieceCidfromCid(opts *bind.CallOpts, _cid []byte) ([32]byte, error) {
	var out []interface{}
	err := _Web3.contract.Call(opts, &out, "getPieceCidfromCid", _cid)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetPieceCidfromCid is a free data retrieval call binding the contract method 0xaa746068.
//
// Solidity: function getPieceCidfromCid(bytes _cid) pure returns(bytes32)
func (_Web3 *Web3Session) GetPieceCidfromCid(_cid []byte) ([32]byte, error) {
	return _Web3.Contract.GetPieceCidfromCid(&_Web3.CallOpts, _cid)
}

// GetPieceCidfromCid is a free data retrieval call binding the contract method 0xaa746068.
//
// Solidity: function getPieceCidfromCid(bytes _cid) pure returns(bytes32)
func (_Web3 *Web3CallerSession) GetPieceCidfromCid(_cid []byte) ([32]byte, error) {
	return _Web3.Contract.GetPieceCidfromCid(&_Web3.CallOpts, _cid)
}

// GetProposalIdSet is a free data retrieval call binding the contract method 0xf44a8903.
//
// Solidity: function getProposalIdSet(bytes cid) view returns((bytes32,bool))
func (_Web3 *Web3Caller) GetProposalIdSet(opts *bind.CallOpts, cid []byte) (RequestId, error) {
	var out []interface{}
	err := _Web3.contract.Call(opts, &out, "getProposalIdSet", cid)

	if err != nil {
		return *new(RequestId), err
	}

	out0 := *abi.ConvertType(out[0], new(RequestId)).(*RequestId)

	return out0, err

}

// GetProposalIdSet is a free data retrieval call binding the contract method 0xf44a8903.
//
// Solidity: function getProposalIdSet(bytes cid) view returns((bytes32,bool))
func (_Web3 *Web3Session) GetProposalIdSet(cid []byte) (RequestId, error) {
	return _Web3.Contract.GetProposalIdSet(&_Web3.CallOpts, cid)
}

// GetProposalIdSet is a free data retrieval call binding the contract method 0xf44a8903.
//
// Solidity: function getProposalIdSet(bytes cid) view returns((bytes32,bool))
func (_Web3 *Web3CallerSession) GetProposalIdSet(cid []byte) (RequestId, error) {
	return _Web3.Contract.GetProposalIdSet(&_Web3.CallOpts, cid)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Web3 *Web3Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Web3.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Web3 *Web3Session) Owner() (common.Address, error) {
	return _Web3.Contract.Owner(&_Web3.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Web3 *Web3CallerSession) Owner() (common.Address, error) {
	return _Web3.Contract.Owner(&_Web3.CallOpts)
}

// PieceRequests is a free data retrieval call binding the contract method 0x4e8a3ca4.
//
// Solidity: function pieceRequests(bytes ) view returns(bytes32 requestId, bool valid)
func (_Web3 *Web3Caller) PieceRequests(opts *bind.CallOpts, arg0 []byte) (struct {
	RequestId [32]byte
	Valid     bool
}, error) {
	var out []interface{}
	err := _Web3.contract.Call(opts, &out, "pieceRequests", arg0)

	outstruct := new(struct {
		RequestId [32]byte
		Valid     bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RequestId = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Valid = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// PieceRequests is a free data retrieval call binding the contract method 0x4e8a3ca4.
//
// Solidity: function pieceRequests(bytes ) view returns(bytes32 requestId, bool valid)
func (_Web3 *Web3Session) PieceRequests(arg0 []byte) (struct {
	RequestId [32]byte
	Valid     bool
}, error) {
	return _Web3.Contract.PieceRequests(&_Web3.CallOpts, arg0)
}

// PieceRequests is a free data retrieval call binding the contract method 0x4e8a3ca4.
//
// Solidity: function pieceRequests(bytes ) view returns(bytes32 requestId, bool valid)
func (_Web3 *Web3CallerSession) PieceRequests(arg0 []byte) (struct {
	RequestId [32]byte
	Valid     bool
}, error) {
	return _Web3.Contract.PieceRequests(&_Web3.CallOpts, arg0)
}

// PieceStatus is a free data retrieval call binding the contract method 0xa6d1b7b8.
//
// Solidity: function pieceStatus(bytes ) view returns(uint8)
func (_Web3 *Web3Caller) PieceStatus(opts *bind.CallOpts, arg0 []byte) (uint8, error) {
	var out []interface{}
	err := _Web3.contract.Call(opts, &out, "pieceStatus", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// PieceStatus is a free data retrieval call binding the contract method 0xa6d1b7b8.
//
// Solidity: function pieceStatus(bytes ) view returns(uint8)
func (_Web3 *Web3Session) PieceStatus(arg0 []byte) (uint8, error) {
	return _Web3.Contract.PieceStatus(&_Web3.CallOpts, arg0)
}

// PieceStatus is a free data retrieval call binding the contract method 0xa6d1b7b8.
//
// Solidity: function pieceStatus(bytes ) view returns(uint8)
func (_Web3 *Web3CallerSession) PieceStatus(arg0 []byte) (uint8, error) {
	return _Web3.Contract.PieceStatus(&_Web3.CallOpts, arg0)
}

// PieceToBatchId is a free data retrieval call binding the contract method 0x001822eb.
//
// Solidity: function pieceToBatchId(bytes ) view returns(uint256)
func (_Web3 *Web3Caller) PieceToBatchId(opts *bind.CallOpts, arg0 []byte) (*big.Int, error) {
	var out []interface{}
	err := _Web3.contract.Call(opts, &out, "pieceToBatchId", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PieceToBatchId is a free data retrieval call binding the contract method 0x001822eb.
//
// Solidity: function pieceToBatchId(bytes ) view returns(uint256)
func (_Web3 *Web3Session) PieceToBatchId(arg0 []byte) (*big.Int, error) {
	return _Web3.Contract.PieceToBatchId(&_Web3.CallOpts, arg0)
}

// PieceToBatchId is a free data retrieval call binding the contract method 0x001822eb.
//
// Solidity: function pieceToBatchId(bytes ) view returns(uint256)
func (_Web3 *Web3CallerSession) PieceToBatchId(arg0 []byte) (*big.Int, error) {
	return _Web3.Contract.PieceToBatchId(&_Web3.CallOpts, arg0)
}

// Complete is a paid mutator transaction binding the contract method 0xfff54579.
//
// Solidity: function complete(bytes _marketActorId, bytes _dealId, bytes _providerId, uint256 commPcStartIndex, uint256 commPcEndIndex) returns(bytes32 _root, uint256 _count, uint256 batchId)
func (_Web3 *Web3Transactor) Complete(opts *bind.TransactOpts, _marketActorId []byte, _dealId []byte, _providerId []byte, commPcStartIndex *big.Int, commPcEndIndex *big.Int) (*types.Transaction, error) {
	return _Web3.contract.Transact(opts, "complete", _marketActorId, _dealId, _providerId, commPcStartIndex, commPcEndIndex)
}

// Complete is a paid mutator transaction binding the contract method 0xfff54579.
//
// Solidity: function complete(bytes _marketActorId, bytes _dealId, bytes _providerId, uint256 commPcStartIndex, uint256 commPcEndIndex) returns(bytes32 _root, uint256 _count, uint256 batchId)
func (_Web3 *Web3Session) Complete(_marketActorId []byte, _dealId []byte, _providerId []byte, commPcStartIndex *big.Int, commPcEndIndex *big.Int) (*types.Transaction, error) {
	return _Web3.Contract.Complete(&_Web3.TransactOpts, _marketActorId, _dealId, _providerId, commPcStartIndex, commPcEndIndex)
}

// Complete is a paid mutator transaction binding the contract method 0xfff54579.
//
// Solidity: function complete(bytes _marketActorId, bytes _dealId, bytes _providerId, uint256 commPcStartIndex, uint256 commPcEndIndex) returns(bytes32 _root, uint256 _count, uint256 batchId)
func (_Web3 *Web3TransactorSession) Complete(_marketActorId []byte, _dealId []byte, _providerId []byte, commPcStartIndex *big.Int, commPcEndIndex *big.Int) (*types.Transaction, error) {
	return _Web3.Contract.Complete(&_Web3.TransactOpts, _marketActorId, _dealId, _providerId, commPcStartIndex, commPcEndIndex)
}

// SetAggregator is a paid mutator transaction binding the contract method 0xf9120af6.
//
// Solidity: function setAggregator(address _newAggregator) returns()
func (_Web3 *Web3Transactor) SetAggregator(opts *bind.TransactOpts, _newAggregator common.Address) (*types.Transaction, error) {
	return _Web3.contract.Transact(opts, "setAggregator", _newAggregator)
}

// SetAggregator is a paid mutator transaction binding the contract method 0xf9120af6.
//
// Solidity: function setAggregator(address _newAggregator) returns()
func (_Web3 *Web3Session) SetAggregator(_newAggregator common.Address) (*types.Transaction, error) {
	return _Web3.Contract.SetAggregator(&_Web3.TransactOpts, _newAggregator)
}

// SetAggregator is a paid mutator transaction binding the contract method 0xf9120af6.
//
// Solidity: function setAggregator(address _newAggregator) returns()
func (_Web3 *Web3TransactorSession) SetAggregator(_newAggregator common.Address) (*types.Transaction, error) {
	return _Web3.Contract.SetAggregator(&_Web3.TransactOpts, _newAggregator)
}

// Submit is a paid mutator transaction binding the contract method 0xe1333b85.
//
// Solidity: function submit(bytes _cid, string label, string fetchLink) returns(bytes32)
func (_Web3 *Web3Transactor) Submit(opts *bind.TransactOpts, _cid []byte, label string, fetchLink string) (*types.Transaction, error) {
	return _Web3.contract.Transact(opts, "submit", _cid, label, fetchLink)
}

// Submit is a paid mutator transaction binding the contract method 0xe1333b85.
//
// Solidity: function submit(bytes _cid, string label, string fetchLink) returns(bytes32)
func (_Web3 *Web3Session) Submit(_cid []byte, label string, fetchLink string) (*types.Transaction, error) {
	return _Web3.Contract.Submit(&_Web3.TransactOpts, _cid, label, fetchLink)
}

// Submit is a paid mutator transaction binding the contract method 0xe1333b85.
//
// Solidity: function submit(bytes _cid, string label, string fetchLink) returns(bytes32)
func (_Web3 *Web3TransactorSession) Submit(_cid []byte, label string, fetchLink string) (*types.Transaction, error) {
	return _Web3.Contract.Submit(&_Web3.TransactOpts, _cid, label, fetchLink)
}

// Web3BatchIdLeafNodesIterator is returned from FilterBatchIdLeafNodes and is used to iterate over the raw logs and unpacked data for BatchIdLeafNodes events raised by the Web3 contract.
type Web3BatchIdLeafNodesIterator struct {
	Event *Web3BatchIdLeafNodes // Event containing the contract specifics and raw log

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
func (it *Web3BatchIdLeafNodesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Web3BatchIdLeafNodes)
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
		it.Event = new(Web3BatchIdLeafNodes)
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
func (it *Web3BatchIdLeafNodesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Web3BatchIdLeafNodesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Web3BatchIdLeafNodes represents a BatchIdLeafNodes event raised by the Web3 contract.
type Web3BatchIdLeafNodes struct {
	BatchId       *big.Int
	Node          [32]byte
	PieceCidLeft  [32]byte
	PieceCidRight [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBatchIdLeafNodes is a free log retrieval operation binding the contract event 0x995210ceeab1d853fa4111597adac51f81a63f8efc6ba35d4016b4db15d3f67a.
//
// Solidity: event BatchIdLeafNodes(uint256 indexed batchId, bytes32 indexed node, bytes32 pieceCidLeft, bytes32 pieceCidRight)
func (_Web3 *Web3Filterer) FilterBatchIdLeafNodes(opts *bind.FilterOpts, batchId []*big.Int, node [][32]byte) (*Web3BatchIdLeafNodesIterator, error) {

	var batchIdRule []interface{}
	for _, batchIdItem := range batchId {
		batchIdRule = append(batchIdRule, batchIdItem)
	}
	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _Web3.contract.FilterLogs(opts, "BatchIdLeafNodes", batchIdRule, nodeRule)
	if err != nil {
		return nil, err
	}
	return &Web3BatchIdLeafNodesIterator{contract: _Web3.contract, event: "BatchIdLeafNodes", logs: logs, sub: sub}, nil
}

// WatchBatchIdLeafNodes is a free log subscription operation binding the contract event 0x995210ceeab1d853fa4111597adac51f81a63f8efc6ba35d4016b4db15d3f67a.
//
// Solidity: event BatchIdLeafNodes(uint256 indexed batchId, bytes32 indexed node, bytes32 pieceCidLeft, bytes32 pieceCidRight)
func (_Web3 *Web3Filterer) WatchBatchIdLeafNodes(opts *bind.WatchOpts, sink chan<- *Web3BatchIdLeafNodes, batchId []*big.Int, node [][32]byte) (event.Subscription, error) {

	var batchIdRule []interface{}
	for _, batchIdItem := range batchId {
		batchIdRule = append(batchIdRule, batchIdItem)
	}
	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _Web3.contract.WatchLogs(opts, "BatchIdLeafNodes", batchIdRule, nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Web3BatchIdLeafNodes)
				if err := _Web3.contract.UnpackLog(event, "BatchIdLeafNodes", log); err != nil {
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

// ParseBatchIdLeafNodes is a log parse operation binding the contract event 0x995210ceeab1d853fa4111597adac51f81a63f8efc6ba35d4016b4db15d3f67a.
//
// Solidity: event BatchIdLeafNodes(uint256 indexed batchId, bytes32 indexed node, bytes32 pieceCidLeft, bytes32 pieceCidRight)
func (_Web3 *Web3Filterer) ParseBatchIdLeafNodes(log types.Log) (*Web3BatchIdLeafNodes, error) {
	event := new(Web3BatchIdLeafNodes)
	if err := _Web3.contract.UnpackLog(event, "BatchIdLeafNodes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Web3DealProposalCreateIterator is returned from FilterDealProposalCreate and is used to iterate over the raw logs and unpacked data for DealProposalCreate events raised by the Web3 contract.
type Web3DealProposalCreateIterator struct {
	Event *Web3DealProposalCreate // Event containing the contract specifics and raw log

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
func (it *Web3DealProposalCreateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Web3DealProposalCreate)
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
		it.Event = new(Web3DealProposalCreate)
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
func (it *Web3DealProposalCreateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Web3DealProposalCreateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Web3DealProposalCreate represents a DealProposalCreate event raised by the Web3 contract.
type Web3DealProposalCreate struct {
	Id       [32]byte
	PieceCid []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDealProposalCreate is a free log retrieval operation binding the contract event 0xac465338b6723a390b1d45f3a8026396f860574c565c5ac7a812f63ba0a831bf.
//
// Solidity: event DealProposalCreate(bytes32 indexed id, bytes pieceCid)
func (_Web3 *Web3Filterer) FilterDealProposalCreate(opts *bind.FilterOpts, id [][32]byte) (*Web3DealProposalCreateIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Web3.contract.FilterLogs(opts, "DealProposalCreate", idRule)
	if err != nil {
		return nil, err
	}
	return &Web3DealProposalCreateIterator{contract: _Web3.contract, event: "DealProposalCreate", logs: logs, sub: sub}, nil
}

// WatchDealProposalCreate is a free log subscription operation binding the contract event 0xac465338b6723a390b1d45f3a8026396f860574c565c5ac7a812f63ba0a831bf.
//
// Solidity: event DealProposalCreate(bytes32 indexed id, bytes pieceCid)
func (_Web3 *Web3Filterer) WatchDealProposalCreate(opts *bind.WatchOpts, sink chan<- *Web3DealProposalCreate, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Web3.contract.WatchLogs(opts, "DealProposalCreate", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Web3DealProposalCreate)
				if err := _Web3.contract.UnpackLog(event, "DealProposalCreate", log); err != nil {
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

// ParseDealProposalCreate is a log parse operation binding the contract event 0xac465338b6723a390b1d45f3a8026396f860574c565c5ac7a812f63ba0a831bf.
//
// Solidity: event DealProposalCreate(bytes32 indexed id, bytes pieceCid)
func (_Web3 *Web3Filterer) ParseDealProposalCreate(log types.Log) (*Web3DealProposalCreate, error) {
	event := new(Web3DealProposalCreate)
	if err := _Web3.contract.UnpackLog(event, "DealProposalCreate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Web3ReceivedDataCapIterator is returned from FilterReceivedDataCap and is used to iterate over the raw logs and unpacked data for ReceivedDataCap events raised by the Web3 contract.
type Web3ReceivedDataCapIterator struct {
	Event *Web3ReceivedDataCap // Event containing the contract specifics and raw log

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
func (it *Web3ReceivedDataCapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Web3ReceivedDataCap)
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
		it.Event = new(Web3ReceivedDataCap)
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
func (it *Web3ReceivedDataCapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Web3ReceivedDataCapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Web3ReceivedDataCap represents a ReceivedDataCap event raised by the Web3 contract.
type Web3ReceivedDataCap struct {
	Received string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterReceivedDataCap is a free log retrieval operation binding the contract event 0x10aa319ed8cad9bceb033c0c2788c4ae17469ac844e4c6e2c2e20e74ca8a7be8.
//
// Solidity: event ReceivedDataCap(string received)
func (_Web3 *Web3Filterer) FilterReceivedDataCap(opts *bind.FilterOpts) (*Web3ReceivedDataCapIterator, error) {

	logs, sub, err := _Web3.contract.FilterLogs(opts, "ReceivedDataCap")
	if err != nil {
		return nil, err
	}
	return &Web3ReceivedDataCapIterator{contract: _Web3.contract, event: "ReceivedDataCap", logs: logs, sub: sub}, nil
}

// WatchReceivedDataCap is a free log subscription operation binding the contract event 0x10aa319ed8cad9bceb033c0c2788c4ae17469ac844e4c6e2c2e20e74ca8a7be8.
//
// Solidity: event ReceivedDataCap(string received)
func (_Web3 *Web3Filterer) WatchReceivedDataCap(opts *bind.WatchOpts, sink chan<- *Web3ReceivedDataCap) (event.Subscription, error) {

	logs, sub, err := _Web3.contract.WatchLogs(opts, "ReceivedDataCap")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Web3ReceivedDataCap)
				if err := _Web3.contract.UnpackLog(event, "ReceivedDataCap", log); err != nil {
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

// ParseReceivedDataCap is a log parse operation binding the contract event 0x10aa319ed8cad9bceb033c0c2788c4ae17469ac844e4c6e2c2e20e74ca8a7be8.
//
// Solidity: event ReceivedDataCap(string received)
func (_Web3 *Web3Filterer) ParseReceivedDataCap(log types.Log) (*Web3ReceivedDataCap, error) {
	event := new(Web3ReceivedDataCap)
	if err := _Web3.contract.UnpackLog(event, "ReceivedDataCap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
