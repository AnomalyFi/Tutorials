// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package hyperc20

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

// GasRouterGasRouterConfig is an auto generated low-level Go binding around an user-defined struct.
type GasRouterGasRouterConfig struct {
	Domain uint32
	Gas    *big.Int
}

// HypERC20MetaData contains all meta data concerning the HypERC20 contract.
var HypERC20MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"__decimals\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_mailbox\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"origin\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"recipient\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ReceivedTransferRemote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"destination\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"recipient\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SentTransferRemote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"destinationGas\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domains\",\"outputs\":[{\"internalType\":\"uint32[]\",\"name\":\"\",\"type\":\"uint32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_router\",\"type\":\"bytes32\"}],\"name\":\"enrollRemoteRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32[]\",\"name\":\"_domains\",\"type\":\"uint32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_addresses\",\"type\":\"bytes32[]\"}],\"name\":\"enrollRemoteRouters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_origin\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_sender\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"handle\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hook\",\"outputs\":[{\"internalType\":\"contractIPostDispatchHook\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_totalSupply\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"interchainSecurityModule\",\"outputs\":[{\"internalType\":\"contractIInterchainSecurityModule\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mailbox\",\"outputs\":[{\"internalType\":\"contractIMailbox\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_destinationDomain\",\"type\":\"uint32\"}],\"name\":\"quoteGasPayment\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_gasPayment\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"}],\"name\":\"routers\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"}],\"name\":\"setDestinationGas\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"}],\"internalType\":\"structGasRouter.GasRouterConfig[]\",\"name\":\"gasConfigs\",\"type\":\"tuple[]\"}],\"name\":\"setDestinationGas\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_hook\",\"type\":\"address\"}],\"name\":\"setHook\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_module\",\"type\":\"address\"}],\"name\":\"setInterchainSecurityModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_destination\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_recipient\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_amountOrId\",\"type\":\"uint256\"}],\"name\":\"transferRemote\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"}],\"name\":\"unenrollRemoteRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32[]\",\"name\":\"_domains\",\"type\":\"uint32[]\"}],\"name\":\"unenrollRemoteRouters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// HypERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use HypERC20MetaData.ABI instead.
var HypERC20ABI = HypERC20MetaData.ABI

// HypERC20 is an auto generated Go binding around an Ethereum contract.
type HypERC20 struct {
	HypERC20Caller     // Read-only binding to the contract
	HypERC20Transactor // Write-only binding to the contract
	HypERC20Filterer   // Log filterer for contract events
}

// HypERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type HypERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HypERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type HypERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HypERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HypERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HypERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HypERC20Session struct {
	Contract     *HypERC20         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HypERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HypERC20CallerSession struct {
	Contract *HypERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// HypERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HypERC20TransactorSession struct {
	Contract     *HypERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// HypERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type HypERC20Raw struct {
	Contract *HypERC20 // Generic contract binding to access the raw methods on
}

// HypERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HypERC20CallerRaw struct {
	Contract *HypERC20Caller // Generic read-only contract binding to access the raw methods on
}

// HypERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HypERC20TransactorRaw struct {
	Contract *HypERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewHypERC20 creates a new instance of HypERC20, bound to a specific deployed contract.
func NewHypERC20(address common.Address, backend bind.ContractBackend) (*HypERC20, error) {
	contract, err := bindHypERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HypERC20{HypERC20Caller: HypERC20Caller{contract: contract}, HypERC20Transactor: HypERC20Transactor{contract: contract}, HypERC20Filterer: HypERC20Filterer{contract: contract}}, nil
}

// NewHypERC20Caller creates a new read-only instance of HypERC20, bound to a specific deployed contract.
func NewHypERC20Caller(address common.Address, caller bind.ContractCaller) (*HypERC20Caller, error) {
	contract, err := bindHypERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HypERC20Caller{contract: contract}, nil
}

// NewHypERC20Transactor creates a new write-only instance of HypERC20, bound to a specific deployed contract.
func NewHypERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*HypERC20Transactor, error) {
	contract, err := bindHypERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HypERC20Transactor{contract: contract}, nil
}

// NewHypERC20Filterer creates a new log filterer instance of HypERC20, bound to a specific deployed contract.
func NewHypERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*HypERC20Filterer, error) {
	contract, err := bindHypERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HypERC20Filterer{contract: contract}, nil
}

