// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mailbox

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

// MailboxMetaData contains all meta data concerning the Mailbox contract.
var MailboxMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_localDomain\",\"type\":\"uint32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"hook\",\"type\":\"address\"}],\"name\":\"DefaultHookSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"DefaultIsmSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"destination\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"recipient\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"Dispatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"name\":\"DispatchId\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"origin\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"Process\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"name\":\"ProcessId\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"hook\",\"type\":\"address\"}],\"name\":\"RequiredHookSet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultHook\",\"outputs\":[{\"internalType\":\"contractIPostDispatchHook\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultIsm\",\"outputs\":[{\"internalType\":\"contractIInterchainSecurityModule\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"delivered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deployedBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destinationDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"recipientAddress\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"messageBody\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"},{\"internalType\":\"contractIPostDispatchHook\",\"name\":\"hook\",\"type\":\"address\"}],\"name\":\"dispatch\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destinationDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"recipientAddress\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"messageBody\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"hookMetadata\",\"type\":\"bytes\"}],\"name\":\"dispatch\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_destinationDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_recipientAddress\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_messageBody\",\"type\":\"bytes\"}],\"name\":\"dispatch\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_defaultIsm\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_defaultHook\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_requiredHook\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestDispatchedId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nonce\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_metadata\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"process\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"processedAt\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"processor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destinationDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"recipientAddress\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"messageBody\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"},{\"internalType\":\"contractIPostDispatchHook\",\"name\":\"hook\",\"type\":\"address\"}],\"name\":\"quoteDispatch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destinationDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"recipientAddress\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"messageBody\",\"type\":\"bytes\"}],\"name\":\"quoteDispatch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destinationDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"recipientAddress\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"messageBody\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"defaultHookMetadata\",\"type\":\"bytes\"}],\"name\":\"quoteDispatch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"recipientIsm\",\"outputs\":[{\"internalType\":\"contractIInterchainSecurityModule\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requiredHook\",\"outputs\":[{\"internalType\":\"contractIPostDispatchHook\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_hook\",\"type\":\"address\"}],\"name\":\"setDefaultHook\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_module\",\"type\":\"address\"}],\"name\":\"setDefaultIsm\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_hook\",\"type\":\"address\"}],\"name\":\"setRequiredHook\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MailboxABI is the input ABI used to generate the binding from.
// Deprecated: Use MailboxMetaData.ABI instead.
var MailboxABI = MailboxMetaData.ABI

// Mailbox is an auto generated Go binding around an Ethereum contract.
type Mailbox struct {
	MailboxCaller     // Read-only binding to the contract
	MailboxTransactor // Write-only binding to the contract
	MailboxFilterer   // Log filterer for contract events
}

