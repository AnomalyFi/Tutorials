// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tokenrouter

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

// TokenRouterMetaData contains all meta data concerning the TokenRouter contract.
var TokenRouterMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"origin\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"recipient\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ReceivedTransferRemote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"destination\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"recipient\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SentTransferRemote\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"destinationGas\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domains\",\"outputs\":[{\"internalType\":\"uint32[]\",\"name\":\"\",\"type\":\"uint32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_router\",\"type\":\"bytes32\"}],\"name\":\"enrollRemoteRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32[]\",\"name\":\"_domains\",\"type\":\"uint32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_addresses\",\"type\":\"bytes32[]\"}],\"name\":\"enrollRemoteRouters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_origin\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_sender\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"handle\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hook\",\"outputs\":[{\"internalType\":\"contractIPostDispatchHook\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"interchainSecurityModule\",\"outputs\":[{\"internalType\":\"contractIInterchainSecurityModule\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mailbox\",\"outputs\":[{\"internalType\":\"contractIMailbox\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_destinationDomain\",\"type\":\"uint32\"}],\"name\":\"quoteGasPayment\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_gasPayment\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"}],\"name\":\"routers\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"}],\"name\":\"setDestinationGas\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"}],\"internalType\":\"structGasRouter.GasRouterConfig[]\",\"name\":\"gasConfigs\",\"type\":\"tuple[]\"}],\"name\":\"setDestinationGas\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_hook\",\"type\":\"address\"}],\"name\":\"setHook\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_module\",\"type\":\"address\"}],\"name\":\"setInterchainSecurityModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_destination\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_recipient\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_amountOrId\",\"type\":\"uint256\"}],\"name\":\"transferRemote\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"}],\"name\":\"unenrollRemoteRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32[]\",\"name\":\"_domains\",\"type\":\"uint32[]\"}],\"name\":\"unenrollRemoteRouters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TokenRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenRouterMetaData.ABI instead.
var TokenRouterABI = TokenRouterMetaData.ABI

// TokenRouter is an auto generated Go binding around an Ethereum contract.
type TokenRouter struct {
	TokenRouterCaller     // Read-only binding to the contract
	TokenRouterTransactor // Write-only binding to the contract
	TokenRouterFilterer   // Log filterer for contract events
}

// TokenRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenRouterSession struct {
	Contract     *TokenRouter      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenRouterCallerSession struct {
	Contract *TokenRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// TokenRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenRouterTransactorSession struct {
	Contract     *TokenRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// TokenRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenRouterRaw struct {
	Contract *TokenRouter // Generic contract binding to access the raw methods on
}

// TokenRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenRouterCallerRaw struct {
	Contract *TokenRouterCaller // Generic read-only contract binding to access the raw methods on
}

// TokenRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenRouterTransactorRaw struct {
	Contract *TokenRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenRouter creates a new instance of TokenRouter, bound to a specific deployed contract.
func NewTokenRouter(address common.Address, backend bind.ContractBackend) (*TokenRouter, error) {
	contract, err := bindTokenRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenRouter{TokenRouterCaller: TokenRouterCaller{contract: contract}, TokenRouterTransactor: TokenRouterTransactor{contract: contract}, TokenRouterFilterer: TokenRouterFilterer{contract: contract}}, nil
}

// NewTokenRouterCaller creates a new read-only instance of TokenRouter, bound to a specific deployed contract.
func NewTokenRouterCaller(address common.Address, caller bind.ContractCaller) (*TokenRouterCaller, error) {
	contract, err := bindTokenRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenRouterCaller{contract: contract}, nil
}

// NewTokenRouterTransactor creates a new write-only instance of TokenRouter, bound to a specific deployed contract.
func NewTokenRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenRouterTransactor, error) {
	contract, err := bindTokenRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenRouterTransactor{contract: contract}, nil
}

// NewTokenRouterFilterer creates a new log filterer instance of TokenRouter, bound to a specific deployed contract.
func NewTokenRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenRouterFilterer, error) {
	contract, err := bindTokenRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenRouterFilterer{contract: contract}, nil
}

// bindTokenRouter binds a generic wrapper to an already deployed contract.
func bindTokenRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TokenRouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenRouter *TokenRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenRouter.Contract.TokenRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenRouter *TokenRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenRouter.Contract.TokenRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenRouter *TokenRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenRouter.Contract.TokenRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenRouter *TokenRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenRouter *TokenRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenRouter *TokenRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenRouter.Contract.contract.Transact(opts, method, params...)
}