// bindHypERC20 binds a generic wrapper to an already deployed contract.
func bindHypERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HypERC20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HypERC20 *HypERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HypERC20.Contract.HypERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HypERC20 *HypERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HypERC20.Contract.HypERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HypERC20 *HypERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HypERC20.Contract.HypERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HypERC20 *HypERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HypERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HypERC20 *HypERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HypERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HypERC20 *HypERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HypERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_HypERC20 *HypERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _HypERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_HypERC20 *HypERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _HypERC20.Contract.Allowance(&_HypERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_HypERC20 *HypERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _HypERC20.Contract.Allowance(&_HypERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_HypERC20 *HypERC20Caller) BalanceOf(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _HypERC20.contract.Call(opts, &out, "balanceOf", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_HypERC20 *HypERC20Session) BalanceOf(_account common.Address) (*big.Int, error) {
	return _HypERC20.Contract.BalanceOf(&_HypERC20.CallOpts, _account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_HypERC20 *HypERC20CallerSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _HypERC20.Contract.BalanceOf(&_HypERC20.CallOpts, _account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_HypERC20 *HypERC20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _HypERC20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_HypERC20 *HypERC20Session) Decimals() (uint8, error) {
	return _HypERC20.Contract.Decimals(&_HypERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_HypERC20 *HypERC20CallerSession) Decimals() (uint8, error) {
	return _HypERC20.Contract.Decimals(&_HypERC20.CallOpts)
}

// DestinationGas is a free data retrieval call binding the contract method 0x775313a1.
//
// Solidity: function destinationGas(uint32 ) view returns(uint256)
func (_HypERC20 *HypERC20Caller) DestinationGas(opts *bind.CallOpts, arg0 uint32) (*big.Int, error) {
	var out []interface{}
	err := _HypERC20.contract.Call(opts, &out, "destinationGas", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DestinationGas is a free data retrieval call binding the contract method 0x775313a1.
//
// Solidity: function destinationGas(uint32 ) view returns(uint256)
func (_HypERC20 *HypERC20Session) DestinationGas(arg0 uint32) (*big.Int, error) {
	return _HypERC20.Contract.DestinationGas(&_HypERC20.CallOpts, arg0)
}

// DestinationGas is a free data retrieval call binding the contract method 0x775313a1.
//
// Solidity: function destinationGas(uint32 ) view returns(uint256)
func (_HypERC20 *HypERC20CallerSession) DestinationGas(arg0 uint32) (*big.Int, error) {
	return _HypERC20.Contract.DestinationGas(&_HypERC20.CallOpts, arg0)
}

// Domains is a free data retrieval call binding the contract method 0x440df4f4.
//
// Solidity: function domains() view returns(uint32[])
func (_HypERC20 *HypERC20Caller) Domains(opts *bind.CallOpts) ([]uint32, error) {
	var out []interface{}
	err := _HypERC20.contract.Call(opts, &out, "domains")

	if err != nil {
		return *new([]uint32), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint32)).(*[]uint32)

	return out0, err

}

// Domains is a free data retrieval call binding the contract method 0x440df4f4.
//
// Solidity: function domains() view returns(uint32[])
func (_HypERC20 *HypERC20Session) Domains() ([]uint32, error) {
	return _HypERC20.Contract.Domains(&_HypERC20.CallOpts)
}

// Domains is a free data retrieval call binding the contract method 0x440df4f4.
//
// Solidity: function domains() view returns(uint32[])
func (_HypERC20 *HypERC20CallerSession) Domains() ([]uint32, error) {
	return _HypERC20.Contract.Domains(&_HypERC20.CallOpts)
}

// Hook is a free data retrieval call binding the contract method 0x7f5a7c7b.
//
// Solidity: function hook() view returns(address)
func (_HypERC20 *HypERC20Caller) Hook(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HypERC20.contract.Call(opts, &out, "hook")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Hook is a free data retrieval call binding the contract method 0x7f5a7c7b.
//
// Solidity: function hook() view returns(address)
func (_HypERC20 *HypERC20Session) Hook() (common.Address, error) {
	return _HypERC20.Contract.Hook(&_HypERC20.CallOpts)
}

// Hook is a free data retrieval call binding the contract method 0x7f5a7c7b.
//
// Solidity: function hook() view returns(address)
func (_HypERC20 *HypERC20CallerSession) Hook() (common.Address, error) {
	return _HypERC20.Contract.Hook(&_HypERC20.CallOpts)
}

// InterchainSecurityModule is a free data retrieval call binding the contract method 0xde523cf3.
//
// Solidity: function interchainSecurityModule() view returns(address)
func (_HypERC20 *HypERC20Caller) InterchainSecurityModule(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HypERC20.contract.Call(opts, &out, "interchainSecurityModule")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InterchainSecurityModule is a free data retrieval call binding the contract method 0xde523cf3.
//
// Solidity: function interchainSecurityModule() view returns(address)
func (_HypERC20 *HypERC20Session) InterchainSecurityModule() (common.Address, error) {
	return _HypERC20.Contract.InterchainSecurityModule(&_HypERC20.CallOpts)
}

// InterchainSecurityModule is a free data retrieval call binding the contract method 0xde523cf3.
//
// Solidity: function interchainSecurityModule() view returns(address)
func (_HypERC20 *HypERC20CallerSession) InterchainSecurityModule() (common.Address, error) {
	return _HypERC20.Contract.InterchainSecurityModule(&_HypERC20.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_HypERC20 *HypERC20Caller) LocalDomain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _HypERC20.contract.Call(opts, &out, "localDomain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_HypERC20 *HypERC20Session) LocalDomain() (uint32, error) {
	return _HypERC20.Contract.LocalDomain(&_HypERC20.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_HypERC20 *HypERC20CallerSession) LocalDomain() (uint32, error) {
	return _HypERC20.Contract.LocalDomain(&_HypERC20.CallOpts)
}

// Mailbox is a free data retrieval call binding the contract method 0xd5438eae.
//
// Solidity: function mailbox() view returns(address)
func (_HypERC20 *HypERC20Caller) Mailbox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HypERC20.contract.Call(opts, &out, "mailbox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Mailbox is a free data retrieval call binding the contract method 0xd5438eae.
//
// Solidity: function mailbox() view returns(address)
func (_HypERC20 *HypERC20Session) Mailbox() (common.Address, error) {
	return _HypERC20.Contract.Mailbox(&_HypERC20.CallOpts)
}

// Mailbox is a free data retrieval call binding the contract method 0xd5438eae.
//
// Solidity: function mailbox() view returns(address)
func (_HypERC20 *HypERC20CallerSession) Mailbox() (common.Address, error) {
	return _HypERC20.Contract.Mailbox(&_HypERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_HypERC20 *HypERC20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _HypERC20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_HypERC20 *HypERC20Session) Name() (string, error) {
	return _HypERC20.Contract.Name(&_HypERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_HypERC20 *HypERC20CallerSession) Name() (string, error) {
	return _HypERC20.Contract.Name(&_HypERC20.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_HypERC20 *HypERC20Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HypERC20.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_HypERC20 *HypERC20Session) Owner() (common.Address, error) {
	return _HypERC20.Contract.Owner(&_HypERC20.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_HypERC20 *HypERC20CallerSession) Owner() (common.Address, error) {
	return _HypERC20.Contract.Owner(&_HypERC20.CallOpts)
}

// QuoteGasPayment is a free data retrieval call binding the contract method 0xf2ed8c53.
//
// Solidity: function quoteGasPayment(uint32 _destinationDomain) view returns(uint256 _gasPayment)
func (_HypERC20 *HypERC20Caller) QuoteGasPayment(opts *bind.CallOpts, _destinationDomain uint32) (*big.Int, error) {
	var out []interface{}
	err := _HypERC20.contract.Call(opts, &out, "quoteGasPayment", _destinationDomain)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuoteGasPayment is a free data retrieval call binding the contract method 0xf2ed8c53.
//
// Solidity: function quoteGasPayment(uint32 _destinationDomain) view returns(uint256 _gasPayment)
func (_HypERC20 *HypERC20Session) QuoteGasPayment(_destinationDomain uint32) (*big.Int, error) {
	return _HypERC20.Contract.QuoteGasPayment(&_HypERC20.CallOpts, _destinationDomain)
}

// QuoteGasPayment is a free data retrieval call binding the contract method 0xf2ed8c53.
//
// Solidity: function quoteGasPayment(uint32 _destinationDomain) view returns(uint256 _gasPayment)
func (_HypERC20 *HypERC20CallerSession) QuoteGasPayment(_destinationDomain uint32) (*big.Int, error) {
	return _HypERC20.Contract.QuoteGasPayment(&_HypERC20.CallOpts, _destinationDomain)
}

// Routers is a free data retrieval call binding the contract method 0x2ead72f6.
//
// Solidity: function routers(uint32 _domain) view returns(bytes32)
func (_HypERC20 *HypERC20Caller) Routers(opts *bind.CallOpts, _domain uint32) ([32]byte, error) {
	var out []interface{}
	err := _HypERC20.contract.Call(opts, &out, "routers", _domain)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Routers is a free data retrieval call binding the contract method 0x2ead72f6.
//
// Solidity: function routers(uint32 _domain) view returns(bytes32)
func (_HypERC20 *HypERC20Session) Routers(_domain uint32) ([32]byte, error) {
	return _HypERC20.Contract.Routers(&_HypERC20.CallOpts, _domain)
}

// Routers is a free data retrieval call binding the contract method 0x2ead72f6.
//
// Solidity: function routers(uint32 _domain) view returns(bytes32)
func (_HypERC20 *HypERC20CallerSession) Routers(_domain uint32) ([32]byte, error) {
	return _HypERC20.Contract.Routers(&_HypERC20.CallOpts, _domain)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_HypERC20 *HypERC20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _HypERC20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_HypERC20 *HypERC20Session) Symbol() (string, error) {
	return _HypERC20.Contract.Symbol(&_HypERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_HypERC20 *HypERC20CallerSession) Symbol() (string, error) {
	return _HypERC20.Contract.Symbol(&_HypERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_HypERC20 *HypERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _HypERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_HypERC20 *HypERC20Session) TotalSupply() (*big.Int, error) {
	return _HypERC20.Contract.TotalSupply(&_HypERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_HypERC20 *HypERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _HypERC20.Contract.TotalSupply(&_HypERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_HypERC20 *HypERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _HypERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_HypERC20 *HypERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _HypERC20.Contract.Approve(&_HypERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_HypERC20 *HypERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _HypERC20.Contract.Approve(&_HypERC20.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_HypERC20 *HypERC20Transactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _HypERC20.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_HypERC20 *HypERC20Session) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _HypERC20.Contract.DecreaseAllowance(&_HypERC20.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_HypERC20 *HypERC20TransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _HypERC20.Contract.DecreaseAllowance(&_HypERC20.TransactOpts, spender, subtractedValue)
}

// EnrollRemoteRouter is a paid mutator transaction binding the contract method 0xb49c53a7.
//
// Solidity: function enrollRemoteRouter(uint32 _domain, bytes32 _router) returns()
func (_HypERC20 *HypERC20Transactor) EnrollRemoteRouter(opts *bind.TransactOpts, _domain uint32, _router [32]byte) (*types.Transaction, error) {
	return _HypERC20.contract.Transact(opts, "enrollRemoteRouter", _domain, _router)
}

// EnrollRemoteRouter is a paid mutator transaction binding the contract method 0xb49c53a7.
//
// Solidity: function enrollRemoteRouter(uint32 _domain, bytes32 _router) returns()
func (_HypERC20 *HypERC20Session) EnrollRemoteRouter(_domain uint32, _router [32]byte) (*types.Transaction, error) {
	return _HypERC20.Contract.EnrollRemoteRouter(&_HypERC20.TransactOpts, _domain, _router)
}

// EnrollRemoteRouter is a paid mutator transaction binding the contract method 0xb49c53a7.
//
// Solidity: function enrollRemoteRouter(uint32 _domain, bytes32 _router) returns()
func (_HypERC20 *HypERC20TransactorSession) EnrollRemoteRouter(_domain uint32, _router [32]byte) (*types.Transaction, error) {
	return _HypERC20.Contract.EnrollRemoteRouter(&_HypERC20.TransactOpts, _domain, _router)
}

// EnrollRemoteRouters is a paid mutator transaction binding the contract method 0xe9198bf9.
//
// Solidity: function enrollRemoteRouters(uint32[] _domains, bytes32[] _addresses) returns()
func (_HypERC20 *HypERC20Transactor) EnrollRemoteRouters(opts *bind.TransactOpts, _domains []uint32, _addresses [][32]byte) (*types.Transaction, error) {
	return _HypERC20.contract.Transact(opts, "enrollRemoteRouters", _domains, _addresses)
}

// EnrollRemoteRouters is a paid mutator transaction binding the contract method 0xe9198bf9.
//
// Solidity: function enrollRemoteRouters(uint32[] _domains, bytes32[] _addresses) returns()
func (_HypERC20 *HypERC20Session) EnrollRemoteRouters(_domains []uint32, _addresses [][32]byte) (*types.Transaction, error) {
	return _HypERC20.Contract.EnrollRemoteRouters(&_HypERC20.TransactOpts, _domains, _addresses)
}

// EnrollRemoteRouters is a paid mutator transaction binding the contract method 0xe9198bf9.
//
// Solidity: function enrollRemoteRouters(uint32[] _domains, bytes32[] _addresses) returns()
func (_HypERC20 *HypERC20TransactorSession) EnrollRemoteRouters(_domains []uint32, _addresses [][32]byte) (*types.Transaction, error) {
	return _HypERC20.Contract.EnrollRemoteRouters(&_HypERC20.TransactOpts, _domains, _addresses)
}

// Handle is a paid mutator transaction binding the contract method 0x56d5d475.
//
// Solidity: function handle(uint32 _origin, bytes32 _sender, bytes _message) payable returns()
func (_HypERC20 *HypERC20Transactor) Handle(opts *bind.TransactOpts, _origin uint32, _sender [32]byte, _message []byte) (*types.Transaction, error) {
	return _HypERC20.contract.Transact(opts, "handle", _origin, _sender, _message)
}

// Handle is a paid mutator transaction binding the contract method 0x56d5d475.
//
// Solidity: function handle(uint32 _origin, bytes32 _sender, bytes _message) payable returns()
func (_HypERC20 *HypERC20Session) Handle(_origin uint32, _sender [32]byte, _message []byte) (*types.Transaction, error) {
	return _HypERC20.Contract.Handle(&_HypERC20.TransactOpts, _origin, _sender, _message)
}

// Handle is a paid mutator transaction binding the contract method 0x56d5d475.
//
// Solidity: function handle(uint32 _origin, bytes32 _sender, bytes _message) payable returns()
func (_HypERC20 *HypERC20TransactorSession) Handle(_origin uint32, _sender [32]byte, _message []byte) (*types.Transaction, error) {
	return _HypERC20.Contract.Handle(&_HypERC20.TransactOpts, _origin, _sender, _message)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_HypERC20 *HypERC20Transactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _HypERC20.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_HypERC20 *HypERC20Session) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _HypERC20.Contract.IncreaseAllowance(&_HypERC20.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_HypERC20 *HypERC20TransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _HypERC20.Contract.IncreaseAllowance(&_HypERC20.TransactOpts, spender, addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0xeedfca5f.
//
// Solidity: function initialize(uint256 _totalSupply, string _name, string _symbol) returns()
func (_HypERC20 *HypERC20Transactor) Initialize(opts *bind.TransactOpts, _totalSupply *big.Int, _name string, _symbol string) (*types.Transaction, error) {
	return _HypERC20.contract.Transact(opts, "initialize", _totalSupply, _name, _symbol)
}

// Initialize is a paid mutator transaction binding the contract method 0xeedfca5f.
//
// Solidity: function initialize(uint256 _totalSupply, string _name, string _symbol) returns()
func (_HypERC20 *HypERC20Session) Initialize(_totalSupply *big.Int, _name string, _symbol string) (*types.Transaction, error) {
	return _HypERC20.Contract.Initialize(&_HypERC20.TransactOpts, _totalSupply, _name, _symbol)
}

// Initialize is a paid mutator transaction binding the contract method 0xeedfca5f.
//
// Solidity: function initialize(uint256 _totalSupply, string _name, string _symbol) returns()
func (_HypERC20 *HypERC20TransactorSession) Initialize(_totalSupply *big.Int, _name string, _symbol string) (*types.Transaction, error) {
	return _HypERC20.Contract.Initialize(&_HypERC20.TransactOpts, _totalSupply, _name, _symbol)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_HypERC20 *HypERC20Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HypERC20.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_HypERC20 *HypERC20Session) RenounceOwnership() (*types.Transaction, error) {
	return _HypERC20.Contract.RenounceOwnership(&_HypERC20.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_HypERC20 *HypERC20TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _HypERC20.Contract.RenounceOwnership(&_HypERC20.TransactOpts)
}

// SetDestinationGas is a paid mutator transaction binding the contract method 0x49d462ef.
//
// Solidity: function setDestinationGas(uint32 domain, uint256 gas) returns()
func (_HypERC20 *HypERC20Transactor) SetDestinationGas(opts *bind.TransactOpts, domain uint32, gas *big.Int) (*types.Transaction, error) {
	return _HypERC20.contract.Transact(opts, "setDestinationGas", domain, gas)
}

// SetDestinationGas is a paid mutator transaction binding the contract method 0x49d462ef.
//
// Solidity: function setDestinationGas(uint32 domain, uint256 gas) returns()
func (_HypERC20 *HypERC20Session) SetDestinationGas(domain uint32, gas *big.Int) (*types.Transaction, error) {
	return _HypERC20.Contract.SetDestinationGas(&_HypERC20.TransactOpts, domain, gas)
}

// SetDestinationGas is a paid mutator transaction binding the contract method 0x49d462ef.
//
// Solidity: function setDestinationGas(uint32 domain, uint256 gas) returns()
func (_HypERC20 *HypERC20TransactorSession) SetDestinationGas(domain uint32, gas *big.Int) (*types.Transaction, error) {
	return _HypERC20.Contract.SetDestinationGas(&_HypERC20.TransactOpts, domain, gas)
}

// SetDestinationGas0 is a paid mutator transaction binding the contract method 0xb1bd6436.
//
// Solidity: function setDestinationGas((uint32,uint256)[] gasConfigs) returns()
func (_HypERC20 *HypERC20Transactor) SetDestinationGas0(opts *bind.TransactOpts, gasConfigs []GasRouterGasRouterConfig) (*types.Transaction, error) {
	return _HypERC20.contract.Transact(opts, "setDestinationGas0", gasConfigs)
}

// SetDestinationGas0 is a paid mutator transaction binding the contract method 0xb1bd6436.
//
// Solidity: function setDestinationGas((uint32,uint256)[] gasConfigs) returns()
func (_HypERC20 *HypERC20Session) SetDestinationGas0(gasConfigs []GasRouterGasRouterConfig) (*types.Transaction, error) {
	return _HypERC20.Contract.SetDestinationGas0(&_HypERC20.TransactOpts, gasConfigs)
}

// SetDestinationGas0 is a paid mutator transaction binding the contract method 0xb1bd6436.
//
// Solidity: function setDestinationGas((uint32,uint256)[] gasConfigs) returns()
func (_HypERC20 *HypERC20TransactorSession) SetDestinationGas0(gasConfigs []GasRouterGasRouterConfig) (*types.Transaction, error) {
	return _HypERC20.Contract.SetDestinationGas0(&_HypERC20.TransactOpts, gasConfigs)
}

// SetHook is a paid mutator transaction binding the contract method 0x3dfd3873.
//
// Solidity: function setHook(address _hook) returns()
func (_HypERC20 *HypERC20Transactor) SetHook(opts *bind.TransactOpts, _hook common.Address) (*types.Transaction, error) {
	return _HypERC20.contract.Transact(opts, "setHook", _hook)
}

// SetHook is a paid mutator transaction binding the contract method 0x3dfd3873.
//
// Solidity: function setHook(address _hook) returns()
func (_HypERC20 *HypERC20Session) SetHook(_hook common.Address) (*types.Transaction, error) {
	return _HypERC20.Contract.SetHook(&_HypERC20.TransactOpts, _hook)
}

// SetHook is a paid mutator transaction binding the contract method 0x3dfd3873.
//
// Solidity: function setHook(address _hook) returns()
func (_HypERC20 *HypERC20TransactorSession) SetHook(_hook common.Address) (*types.Transaction, error) {
	return _HypERC20.Contract.SetHook(&_HypERC20.TransactOpts, _hook)
}

// SetInterchainSecurityModule is a paid mutator transaction binding the contract method 0x0e72cc06.
//
// Solidity: function setInterchainSecurityModule(address _module) returns()
func (_HypERC20 *HypERC20Transactor) SetInterchainSecurityModule(opts *bind.TransactOpts, _module common.Address) (*types.Transaction, error) {
	return _HypERC20.contract.Transact(opts, "setInterchainSecurityModule", _module)
}

// SetInterchainSecurityModule is a paid mutator transaction binding the contract method 0x0e72cc06.
//
// Solidity: function setInterchainSecurityModule(address _module) returns()
func (_HypERC20 *HypERC20Session) SetInterchainSecurityModule(_module common.Address) (*types.Transaction, error) {
	return _HypERC20.Contract.SetInterchainSecurityModule(&_HypERC20.TransactOpts, _module)
}

// SetInterchainSecurityModule is a paid mutator transaction binding the contract method 0x0e72cc06.
//
// Solidity: function setInterchainSecurityModule(address _module) returns()
func (_HypERC20 *HypERC20TransactorSession) SetInterchainSecurityModule(_module common.Address) (*types.Transaction, error) {
	return _HypERC20.Contract.SetInterchainSecurityModule(&_HypERC20.TransactOpts, _module)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_HypERC20 *HypERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _HypERC20.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_HypERC20 *HypERC20Session) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _HypERC20.Contract.Transfer(&_HypERC20.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_HypERC20 *HypERC20TransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _HypERC20.Contract.Transfer(&_HypERC20.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_HypERC20 *HypERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _HypERC20.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_HypERC20 *HypERC20Session) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _HypERC20.Contract.TransferFrom(&_HypERC20.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_HypERC20 *HypERC20TransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _HypERC20.Contract.TransferFrom(&_HypERC20.TransactOpts, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_HypERC20 *HypERC20Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _HypERC20.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_HypERC20 *HypERC20Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _HypERC20.Contract.TransferOwnership(&_HypERC20.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_HypERC20 *HypERC20TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _HypERC20.Contract.TransferOwnership(&_HypERC20.TransactOpts, newOwner)
}

// TransferRemote is a paid mutator transaction binding the contract method 0x81b4e8b4.
//
// Solidity: function transferRemote(uint32 _destination, bytes32 _recipient, uint256 _amountOrId) payable returns(bytes32 messageId)
func (_HypERC20 *HypERC20Transactor) TransferRemote(opts *bind.TransactOpts, _destination uint32, _recipient [32]byte, _amountOrId *big.Int) (*types.Transaction, error) {
	return _HypERC20.contract.Transact(opts, "transferRemote", _destination, _recipient, _amountOrId)
}

// TransferRemote is a paid mutator transaction binding the contract method 0x81b4e8b4.
//
// Solidity: function transferRemote(uint32 _destination, bytes32 _recipient, uint256 _amountOrId) payable returns(bytes32 messageId)
func (_HypERC20 *HypERC20Session) TransferRemote(_destination uint32, _recipient [32]byte, _amountOrId *big.Int) (*types.Transaction, error) {
	return _HypERC20.Contract.TransferRemote(&_HypERC20.TransactOpts, _destination, _recipient, _amountOrId)
}

// TransferRemote is a paid mutator transaction binding the contract method 0x81b4e8b4.
//
// Solidity: function transferRemote(uint32 _destination, bytes32 _recipient, uint256 _amountOrId) payable returns(bytes32 messageId)
func (_HypERC20 *HypERC20TransactorSession) TransferRemote(_destination uint32, _recipient [32]byte, _amountOrId *big.Int) (*types.Transaction, error) {
	return _HypERC20.Contract.TransferRemote(&_HypERC20.TransactOpts, _destination, _recipient, _amountOrId)
}

// UnenrollRemoteRouter is a paid mutator transaction binding the contract method 0xefae508a.
//
// Solidity: function unenrollRemoteRouter(uint32 _domain) returns()
func (_HypERC20 *HypERC20Transactor) UnenrollRemoteRouter(opts *bind.TransactOpts, _domain uint32) (*types.Transaction, error) {
	return _HypERC20.contract.Transact(opts, "unenrollRemoteRouter", _domain)
}

// UnenrollRemoteRouter is a paid mutator transaction binding the contract method 0xefae508a.
//
// Solidity: function unenrollRemoteRouter(uint32 _domain) returns()
func (_HypERC20 *HypERC20Session) UnenrollRemoteRouter(_domain uint32) (*types.Transaction, error) {
	return _HypERC20.Contract.UnenrollRemoteRouter(&_HypERC20.TransactOpts, _domain)
}

// UnenrollRemoteRouter is a paid mutator transaction binding the contract method 0xefae508a.
//
// Solidity: function unenrollRemoteRouter(uint32 _domain) returns()
func (_HypERC20 *HypERC20TransactorSession) UnenrollRemoteRouter(_domain uint32) (*types.Transaction, error) {
	return _HypERC20.Contract.UnenrollRemoteRouter(&_HypERC20.TransactOpts, _domain)
}

// UnenrollRemoteRouters is a paid mutator transaction binding the contract method 0x71a15b38.
//
// Solidity: function unenrollRemoteRouters(uint32[] _domains) returns()
func (_HypERC20 *HypERC20Transactor) UnenrollRemoteRouters(opts *bind.TransactOpts, _domains []uint32) (*types.Transaction, error) {
	return _HypERC20.contract.Transact(opts, "unenrollRemoteRouters", _domains)
}

// UnenrollRemoteRouters is a paid mutator transaction binding the contract method 0x71a15b38.
//
// Solidity: function unenrollRemoteRouters(uint32[] _domains) returns()
func (_HypERC20 *HypERC20Session) UnenrollRemoteRouters(_domains []uint32) (*types.Transaction, error) {
	return _HypERC20.Contract.UnenrollRemoteRouters(&_HypERC20.TransactOpts, _domains)
}

// UnenrollRemoteRouters is a paid mutator transaction binding the contract method 0x71a15b38.
//
// Solidity: function unenrollRemoteRouters(uint32[] _domains) returns()
func (_HypERC20 *HypERC20TransactorSession) UnenrollRemoteRouters(_domains []uint32) (*types.Transaction, error) {
	return _HypERC20.Contract.UnenrollRemoteRouters(&_HypERC20.TransactOpts, _domains)
}

// HypERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the HypERC20 contract.
type HypERC20ApprovalIterator struct {
	Event *HypERC20Approval // Event containing the contract specifics and raw log

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
func (it *HypERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HypERC20Approval)
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
		it.Event = new(HypERC20Approval)
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
func (it *HypERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HypERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HypERC20Approval represents a Approval event raised by the HypERC20 contract.
type HypERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_HypERC20 *HypERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*HypERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _HypERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &HypERC20ApprovalIterator{contract: _HypERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_HypERC20 *HypERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *HypERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _HypERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HypERC20Approval)
				if err := _HypERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_HypERC20 *HypERC20Filterer) ParseApproval(log types.Log) (*HypERC20Approval, error) {
	event := new(HypERC20Approval)
	if err := _HypERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HypERC20InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the HypERC20 contract.
type HypERC20InitializedIterator struct {
	Event *HypERC20Initialized // Event containing the contract specifics and raw log

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
func (it *HypERC20InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HypERC20Initialized)
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
		it.Event = new(HypERC20Initialized)
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
func (it *HypERC20InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HypERC20InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HypERC20Initialized represents a Initialized event raised by the HypERC20 contract.
type HypERC20Initialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_HypERC20 *HypERC20Filterer) FilterInitialized(opts *bind.FilterOpts) (*HypERC20InitializedIterator, error) {

	logs, sub, err := _HypERC20.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &HypERC20InitializedIterator{contract: _HypERC20.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_HypERC20 *HypERC20Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *HypERC20Initialized) (event.Subscription, error) {

	logs, sub, err := _HypERC20.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HypERC20Initialized)
				if err := _HypERC20.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_HypERC20 *HypERC20Filterer) ParseInitialized(log types.Log) (*HypERC20Initialized, error) {
	event := new(HypERC20Initialized)
	if err := _HypERC20.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HypERC20OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the HypERC20 contract.
type HypERC20OwnershipTransferredIterator struct {
	Event *HypERC20OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *HypERC20OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HypERC20OwnershipTransferred)
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
		it.Event = new(HypERC20OwnershipTransferred)
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
func (it *HypERC20OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HypERC20OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HypERC20OwnershipTransferred represents a OwnershipTransferred event raised by the HypERC20 contract.
type HypERC20OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_HypERC20 *HypERC20Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*HypERC20OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _HypERC20.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &HypERC20OwnershipTransferredIterator{contract: _HypERC20.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_HypERC20 *HypERC20Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *HypERC20OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _HypERC20.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HypERC20OwnershipTransferred)
				if err := _HypERC20.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_HypERC20 *HypERC20Filterer) ParseOwnershipTransferred(log types.Log) (*HypERC20OwnershipTransferred, error) {
	event := new(HypERC20OwnershipTransferred)
	if err := _HypERC20.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HypERC20ReceivedTransferRemoteIterator is returned from FilterReceivedTransferRemote and is used to iterate over the raw logs and unpacked data for ReceivedTransferRemote events raised by the HypERC20 contract.
type HypERC20ReceivedTransferRemoteIterator struct {
	Event *HypERC20ReceivedTransferRemote // Event containing the contract specifics and raw log

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
func (it *HypERC20ReceivedTransferRemoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HypERC20ReceivedTransferRemote)
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
		it.Event = new(HypERC20ReceivedTransferRemote)
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
func (it *HypERC20ReceivedTransferRemoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HypERC20ReceivedTransferRemoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HypERC20ReceivedTransferRemote represents a ReceivedTransferRemote event raised by the HypERC20 contract.
type HypERC20ReceivedTransferRemote struct {
	Origin    uint32
	Recipient [32]byte
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterReceivedTransferRemote is a free log retrieval operation binding the contract event 0xba20947a325f450d232530e5f5fce293e7963499d5309a07cee84a269f2f15a6.
//
// Solidity: event ReceivedTransferRemote(uint32 indexed origin, bytes32 indexed recipient, uint256 amount)
func (_HypERC20 *HypERC20Filterer) FilterReceivedTransferRemote(opts *bind.FilterOpts, origin []uint32, recipient [][32]byte) (*HypERC20ReceivedTransferRemoteIterator, error) {

	var originRule []interface{}
	for _, originItem := range origin {
		originRule = append(originRule, originItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _HypERC20.contract.FilterLogs(opts, "ReceivedTransferRemote", originRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &HypERC20ReceivedTransferRemoteIterator{contract: _HypERC20.contract, event: "ReceivedTransferRemote", logs: logs, sub: sub}, nil
}

// WatchReceivedTransferRemote is a free log subscription operation binding the contract event 0xba20947a325f450d232530e5f5fce293e7963499d5309a07cee84a269f2f15a6.
//
// Solidity: event ReceivedTransferRemote(uint32 indexed origin, bytes32 indexed recipient, uint256 amount)
func (_HypERC20 *HypERC20Filterer) WatchReceivedTransferRemote(opts *bind.WatchOpts, sink chan<- *HypERC20ReceivedTransferRemote, origin []uint32, recipient [][32]byte) (event.Subscription, error) {

	var originRule []interface{}
	for _, originItem := range origin {
		originRule = append(originRule, originItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _HypERC20.contract.WatchLogs(opts, "ReceivedTransferRemote", originRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HypERC20ReceivedTransferRemote)
				if err := _HypERC20.contract.UnpackLog(event, "ReceivedTransferRemote", log); err != nil {
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

// ParseReceivedTransferRemote is a log parse operation binding the contract event 0xba20947a325f450d232530e5f5fce293e7963499d5309a07cee84a269f2f15a6.
//
// Solidity: event ReceivedTransferRemote(uint32 indexed origin, bytes32 indexed recipient, uint256 amount)
func (_HypERC20 *HypERC20Filterer) ParseReceivedTransferRemote(log types.Log) (*HypERC20ReceivedTransferRemote, error) {
	event := new(HypERC20ReceivedTransferRemote)
	if err := _HypERC20.contract.UnpackLog(event, "ReceivedTransferRemote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HypERC20SentTransferRemoteIterator is returned from FilterSentTransferRemote and is used to iterate over the raw logs and unpacked data for SentTransferRemote events raised by the HypERC20 contract.
type HypERC20SentTransferRemoteIterator struct {
	Event *HypERC20SentTransferRemote // Event containing the contract specifics and raw log

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
func (it *HypERC20SentTransferRemoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HypERC20SentTransferRemote)
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
		it.Event = new(HypERC20SentTransferRemote)
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
func (it *HypERC20SentTransferRemoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HypERC20SentTransferRemoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HypERC20SentTransferRemote represents a SentTransferRemote event raised by the HypERC20 contract.
type HypERC20SentTransferRemote struct {
	Destination uint32
	Recipient   [32]byte
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSentTransferRemote is a free log retrieval operation binding the contract event 0xd229aacb94204188fe8042965fa6b269c62dc5818b21238779ab64bdd17efeec.
//
// Solidity: event SentTransferRemote(uint32 indexed destination, bytes32 indexed recipient, uint256 amount)
func (_HypERC20 *HypERC20Filterer) FilterSentTransferRemote(opts *bind.FilterOpts, destination []uint32, recipient [][32]byte) (*HypERC20SentTransferRemoteIterator, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _HypERC20.contract.FilterLogs(opts, "SentTransferRemote", destinationRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &HypERC20SentTransferRemoteIterator{contract: _HypERC20.contract, event: "SentTransferRemote", logs: logs, sub: sub}, nil
}

// WatchSentTransferRemote is a free log subscription operation binding the contract event 0xd229aacb94204188fe8042965fa6b269c62dc5818b21238779ab64bdd17efeec.
//
// Solidity: event SentTransferRemote(uint32 indexed destination, bytes32 indexed recipient, uint256 amount)
func (_HypERC20 *HypERC20Filterer) WatchSentTransferRemote(opts *bind.WatchOpts, sink chan<- *HypERC20SentTransferRemote, destination []uint32, recipient [][32]byte) (event.Subscription, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _HypERC20.contract.WatchLogs(opts, "SentTransferRemote", destinationRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HypERC20SentTransferRemote)
				if err := _HypERC20.contract.UnpackLog(event, "SentTransferRemote", log); err != nil {
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

// ParseSentTransferRemote is a log parse operation binding the contract event 0xd229aacb94204188fe8042965fa6b269c62dc5818b21238779ab64bdd17efeec.
//
// Solidity: event SentTransferRemote(uint32 indexed destination, bytes32 indexed recipient, uint256 amount)
func (_HypERC20 *HypERC20Filterer) ParseSentTransferRemote(log types.Log) (*HypERC20SentTransferRemote, error) {
	event := new(HypERC20SentTransferRemote)
	if err := _HypERC20.contract.UnpackLog(event, "SentTransferRemote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HypERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the HypERC20 contract.
type HypERC20TransferIterator struct {
	Event *HypERC20Transfer // Event containing the contract specifics and raw log

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
func (it *HypERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HypERC20Transfer)
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
		it.Event = new(HypERC20Transfer)
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
func (it *HypERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HypERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HypERC20Transfer represents a Transfer event raised by the HypERC20 contract.
type HypERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_HypERC20 *HypERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*HypERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _HypERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &HypERC20TransferIterator{contract: _HypERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_HypERC20 *HypERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *HypERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _HypERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HypERC20Transfer)
				if err := _HypERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_HypERC20 *HypERC20Filterer) ParseTransfer(log types.Log) (*HypERC20Transfer, error) {
	event := new(HypERC20Transfer)
	if err := _HypERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
