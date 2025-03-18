// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package wrappers

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
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// HyperionABI is the input ABI used to generate the binding from.
const HyperionABI = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hyperionId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_powerThreshold\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_validators\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_tokenContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_destination\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_eventNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"}],\"name\":\"SendToCosmosEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_batchNonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_eventNonce\",\"type\":\"uint256\"}],\"name\":\"TransactionBatchExecutedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_newValsetNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"_validators\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"ValsetUpdatedEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_erc20Address\",\"type\":\"address\"}],\"name\":\"lastBatchNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"seenTokens\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"sendToCosmos\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"state_lastBatchNonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state_lastEventNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state_lastValsetCheckpoint\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state_lastValsetNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state_hyperionId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state_powerThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_currentValidators\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_currentPowers\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_currentValsetNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint8[]\",\"name\":\"_v\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_r\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_s\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"_destinations\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_fees\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_batchNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_tokenContract\",\"type\":\"address\"}],\"name\":\"submitBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_currentValidators\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_currentPowers\",\"type\":\"uint256[]\"},{\"internalType\":\"uint8[]\",\"name\":\"_v\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_r\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_s\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"_theHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_powerThreshold\",\"type\":\"uint256\"}],\"name\":\"testCheckValidatorSignatures\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_validators\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_valsetNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_hyperionId\",\"type\":\"bytes32\"}],\"name\":\"testMakeCheckpoint\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenDecimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenNames\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenSymbols\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_newValidators\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_newPowers\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_newValsetNonce\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_currentValidators\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_currentPowers\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_currentValsetNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint8[]\",\"name\":\"_v\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_r\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_s\",\"type\":\"bytes32[]\"}],\"name\":\"updateValset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Hyperion is an auto generated Go binding around an Ethereum contract.
type Hyperion struct {
	HyperionCaller     // Read-only binding to the contract
	HyperionTransactor // Write-only binding to the contract
	HyperionFilterer   // Log filterer for contract events
}