// DestinationGas is a free data retrieval call binding the contract method 0x775313a1.
//
// Solidity: function destinationGas(uint32 ) view returns(uint256)
func (_TokenRouter *TokenRouterCaller) DestinationGas(opts *bind.CallOpts, arg0 uint32) (*big.Int, error) {
	var out []interface{}
	err := _TokenRouter.contract.Call(opts, &out, "destinationGas", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DestinationGas is a free data retrieval call binding the contract method 0x775313a1.
//
// Solidity: function destinationGas(uint32 ) view returns(uint256)
func (_TokenRouter *TokenRouterSession) DestinationGas(arg0 uint32) (*big.Int, error) {
	return _TokenRouter.Contract.DestinationGas(&_TokenRouter.CallOpts, arg0)
}

// DestinationGas is a free data retrieval call binding the contract method 0x775313a1.
//
// Solidity: function destinationGas(uint32 ) view returns(uint256)
func (_TokenRouter *TokenRouterCallerSession) DestinationGas(arg0 uint32) (*big.Int, error) {
	return _TokenRouter.Contract.DestinationGas(&_TokenRouter.CallOpts, arg0)
}

// Domains is a free data retrieval call binding the contract method 0x440df4f4.
//
// Solidity: function domains() view returns(uint32[])
func (_TokenRouter *TokenRouterCaller) Domains(opts *bind.CallOpts) ([]uint32, error) {
	var out []interface{}
	err := _TokenRouter.contract.Call(opts, &out, "domains")

	if err != nil {
		return *new([]uint32), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint32)).(*[]uint32)

	return out0, err

}

// Domains is a free data retrieval call binding the contract method 0x440df4f4.
//
// Solidity: function domains() view returns(uint32[])
func (_TokenRouter *TokenRouterSession) Domains() ([]uint32, error) {
	return _TokenRouter.Contract.Domains(&_TokenRouter.CallOpts)
}

// Domains is a free data retrieval call binding the contract method 0x440df4f4.
//
// Solidity: function domains() view returns(uint32[])
func (_TokenRouter *TokenRouterCallerSession) Domains() ([]uint32, error) {
	return _TokenRouter.Contract.Domains(&_TokenRouter.CallOpts)
}

// Hook is a free data retrieval call binding the contract method 0x7f5a7c7b.
//
// Solidity: function hook() view returns(address)
func (_TokenRouter *TokenRouterCaller) Hook(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenRouter.contract.Call(opts, &out, "hook")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Hook is a free data retrieval call binding the contract method 0x7f5a7c7b.
//
// Solidity: function hook() view returns(address)
func (_TokenRouter *TokenRouterSession) Hook() (common.Address, error) {
	return _TokenRouter.Contract.Hook(&_TokenRouter.CallOpts)
}

// Hook is a free data retrieval call binding the contract method 0x7f5a7c7b.
//
// Solidity: function hook() view returns(address)
func (_TokenRouter *TokenRouterCallerSession) Hook() (common.Address, error) {
	return _TokenRouter.Contract.Hook(&_TokenRouter.CallOpts)
}

// InterchainSecurityModule is a free data retrieval call binding the contract method 0xde523cf3.
//
// Solidity: function interchainSecurityModule() view returns(address)
func (_TokenRouter *TokenRouterCaller) InterchainSecurityModule(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenRouter.contract.Call(opts, &out, "interchainSecurityModule")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InterchainSecurityModule is a free data retrieval call binding the contract method 0xde523cf3.
//
// Solidity: function interchainSecurityModule() view returns(address)
func (_TokenRouter *TokenRouterSession) InterchainSecurityModule() (common.Address, error) {
	return _TokenRouter.Contract.InterchainSecurityModule(&_TokenRouter.CallOpts)
}

// InterchainSecurityModule is a free data retrieval call binding the contract method 0xde523cf3.
//
// Solidity: function interchainSecurityModule() view returns(address)
func (_TokenRouter *TokenRouterCallerSession) InterchainSecurityModule() (common.Address, error) {
	return _TokenRouter.Contract.InterchainSecurityModule(&_TokenRouter.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_TokenRouter *TokenRouterCaller) LocalDomain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _TokenRouter.contract.Call(opts, &out, "localDomain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_TokenRouter *TokenRouterSession) LocalDomain() (uint32, error) {
	return _TokenRouter.Contract.LocalDomain(&_TokenRouter.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_TokenRouter *TokenRouterCallerSession) LocalDomain() (uint32, error) {
	return _TokenRouter.Contract.LocalDomain(&_TokenRouter.CallOpts)
}

// Mailbox is a free data retrieval call binding the contract method 0xd5438eae.
//
// Solidity: function mailbox() view returns(address)
func (_TokenRouter *TokenRouterCaller) Mailbox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenRouter.contract.Call(opts, &out, "mailbox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Mailbox is a free data retrieval call binding the contract method 0xd5438eae.
//
// Solidity: function mailbox() view returns(address)
func (_TokenRouter *TokenRouterSession) Mailbox() (common.Address, error) {
	return _TokenRouter.Contract.Mailbox(&_TokenRouter.CallOpts)
}

// Mailbox is a free data retrieval call binding the contract method 0xd5438eae.
//
// Solidity: function mailbox() view returns(address)
func (_TokenRouter *TokenRouterCallerSession) Mailbox() (common.Address, error) {
	return _TokenRouter.Contract.Mailbox(&_TokenRouter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TokenRouter *TokenRouterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenRouter.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TokenRouter *TokenRouterSession) Owner() (common.Address, error) {
	return _TokenRouter.Contract.Owner(&_TokenRouter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TokenRouter *TokenRouterCallerSession) Owner() (common.Address, error) {
	return _TokenRouter.Contract.Owner(&_TokenRouter.CallOpts)
}

// QuoteGasPayment is a free data retrieval call binding the contract method 0xf2ed8c53.
//
// Solidity: function quoteGasPayment(uint32 _destinationDomain) view returns(uint256 _gasPayment)
func (_TokenRouter *TokenRouterCaller) QuoteGasPayment(opts *bind.CallOpts, _destinationDomain uint32) (*big.Int, error) {
	var out []interface{}
	err := _TokenRouter.contract.Call(opts, &out, "quoteGasPayment", _destinationDomain)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuoteGasPayment is a free data retrieval call binding the contract method 0xf2ed8c53.
//
// Solidity: function quoteGasPayment(uint32 _destinationDomain) view returns(uint256 _gasPayment)
func (_TokenRouter *TokenRouterSession) QuoteGasPayment(_destinationDomain uint32) (*big.Int, error) {
	return _TokenRouter.Contract.QuoteGasPayment(&_TokenRouter.CallOpts, _destinationDomain)
}

// QuoteGasPayment is a free data retrieval call binding the contract method 0xf2ed8c53.
//
// Solidity: function quoteGasPayment(uint32 _destinationDomain) view returns(uint256 _gasPayment)
func (_TokenRouter *TokenRouterCallerSession) QuoteGasPayment(_destinationDomain uint32) (*big.Int, error) {
	return _TokenRouter.Contract.QuoteGasPayment(&_TokenRouter.CallOpts, _destinationDomain)
}

// Routers is a free data retrieval call binding the contract method 0x2ead72f6.
//
// Solidity: function routers(uint32 _domain) view returns(bytes32)
func (_TokenRouter *TokenRouterCaller) Routers(opts *bind.CallOpts, _domain uint32) ([32]byte, error) {
	var out []interface{}
	err := _TokenRouter.contract.Call(opts, &out, "routers", _domain)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Routers is a free data retrieval call binding the contract method 0x2ead72f6.
//
// Solidity: function routers(uint32 _domain) view returns(bytes32)
func (_TokenRouter *TokenRouterSession) Routers(_domain uint32) ([32]byte, error) {
	return _TokenRouter.Contract.Routers(&_TokenRouter.CallOpts, _domain)
}

// Routers is a free data retrieval call binding the contract method 0x2ead72f6.
//
// Solidity: function routers(uint32 _domain) view returns(bytes32)
func (_TokenRouter *TokenRouterCallerSession) Routers(_domain uint32) ([32]byte, error) {
	return _TokenRouter.Contract.Routers(&_TokenRouter.CallOpts, _domain)
}

// BalanceOf is a paid mutator transaction binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) returns(uint256)
func (_TokenRouter *TokenRouterTransactor) BalanceOf(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _TokenRouter.contract.Transact(opts, "balanceOf", account)
}

// BalanceOf is a paid mutator transaction binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) returns(uint256)
func (_TokenRouter *TokenRouterSession) BalanceOf(account common.Address) (*types.Transaction, error) {
	return _TokenRouter.Contract.BalanceOf(&_TokenRouter.TransactOpts, account)
}

// BalanceOf is a paid mutator transaction binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) returns(uint256)
func (_TokenRouter *TokenRouterTransactorSession) BalanceOf(account common.Address) (*types.Transaction, error) {
	return _TokenRouter.Contract.BalanceOf(&_TokenRouter.TransactOpts, account)
}

// EnrollRemoteRouter is a paid mutator transaction binding the contract method 0xb49c53a7.
//
// Solidity: function enrollRemoteRouter(uint32 _domain, bytes32 _router) returns()
func (_TokenRouter *TokenRouterTransactor) EnrollRemoteRouter(opts *bind.TransactOpts, _domain uint32, _router [32]byte) (*types.Transaction, error) {
	return _TokenRouter.contract.Transact(opts, "enrollRemoteRouter", _domain, _router)
}

// EnrollRemoteRouter is a paid mutator transaction binding the contract method 0xb49c53a7.
//
// Solidity: function enrollRemoteRouter(uint32 _domain, bytes32 _router) returns()
func (_TokenRouter *TokenRouterSession) EnrollRemoteRouter(_domain uint32, _router [32]byte) (*types.Transaction, error) {
	return _TokenRouter.Contract.EnrollRemoteRouter(&_TokenRouter.TransactOpts, _domain, _router)
}

// EnrollRemoteRouter is a paid mutator transaction binding the contract method 0xb49c53a7.
//
// Solidity: function enrollRemoteRouter(uint32 _domain, bytes32 _router) returns()
func (_TokenRouter *TokenRouterTransactorSession) EnrollRemoteRouter(_domain uint32, _router [32]byte) (*types.Transaction, error) {
	return _TokenRouter.Contract.EnrollRemoteRouter(&_TokenRouter.TransactOpts, _domain, _router)
}

// EnrollRemoteRouters is a paid mutator transaction binding the contract method 0xe9198bf9.
//
// Solidity: function enrollRemoteRouters(uint32[] _domains, bytes32[] _addresses) returns()
func (_TokenRouter *TokenRouterTransactor) EnrollRemoteRouters(opts *bind.TransactOpts, _domains []uint32, _addresses [][32]byte) (*types.Transaction, error) {
	return _TokenRouter.contract.Transact(opts, "enrollRemoteRouters", _domains, _addresses)
}

// EnrollRemoteRouters is a paid mutator transaction binding the contract method 0xe9198bf9.
//
// Solidity: function enrollRemoteRouters(uint32[] _domains, bytes32[] _addresses) returns()
func (_TokenRouter *TokenRouterSession) EnrollRemoteRouters(_domains []uint32, _addresses [][32]byte) (*types.Transaction, error) {
	return _TokenRouter.Contract.EnrollRemoteRouters(&_TokenRouter.TransactOpts, _domains, _addresses)
}

// EnrollRemoteRouters is a paid mutator transaction binding the contract method 0xe9198bf9.
//
// Solidity: function enrollRemoteRouters(uint32[] _domains, bytes32[] _addresses) returns()
func (_TokenRouter *TokenRouterTransactorSession) EnrollRemoteRouters(_domains []uint32, _addresses [][32]byte) (*types.Transaction, error) {
	return _TokenRouter.Contract.EnrollRemoteRouters(&_TokenRouter.TransactOpts, _domains, _addresses)
}

// Handle is a paid mutator transaction binding the contract method 0x56d5d475.
//
// Solidity: function handle(uint32 _origin, bytes32 _sender, bytes _message) payable returns()
func (_TokenRouter *TokenRouterTransactor) Handle(opts *bind.TransactOpts, _origin uint32, _sender [32]byte, _message []byte) (*types.Transaction, error) {
	return _TokenRouter.contract.Transact(opts, "handle", _origin, _sender, _message)
}

// Handle is a paid mutator transaction binding the contract method 0x56d5d475.
//
// Solidity: function handle(uint32 _origin, bytes32 _sender, bytes _message) payable returns()
func (_TokenRouter *TokenRouterSession) Handle(_origin uint32, _sender [32]byte, _message []byte) (*types.Transaction, error) {
	return _TokenRouter.Contract.Handle(&_TokenRouter.TransactOpts, _origin, _sender, _message)
}

// Handle is a paid mutator transaction binding the contract method 0x56d5d475.
//
// Solidity: function handle(uint32 _origin, bytes32 _sender, bytes _message) payable returns()
func (_TokenRouter *TokenRouterTransactorSession) Handle(_origin uint32, _sender [32]byte, _message []byte) (*types.Transaction, error) {
	return _TokenRouter.Contract.Handle(&_TokenRouter.TransactOpts, _origin, _sender, _message)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TokenRouter *TokenRouterTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenRouter.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TokenRouter *TokenRouterSession) RenounceOwnership() (*types.Transaction, error) {
	return _TokenRouter.Contract.RenounceOwnership(&_TokenRouter.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TokenRouter *TokenRouterTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TokenRouter.Contract.RenounceOwnership(&_TokenRouter.TransactOpts)
}

// SetDestinationGas is a paid mutator transaction binding the contract method 0x49d462ef.
//
// Solidity: function setDestinationGas(uint32 domain, uint256 gas) returns()
func (_TokenRouter *TokenRouterTransactor) SetDestinationGas(opts *bind.TransactOpts, domain uint32, gas *big.Int) (*types.Transaction, error) {
	return _TokenRouter.contract.Transact(opts, "setDestinationGas", domain, gas)
}

// SetDestinationGas is a paid mutator transaction binding the contract method 0x49d462ef.
//
// Solidity: function setDestinationGas(uint32 domain, uint256 gas) returns()
func (_TokenRouter *TokenRouterSession) SetDestinationGas(domain uint32, gas *big.Int) (*types.Transaction, error) {
	return _TokenRouter.Contract.SetDestinationGas(&_TokenRouter.TransactOpts, domain, gas)
}

// SetDestinationGas is a paid mutator transaction binding the contract method 0x49d462ef.
//
// Solidity: function setDestinationGas(uint32 domain, uint256 gas) returns()
func (_TokenRouter *TokenRouterTransactorSession) SetDestinationGas(domain uint32, gas *big.Int) (*types.Transaction, error) {
	return _TokenRouter.Contract.SetDestinationGas(&_TokenRouter.TransactOpts, domain, gas)
}

// SetDestinationGas0 is a paid mutator transaction binding the contract method 0xb1bd6436.
//
// Solidity: function setDestinationGas((uint32,uint256)[] gasConfigs) returns()
func (_TokenRouter *TokenRouterTransactor) SetDestinationGas0(opts *bind.TransactOpts, gasConfigs []GasRouterGasRouterConfig) (*types.Transaction, error) {
	return _TokenRouter.contract.Transact(opts, "setDestinationGas0", gasConfigs)
}

// SetDestinationGas0 is a paid mutator transaction binding the contract method 0xb1bd6436.
//
// Solidity: function setDestinationGas((uint32,uint256)[] gasConfigs) returns()
func (_TokenRouter *TokenRouterSession) SetDestinationGas0(gasConfigs []GasRouterGasRouterConfig) (*types.Transaction, error) {
	return _TokenRouter.Contract.SetDestinationGas0(&_TokenRouter.TransactOpts, gasConfigs)
}

// SetDestinationGas0 is a paid mutator transaction binding the contract method 0xb1bd6436.
//
// Solidity: function setDestinationGas((uint32,uint256)[] gasConfigs) returns()
func (_TokenRouter *TokenRouterTransactorSession) SetDestinationGas0(gasConfigs []GasRouterGasRouterConfig) (*types.Transaction, error) {
	return _TokenRouter.Contract.SetDestinationGas0(&_TokenRouter.TransactOpts, gasConfigs)
}

// SetHook is a paid mutator transaction binding the contract method 0x3dfd3873.
//
// Solidity: function setHook(address _hook) returns()
func (_TokenRouter *TokenRouterTransactor) SetHook(opts *bind.TransactOpts, _hook common.Address) (*types.Transaction, error) {
	return _TokenRouter.contract.Transact(opts, "setHook", _hook)
}

// SetHook is a paid mutator transaction binding the contract method 0x3dfd3873.
//
// Solidity: function setHook(address _hook) returns()
func (_TokenRouter *TokenRouterSession) SetHook(_hook common.Address) (*types.Transaction, error) {
	return _TokenRouter.Contract.SetHook(&_TokenRouter.TransactOpts, _hook)
}

// SetHook is a paid mutator transaction binding the contract method 0x3dfd3873.
//
// Solidity: function setHook(address _hook) returns()
func (_TokenRouter *TokenRouterTransactorSession) SetHook(_hook common.Address) (*types.Transaction, error) {
	return _TokenRouter.Contract.SetHook(&_TokenRouter.TransactOpts, _hook)
}

// SetInterchainSecurityModule is a paid mutator transaction binding the contract method 0x0e72cc06.
//
// Solidity: function setInterchainSecurityModule(address _module) returns()
func (_TokenRouter *TokenRouterTransactor) SetInterchainSecurityModule(opts *bind.TransactOpts, _module common.Address) (*types.Transaction, error) {
	return _TokenRouter.contract.Transact(opts, "setInterchainSecurityModule", _module)
}

// SetInterchainSecurityModule is a paid mutator transaction binding the contract method 0x0e72cc06.
//
// Solidity: function setInterchainSecurityModule(address _module) returns()
func (_TokenRouter *TokenRouterSession) SetInterchainSecurityModule(_module common.Address) (*types.Transaction, error) {
	return _TokenRouter.Contract.SetInterchainSecurityModule(&_TokenRouter.TransactOpts, _module)
}

// SetInterchainSecurityModule is a paid mutator transaction binding the contract method 0x0e72cc06.
//
// Solidity: function setInterchainSecurityModule(address _module) returns()
func (_TokenRouter *TokenRouterTransactorSession) SetInterchainSecurityModule(_module common.Address) (*types.Transaction, error) {
	return _TokenRouter.Contract.SetInterchainSecurityModule(&_TokenRouter.TransactOpts, _module)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TokenRouter *TokenRouterTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TokenRouter.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TokenRouter *TokenRouterSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TokenRouter.Contract.TransferOwnership(&_TokenRouter.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TokenRouter *TokenRouterTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TokenRouter.Contract.TransferOwnership(&_TokenRouter.TransactOpts, newOwner)
}

// TransferRemote is a paid mutator transaction binding the contract method 0x81b4e8b4.
//
// Solidity: function transferRemote(uint32 _destination, bytes32 _recipient, uint256 _amountOrId) payable returns(bytes32 messageId)
func (_TokenRouter *TokenRouterTransactor) TransferRemote(opts *bind.TransactOpts, _destination uint32, _recipient [32]byte, _amountOrId *big.Int) (*types.Transaction, error) {
	return _TokenRouter.contract.Transact(opts, "transferRemote", _destination, _recipient, _amountOrId)
}

// TransferRemote is a paid mutator transaction binding the contract method 0x81b4e8b4.
//
// Solidity: function transferRemote(uint32 _destination, bytes32 _recipient, uint256 _amountOrId) payable returns(bytes32 messageId)
func (_TokenRouter *TokenRouterSession) TransferRemote(_destination uint32, _recipient [32]byte, _amountOrId *big.Int) (*types.Transaction, error) {
	return _TokenRouter.Contract.TransferRemote(&_TokenRouter.TransactOpts, _destination, _recipient, _amountOrId)
}

// TransferRemote is a paid mutator transaction binding the contract method 0x81b4e8b4.
//
// Solidity: function transferRemote(uint32 _destination, bytes32 _recipient, uint256 _amountOrId) payable returns(bytes32 messageId)
func (_TokenRouter *TokenRouterTransactorSession) TransferRemote(_destination uint32, _recipient [32]byte, _amountOrId *big.Int) (*types.Transaction, error) {
	return _TokenRouter.Contract.TransferRemote(&_TokenRouter.TransactOpts, _destination, _recipient, _amountOrId)
}

// UnenrollRemoteRouter is a paid mutator transaction binding the contract method 0xefae508a.
//
// Solidity: function unenrollRemoteRouter(uint32 _domain) returns()
func (_TokenRouter *TokenRouterTransactor) UnenrollRemoteRouter(opts *bind.TransactOpts, _domain uint32) (*types.Transaction, error) {
	return _TokenRouter.contract.Transact(opts, "unenrollRemoteRouter", _domain)
}

// UnenrollRemoteRouter is a paid mutator transaction binding the contract method 0xefae508a.
//
// Solidity: function unenrollRemoteRouter(uint32 _domain) returns()
func (_TokenRouter *TokenRouterSession) UnenrollRemoteRouter(_domain uint32) (*types.Transaction, error) {
	return _TokenRouter.Contract.UnenrollRemoteRouter(&_TokenRouter.TransactOpts, _domain)
}

// UnenrollRemoteRouter is a paid mutator transaction binding the contract method 0xefae508a.
//
// Solidity: function unenrollRemoteRouter(uint32 _domain) returns()
func (_TokenRouter *TokenRouterTransactorSession) UnenrollRemoteRouter(_domain uint32) (*types.Transaction, error) {
	return _TokenRouter.Contract.UnenrollRemoteRouter(&_TokenRouter.TransactOpts, _domain)
}

// UnenrollRemoteRouters is a paid mutator transaction binding the contract method 0x71a15b38.
//
// Solidity: function unenrollRemoteRouters(uint32[] _domains) returns()
func (_TokenRouter *TokenRouterTransactor) UnenrollRemoteRouters(opts *bind.TransactOpts, _domains []uint32) (*types.Transaction, error) {
	return _TokenRouter.contract.Transact(opts, "unenrollRemoteRouters", _domains)
}

// UnenrollRemoteRouters is a paid mutator transaction binding the contract method 0x71a15b38.
//
// Solidity: function unenrollRemoteRouters(uint32[] _domains) returns()
func (_TokenRouter *TokenRouterSession) UnenrollRemoteRouters(_domains []uint32) (*types.Transaction, error) {
	return _TokenRouter.Contract.UnenrollRemoteRouters(&_TokenRouter.TransactOpts, _domains)
}

// UnenrollRemoteRouters is a paid mutator transaction binding the contract method 0x71a15b38.
//
// Solidity: function unenrollRemoteRouters(uint32[] _domains) returns()
func (_TokenRouter *TokenRouterTransactorSession) UnenrollRemoteRouters(_domains []uint32) (*types.Transaction, error) {
	return _TokenRouter.Contract.UnenrollRemoteRouters(&_TokenRouter.TransactOpts, _domains)
}

// TokenRouterInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the TokenRouter contract.
type TokenRouterInitializedIterator struct {
	Event *TokenRouterInitialized // Event containing the contract specifics and raw log

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
func (it *TokenRouterInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenRouterInitialized)
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
		it.Event = new(TokenRouterInitialized)
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
func (it *TokenRouterInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenRouterInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenRouterInitialized represents a Initialized event raised by the TokenRouter contract.
type TokenRouterInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_TokenRouter *TokenRouterFilterer) FilterInitialized(opts *bind.FilterOpts) (*TokenRouterInitializedIterator, error) {

	logs, sub, err := _TokenRouter.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &TokenRouterInitializedIterator{contract: _TokenRouter.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_TokenRouter *TokenRouterFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *TokenRouterInitialized) (event.Subscription, error) {

	logs, sub, err := _TokenRouter.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenRouterInitialized)
				if err := _TokenRouter.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_TokenRouter *TokenRouterFilterer) ParseInitialized(log types.Log) (*TokenRouterInitialized, error) {
	event := new(TokenRouterInitialized)
	if err := _TokenRouter.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenRouterOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TokenRouter contract.
type TokenRouterOwnershipTransferredIterator struct {
	Event *TokenRouterOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TokenRouterOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenRouterOwnershipTransferred)
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
		it.Event = new(TokenRouterOwnershipTransferred)
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
func (it *TokenRouterOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenRouterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenRouterOwnershipTransferred represents a OwnershipTransferred event raised by the TokenRouter contract.
type TokenRouterOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TokenRouter *TokenRouterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TokenRouterOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TokenRouter.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TokenRouterOwnershipTransferredIterator{contract: _TokenRouter.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TokenRouter *TokenRouterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TokenRouterOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TokenRouter.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenRouterOwnershipTransferred)
				if err := _TokenRouter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_TokenRouter *TokenRouterFilterer) ParseOwnershipTransferred(log types.Log) (*TokenRouterOwnershipTransferred, error) {
	event := new(TokenRouterOwnershipTransferred)
	if err := _TokenRouter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenRouterReceivedTransferRemoteIterator is returned from FilterReceivedTransferRemote and is used to iterate over the raw logs and unpacked data for ReceivedTransferRemote events raised by the TokenRouter contract.
type TokenRouterReceivedTransferRemoteIterator struct {
	Event *TokenRouterReceivedTransferRemote // Event containing the contract specifics and raw log

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
func (it *TokenRouterReceivedTransferRemoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenRouterReceivedTransferRemote)
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
		it.Event = new(TokenRouterReceivedTransferRemote)
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
func (it *TokenRouterReceivedTransferRemoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenRouterReceivedTransferRemoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenRouterReceivedTransferRemote represents a ReceivedTransferRemote event raised by the TokenRouter contract.
type TokenRouterReceivedTransferRemote struct {
	Origin    uint32
	Recipient [32]byte
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterReceivedTransferRemote is a free log retrieval operation binding the contract event 0xba20947a325f450d232530e5f5fce293e7963499d5309a07cee84a269f2f15a6.
//
// Solidity: event ReceivedTransferRemote(uint32 indexed origin, bytes32 indexed recipient, uint256 amount)
func (_TokenRouter *TokenRouterFilterer) FilterReceivedTransferRemote(opts *bind.FilterOpts, origin []uint32, recipient [][32]byte) (*TokenRouterReceivedTransferRemoteIterator, error) {

	var originRule []interface{}
	for _, originItem := range origin {
		originRule = append(originRule, originItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _TokenRouter.contract.FilterLogs(opts, "ReceivedTransferRemote", originRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &TokenRouterReceivedTransferRemoteIterator{contract: _TokenRouter.contract, event: "ReceivedTransferRemote", logs: logs, sub: sub}, nil
}

// WatchReceivedTransferRemote is a free log subscription operation binding the contract event 0xba20947a325f450d232530e5f5fce293e7963499d5309a07cee84a269f2f15a6.
//
// Solidity: event ReceivedTransferRemote(uint32 indexed origin, bytes32 indexed recipient, uint256 amount)
func (_TokenRouter *TokenRouterFilterer) WatchReceivedTransferRemote(opts *bind.WatchOpts, sink chan<- *TokenRouterReceivedTransferRemote, origin []uint32, recipient [][32]byte) (event.Subscription, error) {

	var originRule []interface{}
	for _, originItem := range origin {
		originRule = append(originRule, originItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _TokenRouter.contract.WatchLogs(opts, "ReceivedTransferRemote", originRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenRouterReceivedTransferRemote)
				if err := _TokenRouter.contract.UnpackLog(event, "ReceivedTransferRemote", log); err != nil {
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
func (_TokenRouter *TokenRouterFilterer) ParseReceivedTransferRemote(log types.Log) (*TokenRouterReceivedTransferRemote, error) {
	event := new(TokenRouterReceivedTransferRemote)
	if err := _TokenRouter.contract.UnpackLog(event, "ReceivedTransferRemote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenRouterSentTransferRemoteIterator is returned from FilterSentTransferRemote and is used to iterate over the raw logs and unpacked data for SentTransferRemote events raised by the TokenRouter contract.
type TokenRouterSentTransferRemoteIterator struct {
	Event *TokenRouterSentTransferRemote // Event containing the contract specifics and raw log

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
func (it *TokenRouterSentTransferRemoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenRouterSentTransferRemote)
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
		it.Event = new(TokenRouterSentTransferRemote)
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
func (it *TokenRouterSentTransferRemoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenRouterSentTransferRemoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenRouterSentTransferRemote represents a SentTransferRemote event raised by the TokenRouter contract.
type TokenRouterSentTransferRemote struct {
	Destination uint32
	Recipient   [32]byte
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSentTransferRemote is a free log retrieval operation binding the contract event 0xd229aacb94204188fe8042965fa6b269c62dc5818b21238779ab64bdd17efeec.
//
// Solidity: event SentTransferRemote(uint32 indexed destination, bytes32 indexed recipient, uint256 amount)
func (_TokenRouter *TokenRouterFilterer) FilterSentTransferRemote(opts *bind.FilterOpts, destination []uint32, recipient [][32]byte) (*TokenRouterSentTransferRemoteIterator, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _TokenRouter.contract.FilterLogs(opts, "SentTransferRemote", destinationRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &TokenRouterSentTransferRemoteIterator{contract: _TokenRouter.contract, event: "SentTransferRemote", logs: logs, sub: sub}, nil
}

// WatchSentTransferRemote is a free log subscription operation binding the contract event 0xd229aacb94204188fe8042965fa6b269c62dc5818b21238779ab64bdd17efeec.
//
// Solidity: event SentTransferRemote(uint32 indexed destination, bytes32 indexed recipient, uint256 amount)
func (_TokenRouter *TokenRouterFilterer) WatchSentTransferRemote(opts *bind.WatchOpts, sink chan<- *TokenRouterSentTransferRemote, destination []uint32, recipient [][32]byte) (event.Subscription, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _TokenRouter.contract.WatchLogs(opts, "SentTransferRemote", destinationRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenRouterSentTransferRemote)
				if err := _TokenRouter.contract.UnpackLog(event, "SentTransferRemote", log); err != nil {
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
func (_TokenRouter *TokenRouterFilterer) ParseSentTransferRemote(log types.Log) (*TokenRouterSentTransferRemote, error) {
	event := new(TokenRouterSentTransferRemote)
	if err := _TokenRouter.contract.UnpackLog(event, "SentTransferRemote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