// MailboxCaller is an auto generated read-only Go binding around an Ethereum contract.
type MailboxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MailboxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MailboxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MailboxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MailboxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MailboxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MailboxSession struct {
	Contract     *Mailbox          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MailboxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MailboxCallerSession struct {
	Contract *MailboxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// MailboxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MailboxTransactorSession struct {
	Contract     *MailboxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// MailboxRaw is an auto generated low-level Go binding around an Ethereum contract.
type MailboxRaw struct {
	Contract *Mailbox // Generic contract binding to access the raw methods on
}

// MailboxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MailboxCallerRaw struct {
	Contract *MailboxCaller // Generic read-only contract binding to access the raw methods on
}

// MailboxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MailboxTransactorRaw struct {
	Contract *MailboxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMailbox creates a new instance of Mailbox, bound to a specific deployed contract.
func NewMailbox(address common.Address, backend bind.ContractBackend) (*Mailbox, error) {
	contract, err := bindMailbox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Mailbox{MailboxCaller: MailboxCaller{contract: contract}, MailboxTransactor: MailboxTransactor{contract: contract}, MailboxFilterer: MailboxFilterer{contract: contract}}, nil
}

// NewMailboxCaller creates a new read-only instance of Mailbox, bound to a specific deployed contract.
func NewMailboxCaller(address common.Address, caller bind.ContractCaller) (*MailboxCaller, error) {
	contract, err := bindMailbox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MailboxCaller{contract: contract}, nil
}

// NewMailboxTransactor creates a new write-only instance of Mailbox, bound to a specific deployed contract.
func NewMailboxTransactor(address common.Address, transactor bind.ContractTransactor) (*MailboxTransactor, error) {
	contract, err := bindMailbox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MailboxTransactor{contract: contract}, nil
}

// NewMailboxFilterer creates a new log filterer instance of Mailbox, bound to a specific deployed contract.
func NewMailboxFilterer(address common.Address, filterer bind.ContractFilterer) (*MailboxFilterer, error) {
	contract, err := bindMailbox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MailboxFilterer{contract: contract}, nil
}

// bindMailbox binds a generic wrapper to an already deployed contract.
func bindMailbox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MailboxMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mailbox *MailboxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mailbox.Contract.MailboxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mailbox *MailboxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mailbox.Contract.MailboxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mailbox *MailboxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mailbox.Contract.MailboxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mailbox *MailboxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mailbox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mailbox *MailboxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mailbox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mailbox *MailboxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mailbox.Contract.contract.Transact(opts, method, params...)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint8)
func (_Mailbox *MailboxCaller) VERSION(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Mailbox.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint8)
func (_Mailbox *MailboxSession) VERSION() (uint8, error) {
	return _Mailbox.Contract.VERSION(&_Mailbox.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint8)
func (_Mailbox *MailboxCallerSession) VERSION() (uint8, error) {
	return _Mailbox.Contract.VERSION(&_Mailbox.CallOpts)
}

// DefaultHook is a free data retrieval call binding the contract method 0x3d1250b7.
//
// Solidity: function defaultHook() view returns(address)
func (_Mailbox *MailboxCaller) DefaultHook(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mailbox.contract.Call(opts, &out, "defaultHook")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DefaultHook is a free data retrieval call binding the contract method 0x3d1250b7.
//
// Solidity: function defaultHook() view returns(address)
func (_Mailbox *MailboxSession) DefaultHook() (common.Address, error) {
	return _Mailbox.Contract.DefaultHook(&_Mailbox.CallOpts)
}

// DefaultHook is a free data retrieval call binding the contract method 0x3d1250b7.
//
// Solidity: function defaultHook() view returns(address)
func (_Mailbox *MailboxCallerSession) DefaultHook() (common.Address, error) {
	return _Mailbox.Contract.DefaultHook(&_Mailbox.CallOpts)
}

// DefaultIsm is a free data retrieval call binding the contract method 0x6e5f516e.
//
// Solidity: function defaultIsm() view returns(address)
func (_Mailbox *MailboxCaller) DefaultIsm(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mailbox.contract.Call(opts, &out, "defaultIsm")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DefaultIsm is a free data retrieval call binding the contract method 0x6e5f516e.
//
// Solidity: function defaultIsm() view returns(address)
func (_Mailbox *MailboxSession) DefaultIsm() (common.Address, error) {
	return _Mailbox.Contract.DefaultIsm(&_Mailbox.CallOpts)
}

// DefaultIsm is a free data retrieval call binding the contract method 0x6e5f516e.
//
// Solidity: function defaultIsm() view returns(address)
func (_Mailbox *MailboxCallerSession) DefaultIsm() (common.Address, error) {
	return _Mailbox.Contract.DefaultIsm(&_Mailbox.CallOpts)
}

// Delivered is a free data retrieval call binding the contract method 0xe495f1d4.
//
// Solidity: function delivered(bytes32 _id) view returns(bool)
func (_Mailbox *MailboxCaller) Delivered(opts *bind.CallOpts, _id [32]byte) (bool, error) {
	var out []interface{}
	err := _Mailbox.contract.Call(opts, &out, "delivered", _id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Delivered is a free data retrieval call binding the contract method 0xe495f1d4.
//
// Solidity: function delivered(bytes32 _id) view returns(bool)
func (_Mailbox *MailboxSession) Delivered(_id [32]byte) (bool, error) {
	return _Mailbox.Contract.Delivered(&_Mailbox.CallOpts, _id)
}

// Delivered is a free data retrieval call binding the contract method 0xe495f1d4.
//
// Solidity: function delivered(bytes32 _id) view returns(bool)
func (_Mailbox *MailboxCallerSession) Delivered(_id [32]byte) (bool, error) {
	return _Mailbox.Contract.Delivered(&_Mailbox.CallOpts, _id)
}

// DeployedBlock is a free data retrieval call binding the contract method 0x82ea7bfe.
//
// Solidity: function deployedBlock() view returns(uint256)
func (_Mailbox *MailboxCaller) DeployedBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mailbox.contract.Call(opts, &out, "deployedBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DeployedBlock is a free data retrieval call binding the contract method 0x82ea7bfe.
//
// Solidity: function deployedBlock() view returns(uint256)
func (_Mailbox *MailboxSession) DeployedBlock() (*big.Int, error) {
	return _Mailbox.Contract.DeployedBlock(&_Mailbox.CallOpts)
}

// DeployedBlock is a free data retrieval call binding the contract method 0x82ea7bfe.
//
// Solidity: function deployedBlock() view returns(uint256)
func (_Mailbox *MailboxCallerSession) DeployedBlock() (*big.Int, error) {
	return _Mailbox.Contract.DeployedBlock(&_Mailbox.CallOpts)
}

// LatestDispatchedId is a free data retrieval call binding the contract method 0x134fbb4f.
//
// Solidity: function latestDispatchedId() view returns(bytes32)
func (_Mailbox *MailboxCaller) LatestDispatchedId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Mailbox.contract.Call(opts, &out, "latestDispatchedId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LatestDispatchedId is a free data retrieval call binding the contract method 0x134fbb4f.
//
// Solidity: function latestDispatchedId() view returns(bytes32)
func (_Mailbox *MailboxSession) LatestDispatchedId() ([32]byte, error) {
	return _Mailbox.Contract.LatestDispatchedId(&_Mailbox.CallOpts)
}

// LatestDispatchedId is a free data retrieval call binding the contract method 0x134fbb4f.
//
// Solidity: function latestDispatchedId() view returns(bytes32)
func (_Mailbox *MailboxCallerSession) LatestDispatchedId() ([32]byte, error) {
	return _Mailbox.Contract.LatestDispatchedId(&_Mailbox.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_Mailbox *MailboxCaller) LocalDomain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Mailbox.contract.Call(opts, &out, "localDomain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_Mailbox *MailboxSession) LocalDomain() (uint32, error) {
	return _Mailbox.Contract.LocalDomain(&_Mailbox.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_Mailbox *MailboxCallerSession) LocalDomain() (uint32, error) {
	return _Mailbox.Contract.LocalDomain(&_Mailbox.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint32)
func (_Mailbox *MailboxCaller) Nonce(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Mailbox.contract.Call(opts, &out, "nonce")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint32)
func (_Mailbox *MailboxSession) Nonce() (uint32, error) {
	return _Mailbox.Contract.Nonce(&_Mailbox.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint32)
func (_Mailbox *MailboxCallerSession) Nonce() (uint32, error) {
	return _Mailbox.Contract.Nonce(&_Mailbox.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mailbox *MailboxCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mailbox.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mailbox *MailboxSession) Owner() (common.Address, error) {
	return _Mailbox.Contract.Owner(&_Mailbox.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mailbox *MailboxCallerSession) Owner() (common.Address, error) {
	return _Mailbox.Contract.Owner(&_Mailbox.CallOpts)
}

// ProcessedAt is a free data retrieval call binding the contract method 0x07a2fda1.
//
// Solidity: function processedAt(bytes32 _id) view returns(uint48)
func (_Mailbox *MailboxCaller) ProcessedAt(opts *bind.CallOpts, _id [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Mailbox.contract.Call(opts, &out, "processedAt", _id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProcessedAt is a free data retrieval call binding the contract method 0x07a2fda1.
//
// Solidity: function processedAt(bytes32 _id) view returns(uint48)
func (_Mailbox *MailboxSession) ProcessedAt(_id [32]byte) (*big.Int, error) {
	return _Mailbox.Contract.ProcessedAt(&_Mailbox.CallOpts, _id)
}

// ProcessedAt is a free data retrieval call binding the contract method 0x07a2fda1.
//
// Solidity: function processedAt(bytes32 _id) view returns(uint48)
func (_Mailbox *MailboxCallerSession) ProcessedAt(_id [32]byte) (*big.Int, error) {
	return _Mailbox.Contract.ProcessedAt(&_Mailbox.CallOpts, _id)
}

// Processor is a free data retrieval call binding the contract method 0x5d1fe5a9.
//
// Solidity: function processor(bytes32 _id) view returns(address)
func (_Mailbox *MailboxCaller) Processor(opts *bind.CallOpts, _id [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Mailbox.contract.Call(opts, &out, "processor", _id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Processor is a free data retrieval call binding the contract method 0x5d1fe5a9.
//
// Solidity: function processor(bytes32 _id) view returns(address)
func (_Mailbox *MailboxSession) Processor(_id [32]byte) (common.Address, error) {
	return _Mailbox.Contract.Processor(&_Mailbox.CallOpts, _id)
}

// Processor is a free data retrieval call binding the contract method 0x5d1fe5a9.
//
// Solidity: function processor(bytes32 _id) view returns(address)
func (_Mailbox *MailboxCallerSession) Processor(_id [32]byte) (common.Address, error) {
	return _Mailbox.Contract.Processor(&_Mailbox.CallOpts, _id)
}

// QuoteDispatch is a free data retrieval call binding the contract method 0x81d2ea95.
//
// Solidity: function quoteDispatch(uint32 destinationDomain, bytes32 recipientAddress, bytes messageBody, bytes metadata, address hook) view returns(uint256 fee)
func (_Mailbox *MailboxCaller) QuoteDispatch(opts *bind.CallOpts, destinationDomain uint32, recipientAddress [32]byte, messageBody []byte, metadata []byte, hook common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Mailbox.contract.Call(opts, &out, "quoteDispatch", destinationDomain, recipientAddress, messageBody, metadata, hook)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuoteDispatch is a free data retrieval call binding the contract method 0x81d2ea95.
//
// Solidity: function quoteDispatch(uint32 destinationDomain, bytes32 recipientAddress, bytes messageBody, bytes metadata, address hook) view returns(uint256 fee)
func (_Mailbox *MailboxSession) QuoteDispatch(destinationDomain uint32, recipientAddress [32]byte, messageBody []byte, metadata []byte, hook common.Address) (*big.Int, error) {
	return _Mailbox.Contract.QuoteDispatch(&_Mailbox.CallOpts, destinationDomain, recipientAddress, messageBody, metadata, hook)
}

// QuoteDispatch is a free data retrieval call binding the contract method 0x81d2ea95.
//
// Solidity: function quoteDispatch(uint32 destinationDomain, bytes32 recipientAddress, bytes messageBody, bytes metadata, address hook) view returns(uint256 fee)
func (_Mailbox *MailboxCallerSession) QuoteDispatch(destinationDomain uint32, recipientAddress [32]byte, messageBody []byte, metadata []byte, hook common.Address) (*big.Int, error) {
	return _Mailbox.Contract.QuoteDispatch(&_Mailbox.CallOpts, destinationDomain, recipientAddress, messageBody, metadata, hook)
}

// QuoteDispatch0 is a free data retrieval call binding the contract method 0x9c42bd18.
//
// Solidity: function quoteDispatch(uint32 destinationDomain, bytes32 recipientAddress, bytes messageBody) view returns(uint256 fee)
func (_Mailbox *MailboxCaller) QuoteDispatch0(opts *bind.CallOpts, destinationDomain uint32, recipientAddress [32]byte, messageBody []byte) (*big.Int, error) {
	var out []interface{}
	err := _Mailbox.contract.Call(opts, &out, "quoteDispatch0", destinationDomain, recipientAddress, messageBody)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuoteDispatch0 is a free data retrieval call binding the contract method 0x9c42bd18.
//
// Solidity: function quoteDispatch(uint32 destinationDomain, bytes32 recipientAddress, bytes messageBody) view returns(uint256 fee)
func (_Mailbox *MailboxSession) QuoteDispatch0(destinationDomain uint32, recipientAddress [32]byte, messageBody []byte) (*big.Int, error) {
	return _Mailbox.Contract.QuoteDispatch0(&_Mailbox.CallOpts, destinationDomain, recipientAddress, messageBody)
}

// QuoteDispatch0 is a free data retrieval call binding the contract method 0x9c42bd18.
//
// Solidity: function quoteDispatch(uint32 destinationDomain, bytes32 recipientAddress, bytes messageBody) view returns(uint256 fee)
func (_Mailbox *MailboxCallerSession) QuoteDispatch0(destinationDomain uint32, recipientAddress [32]byte, messageBody []byte) (*big.Int, error) {
	return _Mailbox.Contract.QuoteDispatch0(&_Mailbox.CallOpts, destinationDomain, recipientAddress, messageBody)
}

// QuoteDispatch1 is a free data retrieval call binding the contract method 0xf7ccd321.
//
// Solidity: function quoteDispatch(uint32 destinationDomain, bytes32 recipientAddress, bytes messageBody, bytes defaultHookMetadata) view returns(uint256 fee)
func (_Mailbox *MailboxCaller) QuoteDispatch1(opts *bind.CallOpts, destinationDomain uint32, recipientAddress [32]byte, messageBody []byte, defaultHookMetadata []byte) (*big.Int, error) {
	var out []interface{}
	err := _Mailbox.contract.Call(opts, &out, "quoteDispatch1", destinationDomain, recipientAddress, messageBody, defaultHookMetadata)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuoteDispatch1 is a free data retrieval call binding the contract method 0xf7ccd321.
//
// Solidity: function quoteDispatch(uint32 destinationDomain, bytes32 recipientAddress, bytes messageBody, bytes defaultHookMetadata) view returns(uint256 fee)
func (_Mailbox *MailboxSession) QuoteDispatch1(destinationDomain uint32, recipientAddress [32]byte, messageBody []byte, defaultHookMetadata []byte) (*big.Int, error) {
	return _Mailbox.Contract.QuoteDispatch1(&_Mailbox.CallOpts, destinationDomain, recipientAddress, messageBody, defaultHookMetadata)
}

// QuoteDispatch1 is a free data retrieval call binding the contract method 0xf7ccd321.
//
// Solidity: function quoteDispatch(uint32 destinationDomain, bytes32 recipientAddress, bytes messageBody, bytes defaultHookMetadata) view returns(uint256 fee)
func (_Mailbox *MailboxCallerSession) QuoteDispatch1(destinationDomain uint32, recipientAddress [32]byte, messageBody []byte, defaultHookMetadata []byte) (*big.Int, error) {
	return _Mailbox.Contract.QuoteDispatch1(&_Mailbox.CallOpts, destinationDomain, recipientAddress, messageBody, defaultHookMetadata)
}

// RecipientIsm is a free data retrieval call binding the contract method 0xe70f48ac.
//
// Solidity: function recipientIsm(address _recipient) view returns(address)
func (_Mailbox *MailboxCaller) RecipientIsm(opts *bind.CallOpts, _recipient common.Address) (common.Address, error) {
	var out []interface{}
	err := _Mailbox.contract.Call(opts, &out, "recipientIsm", _recipient)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RecipientIsm is a free data retrieval call binding the contract method 0xe70f48ac.
//
// Solidity: function recipientIsm(address _recipient) view returns(address)
func (_Mailbox *MailboxSession) RecipientIsm(_recipient common.Address) (common.Address, error) {
	return _Mailbox.Contract.RecipientIsm(&_Mailbox.CallOpts, _recipient)
}

// RecipientIsm is a free data retrieval call binding the contract method 0xe70f48ac.
//
// Solidity: function recipientIsm(address _recipient) view returns(address)
func (_Mailbox *MailboxCallerSession) RecipientIsm(_recipient common.Address) (common.Address, error) {
	return _Mailbox.Contract.RecipientIsm(&_Mailbox.CallOpts, _recipient)
}

// RequiredHook is a free data retrieval call binding the contract method 0xd6d08a09.
//
// Solidity: function requiredHook() view returns(address)
func (_Mailbox *MailboxCaller) RequiredHook(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mailbox.contract.Call(opts, &out, "requiredHook")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RequiredHook is a free data retrieval call binding the contract method 0xd6d08a09.
//
// Solidity: function requiredHook() view returns(address)
func (_Mailbox *MailboxSession) RequiredHook() (common.Address, error) {
	return _Mailbox.Contract.RequiredHook(&_Mailbox.CallOpts)
}

// RequiredHook is a free data retrieval call binding the contract method 0xd6d08a09.
//
// Solidity: function requiredHook() view returns(address)
func (_Mailbox *MailboxCallerSession) RequiredHook() (common.Address, error) {
	return _Mailbox.Contract.RequiredHook(&_Mailbox.CallOpts)
}

// Dispatch is a paid mutator transaction binding the contract method 0x10b83dc0.
//
// Solidity: function dispatch(uint32 destinationDomain, bytes32 recipientAddress, bytes messageBody, bytes metadata, address hook) payable returns(bytes32)
func (_Mailbox *MailboxTransactor) Dispatch(opts *bind.TransactOpts, destinationDomain uint32, recipientAddress [32]byte, messageBody []byte, metadata []byte, hook common.Address) (*types.Transaction, error) {
	return _Mailbox.contract.Transact(opts, "dispatch", destinationDomain, recipientAddress, messageBody, metadata, hook)
}

// Dispatch is a paid mutator transaction binding the contract method 0x10b83dc0.
//
// Solidity: function dispatch(uint32 destinationDomain, bytes32 recipientAddress, bytes messageBody, bytes metadata, address hook) payable returns(bytes32)
func (_Mailbox *MailboxSession) Dispatch(destinationDomain uint32, recipientAddress [32]byte, messageBody []byte, metadata []byte, hook common.Address) (*types.Transaction, error) {
	return _Mailbox.Contract.Dispatch(&_Mailbox.TransactOpts, destinationDomain, recipientAddress, messageBody, metadata, hook)
}

// Dispatch is a paid mutator transaction binding the contract method 0x10b83dc0.
//
// Solidity: function dispatch(uint32 destinationDomain, bytes32 recipientAddress, bytes messageBody, bytes metadata, address hook) payable returns(bytes32)
func (_Mailbox *MailboxTransactorSession) Dispatch(destinationDomain uint32, recipientAddress [32]byte, messageBody []byte, metadata []byte, hook common.Address) (*types.Transaction, error) {
	return _Mailbox.Contract.Dispatch(&_Mailbox.TransactOpts, destinationDomain, recipientAddress, messageBody, metadata, hook)
}

// Dispatch0 is a paid mutator transaction binding the contract method 0x48aee8d4.
//
// Solidity: function dispatch(uint32 destinationDomain, bytes32 recipientAddress, bytes messageBody, bytes hookMetadata) payable returns(bytes32)
func (_Mailbox *MailboxTransactor) Dispatch0(opts *bind.TransactOpts, destinationDomain uint32, recipientAddress [32]byte, messageBody []byte, hookMetadata []byte) (*types.Transaction, error) {
	return _Mailbox.contract.Transact(opts, "dispatch0", destinationDomain, recipientAddress, messageBody, hookMetadata)
}

// Dispatch0 is a paid mutator transaction binding the contract method 0x48aee8d4.
//
// Solidity: function dispatch(uint32 destinationDomain, bytes32 recipientAddress, bytes messageBody, bytes hookMetadata) payable returns(bytes32)
func (_Mailbox *MailboxSession) Dispatch0(destinationDomain uint32, recipientAddress [32]byte, messageBody []byte, hookMetadata []byte) (*types.Transaction, error) {
	return _Mailbox.Contract.Dispatch0(&_Mailbox.TransactOpts, destinationDomain, recipientAddress, messageBody, hookMetadata)
}

// Dispatch0 is a paid mutator transaction binding the contract method 0x48aee8d4.
//
// Solidity: function dispatch(uint32 destinationDomain, bytes32 recipientAddress, bytes messageBody, bytes hookMetadata) payable returns(bytes32)
func (_Mailbox *MailboxTransactorSession) Dispatch0(destinationDomain uint32, recipientAddress [32]byte, messageBody []byte, hookMetadata []byte) (*types.Transaction, error) {
	return _Mailbox.Contract.Dispatch0(&_Mailbox.TransactOpts, destinationDomain, recipientAddress, messageBody, hookMetadata)
}

// Dispatch1 is a paid mutator transaction binding the contract method 0xfa31de01.
//
// Solidity: function dispatch(uint32 _destinationDomain, bytes32 _recipientAddress, bytes _messageBody) payable returns(bytes32)
func (_Mailbox *MailboxTransactor) Dispatch1(opts *bind.TransactOpts, _destinationDomain uint32, _recipientAddress [32]byte, _messageBody []byte) (*types.Transaction, error) {
	return _Mailbox.contract.Transact(opts, "dispatch1", _destinationDomain, _recipientAddress, _messageBody)
}

// Dispatch1 is a paid mutator transaction binding the contract method 0xfa31de01.
//
// Solidity: function dispatch(uint32 _destinationDomain, bytes32 _recipientAddress, bytes _messageBody) payable returns(bytes32)
func (_Mailbox *MailboxSession) Dispatch1(_destinationDomain uint32, _recipientAddress [32]byte, _messageBody []byte) (*types.Transaction, error) {
	return _Mailbox.Contract.Dispatch1(&_Mailbox.TransactOpts, _destinationDomain, _recipientAddress, _messageBody)
}

// Dispatch1 is a paid mutator transaction binding the contract method 0xfa31de01.
//
// Solidity: function dispatch(uint32 _destinationDomain, bytes32 _recipientAddress, bytes _messageBody) payable returns(bytes32)
func (_Mailbox *MailboxTransactorSession) Dispatch1(_destinationDomain uint32, _recipientAddress [32]byte, _messageBody []byte) (*types.Transaction, error) {
	return _Mailbox.Contract.Dispatch1(&_Mailbox.TransactOpts, _destinationDomain, _recipientAddress, _messageBody)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _owner, address _defaultIsm, address _defaultHook, address _requiredHook) returns()
func (_Mailbox *MailboxTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address, _defaultIsm common.Address, _defaultHook common.Address, _requiredHook common.Address) (*types.Transaction, error) {
	return _Mailbox.contract.Transact(opts, "initialize", _owner, _defaultIsm, _defaultHook, _requiredHook)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _owner, address _defaultIsm, address _defaultHook, address _requiredHook) returns()
func (_Mailbox *MailboxSession) Initialize(_owner common.Address, _defaultIsm common.Address, _defaultHook common.Address, _requiredHook common.Address) (*types.Transaction, error) {
	return _Mailbox.Contract.Initialize(&_Mailbox.TransactOpts, _owner, _defaultIsm, _defaultHook, _requiredHook)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _owner, address _defaultIsm, address _defaultHook, address _requiredHook) returns()
func (_Mailbox *MailboxTransactorSession) Initialize(_owner common.Address, _defaultIsm common.Address, _defaultHook common.Address, _requiredHook common.Address) (*types.Transaction, error) {
	return _Mailbox.Contract.Initialize(&_Mailbox.TransactOpts, _owner, _defaultIsm, _defaultHook, _requiredHook)
}

// Process is a paid mutator transaction binding the contract method 0x7c39d130.
//
// Solidity: function process(bytes _metadata, bytes _message) payable returns()
func (_Mailbox *MailboxTransactor) Process(opts *bind.TransactOpts, _metadata []byte, _message []byte) (*types.Transaction, error) {
	return _Mailbox.contract.Transact(opts, "process", _metadata, _message)
}

// Process is a paid mutator transaction binding the contract method 0x7c39d130.
//
// Solidity: function process(bytes _metadata, bytes _message) payable returns()
func (_Mailbox *MailboxSession) Process(_metadata []byte, _message []byte) (*types.Transaction, error) {
	return _Mailbox.Contract.Process(&_Mailbox.TransactOpts, _metadata, _message)
}

// Process is a paid mutator transaction binding the contract method 0x7c39d130.
//
// Solidity: function process(bytes _metadata, bytes _message) payable returns()
func (_Mailbox *MailboxTransactorSession) Process(_metadata []byte, _message []byte) (*types.Transaction, error) {
	return _Mailbox.Contract.Process(&_Mailbox.TransactOpts, _metadata, _message)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mailbox *MailboxTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mailbox.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mailbox *MailboxSession) RenounceOwnership() (*types.Transaction, error) {
	return _Mailbox.Contract.RenounceOwnership(&_Mailbox.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mailbox *MailboxTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Mailbox.Contract.RenounceOwnership(&_Mailbox.TransactOpts)
}

// SetDefaultHook is a paid mutator transaction binding the contract method 0x99b04809.
//
// Solidity: function setDefaultHook(address _hook) returns()
func (_Mailbox *MailboxTransactor) SetDefaultHook(opts *bind.TransactOpts, _hook common.Address) (*types.Transaction, error) {
	return _Mailbox.contract.Transact(opts, "setDefaultHook", _hook)
}

// SetDefaultHook is a paid mutator transaction binding the contract method 0x99b04809.
//
// Solidity: function setDefaultHook(address _hook) returns()
func (_Mailbox *MailboxSession) SetDefaultHook(_hook common.Address) (*types.Transaction, error) {
	return _Mailbox.Contract.SetDefaultHook(&_Mailbox.TransactOpts, _hook)
}

// SetDefaultHook is a paid mutator transaction binding the contract method 0x99b04809.
//
// Solidity: function setDefaultHook(address _hook) returns()
func (_Mailbox *MailboxTransactorSession) SetDefaultHook(_hook common.Address) (*types.Transaction, error) {
	return _Mailbox.Contract.SetDefaultHook(&_Mailbox.TransactOpts, _hook)
}

// SetDefaultIsm is a paid mutator transaction binding the contract method 0xf794687a.
//
// Solidity: function setDefaultIsm(address _module) returns()
func (_Mailbox *MailboxTransactor) SetDefaultIsm(opts *bind.TransactOpts, _module common.Address) (*types.Transaction, error) {
	return _Mailbox.contract.Transact(opts, "setDefaultIsm", _module)
}

// SetDefaultIsm is a paid mutator transaction binding the contract method 0xf794687a.
//
// Solidity: function setDefaultIsm(address _module) returns()
func (_Mailbox *MailboxSession) SetDefaultIsm(_module common.Address) (*types.Transaction, error) {
	return _Mailbox.Contract.SetDefaultIsm(&_Mailbox.TransactOpts, _module)
}

// SetDefaultIsm is a paid mutator transaction binding the contract method 0xf794687a.
//
// Solidity: function setDefaultIsm(address _module) returns()
func (_Mailbox *MailboxTransactorSession) SetDefaultIsm(_module common.Address) (*types.Transaction, error) {
	return _Mailbox.Contract.SetDefaultIsm(&_Mailbox.TransactOpts, _module)
}

// SetRequiredHook is a paid mutator transaction binding the contract method 0x1426b7f4.
//
// Solidity: function setRequiredHook(address _hook) returns()
func (_Mailbox *MailboxTransactor) SetRequiredHook(opts *bind.TransactOpts, _hook common.Address) (*types.Transaction, error) {
	return _Mailbox.contract.Transact(opts, "setRequiredHook", _hook)
}

// SetRequiredHook is a paid mutator transaction binding the contract method 0x1426b7f4.
//
// Solidity: function setRequiredHook(address _hook) returns()
func (_Mailbox *MailboxSession) SetRequiredHook(_hook common.Address) (*types.Transaction, error) {
	return _Mailbox.Contract.SetRequiredHook(&_Mailbox.TransactOpts, _hook)
}

// SetRequiredHook is a paid mutator transaction binding the contract method 0x1426b7f4.
//
// Solidity: function setRequiredHook(address _hook) returns()
func (_Mailbox *MailboxTransactorSession) SetRequiredHook(_hook common.Address) (*types.Transaction, error) {
	return _Mailbox.Contract.SetRequiredHook(&_Mailbox.TransactOpts, _hook)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mailbox *MailboxTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Mailbox.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mailbox *MailboxSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Mailbox.Contract.TransferOwnership(&_Mailbox.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mailbox *MailboxTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Mailbox.Contract.TransferOwnership(&_Mailbox.TransactOpts, newOwner)
}

// MailboxDefaultHookSetIterator is returned from FilterDefaultHookSet and is used to iterate over the raw logs and unpacked data for DefaultHookSet events raised by the Mailbox contract.
type MailboxDefaultHookSetIterator struct {
	Event *MailboxDefaultHookSet // Event containing the contract specifics and raw log

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
func (it *MailboxDefaultHookSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MailboxDefaultHookSet)
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
		it.Event = new(MailboxDefaultHookSet)
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
func (it *MailboxDefaultHookSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MailboxDefaultHookSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MailboxDefaultHookSet represents a DefaultHookSet event raised by the Mailbox contract.
type MailboxDefaultHookSet struct {
	Hook common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterDefaultHookSet is a free log retrieval operation binding the contract event 0x65a63e5066ee2fcdf9d32a7f1bf7ce71c76066f19d0609dddccd334ab87237d7.
//
// Solidity: event DefaultHookSet(address indexed hook)
func (_Mailbox *MailboxFilterer) FilterDefaultHookSet(opts *bind.FilterOpts, hook []common.Address) (*MailboxDefaultHookSetIterator, error) {

	var hookRule []interface{}
	for _, hookItem := range hook {
		hookRule = append(hookRule, hookItem)
	}

	logs, sub, err := _Mailbox.contract.FilterLogs(opts, "DefaultHookSet", hookRule)
	if err != nil {
		return nil, err
	}
	return &MailboxDefaultHookSetIterator{contract: _Mailbox.contract, event: "DefaultHookSet", logs: logs, sub: sub}, nil
}

// WatchDefaultHookSet is a free log subscription operation binding the contract event 0x65a63e5066ee2fcdf9d32a7f1bf7ce71c76066f19d0609dddccd334ab87237d7.
//
// Solidity: event DefaultHookSet(address indexed hook)
func (_Mailbox *MailboxFilterer) WatchDefaultHookSet(opts *bind.WatchOpts, sink chan<- *MailboxDefaultHookSet, hook []common.Address) (event.Subscription, error) {

	var hookRule []interface{}
	for _, hookItem := range hook {
		hookRule = append(hookRule, hookItem)
	}

	logs, sub, err := _Mailbox.contract.WatchLogs(opts, "DefaultHookSet", hookRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MailboxDefaultHookSet)
				if err := _Mailbox.contract.UnpackLog(event, "DefaultHookSet", log); err != nil {
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

// ParseDefaultHookSet is a log parse operation binding the contract event 0x65a63e5066ee2fcdf9d32a7f1bf7ce71c76066f19d0609dddccd334ab87237d7.
//
// Solidity: event DefaultHookSet(address indexed hook)
func (_Mailbox *MailboxFilterer) ParseDefaultHookSet(log types.Log) (*MailboxDefaultHookSet, error) {
	event := new(MailboxDefaultHookSet)
	if err := _Mailbox.contract.UnpackLog(event, "DefaultHookSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MailboxDefaultIsmSetIterator is returned from FilterDefaultIsmSet and is used to iterate over the raw logs and unpacked data for DefaultIsmSet events raised by the Mailbox contract.
type MailboxDefaultIsmSetIterator struct {
	Event *MailboxDefaultIsmSet // Event containing the contract specifics and raw log

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
func (it *MailboxDefaultIsmSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MailboxDefaultIsmSet)
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
		it.Event = new(MailboxDefaultIsmSet)
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
func (it *MailboxDefaultIsmSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MailboxDefaultIsmSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MailboxDefaultIsmSet represents a DefaultIsmSet event raised by the Mailbox contract.
type MailboxDefaultIsmSet struct {
	Module common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDefaultIsmSet is a free log retrieval operation binding the contract event 0xa76ad0adbf45318f8633aa0210f711273d50fbb6fef76ed95bbae97082c75daa.
//
// Solidity: event DefaultIsmSet(address indexed module)
func (_Mailbox *MailboxFilterer) FilterDefaultIsmSet(opts *bind.FilterOpts, module []common.Address) (*MailboxDefaultIsmSetIterator, error) {

	var moduleRule []interface{}
	for _, moduleItem := range module {
		moduleRule = append(moduleRule, moduleItem)
	}

	logs, sub, err := _Mailbox.contract.FilterLogs(opts, "DefaultIsmSet", moduleRule)
	if err != nil {
		return nil, err
	}
	return &MailboxDefaultIsmSetIterator{contract: _Mailbox.contract, event: "DefaultIsmSet", logs: logs, sub: sub}, nil
}

// WatchDefaultIsmSet is a free log subscription operation binding the contract event 0xa76ad0adbf45318f8633aa0210f711273d50fbb6fef76ed95bbae97082c75daa.
//
// Solidity: event DefaultIsmSet(address indexed module)
func (_Mailbox *MailboxFilterer) WatchDefaultIsmSet(opts *bind.WatchOpts, sink chan<- *MailboxDefaultIsmSet, module []common.Address) (event.Subscription, error) {

	var moduleRule []interface{}
	for _, moduleItem := range module {
		moduleRule = append(moduleRule, moduleItem)
	}

	logs, sub, err := _Mailbox.contract.WatchLogs(opts, "DefaultIsmSet", moduleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MailboxDefaultIsmSet)
				if err := _Mailbox.contract.UnpackLog(event, "DefaultIsmSet", log); err != nil {
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

// ParseDefaultIsmSet is a log parse operation binding the contract event 0xa76ad0adbf45318f8633aa0210f711273d50fbb6fef76ed95bbae97082c75daa.
//
// Solidity: event DefaultIsmSet(address indexed module)
func (_Mailbox *MailboxFilterer) ParseDefaultIsmSet(log types.Log) (*MailboxDefaultIsmSet, error) {
	event := new(MailboxDefaultIsmSet)
	if err := _Mailbox.contract.UnpackLog(event, "DefaultIsmSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MailboxDispatchIterator is returned from FilterDispatch and is used to iterate over the raw logs and unpacked data for Dispatch events raised by the Mailbox contract.
type MailboxDispatchIterator struct {
	Event *MailboxDispatch // Event containing the contract specifics and raw log

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
func (it *MailboxDispatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MailboxDispatch)
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
		it.Event = new(MailboxDispatch)
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
func (it *MailboxDispatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MailboxDispatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MailboxDispatch represents a Dispatch event raised by the Mailbox contract.
type MailboxDispatch struct {
	Sender      common.Address
	Destination uint32
	Recipient   [32]byte
	Message     []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDispatch is a free log retrieval operation binding the contract event 0x769f711d20c679153d382254f59892613b58a97cc876b249134ac25c80f9c814.
//
// Solidity: event Dispatch(address indexed sender, uint32 indexed destination, bytes32 indexed recipient, bytes message)
func (_Mailbox *MailboxFilterer) FilterDispatch(opts *bind.FilterOpts, sender []common.Address, destination []uint32, recipient [][32]byte) (*MailboxDispatchIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Mailbox.contract.FilterLogs(opts, "Dispatch", senderRule, destinationRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &MailboxDispatchIterator{contract: _Mailbox.contract, event: "Dispatch", logs: logs, sub: sub}, nil
}

// WatchDispatch is a free log subscription operation binding the contract event 0x769f711d20c679153d382254f59892613b58a97cc876b249134ac25c80f9c814.
//
// Solidity: event Dispatch(address indexed sender, uint32 indexed destination, bytes32 indexed recipient, bytes message)
func (_Mailbox *MailboxFilterer) WatchDispatch(opts *bind.WatchOpts, sink chan<- *MailboxDispatch, sender []common.Address, destination []uint32, recipient [][32]byte) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Mailbox.contract.WatchLogs(opts, "Dispatch", senderRule, destinationRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MailboxDispatch)
				if err := _Mailbox.contract.UnpackLog(event, "Dispatch", log); err != nil {
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

// ParseDispatch is a log parse operation binding the contract event 0x769f711d20c679153d382254f59892613b58a97cc876b249134ac25c80f9c814.
//
// Solidity: event Dispatch(address indexed sender, uint32 indexed destination, bytes32 indexed recipient, bytes message)
func (_Mailbox *MailboxFilterer) ParseDispatch(log types.Log) (*MailboxDispatch, error) {
	event := new(MailboxDispatch)
	if err := _Mailbox.contract.UnpackLog(event, "Dispatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MailboxDispatchIdIterator is returned from FilterDispatchId and is used to iterate over the raw logs and unpacked data for DispatchId events raised by the Mailbox contract.
type MailboxDispatchIdIterator struct {
	Event *MailboxDispatchId // Event containing the contract specifics and raw log

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
func (it *MailboxDispatchIdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MailboxDispatchId)
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
		it.Event = new(MailboxDispatchId)
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
func (it *MailboxDispatchIdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MailboxDispatchIdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MailboxDispatchId represents a DispatchId event raised by the Mailbox contract.
type MailboxDispatchId struct {
	MessageId [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDispatchId is a free log retrieval operation binding the contract event 0x788dbc1b7152732178210e7f4d9d010ef016f9eafbe66786bd7169f56e0c353a.
//
// Solidity: event DispatchId(bytes32 indexed messageId)
func (_Mailbox *MailboxFilterer) FilterDispatchId(opts *bind.FilterOpts, messageId [][32]byte) (*MailboxDispatchIdIterator, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _Mailbox.contract.FilterLogs(opts, "DispatchId", messageIdRule)
	if err != nil {
		return nil, err
	}
	return &MailboxDispatchIdIterator{contract: _Mailbox.contract, event: "DispatchId", logs: logs, sub: sub}, nil
}

// WatchDispatchId is a free log subscription operation binding the contract event 0x788dbc1b7152732178210e7f4d9d010ef016f9eafbe66786bd7169f56e0c353a.
//
// Solidity: event DispatchId(bytes32 indexed messageId)
func (_Mailbox *MailboxFilterer) WatchDispatchId(opts *bind.WatchOpts, sink chan<- *MailboxDispatchId, messageId [][32]byte) (event.Subscription, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _Mailbox.contract.WatchLogs(opts, "DispatchId", messageIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MailboxDispatchId)
				if err := _Mailbox.contract.UnpackLog(event, "DispatchId", log); err != nil {
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

// ParseDispatchId is a log parse operation binding the contract event 0x788dbc1b7152732178210e7f4d9d010ef016f9eafbe66786bd7169f56e0c353a.
//
// Solidity: event DispatchId(bytes32 indexed messageId)
func (_Mailbox *MailboxFilterer) ParseDispatchId(log types.Log) (*MailboxDispatchId, error) {
	event := new(MailboxDispatchId)
	if err := _Mailbox.contract.UnpackLog(event, "DispatchId", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MailboxInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Mailbox contract.
type MailboxInitializedIterator struct {
	Event *MailboxInitialized // Event containing the contract specifics and raw log

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
func (it *MailboxInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MailboxInitialized)
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
		it.Event = new(MailboxInitialized)
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
func (it *MailboxInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MailboxInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MailboxInitialized represents a Initialized event raised by the Mailbox contract.
type MailboxInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Mailbox *MailboxFilterer) FilterInitialized(opts *bind.FilterOpts) (*MailboxInitializedIterator, error) {

	logs, sub, err := _Mailbox.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &MailboxInitializedIterator{contract: _Mailbox.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Mailbox *MailboxFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *MailboxInitialized) (event.Subscription, error) {

	logs, sub, err := _Mailbox.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MailboxInitialized)
				if err := _Mailbox.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Mailbox *MailboxFilterer) ParseInitialized(log types.Log) (*MailboxInitialized, error) {
	event := new(MailboxInitialized)
	if err := _Mailbox.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MailboxOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Mailbox contract.
type MailboxOwnershipTransferredIterator struct {
	Event *MailboxOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MailboxOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MailboxOwnershipTransferred)
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
		it.Event = new(MailboxOwnershipTransferred)
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
func (it *MailboxOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MailboxOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MailboxOwnershipTransferred represents a OwnershipTransferred event raised by the Mailbox contract.
type MailboxOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Mailbox *MailboxFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MailboxOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Mailbox.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MailboxOwnershipTransferredIterator{contract: _Mailbox.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Mailbox *MailboxFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MailboxOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Mailbox.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MailboxOwnershipTransferred)
				if err := _Mailbox.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Mailbox *MailboxFilterer) ParseOwnershipTransferred(log types.Log) (*MailboxOwnershipTransferred, error) {
	event := new(MailboxOwnershipTransferred)
	if err := _Mailbox.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MailboxProcessIterator is returned from FilterProcess and is used to iterate over the raw logs and unpacked data for Process events raised by the Mailbox contract.
type MailboxProcessIterator struct {
	Event *MailboxProcess // Event containing the contract specifics and raw log

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
func (it *MailboxProcessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MailboxProcess)
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
		it.Event = new(MailboxProcess)
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
func (it *MailboxProcessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MailboxProcessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MailboxProcess represents a Process event raised by the Mailbox contract.
type MailboxProcess struct {
	Origin    uint32
	Sender    [32]byte
	Recipient common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProcess is a free log retrieval operation binding the contract event 0x0d381c2a574ae8f04e213db7cfb4df8df712cdbd427d9868ffef380660ca6574.
//
// Solidity: event Process(uint32 indexed origin, bytes32 indexed sender, address indexed recipient)
func (_Mailbox *MailboxFilterer) FilterProcess(opts *bind.FilterOpts, origin []uint32, sender [][32]byte, recipient []common.Address) (*MailboxProcessIterator, error) {

	var originRule []interface{}
	for _, originItem := range origin {
		originRule = append(originRule, originItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Mailbox.contract.FilterLogs(opts, "Process", originRule, senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &MailboxProcessIterator{contract: _Mailbox.contract, event: "Process", logs: logs, sub: sub}, nil
}

// WatchProcess is a free log subscription operation binding the contract event 0x0d381c2a574ae8f04e213db7cfb4df8df712cdbd427d9868ffef380660ca6574.
//
// Solidity: event Process(uint32 indexed origin, bytes32 indexed sender, address indexed recipient)
func (_Mailbox *MailboxFilterer) WatchProcess(opts *bind.WatchOpts, sink chan<- *MailboxProcess, origin []uint32, sender [][32]byte, recipient []common.Address) (event.Subscription, error) {

	var originRule []interface{}
	for _, originItem := range origin {
		originRule = append(originRule, originItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Mailbox.contract.WatchLogs(opts, "Process", originRule, senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MailboxProcess)
				if err := _Mailbox.contract.UnpackLog(event, "Process", log); err != nil {
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

// ParseProcess is a log parse operation binding the contract event 0x0d381c2a574ae8f04e213db7cfb4df8df712cdbd427d9868ffef380660ca6574.
//
// Solidity: event Process(uint32 indexed origin, bytes32 indexed sender, address indexed recipient)
func (_Mailbox *MailboxFilterer) ParseProcess(log types.Log) (*MailboxProcess, error) {
	event := new(MailboxProcess)
	if err := _Mailbox.contract.UnpackLog(event, "Process", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MailboxProcessIdIterator is returned from FilterProcessId and is used to iterate over the raw logs and unpacked data for ProcessId events raised by the Mailbox contract.
type MailboxProcessIdIterator struct {
	Event *MailboxProcessId // Event containing the contract specifics and raw log

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
func (it *MailboxProcessIdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MailboxProcessId)
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
		it.Event = new(MailboxProcessId)
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
func (it *MailboxProcessIdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MailboxProcessIdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MailboxProcessId represents a ProcessId event raised by the Mailbox contract.
type MailboxProcessId struct {
	MessageId [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProcessId is a free log retrieval operation binding the contract event 0x1cae38cdd3d3919489272725a5ae62a4f48b2989b0dae843d3c279fee18073a9.
//
// Solidity: event ProcessId(bytes32 indexed messageId)
func (_Mailbox *MailboxFilterer) FilterProcessId(opts *bind.FilterOpts, messageId [][32]byte) (*MailboxProcessIdIterator, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _Mailbox.contract.FilterLogs(opts, "ProcessId", messageIdRule)
	if err != nil {
		return nil, err
	}
	return &MailboxProcessIdIterator{contract: _Mailbox.contract, event: "ProcessId", logs: logs, sub: sub}, nil
}

// WatchProcessId is a free log subscription operation binding the contract event 0x1cae38cdd3d3919489272725a5ae62a4f48b2989b0dae843d3c279fee18073a9.
//
// Solidity: event ProcessId(bytes32 indexed messageId)
func (_Mailbox *MailboxFilterer) WatchProcessId(opts *bind.WatchOpts, sink chan<- *MailboxProcessId, messageId [][32]byte) (event.Subscription, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _Mailbox.contract.WatchLogs(opts, "ProcessId", messageIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MailboxProcessId)
				if err := _Mailbox.contract.UnpackLog(event, "ProcessId", log); err != nil {
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

// ParseProcessId is a log parse operation binding the contract event 0x1cae38cdd3d3919489272725a5ae62a4f48b2989b0dae843d3c279fee18073a9.
//
// Solidity: event ProcessId(bytes32 indexed messageId)
func (_Mailbox *MailboxFilterer) ParseProcessId(log types.Log) (*MailboxProcessId, error) {
	event := new(MailboxProcessId)
	if err := _Mailbox.contract.UnpackLog(event, "ProcessId", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MailboxRequiredHookSetIterator is returned from FilterRequiredHookSet and is used to iterate over the raw logs and unpacked data for RequiredHookSet events raised by the Mailbox contract.
type MailboxRequiredHookSetIterator struct {
	Event *MailboxRequiredHookSet // Event containing the contract specifics and raw log

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
func (it *MailboxRequiredHookSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MailboxRequiredHookSet)
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
		it.Event = new(MailboxRequiredHookSet)
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
func (it *MailboxRequiredHookSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MailboxRequiredHookSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MailboxRequiredHookSet represents a RequiredHookSet event raised by the Mailbox contract.
type MailboxRequiredHookSet struct {
	Hook common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRequiredHookSet is a free log retrieval operation binding the contract event 0x329ec8e2438a73828ecf31a6568d7a91d7b1d79e342b0692914fd053d1a002b1.
//
// Solidity: event RequiredHookSet(address indexed hook)
func (_Mailbox *MailboxFilterer) FilterRequiredHookSet(opts *bind.FilterOpts, hook []common.Address) (*MailboxRequiredHookSetIterator, error) {

	var hookRule []interface{}
	for _, hookItem := range hook {
		hookRule = append(hookRule, hookItem)
	}

	logs, sub, err := _Mailbox.contract.FilterLogs(opts, "RequiredHookSet", hookRule)
	if err != nil {
		return nil, err
	}
	return &MailboxRequiredHookSetIterator{contract: _Mailbox.contract, event: "RequiredHookSet", logs: logs, sub: sub}, nil
}

// WatchRequiredHookSet is a free log subscription operation binding the contract event 0x329ec8e2438a73828ecf31a6568d7a91d7b1d79e342b0692914fd053d1a002b1.
//
// Solidity: event RequiredHookSet(address indexed hook)
func (_Mailbox *MailboxFilterer) WatchRequiredHookSet(opts *bind.WatchOpts, sink chan<- *MailboxRequiredHookSet, hook []common.Address) (event.Subscription, error) {

	var hookRule []interface{}
	for _, hookItem := range hook {
		hookRule = append(hookRule, hookItem)
	}

	logs, sub, err := _Mailbox.contract.WatchLogs(opts, "RequiredHookSet", hookRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MailboxRequiredHookSet)
				if err := _Mailbox.contract.UnpackLog(event, "RequiredHookSet", log); err != nil {
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

// ParseRequiredHookSet is a log parse operation binding the contract event 0x329ec8e2438a73828ecf31a6568d7a91d7b1d79e342b0692914fd053d1a002b1.
//
// Solidity: event RequiredHookSet(address indexed hook)
func (_Mailbox *MailboxFilterer) ParseRequiredHookSet(log types.Log) (*MailboxRequiredHookSet, error) {
	event := new(MailboxRequiredHookSet)
	if err := _Mailbox.contract.UnpackLog(event, "RequiredHookSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