// HyperionCaller is an auto generated read-only Go binding around an Ethereum contract.
type HyperionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HyperionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HyperionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HyperionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HyperionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HyperionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HyperionSession struct {
	Contract     *Hyperion            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HyperionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HyperionCallerSession struct {
	Contract *HyperionCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// HyperionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HyperionTransactorSession struct {
	Contract     *HyperionTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HyperionRaw is an auto generated low-level Go binding around an Ethereum contract.
type HyperionRaw struct {
	Contract *Hyperion // Generic contract binding to access the raw methods on
}

// HyperionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HyperionCallerRaw struct {
	Contract *HyperionCaller // Generic read-only contract binding to access the raw methods on
}

// HyperionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HyperionTransactorRaw struct {
	Contract *HyperionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHyperion creates a new instance of Hyperion, bound to a specific deployed contract.
func NewHyperion(address common.Address, backend bind.ContractBackend) (*Hyperion, error) {
	contract, err := bindHyperion(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Hyperion{HyperionCaller: HyperionCaller{contract: contract}, HyperionTransactor: HyperionTransactor{contract: contract}, HyperionFilterer: HyperionFilterer{contract: contract}}, nil
}

// NewHyperionCaller creates a new read-only instance of Hyperion, bound to a specific deployed contract.
func NewHyperionCaller(address common.Address, caller bind.ContractCaller) (*HyperionCaller, error) {
	contract, err := bindHyperion(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HyperionCaller{contract: contract}, nil
}

// NewHyperionTransactor creates a new write-only instance of Hyperion, bound to a specific deployed contract.
func NewHyperionTransactor(address common.Address, transactor bind.ContractTransactor) (*HyperionTransactor, error) {
	contract, err := bindHyperion(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HyperionTransactor{contract: contract}, nil
}

// NewHyperionFilterer creates a new log filterer instance of Hyperion, bound to a specific deployed contract.
func NewHyperionFilterer(address common.Address, filterer bind.ContractFilterer) (*HyperionFilterer, error) {
	contract, err := bindHyperion(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HyperionFilterer{contract: contract}, nil
}

// bindHyperion binds a generic wrapper to an already deployed contract.
func bindHyperion(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HyperionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hyperion *HyperionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Hyperion.Contract.HyperionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hyperion *HyperionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hyperion.Contract.HyperionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hyperion *HyperionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hyperion.Contract.HyperionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hyperion *HyperionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Hyperion.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hyperion *HyperionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hyperion.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hyperion *HyperionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hyperion.Contract.contract.Transact(opts, method, params...)
}

// LastBatchNonce is a free data retrieval call binding the contract method 0x011b2174.
//
// Solidity: function lastBatchNonce(address _erc20Address) view returns(uint256)
func (_Hyperion *HyperionCaller) LastBatchNonce(opts *bind.CallOpts, _erc20Address common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Hyperion.contract.Call(opts, &out, "lastBatchNonce", _erc20Address)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastBatchNonce is a free data retrieval call binding the contract method 0x011b2174.
//
// Solidity: function lastBatchNonce(address _erc20Address) view returns(uint256)
func (_Hyperion *HyperionSession) LastBatchNonce(_erc20Address common.Address) (*big.Int, error) {
	return _Hyperion.Contract.LastBatchNonce(&_Hyperion.CallOpts, _erc20Address)
}

// LastBatchNonce is a free data retrieval call binding the contract method 0x011b2174.
//
// Solidity: function lastBatchNonce(address _erc20Address) view returns(uint256)
func (_Hyperion *HyperionCallerSession) LastBatchNonce(_erc20Address common.Address) (*big.Int, error) {
	return _Hyperion.Contract.LastBatchNonce(&_Hyperion.CallOpts, _erc20Address)
}

// SeenTokens is a free data retrieval call binding the contract method 0x13100091.
//
// Solidity: function seenTokens(address ) view returns(bool)
func (_Hyperion *HyperionCaller) SeenTokens(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Hyperion.contract.Call(opts, &out, "seenTokens", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SeenTokens is a free data retrieval call binding the contract method 0x13100091.
//
// Solidity: function seenTokens(address ) view returns(bool)
func (_Hyperion *HyperionSession) SeenTokens(arg0 common.Address) (bool, error) {
	return _Hyperion.Contract.SeenTokens(&_Hyperion.CallOpts, arg0)
}

// SeenTokens is a free data retrieval call binding the contract method 0x13100091.
//
// Solidity: function seenTokens(address ) view returns(bool)
func (_Hyperion *HyperionCallerSession) SeenTokens(arg0 common.Address) (bool, error) {
	return _Hyperion.Contract.SeenTokens(&_Hyperion.CallOpts, arg0)
}

// StateLastBatchNonces is a free data retrieval call binding the contract method 0xdf97174b.
//
// Solidity: function state_lastBatchNonces(address ) view returns(uint256)
func (_Hyperion *HyperionCaller) StateLastBatchNonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Hyperion.contract.Call(opts, &out, "state_lastBatchNonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StateLastBatchNonces is a free data retrieval call binding the contract method 0xdf97174b.
//
// Solidity: function state_lastBatchNonces(address ) view returns(uint256)
func (_Hyperion *HyperionSession) StateLastBatchNonces(arg0 common.Address) (*big.Int, error) {
	return _Hyperion.Contract.StateLastBatchNonces(&_Hyperion.CallOpts, arg0)
}

// StateLastBatchNonces is a free data retrieval call binding the contract method 0xdf97174b.
//
// Solidity: function state_lastBatchNonces(address ) view returns(uint256)
func (_Hyperion *HyperionCallerSession) StateLastBatchNonces(arg0 common.Address) (*big.Int, error) {
	return _Hyperion.Contract.StateLastBatchNonces(&_Hyperion.CallOpts, arg0)
}

// StateLastEventNonce is a free data retrieval call binding the contract method 0x73b20547.
//
// Solidity: function state_lastEventNonce() view returns(uint256)
func (_Hyperion *HyperionCaller) StateLastEventNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hyperion.contract.Call(opts, &out, "state_lastEventNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StateLastEventNonce is a free data retrieval call binding the contract method 0x73b20547.
//
// Solidity: function state_lastEventNonce() view returns(uint256)
func (_Hyperion *HyperionSession) StateLastEventNonce() (*big.Int, error) {
	return _Hyperion.Contract.StateLastEventNonce(&_Hyperion.CallOpts)
}

// StateLastEventNonce is a free data retrieval call binding the contract method 0x73b20547.
//
// Solidity: function state_lastEventNonce() view returns(uint256)
func (_Hyperion *HyperionCallerSession) StateLastEventNonce() (*big.Int, error) {
	return _Hyperion.Contract.StateLastEventNonce(&_Hyperion.CallOpts)
}

// StateLastValsetCheckpoint is a free data retrieval call binding the contract method 0xf2b53307.
//
// Solidity: function state_lastValsetCheckpoint() view returns(bytes32)
func (_Hyperion *HyperionCaller) StateLastValsetCheckpoint(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Hyperion.contract.Call(opts, &out, "state_lastValsetCheckpoint")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// StateLastValsetCheckpoint is a free data retrieval call binding the contract method 0xf2b53307.
//
// Solidity: function state_lastValsetCheckpoint() view returns(bytes32)
func (_Hyperion *HyperionSession) StateLastValsetCheckpoint() ([32]byte, error) {
	return _Hyperion.Contract.StateLastValsetCheckpoint(&_Hyperion.CallOpts)
}

// StateLastValsetCheckpoint is a free data retrieval call binding the contract method 0xf2b53307.
//
// Solidity: function state_lastValsetCheckpoint() view returns(bytes32)
func (_Hyperion *HyperionCallerSession) StateLastValsetCheckpoint() ([32]byte, error) {
	return _Hyperion.Contract.StateLastValsetCheckpoint(&_Hyperion.CallOpts)
}

// StateLastValsetNonce is a free data retrieval call binding the contract method 0xb56561fe.
//
// Solidity: function state_lastValsetNonce() view returns(uint256)
func (_Hyperion *HyperionCaller) StateLastValsetNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hyperion.contract.Call(opts, &out, "state_lastValsetNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StateLastValsetNonce is a free data retrieval call binding the contract method 0xb56561fe.
//
// Solidity: function state_lastValsetNonce() view returns(uint256)
func (_Hyperion *HyperionSession) StateLastValsetNonce() (*big.Int, error) {
	return _Hyperion.Contract.StateLastValsetNonce(&_Hyperion.CallOpts)
}

// StateLastValsetNonce is a free data retrieval call binding the contract method 0xb56561fe.
//
// Solidity: function state_lastValsetNonce() view returns(uint256)
func (_Hyperion *HyperionCallerSession) StateLastValsetNonce() (*big.Int, error) {
	return _Hyperion.Contract.StateLastValsetNonce(&_Hyperion.CallOpts)
}

// StateHyperionId is a free data retrieval call binding the contract method 0x69dd3908.
//
// Solidity: function state_hyperionId() view returns(bytes32)
func (_Hyperion *HyperionCaller) StateHyperionId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Hyperion.contract.Call(opts, &out, "state_hyperionId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// StateHyperionId is a free data retrieval call binding the contract method 0x69dd3908.
//
// Solidity: function state_hyperionId() view returns(bytes32)
func (_Hyperion *HyperionSession) StateHyperionId() ([32]byte, error) {
	return _Hyperion.Contract.StateHyperionId(&_Hyperion.CallOpts)
}

// StateHyperionId is a free data retrieval call binding the contract method 0x69dd3908.
//
// Solidity: function state_hyperionId() view returns(bytes32)
func (_Hyperion *HyperionCallerSession) StateHyperionId() ([32]byte, error) {
	return _Hyperion.Contract.StateHyperionId(&_Hyperion.CallOpts)
}

// StatePowerThreshold is a free data retrieval call binding the contract method 0xe5a2b5d2.
//
// Solidity: function state_powerThreshold() view returns(uint256)
func (_Hyperion *HyperionCaller) StatePowerThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hyperion.contract.Call(opts, &out, "state_powerThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StatePowerThreshold is a free data retrieval call binding the contract method 0xe5a2b5d2.
//
// Solidity: function state_powerThreshold() view returns(uint256)
func (_Hyperion *HyperionSession) StatePowerThreshold() (*big.Int, error) {
	return _Hyperion.Contract.StatePowerThreshold(&_Hyperion.CallOpts)
}

// StatePowerThreshold is a free data retrieval call binding the contract method 0xe5a2b5d2.
//
// Solidity: function state_powerThreshold() view returns(uint256)
func (_Hyperion *HyperionCallerSession) StatePowerThreshold() (*big.Int, error) {
	return _Hyperion.Contract.StatePowerThreshold(&_Hyperion.CallOpts)
}

// TestCheckValidatorSignatures is a free data retrieval call binding the contract method 0xdb7c4e57.
//
// Solidity: function testCheckValidatorSignatures(address[] _currentValidators, uint256[] _currentPowers, uint8[] _v, bytes32[] _r, bytes32[] _s, bytes32 _theHash, uint256 _powerThreshold) pure returns()
func (_Hyperion *HyperionCaller) TestCheckValidatorSignatures(opts *bind.CallOpts, _currentValidators []common.Address, _currentPowers []*big.Int, _v []uint8, _r [][32]byte, _s [][32]byte, _theHash [32]byte, _powerThreshold *big.Int) error {
	var out []interface{}
	err := _Hyperion.contract.Call(opts, &out, "testCheckValidatorSignatures", _currentValidators, _currentPowers, _v, _r, _s, _theHash, _powerThreshold)

	if err != nil {
		return err
	}

	return err

}

// TestCheckValidatorSignatures is a free data retrieval call binding the contract method 0xdb7c4e57.
//
// Solidity: function testCheckValidatorSignatures(address[] _currentValidators, uint256[] _currentPowers, uint8[] _v, bytes32[] _r, bytes32[] _s, bytes32 _theHash, uint256 _powerThreshold) pure returns()
func (_Hyperion *HyperionSession) TestCheckValidatorSignatures(_currentValidators []common.Address, _currentPowers []*big.Int, _v []uint8, _r [][32]byte, _s [][32]byte, _theHash [32]byte, _powerThreshold *big.Int) error {
	return _Hyperion.Contract.TestCheckValidatorSignatures(&_Hyperion.CallOpts, _currentValidators, _currentPowers, _v, _r, _s, _theHash, _powerThreshold)
}

// TestCheckValidatorSignatures is a free data retrieval call binding the contract method 0xdb7c4e57.
//
// Solidity: function testCheckValidatorSignatures(address[] _currentValidators, uint256[] _currentPowers, uint8[] _v, bytes32[] _r, bytes32[] _s, bytes32 _theHash, uint256 _powerThreshold) pure returns()
func (_Hyperion *HyperionCallerSession) TestCheckValidatorSignatures(_currentValidators []common.Address, _currentPowers []*big.Int, _v []uint8, _r [][32]byte, _s [][32]byte, _theHash [32]byte, _powerThreshold *big.Int) error {
	return _Hyperion.Contract.TestCheckValidatorSignatures(&_Hyperion.CallOpts, _currentValidators, _currentPowers, _v, _r, _s, _theHash, _powerThreshold)
}

// TestMakeCheckpoint is a free data retrieval call binding the contract method 0xc227c30b.
//
// Solidity: function testMakeCheckpoint(address[] _validators, uint256[] _powers, uint256 _valsetNonce, bytes32 _hyperionId) pure returns()
func (_Hyperion *HyperionCaller) TestMakeCheckpoint(opts *bind.CallOpts, _validators []common.Address, _powers []*big.Int, _valsetNonce *big.Int, _hyperionId [32]byte) error {
	var out []interface{}
	err := _Hyperion.contract.Call(opts, &out, "testMakeCheckpoint", _validators, _powers, _valsetNonce, _hyperionId)

	if err != nil {
		return err
	}

	return err

}

// TestMakeCheckpoint is a free data retrieval call binding the contract method 0xc227c30b.
//
// Solidity: function testMakeCheckpoint(address[] _validators, uint256[] _powers, uint256 _valsetNonce, bytes32 _hyperionId) pure returns()
func (_Hyperion *HyperionSession) TestMakeCheckpoint(_validators []common.Address, _powers []*big.Int, _valsetNonce *big.Int, _hyperionId [32]byte) error {
	return _Hyperion.Contract.TestMakeCheckpoint(&_Hyperion.CallOpts, _validators, _powers, _valsetNonce, _hyperionId)
}

// TestMakeCheckpoint is a free data retrieval call binding the contract method 0xc227c30b.
//
// Solidity: function testMakeCheckpoint(address[] _validators, uint256[] _powers, uint256 _valsetNonce, bytes32 _hyperionId) pure returns()
func (_Hyperion *HyperionCallerSession) TestMakeCheckpoint(_validators []common.Address, _powers []*big.Int, _valsetNonce *big.Int, _hyperionId [32]byte) error {
	return _Hyperion.Contract.TestMakeCheckpoint(&_Hyperion.CallOpts, _validators, _powers, _valsetNonce, _hyperionId)
}

// TokenDecimals is a free data retrieval call binding the contract method 0x8ee573ac.
//
// Solidity: function tokenDecimals(address ) view returns(uint8)
func (_Hyperion *HyperionCaller) TokenDecimals(opts *bind.CallOpts, arg0 common.Address) (uint8, error) {
	var out []interface{}
	err := _Hyperion.contract.Call(opts, &out, "tokenDecimals", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TokenDecimals is a free data retrieval call binding the contract method 0x8ee573ac.
//
// Solidity: function tokenDecimals(address ) view returns(uint8)
func (_Hyperion *HyperionSession) TokenDecimals(arg0 common.Address) (uint8, error) {
	return _Hyperion.Contract.TokenDecimals(&_Hyperion.CallOpts, arg0)
}

// TokenDecimals is a free data retrieval call binding the contract method 0x8ee573ac.
//
// Solidity: function tokenDecimals(address ) view returns(uint8)
func (_Hyperion *HyperionCallerSession) TokenDecimals(arg0 common.Address) (uint8, error) {
	return _Hyperion.Contract.TokenDecimals(&_Hyperion.CallOpts, arg0)
}

// TokenNames is a free data retrieval call binding the contract method 0xa8fc75e1.
//
// Solidity: function tokenNames(address ) view returns(string)
func (_Hyperion *HyperionCaller) TokenNames(opts *bind.CallOpts, arg0 common.Address) (string, error) {
	var out []interface{}
	err := _Hyperion.contract.Call(opts, &out, "tokenNames", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenNames is a free data retrieval call binding the contract method 0xa8fc75e1.
//
// Solidity: function tokenNames(address ) view returns(string)
func (_Hyperion *HyperionSession) TokenNames(arg0 common.Address) (string, error) {
	return _Hyperion.Contract.TokenNames(&_Hyperion.CallOpts, arg0)
}

// TokenNames is a free data retrieval call binding the contract method 0xa8fc75e1.
//
// Solidity: function tokenNames(address ) view returns(string)
func (_Hyperion *HyperionCallerSession) TokenNames(arg0 common.Address) (string, error) {
	return _Hyperion.Contract.TokenNames(&_Hyperion.CallOpts, arg0)
}

// TokenSymbols is a free data retrieval call binding the contract method 0xfb0b2b36.
//
// Solidity: function tokenSymbols(address ) view returns(string)
func (_Hyperion *HyperionCaller) TokenSymbols(opts *bind.CallOpts, arg0 common.Address) (string, error) {
	var out []interface{}
	err := _Hyperion.contract.Call(opts, &out, "tokenSymbols", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenSymbols is a free data retrieval call binding the contract method 0xfb0b2b36.
//
// Solidity: function tokenSymbols(address ) view returns(string)
func (_Hyperion *HyperionSession) TokenSymbols(arg0 common.Address) (string, error) {
	return _Hyperion.Contract.TokenSymbols(&_Hyperion.CallOpts, arg0)
}

// TokenSymbols is a free data retrieval call binding the contract method 0xfb0b2b36.
//
// Solidity: function tokenSymbols(address ) view returns(string)
func (_Hyperion *HyperionCallerSession) TokenSymbols(arg0 common.Address) (string, error) {
	return _Hyperion.Contract.TokenSymbols(&_Hyperion.CallOpts, arg0)
}

// SendToCosmos is a paid mutator transaction binding the contract method 0x5d2428fa.
//
// Solidity: function sendToCosmos(address _tokenContract, address _destination, uint256 _amount) returns()
func (_Hyperion *HyperionTransactor) SendToCosmos(opts *bind.TransactOpts, _tokenContract common.Address, _destination common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Hyperion.contract.Transact(opts, "sendToCosmos", _tokenContract, _destination, _amount)
}

// SendToCosmos is a paid mutator transaction binding the contract method 0x5d2428fa.
//
// Solidity: function sendToCosmos(address _tokenContract, address _destination, uint256 _amount) returns()
func (_Hyperion *HyperionSession) SendToCosmos(_tokenContract common.Address, _destination common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Hyperion.Contract.SendToCosmos(&_Hyperion.TransactOpts, _tokenContract, _destination, _amount)
}

// SendToCosmos is a paid mutator transaction binding the contract method 0x5d2428fa.
//
// Solidity: function sendToCosmos(address _tokenContract, address _destination, uint256 _amount) returns()
func (_Hyperion *HyperionTransactorSession) SendToCosmos(_tokenContract common.Address, _destination common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Hyperion.Contract.SendToCosmos(&_Hyperion.TransactOpts, _tokenContract, _destination, _amount)
}

// SubmitBatch is a paid mutator transaction binding the contract method 0xad131ec5.
//
// Solidity: function submitBatch(address[] _currentValidators, uint256[] _currentPowers, uint256 _currentValsetNonce, uint8[] _v, bytes32[] _r, bytes32[] _s, uint256[] _amounts, address[] _destinations, uint256[] _fees, uint256 _batchNonce, address _tokenContract) returns()
func (_Hyperion *HyperionTransactor) SubmitBatch(opts *bind.TransactOpts, _currentValidators []common.Address, _currentPowers []*big.Int, _currentValsetNonce *big.Int, _v []uint8, _r [][32]byte, _s [][32]byte, _amounts []*big.Int, _destinations []common.Address, _fees []*big.Int, _batchNonce *big.Int, _tokenContract common.Address) (*types.Transaction, error) {
	return _Hyperion.contract.Transact(opts, "submitBatch", _currentValidators, _currentPowers, _currentValsetNonce, _v, _r, _s, _amounts, _destinations, _fees, _batchNonce, _tokenContract)
}

// SubmitBatch is a paid mutator transaction binding the contract method 0xad131ec5.
//
// Solidity: function submitBatch(address[] _currentValidators, uint256[] _currentPowers, uint256 _currentValsetNonce, uint8[] _v, bytes32[] _r, bytes32[] _s, uint256[] _amounts, address[] _destinations, uint256[] _fees, uint256 _batchNonce, address _tokenContract) returns()
func (_Hyperion *HyperionSession) SubmitBatch(_currentValidators []common.Address, _currentPowers []*big.Int, _currentValsetNonce *big.Int, _v []uint8, _r [][32]byte, _s [][32]byte, _amounts []*big.Int, _destinations []common.Address, _fees []*big.Int, _batchNonce *big.Int, _tokenContract common.Address) (*types.Transaction, error) {
	return _Hyperion.Contract.SubmitBatch(&_Hyperion.TransactOpts, _currentValidators, _currentPowers, _currentValsetNonce, _v, _r, _s, _amounts, _destinations, _fees, _batchNonce, _tokenContract)
}

// SubmitBatch is a paid mutator transaction binding the contract method 0xad131ec5.
//
// Solidity: function submitBatch(address[] _currentValidators, uint256[] _currentPowers, uint256 _currentValsetNonce, uint8[] _v, bytes32[] _r, bytes32[] _s, uint256[] _amounts, address[] _destinations, uint256[] _fees, uint256 _batchNonce, address _tokenContract) returns()
func (_Hyperion *HyperionTransactorSession) SubmitBatch(_currentValidators []common.Address, _currentPowers []*big.Int, _currentValsetNonce *big.Int, _v []uint8, _r [][32]byte, _s [][32]byte, _amounts []*big.Int, _destinations []common.Address, _fees []*big.Int, _batchNonce *big.Int, _tokenContract common.Address) (*types.Transaction, error) {
	return _Hyperion.Contract.SubmitBatch(&_Hyperion.TransactOpts, _currentValidators, _currentPowers, _currentValsetNonce, _v, _r, _s, _amounts, _destinations, _fees, _batchNonce, _tokenContract)
}

// UpdateValset is a paid mutator transaction binding the contract method 0xe3cb9f62.
//
// Solidity: function updateValset(address[] _newValidators, uint256[] _newPowers, uint256 _newValsetNonce, address[] _currentValidators, uint256[] _currentPowers, uint256 _currentValsetNonce, uint8[] _v, bytes32[] _r, bytes32[] _s) returns()
func (_Hyperion *HyperionTransactor) UpdateValset(opts *bind.TransactOpts, _newValidators []common.Address, _newPowers []*big.Int, _newValsetNonce *big.Int, _currentValidators []common.Address, _currentPowers []*big.Int, _currentValsetNonce *big.Int, _v []uint8, _r [][32]byte, _s [][32]byte) (*types.Transaction, error) {
	return _Hyperion.contract.Transact(opts, "updateValset", _newValidators, _newPowers, _newValsetNonce, _currentValidators, _currentPowers, _currentValsetNonce, _v, _r, _s)
}

// UpdateValset is a paid mutator transaction binding the contract method 0xe3cb9f62.
//
// Solidity: function updateValset(address[] _newValidators, uint256[] _newPowers, uint256 _newValsetNonce, address[] _currentValidators, uint256[] _currentPowers, uint256 _currentValsetNonce, uint8[] _v, bytes32[] _r, bytes32[] _s) returns()
func (_Hyperion *HyperionSession) UpdateValset(_newValidators []common.Address, _newPowers []*big.Int, _newValsetNonce *big.Int, _currentValidators []common.Address, _currentPowers []*big.Int, _currentValsetNonce *big.Int, _v []uint8, _r [][32]byte, _s [][32]byte) (*types.Transaction, error) {
	return _Hyperion.Contract.UpdateValset(&_Hyperion.TransactOpts, _newValidators, _newPowers, _newValsetNonce, _currentValidators, _currentPowers, _currentValsetNonce, _v, _r, _s)
}

// UpdateValset is a paid mutator transaction binding the contract method 0xe3cb9f62.
//
// Solidity: function updateValset(address[] _newValidators, uint256[] _newPowers, uint256 _newValsetNonce, address[] _currentValidators, uint256[] _currentPowers, uint256 _currentValsetNonce, uint8[] _v, bytes32[] _r, bytes32[] _s) returns()
func (_Hyperion *HyperionTransactorSession) UpdateValset(_newValidators []common.Address, _newPowers []*big.Int, _newValsetNonce *big.Int, _currentValidators []common.Address, _currentPowers []*big.Int, _currentValsetNonce *big.Int, _v []uint8, _r [][32]byte, _s [][32]byte) (*types.Transaction, error) {
	return _Hyperion.Contract.UpdateValset(&_Hyperion.TransactOpts, _newValidators, _newPowers, _newValsetNonce, _currentValidators, _currentPowers, _currentValsetNonce, _v, _r, _s)
}

// HyperionSendToCosmosEventIterator is returned from FilterSendToCosmosEvent and is used to iterate over the raw logs and unpacked data for SendToCosmosEvent events raised by the Hyperion contract.
type HyperionSendToCosmosEventIterator struct {
	Event *HyperionSendToCosmosEvent // Event containing the contract specifics and raw log

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
func (it *HyperionSendToCosmosEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HyperionSendToCosmosEvent)
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
		it.Event = new(HyperionSendToCosmosEvent)
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
func (it *HyperionSendToCosmosEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HyperionSendToCosmosEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HyperionSendToCosmosEvent represents a SendToCosmosEvent event raised by the Hyperion contract.
type HyperionSendToCosmosEvent struct {
	TokenContract common.Address
	Sender        common.Address
	Destination   common.Address
	Amount        *big.Int
	EventNonce    *big.Int
	Name          string
	Symbol        string
	Decimals      uint8
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSendToCosmosEvent is a free log retrieval operation binding the contract event 0x5b0cfcd9629f2864d66b907e3625f822058529e928617605883190ad34ecb96d.
//
// Solidity: event SendToCosmosEvent(address indexed _tokenContract, address indexed _sender, address indexed _destination, uint256 _amount, uint256 _eventNonce, string _name, string _symbol, uint8 _decimals)
func (_Hyperion *HyperionFilterer) FilterSendToCosmosEvent(opts *bind.FilterOpts, _tokenContract []common.Address, _sender []common.Address, _destination []common.Address) (*HyperionSendToCosmosEventIterator, error) {

	var _tokenContractRule []interface{}
	for _, _tokenContractItem := range _tokenContract {
		_tokenContractRule = append(_tokenContractRule, _tokenContractItem)
	}
	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}
	var _destinationRule []interface{}
	for _, _destinationItem := range _destination {
		_destinationRule = append(_destinationRule, _destinationItem)
	}

	logs, sub, err := _Hyperion.contract.FilterLogs(opts, "SendToCosmosEvent", _tokenContractRule, _senderRule, _destinationRule)
	if err != nil {
		return nil, err
	}
	return &HyperionSendToCosmosEventIterator{contract: _Hyperion.contract, event: "SendToCosmosEvent", logs: logs, sub: sub}, nil
}

// WatchSendToCosmosEvent is a free log subscription operation binding the contract event 0x5b0cfcd9629f2864d66b907e3625f822058529e928617605883190ad34ecb96d.
//
// Solidity: event SendToCosmosEvent(address indexed _tokenContract, address indexed _sender, address indexed _destination, uint256 _amount, uint256 _eventNonce, string _name, string _symbol, uint8 _decimals)
func (_Hyperion *HyperionFilterer) WatchSendToCosmosEvent(opts *bind.WatchOpts, sink chan<- *HyperionSendToCosmosEvent, _tokenContract []common.Address, _sender []common.Address, _destination []common.Address) (event.Subscription, error) {

	var _tokenContractRule []interface{}
	for _, _tokenContractItem := range _tokenContract {
		_tokenContractRule = append(_tokenContractRule, _tokenContractItem)
	}
	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}
	var _destinationRule []interface{}
	for _, _destinationItem := range _destination {
		_destinationRule = append(_destinationRule, _destinationItem)
	}

	logs, sub, err := _Hyperion.contract.WatchLogs(opts, "SendToCosmosEvent", _tokenContractRule, _senderRule, _destinationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HyperionSendToCosmosEvent)
				if err := _Hyperion.contract.UnpackLog(event, "SendToCosmosEvent", log); err != nil {
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

// ParseSendToCosmosEvent is a log parse operation binding the contract event 0x5b0cfcd9629f2864d66b907e3625f822058529e928617605883190ad34ecb96d.
//
// Solidity: event SendToCosmosEvent(address indexed _tokenContract, address indexed _sender, address indexed _destination, uint256 _amount, uint256 _eventNonce, string _name, string _symbol, uint8 _decimals)
func (_Hyperion *HyperionFilterer) ParseSendToCosmosEvent(log types.Log) (*HyperionSendToCosmosEvent, error) {
	event := new(HyperionSendToCosmosEvent)
	if err := _Hyperion.contract.UnpackLog(event, "SendToCosmosEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HyperionTransactionBatchExecutedEventIterator is returned from FilterTransactionBatchExecutedEvent and is used to iterate over the raw logs and unpacked data for TransactionBatchExecutedEvent events raised by the Hyperion contract.
type HyperionTransactionBatchExecutedEventIterator struct {
	Event *HyperionTransactionBatchExecutedEvent // Event containing the contract specifics and raw log

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
func (it *HyperionTransactionBatchExecutedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HyperionTransactionBatchExecutedEvent)
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
		it.Event = new(HyperionTransactionBatchExecutedEvent)
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
func (it *HyperionTransactionBatchExecutedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HyperionTransactionBatchExecutedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HyperionTransactionBatchExecutedEvent represents a TransactionBatchExecutedEvent event raised by the Hyperion contract.
type HyperionTransactionBatchExecutedEvent struct {
	BatchNonce *big.Int
	Token      common.Address
	EventNonce *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTransactionBatchExecutedEvent is a free log retrieval operation binding the contract event 0x02c7e81975f8edb86e2a0c038b7b86a49c744236abf0f6177ff5afc6986ab708.
//
// Solidity: event TransactionBatchExecutedEvent(uint256 indexed _batchNonce, address indexed _token, uint256 _eventNonce)
func (_Hyperion *HyperionFilterer) FilterTransactionBatchExecutedEvent(opts *bind.FilterOpts, _batchNonce []*big.Int, _token []common.Address) (*HyperionTransactionBatchExecutedEventIterator, error) {

	var _batchNonceRule []interface{}
	for _, _batchNonceItem := range _batchNonce {
		_batchNonceRule = append(_batchNonceRule, _batchNonceItem)
	}
	var _tokenRule []interface{}
	for _, _tokenItem := range _token {
		_tokenRule = append(_tokenRule, _tokenItem)
	}

	logs, sub, err := _Hyperion.contract.FilterLogs(opts, "TransactionBatchExecutedEvent", _batchNonceRule, _tokenRule)
	if err != nil {
		return nil, err
	}
	return &HyperionTransactionBatchExecutedEventIterator{contract: _Hyperion.contract, event: "TransactionBatchExecutedEvent", logs: logs, sub: sub}, nil
}

// WatchTransactionBatchExecutedEvent is a free log subscription operation binding the contract event 0x02c7e81975f8edb86e2a0c038b7b86a49c744236abf0f6177ff5afc6986ab708.
//
// Solidity: event TransactionBatchExecutedEvent(uint256 indexed _batchNonce, address indexed _token, uint256 _eventNonce)
func (_Hyperion *HyperionFilterer) WatchTransactionBatchExecutedEvent(opts *bind.WatchOpts, sink chan<- *HyperionTransactionBatchExecutedEvent, _batchNonce []*big.Int, _token []common.Address) (event.Subscription, error) {

	var _batchNonceRule []interface{}
	for _, _batchNonceItem := range _batchNonce {
		_batchNonceRule = append(_batchNonceRule, _batchNonceItem)
	}
	var _tokenRule []interface{}
	for _, _tokenItem := range _token {
		_tokenRule = append(_tokenRule, _tokenItem)
	}

	logs, sub, err := _Hyperion.contract.WatchLogs(opts, "TransactionBatchExecutedEvent", _batchNonceRule, _tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HyperionTransactionBatchExecutedEvent)
				if err := _Hyperion.contract.UnpackLog(event, "TransactionBatchExecutedEvent", log); err != nil {
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

// ParseTransactionBatchExecutedEvent is a log parse operation binding the contract event 0x02c7e81975f8edb86e2a0c038b7b86a49c744236abf0f6177ff5afc6986ab708.
//
// Solidity: event TransactionBatchExecutedEvent(uint256 indexed _batchNonce, address indexed _token, uint256 _eventNonce)
func (_Hyperion *HyperionFilterer) ParseTransactionBatchExecutedEvent(log types.Log) (*HyperionTransactionBatchExecutedEvent, error) {
	event := new(HyperionTransactionBatchExecutedEvent)
	if err := _Hyperion.contract.UnpackLog(event, "TransactionBatchExecutedEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HyperionValsetUpdatedEventIterator is returned from FilterValsetUpdatedEvent and is used to iterate over the raw logs and unpacked data for ValsetUpdatedEvent events raised by the Hyperion contract.
type HyperionValsetUpdatedEventIterator struct {
	Event *HyperionValsetUpdatedEvent // Event containing the contract specifics and raw log

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
func (it *HyperionValsetUpdatedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HyperionValsetUpdatedEvent)
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
		it.Event = new(HyperionValsetUpdatedEvent)
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
func (it *HyperionValsetUpdatedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HyperionValsetUpdatedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HyperionValsetUpdatedEvent represents a ValsetUpdatedEvent event raised by the Hyperion contract.
type HyperionValsetUpdatedEvent struct {
	NewValsetNonce *big.Int
	Validators     []common.Address
	Powers         []*big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterValsetUpdatedEvent is a free log retrieval operation binding the contract event 0xc6d025c076bafcdd040f00632d5e280b3a5188963f110f8c70c4f810184b30f3.
//
// Solidity: event ValsetUpdatedEvent(uint256 indexed _newValsetNonce, address[] _validators, uint256[] _powers)
func (_Hyperion *HyperionFilterer) FilterValsetUpdatedEvent(opts *bind.FilterOpts, _newValsetNonce []*big.Int) (*HyperionValsetUpdatedEventIterator, error) {

	var _newValsetNonceRule []interface{}
	for _, _newValsetNonceItem := range _newValsetNonce {
		_newValsetNonceRule = append(_newValsetNonceRule, _newValsetNonceItem)
	}

	logs, sub, err := _Hyperion.contract.FilterLogs(opts, "ValsetUpdatedEvent", _newValsetNonceRule)
	if err != nil {
		return nil, err
	}
	return &HyperionValsetUpdatedEventIterator{contract: _Hyperion.contract, event: "ValsetUpdatedEvent", logs: logs, sub: sub}, nil
}

// WatchValsetUpdatedEvent is a free log subscription operation binding the contract event 0xc6d025c076bafcdd040f00632d5e280b3a5188963f110f8c70c4f810184b30f3.
//
// Solidity: event ValsetUpdatedEvent(uint256 indexed _newValsetNonce, address[] _validators, uint256[] _powers)
func (_Hyperion *HyperionFilterer) WatchValsetUpdatedEvent(opts *bind.WatchOpts, sink chan<- *HyperionValsetUpdatedEvent, _newValsetNonce []*big.Int) (event.Subscription, error) {

	var _newValsetNonceRule []interface{}
	for _, _newValsetNonceItem := range _newValsetNonce {
		_newValsetNonceRule = append(_newValsetNonceRule, _newValsetNonceItem)
	}

	logs, sub, err := _Hyperion.contract.WatchLogs(opts, "ValsetUpdatedEvent", _newValsetNonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HyperionValsetUpdatedEvent)
				if err := _Hyperion.contract.UnpackLog(event, "ValsetUpdatedEvent", log); err != nil {
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

// ParseValsetUpdatedEvent is a log parse operation binding the contract event 0xc6d025c076bafcdd040f00632d5e280b3a5188963f110f8c70c4f810184b30f3.
//
// Solidity: event ValsetUpdatedEvent(uint256 indexed _newValsetNonce, address[] _validators, uint256[] _powers)
func (_Hyperion *HyperionFilterer) ParseValsetUpdatedEvent(log types.Log) (*HyperionValsetUpdatedEvent, error) {
	event := new(HyperionValsetUpdatedEvent)
	if err := _Hyperion.contract.UnpackLog(event, "ValsetUpdatedEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
