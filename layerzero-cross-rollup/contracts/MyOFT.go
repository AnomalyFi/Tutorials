// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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
)

// EnforcedOptionParam is an auto generated low-level Go binding around an user-defined struct.
type EnforcedOptionParam struct {
	Eid     uint32
	MsgType uint16
	Options []byte
}

// InboundPacket is an auto generated low-level Go binding around an user-defined struct.
type InboundPacket struct {
	Origin    Origin
	DstEid    uint32
	Receiver  common.Address
	Guid      [32]byte
	Value     *big.Int
	Executor  common.Address
	Message   []byte
	ExtraData []byte
}

// MessagingFee is an auto generated low-level Go binding around an user-defined struct.
type MessagingFee struct {
	NativeFee  *big.Int
	LzTokenFee *big.Int
}

// MessagingReceipt is an auto generated low-level Go binding around an user-defined struct.
type MessagingReceipt struct {
	Guid  [32]byte
	Nonce uint64
	Fee   MessagingFee
}

// OFTFeeDetail is an auto generated low-level Go binding around an user-defined struct.
type OFTFeeDetail struct {
	FeeAmountLD *big.Int
	Description string
}

// OFTLimit is an auto generated low-level Go binding around an user-defined struct.
type OFTLimit struct {
	MinAmountLD *big.Int
	MaxAmountLD *big.Int
}

// OFTReceipt is an auto generated low-level Go binding around an user-defined struct.
type OFTReceipt struct {
	AmountSentLD     *big.Int
	AmountReceivedLD *big.Int
}

// Origin is an auto generated low-level Go binding around an user-defined struct.
type Origin struct {
	SrcEid uint32
	Sender [32]byte
	Nonce  uint64
}

// SendParam is an auto generated low-level Go binding around an user-defined struct.
type SendParam struct {
	DstEid       uint32
	To           [32]byte
	AmountLD     *big.Int
	MinAmountLD  *big.Int
	ExtraOptions []byte
	ComposeMsg   []byte
	OftCmd       []byte
}

// MyOFTMetaData contains all meta data concerning the MyOFT contract.
var MyOFTMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_lzEndpoint\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_delegate\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"SEND\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SEND_AND_CALL\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allowInitializePath\",\"inputs\":[{\"name\":\"origin\",\"type\":\"tuple\",\"internalType\":\"structOrigin\",\"components\":[{\"name\":\"srcEid\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"sender\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approvalRequired\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"combineOptions\",\"inputs\":[{\"name\":\"_eid\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_msgType\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"_extraOptions\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decimalConversionRate\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"endpoint\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractILayerZeroEndpointV2\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"enforcedOptions\",\"inputs\":[{\"name\":\"eid\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"msgType\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"enforcedOption\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isComposeMsgSender\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOrigin\",\"components\":[{\"name\":\"srcEid\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"sender\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_sender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isPeer\",\"inputs\":[{\"name\":\"_eid\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_peer\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lzReceive\",\"inputs\":[{\"name\":\"_origin\",\"type\":\"tuple\",\"internalType\":\"structOrigin\",\"components\":[{\"name\":\"srcEid\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"sender\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"_guid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_executor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"lzReceiveAndRevert\",\"inputs\":[{\"name\":\"_packets\",\"type\":\"tuple[]\",\"internalType\":\"structInboundPacket[]\",\"components\":[{\"name\":\"origin\",\"type\":\"tuple\",\"internalType\":\"structOrigin\",\"components\":[{\"name\":\"srcEid\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"sender\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"dstEid\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"guid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"executor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"lzReceiveSimulate\",\"inputs\":[{\"name\":\"_origin\",\"type\":\"tuple\",\"internalType\":\"structOrigin\",\"components\":[{\"name\":\"srcEid\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"sender\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"_guid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_executor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"msgInspector\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nextNonce\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"oApp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"oAppVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"senderVersion\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiverVersion\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"oftVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"version\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"peers\",\"inputs\":[{\"name\":\"eid\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"peer\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"preCrime\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"quoteOFT\",\"inputs\":[{\"name\":\"_sendParam\",\"type\":\"tuple\",\"internalType\":\"structSendParam\",\"components\":[{\"name\":\"dstEid\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"to\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"amountLD\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minAmountLD\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"extraOptions\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"composeMsg\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"oftCmd\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"oftLimit\",\"type\":\"tuple\",\"internalType\":\"structOFTLimit\",\"components\":[{\"name\":\"minAmountLD\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxAmountLD\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"oftFeeDetails\",\"type\":\"tuple[]\",\"internalType\":\"structOFTFeeDetail[]\",\"components\":[{\"name\":\"feeAmountLD\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"oftReceipt\",\"type\":\"tuple\",\"internalType\":\"structOFTReceipt\",\"components\":[{\"name\":\"amountSentLD\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amountReceivedLD\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"quoteSend\",\"inputs\":[{\"name\":\"_sendParam\",\"type\":\"tuple\",\"internalType\":\"structSendParam\",\"components\":[{\"name\":\"dstEid\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"to\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"amountLD\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minAmountLD\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"extraOptions\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"composeMsg\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"oftCmd\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"_payInLzToken\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"msgFee\",\"type\":\"tuple\",\"internalType\":\"structMessagingFee\",\"components\":[{\"name\":\"nativeFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lzTokenFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"send\",\"inputs\":[{\"name\":\"_sendParam\",\"type\":\"tuple\",\"internalType\":\"structSendParam\",\"components\":[{\"name\":\"dstEid\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"to\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"amountLD\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minAmountLD\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"extraOptions\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"composeMsg\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"oftCmd\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"_fee\",\"type\":\"tuple\",\"internalType\":\"structMessagingFee\",\"components\":[{\"name\":\"nativeFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lzTokenFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"_refundAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"msgReceipt\",\"type\":\"tuple\",\"internalType\":\"structMessagingReceipt\",\"components\":[{\"name\":\"guid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"fee\",\"type\":\"tuple\",\"internalType\":\"structMessagingFee\",\"components\":[{\"name\":\"nativeFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lzTokenFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]},{\"name\":\"oftReceipt\",\"type\":\"tuple\",\"internalType\":\"structOFTReceipt\",\"components\":[{\"name\":\"amountSentLD\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amountReceivedLD\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"setDelegate\",\"inputs\":[{\"name\":\"_delegate\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setEnforcedOptions\",\"inputs\":[{\"name\":\"_enforcedOptions\",\"type\":\"tuple[]\",\"internalType\":\"structEnforcedOptionParam[]\",\"components\":[{\"name\":\"eid\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"msgType\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"options\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMsgInspector\",\"inputs\":[{\"name\":\"_msgInspector\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPeer\",\"inputs\":[{\"name\":\"_eid\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_peer\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPreCrime\",\"inputs\":[{\"name\":\"_preCrime\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sharedDecimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"token\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EnforcedOptionSet\",\"inputs\":[{\"name\":\"_enforcedOptions\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"structEnforcedOptionParam[]\",\"components\":[{\"name\":\"eid\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"msgType\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"options\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MsgInspectorSet\",\"inputs\":[{\"name\":\"inspector\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OFTReceived\",\"inputs\":[{\"name\":\"guid\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"srcEid\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"toAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amountReceivedLD\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OFTSent\",\"inputs\":[{\"name\":\"guid\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"dstEid\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"fromAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amountSentLD\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"amountReceivedLD\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PeerSet\",\"inputs\":[{\"name\":\"eid\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"peer\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PreCrimeSet\",\"inputs\":[{\"name\":\"preCrimeAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ERC20InsufficientAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientBalance\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidApprover\",\"inputs\":[{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSpender\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidDelegate\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidEndpointCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidLocalDecimals\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOptions\",\"inputs\":[{\"name\":\"options\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"LzTokenUnavailable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoPeer\",\"inputs\":[{\"name\":\"eid\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"NotEnoughNative\",\"inputs\":[{\"name\":\"msgValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OnlyEndpoint\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OnlyPeer\",\"inputs\":[{\"name\":\"eid\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"sender\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"OnlySelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"SimulationResult\",\"inputs\":[{\"name\":\"result\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"SlippageExceeded\",\"inputs\":[{\"name\":\"amountLD\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minAmountLD\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]",
}

// MyOFTABI is the input ABI used to generate the binding from.
// Deprecated: Use MyOFTMetaData.ABI instead.
var MyOFTABI = MyOFTMetaData.ABI

// MyOFT is an auto generated Go binding around an Ethereum contract.
type MyOFT struct {
	MyOFTCaller     // Read-only binding to the contract
	MyOFTTransactor // Write-only binding to the contract
	MyOFTFilterer   // Log filterer for contract events
}

// MyOFTCaller is an auto generated read-only Go binding around an Ethereum contract.
type MyOFTCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyOFTTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MyOFTTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyOFTFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MyOFTFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyOFTSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MyOFTSession struct {
	Contract     *MyOFT            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MyOFTCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MyOFTCallerSession struct {
	Contract *MyOFTCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MyOFTTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MyOFTTransactorSession struct {
	Contract     *MyOFTTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MyOFTRaw is an auto generated low-level Go binding around an Ethereum contract.
type MyOFTRaw struct {
	Contract *MyOFT // Generic contract binding to access the raw methods on
}

// MyOFTCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MyOFTCallerRaw struct {
	Contract *MyOFTCaller // Generic read-only contract binding to access the raw methods on
}

// MyOFTTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MyOFTTransactorRaw struct {
	Contract *MyOFTTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMyOFT creates a new instance of MyOFT, bound to a specific deployed contract.
func NewMyOFT(address common.Address, backend bind.ContractBackend) (*MyOFT, error) {
	contract, err := bindMyOFT(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MyOFT{MyOFTCaller: MyOFTCaller{contract: contract}, MyOFTTransactor: MyOFTTransactor{contract: contract}, MyOFTFilterer: MyOFTFilterer{contract: contract}}, nil
}

// NewMyOFTCaller creates a new read-only instance of MyOFT, bound to a specific deployed contract.
func NewMyOFTCaller(address common.Address, caller bind.ContractCaller) (*MyOFTCaller, error) {
	contract, err := bindMyOFT(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MyOFTCaller{contract: contract}, nil
}

// NewMyOFTTransactor creates a new write-only instance of MyOFT, bound to a specific deployed contract.
func NewMyOFTTransactor(address common.Address, transactor bind.ContractTransactor) (*MyOFTTransactor, error) {
	contract, err := bindMyOFT(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MyOFTTransactor{contract: contract}, nil
}

// NewMyOFTFilterer creates a new log filterer instance of MyOFT, bound to a specific deployed contract.
func NewMyOFTFilterer(address common.Address, filterer bind.ContractFilterer) (*MyOFTFilterer, error) {
	contract, err := bindMyOFT(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MyOFTFilterer{contract: contract}, nil
}

// bindMyOFT binds a generic wrapper to an already deployed contract.
func bindMyOFT(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MyOFTABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MyOFT *MyOFTRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MyOFT.Contract.MyOFTCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MyOFT *MyOFTRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyOFT.Contract.MyOFTTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MyOFT *MyOFTRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MyOFT.Contract.MyOFTTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MyOFT *MyOFTCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MyOFT.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MyOFT *MyOFTTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyOFT.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MyOFT *MyOFTTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MyOFT.Contract.contract.Transact(opts, method, params...)
}

// SEND is a free data retrieval call binding the contract method 0x1f5e1334.
//
// Solidity: function SEND() view returns(uint16)
func (_MyOFT *MyOFTCaller) SEND(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _MyOFT.contract.Call(opts, &out, "SEND")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// SEND is a free data retrieval call binding the contract method 0x1f5e1334.
//
// Solidity: function SEND() view returns(uint16)
func (_MyOFT *MyOFTSession) SEND() (uint16, error) {
	return _MyOFT.Contract.SEND(&_MyOFT.CallOpts)
}

// SEND is a free data retrieval call binding the contract method 0x1f5e1334.
//
// Solidity: function SEND() view returns(uint16)
func (_MyOFT *MyOFTCallerSession) SEND() (uint16, error) {
	return _MyOFT.Contract.SEND(&_MyOFT.CallOpts)
}

// SENDANDCALL is a free data retrieval call binding the contract method 0x134d4f25.
//
// Solidity: function SEND_AND_CALL() view returns(uint16)
func (_MyOFT *MyOFTCaller) SENDANDCALL(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _MyOFT.contract.Call(opts, &out, "SEND_AND_CALL")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// SENDANDCALL is a free data retrieval call binding the contract method 0x134d4f25.
//
// Solidity: function SEND_AND_CALL() view returns(uint16)
func (_MyOFT *MyOFTSession) SENDANDCALL() (uint16, error) {
	return _MyOFT.Contract.SENDANDCALL(&_MyOFT.CallOpts)
}

// SENDANDCALL is a free data retrieval call binding the contract method 0x134d4f25.
//
// Solidity: function SEND_AND_CALL() view returns(uint16)
func (_MyOFT *MyOFTCallerSession) SENDANDCALL() (uint16, error) {
	return _MyOFT.Contract.SENDANDCALL(&_MyOFT.CallOpts)
}

// AllowInitializePath is a free data retrieval call binding the contract method 0xff7bd03d.
//
// Solidity: function allowInitializePath((uint32,bytes32,uint64) origin) view returns(bool)
func (_MyOFT *MyOFTCaller) AllowInitializePath(opts *bind.CallOpts, origin Origin) (bool, error) {
	var out []interface{}
	err := _MyOFT.contract.Call(opts, &out, "allowInitializePath", origin)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowInitializePath is a free data retrieval call binding the contract method 0xff7bd03d.
//
// Solidity: function allowInitializePath((uint32,bytes32,uint64) origin) view returns(bool)
func (_MyOFT *MyOFTSession) AllowInitializePath(origin Origin) (bool, error) {
	return _MyOFT.Contract.AllowInitializePath(&_MyOFT.CallOpts, origin)
}

// AllowInitializePath is a free data retrieval call binding the contract method 0xff7bd03d.
//
// Solidity: function allowInitializePath((uint32,bytes32,uint64) origin) view returns(bool)
func (_MyOFT *MyOFTCallerSession) AllowInitializePath(origin Origin) (bool, error) {
	return _MyOFT.Contract.AllowInitializePath(&_MyOFT.CallOpts, origin)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_MyOFT *MyOFTCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MyOFT.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_MyOFT *MyOFTSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _MyOFT.Contract.Allowance(&_MyOFT.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_MyOFT *MyOFTCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _MyOFT.Contract.Allowance(&_MyOFT.CallOpts, owner, spender)
}

// ApprovalRequired is a free data retrieval call binding the contract method 0x9f68b964.
//
// Solidity: function approvalRequired() pure returns(bool)
func (_MyOFT *MyOFTCaller) ApprovalRequired(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MyOFT.contract.Call(opts, &out, "approvalRequired")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ApprovalRequired is a free data retrieval call binding the contract method 0x9f68b964.
//
// Solidity: function approvalRequired() pure returns(bool)
func (_MyOFT *MyOFTSession) ApprovalRequired() (bool, error) {
	return _MyOFT.Contract.ApprovalRequired(&_MyOFT.CallOpts)
}

// ApprovalRequired is a free data retrieval call binding the contract method 0x9f68b964.
//
// Solidity: function approvalRequired() pure returns(bool)
func (_MyOFT *MyOFTCallerSession) ApprovalRequired() (bool, error) {
	return _MyOFT.Contract.ApprovalRequired(&_MyOFT.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MyOFT *MyOFTCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MyOFT.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MyOFT *MyOFTSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _MyOFT.Contract.BalanceOf(&_MyOFT.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MyOFT *MyOFTCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _MyOFT.Contract.BalanceOf(&_MyOFT.CallOpts, account)
}

// CombineOptions is a free data retrieval call binding the contract method 0xbc70b354.
//
// Solidity: function combineOptions(uint32 _eid, uint16 _msgType, bytes _extraOptions) view returns(bytes)
func (_MyOFT *MyOFTCaller) CombineOptions(opts *bind.CallOpts, _eid uint32, _msgType uint16, _extraOptions []byte) ([]byte, error) {
	var out []interface{}
	err := _MyOFT.contract.Call(opts, &out, "combineOptions", _eid, _msgType, _extraOptions)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// CombineOptions is a free data retrieval call binding the contract method 0xbc70b354.
//
// Solidity: function combineOptions(uint32 _eid, uint16 _msgType, bytes _extraOptions) view returns(bytes)
func (_MyOFT *MyOFTSession) CombineOptions(_eid uint32, _msgType uint16, _extraOptions []byte) ([]byte, error) {
	return _MyOFT.Contract.CombineOptions(&_MyOFT.CallOpts, _eid, _msgType, _extraOptions)
}

// CombineOptions is a free data retrieval call binding the contract method 0xbc70b354.
//
// Solidity: function combineOptions(uint32 _eid, uint16 _msgType, bytes _extraOptions) view returns(bytes)
func (_MyOFT *MyOFTCallerSession) CombineOptions(_eid uint32, _msgType uint16, _extraOptions []byte) ([]byte, error) {
	return _MyOFT.Contract.CombineOptions(&_MyOFT.CallOpts, _eid, _msgType, _extraOptions)
}

// DecimalConversionRate is a free data retrieval call binding the contract method 0x963efcaa.
//
// Solidity: function decimalConversionRate() view returns(uint256)
func (_MyOFT *MyOFTCaller) DecimalConversionRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MyOFT.contract.Call(opts, &out, "decimalConversionRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DecimalConversionRate is a free data retrieval call binding the contract method 0x963efcaa.
//
// Solidity: function decimalConversionRate() view returns(uint256)
func (_MyOFT *MyOFTSession) DecimalConversionRate() (*big.Int, error) {
	return _MyOFT.Contract.DecimalConversionRate(&_MyOFT.CallOpts)
}

// DecimalConversionRate is a free data retrieval call binding the contract method 0x963efcaa.
//
// Solidity: function decimalConversionRate() view returns(uint256)
func (_MyOFT *MyOFTCallerSession) DecimalConversionRate() (*big.Int, error) {
	return _MyOFT.Contract.DecimalConversionRate(&_MyOFT.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MyOFT *MyOFTCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _MyOFT.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MyOFT *MyOFTSession) Decimals() (uint8, error) {
	return _MyOFT.Contract.Decimals(&_MyOFT.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MyOFT *MyOFTCallerSession) Decimals() (uint8, error) {
	return _MyOFT.Contract.Decimals(&_MyOFT.CallOpts)
}

// Endpoint is a free data retrieval call binding the contract method 0x5e280f11.
//
// Solidity: function endpoint() view returns(address)
func (_MyOFT *MyOFTCaller) Endpoint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MyOFT.contract.Call(opts, &out, "endpoint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Endpoint is a free data retrieval call binding the contract method 0x5e280f11.
//
// Solidity: function endpoint() view returns(address)
func (_MyOFT *MyOFTSession) Endpoint() (common.Address, error) {
	return _MyOFT.Contract.Endpoint(&_MyOFT.CallOpts)
}

// Endpoint is a free data retrieval call binding the contract method 0x5e280f11.
//
// Solidity: function endpoint() view returns(address)
func (_MyOFT *MyOFTCallerSession) Endpoint() (common.Address, error) {
	return _MyOFT.Contract.Endpoint(&_MyOFT.CallOpts)
}

// EnforcedOptions is a free data retrieval call binding the contract method 0x5535d461.
//
// Solidity: function enforcedOptions(uint32 eid, uint16 msgType) view returns(bytes enforcedOption)
func (_MyOFT *MyOFTCaller) EnforcedOptions(opts *bind.CallOpts, eid uint32, msgType uint16) ([]byte, error) {
	var out []interface{}
	err := _MyOFT.contract.Call(opts, &out, "enforcedOptions", eid, msgType)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EnforcedOptions is a free data retrieval call binding the contract method 0x5535d461.
//
// Solidity: function enforcedOptions(uint32 eid, uint16 msgType) view returns(bytes enforcedOption)
func (_MyOFT *MyOFTSession) EnforcedOptions(eid uint32, msgType uint16) ([]byte, error) {
	return _MyOFT.Contract.EnforcedOptions(&_MyOFT.CallOpts, eid, msgType)
}

// EnforcedOptions is a free data retrieval call binding the contract method 0x5535d461.
//
// Solidity: function enforcedOptions(uint32 eid, uint16 msgType) view returns(bytes enforcedOption)
func (_MyOFT *MyOFTCallerSession) EnforcedOptions(eid uint32, msgType uint16) ([]byte, error) {
	return _MyOFT.Contract.EnforcedOptions(&_MyOFT.CallOpts, eid, msgType)
}

// IsComposeMsgSender is a free data retrieval call binding the contract method 0x82413eac.
//
// Solidity: function isComposeMsgSender((uint32,bytes32,uint64) , bytes , address _sender) view returns(bool)
func (_MyOFT *MyOFTCaller) IsComposeMsgSender(opts *bind.CallOpts, arg0 Origin, arg1 []byte, _sender common.Address) (bool, error) {
	var out []interface{}
	err := _MyOFT.contract.Call(opts, &out, "isComposeMsgSender", arg0, arg1, _sender)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsComposeMsgSender is a free data retrieval call binding the contract method 0x82413eac.
//
// Solidity: function isComposeMsgSender((uint32,bytes32,uint64) , bytes , address _sender) view returns(bool)
func (_MyOFT *MyOFTSession) IsComposeMsgSender(arg0 Origin, arg1 []byte, _sender common.Address) (bool, error) {
	return _MyOFT.Contract.IsComposeMsgSender(&_MyOFT.CallOpts, arg0, arg1, _sender)
}

// IsComposeMsgSender is a free data retrieval call binding the contract method 0x82413eac.
//
// Solidity: function isComposeMsgSender((uint32,bytes32,uint64) , bytes , address _sender) view returns(bool)
func (_MyOFT *MyOFTCallerSession) IsComposeMsgSender(arg0 Origin, arg1 []byte, _sender common.Address) (bool, error) {
	return _MyOFT.Contract.IsComposeMsgSender(&_MyOFT.CallOpts, arg0, arg1, _sender)
}

// IsPeer is a free data retrieval call binding the contract method 0x5a0dfe4d.
//
// Solidity: function isPeer(uint32 _eid, bytes32 _peer) view returns(bool)
func (_MyOFT *MyOFTCaller) IsPeer(opts *bind.CallOpts, _eid uint32, _peer [32]byte) (bool, error) {
	var out []interface{}
	err := _MyOFT.contract.Call(opts, &out, "isPeer", _eid, _peer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPeer is a free data retrieval call binding the contract method 0x5a0dfe4d.
//
// Solidity: function isPeer(uint32 _eid, bytes32 _peer) view returns(bool)
func (_MyOFT *MyOFTSession) IsPeer(_eid uint32, _peer [32]byte) (bool, error) {
	return _MyOFT.Contract.IsPeer(&_MyOFT.CallOpts, _eid, _peer)
}

// IsPeer is a free data retrieval call binding the contract method 0x5a0dfe4d.
//
// Solidity: function isPeer(uint32 _eid, bytes32 _peer) view returns(bool)
func (_MyOFT *MyOFTCallerSession) IsPeer(_eid uint32, _peer [32]byte) (bool, error) {
	return _MyOFT.Contract.IsPeer(&_MyOFT.CallOpts, _eid, _peer)
}

// MsgInspector is a free data retrieval call binding the contract method 0x111ecdad.
//
// Solidity: function msgInspector() view returns(address)
func (_MyOFT *MyOFTCaller) MsgInspector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MyOFT.contract.Call(opts, &out, "msgInspector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MsgInspector is a free data retrieval call binding the contract method 0x111ecdad.
//
// Solidity: function msgInspector() view returns(address)
func (_MyOFT *MyOFTSession) MsgInspector() (common.Address, error) {
	return _MyOFT.Contract.MsgInspector(&_MyOFT.CallOpts)
}

// MsgInspector is a free data retrieval call binding the contract method 0x111ecdad.
//
// Solidity: function msgInspector() view returns(address)
func (_MyOFT *MyOFTCallerSession) MsgInspector() (common.Address, error) {
	return _MyOFT.Contract.MsgInspector(&_MyOFT.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MyOFT *MyOFTCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MyOFT.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MyOFT *MyOFTSession) Name() (string, error) {
	return _MyOFT.Contract.Name(&_MyOFT.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MyOFT *MyOFTCallerSession) Name() (string, error) {
	return _MyOFT.Contract.Name(&_MyOFT.CallOpts)
}

// NextNonce is a free data retrieval call binding the contract method 0x7d25a05e.
//
// Solidity: function nextNonce(uint32 , bytes32 ) view returns(uint64 nonce)
func (_MyOFT *MyOFTCaller) NextNonce(opts *bind.CallOpts, arg0 uint32, arg1 [32]byte) (uint64, error) {
	var out []interface{}
	err := _MyOFT.contract.Call(opts, &out, "nextNonce", arg0, arg1)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// NextNonce is a free data retrieval call binding the contract method 0x7d25a05e.
//
// Solidity: function nextNonce(uint32 , bytes32 ) view returns(uint64 nonce)
func (_MyOFT *MyOFTSession) NextNonce(arg0 uint32, arg1 [32]byte) (uint64, error) {
	return _MyOFT.Contract.NextNonce(&_MyOFT.CallOpts, arg0, arg1)
}

// NextNonce is a free data retrieval call binding the contract method 0x7d25a05e.
//
// Solidity: function nextNonce(uint32 , bytes32 ) view returns(uint64 nonce)
func (_MyOFT *MyOFTCallerSession) NextNonce(arg0 uint32, arg1 [32]byte) (uint64, error) {
	return _MyOFT.Contract.NextNonce(&_MyOFT.CallOpts, arg0, arg1)
}

// OApp is a free data retrieval call binding the contract method 0x52ae2879.
//
// Solidity: function oApp() view returns(address)
func (_MyOFT *MyOFTCaller) OApp(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MyOFT.contract.Call(opts, &out, "oApp")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OApp is a free data retrieval call binding the contract method 0x52ae2879.
//
// Solidity: function oApp() view returns(address)
func (_MyOFT *MyOFTSession) OApp() (common.Address, error) {
	return _MyOFT.Contract.OApp(&_MyOFT.CallOpts)
}

// OApp is a free data retrieval call binding the contract method 0x52ae2879.
//
// Solidity: function oApp() view returns(address)
func (_MyOFT *MyOFTCallerSession) OApp() (common.Address, error) {
	return _MyOFT.Contract.OApp(&_MyOFT.CallOpts)
}

// OAppVersion is a free data retrieval call binding the contract method 0x17442b70.
//
// Solidity: function oAppVersion() pure returns(uint64 senderVersion, uint64 receiverVersion)
func (_MyOFT *MyOFTCaller) OAppVersion(opts *bind.CallOpts) (struct {
	SenderVersion   uint64
	ReceiverVersion uint64
}, error) {
	var out []interface{}
	err := _MyOFT.contract.Call(opts, &out, "oAppVersion")

	outstruct := new(struct {
		SenderVersion   uint64
		ReceiverVersion uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SenderVersion = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.ReceiverVersion = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// OAppVersion is a free data retrieval call binding the contract method 0x17442b70.
//
// Solidity: function oAppVersion() pure returns(uint64 senderVersion, uint64 receiverVersion)
func (_MyOFT *MyOFTSession) OAppVersion() (struct {
	SenderVersion   uint64
	ReceiverVersion uint64
}, error) {
	return _MyOFT.Contract.OAppVersion(&_MyOFT.CallOpts)
}

// OAppVersion is a free data retrieval call binding the contract method 0x17442b70.
//
// Solidity: function oAppVersion() pure returns(uint64 senderVersion, uint64 receiverVersion)
func (_MyOFT *MyOFTCallerSession) OAppVersion() (struct {
	SenderVersion   uint64
	ReceiverVersion uint64
}, error) {
	return _MyOFT.Contract.OAppVersion(&_MyOFT.CallOpts)
}

// OftVersion is a free data retrieval call binding the contract method 0x156a0d0f.
//
// Solidity: function oftVersion() pure returns(bytes4 interfaceId, uint64 version)
func (_MyOFT *MyOFTCaller) OftVersion(opts *bind.CallOpts) (struct {
	InterfaceId [4]byte
	Version     uint64
}, error) {
	var out []interface{}
	err := _MyOFT.contract.Call(opts, &out, "oftVersion")

	outstruct := new(struct {
		InterfaceId [4]byte
		Version     uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.InterfaceId = *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)
	outstruct.Version = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// OftVersion is a free data retrieval call binding the contract method 0x156a0d0f.
//
// Solidity: function oftVersion() pure returns(bytes4 interfaceId, uint64 version)
func (_MyOFT *MyOFTSession) OftVersion() (struct {
	InterfaceId [4]byte
	Version     uint64
}, error) {
	return _MyOFT.Contract.OftVersion(&_MyOFT.CallOpts)
}

// OftVersion is a free data retrieval call binding the contract method 0x156a0d0f.
//
// Solidity: function oftVersion() pure returns(bytes4 interfaceId, uint64 version)
func (_MyOFT *MyOFTCallerSession) OftVersion() (struct {
	InterfaceId [4]byte
	Version     uint64
}, error) {
	return _MyOFT.Contract.OftVersion(&_MyOFT.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MyOFT *MyOFTCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MyOFT.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MyOFT *MyOFTSession) Owner() (common.Address, error) {
	return _MyOFT.Contract.Owner(&_MyOFT.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MyOFT *MyOFTCallerSession) Owner() (common.Address, error) {
	return _MyOFT.Contract.Owner(&_MyOFT.CallOpts)
}

// Peers is a free data retrieval call binding the contract method 0xbb0b6a53.
//
// Solidity: function peers(uint32 eid) view returns(bytes32 peer)
func (_MyOFT *MyOFTCaller) Peers(opts *bind.CallOpts, eid uint32) ([32]byte, error) {
	var out []interface{}
	err := _MyOFT.contract.Call(opts, &out, "peers", eid)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Peers is a free data retrieval call binding the contract method 0xbb0b6a53.
//
// Solidity: function peers(uint32 eid) view returns(bytes32 peer)
func (_MyOFT *MyOFTSession) Peers(eid uint32) ([32]byte, error) {
	return _MyOFT.Contract.Peers(&_MyOFT.CallOpts, eid)
}

// Peers is a free data retrieval call binding the contract method 0xbb0b6a53.
//
// Solidity: function peers(uint32 eid) view returns(bytes32 peer)
func (_MyOFT *MyOFTCallerSession) Peers(eid uint32) ([32]byte, error) {
	return _MyOFT.Contract.Peers(&_MyOFT.CallOpts, eid)
}

// PreCrime is a free data retrieval call binding the contract method 0xb731ea0a.
//
// Solidity: function preCrime() view returns(address)
func (_MyOFT *MyOFTCaller) PreCrime(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MyOFT.contract.Call(opts, &out, "preCrime")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PreCrime is a free data retrieval call binding the contract method 0xb731ea0a.
//
// Solidity: function preCrime() view returns(address)
func (_MyOFT *MyOFTSession) PreCrime() (common.Address, error) {
	return _MyOFT.Contract.PreCrime(&_MyOFT.CallOpts)
}

// PreCrime is a free data retrieval call binding the contract method 0xb731ea0a.
//
// Solidity: function preCrime() view returns(address)
func (_MyOFT *MyOFTCallerSession) PreCrime() (common.Address, error) {
	return _MyOFT.Contract.PreCrime(&_MyOFT.CallOpts)
}

// QuoteOFT is a free data retrieval call binding the contract method 0x0d35b415.
//
// Solidity: function quoteOFT((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam) view returns((uint256,uint256) oftLimit, (int256,string)[] oftFeeDetails, (uint256,uint256) oftReceipt)
func (_MyOFT *MyOFTCaller) QuoteOFT(opts *bind.CallOpts, _sendParam SendParam) (struct {
	OftLimit      OFTLimit
	OftFeeDetails []OFTFeeDetail
	OftReceipt    OFTReceipt
}, error) {
	var out []interface{}
	err := _MyOFT.contract.Call(opts, &out, "quoteOFT", _sendParam)

	outstruct := new(struct {
		OftLimit      OFTLimit
		OftFeeDetails []OFTFeeDetail
		OftReceipt    OFTReceipt
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OftLimit = *abi.ConvertType(out[0], new(OFTLimit)).(*OFTLimit)
	outstruct.OftFeeDetails = *abi.ConvertType(out[1], new([]OFTFeeDetail)).(*[]OFTFeeDetail)
	outstruct.OftReceipt = *abi.ConvertType(out[2], new(OFTReceipt)).(*OFTReceipt)

	return *outstruct, err

}

// QuoteOFT is a free data retrieval call binding the contract method 0x0d35b415.
//
// Solidity: function quoteOFT((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam) view returns((uint256,uint256) oftLimit, (int256,string)[] oftFeeDetails, (uint256,uint256) oftReceipt)
func (_MyOFT *MyOFTSession) QuoteOFT(_sendParam SendParam) (struct {
	OftLimit      OFTLimit
	OftFeeDetails []OFTFeeDetail
	OftReceipt    OFTReceipt
}, error) {
	return _MyOFT.Contract.QuoteOFT(&_MyOFT.CallOpts, _sendParam)
}

// QuoteOFT is a free data retrieval call binding the contract method 0x0d35b415.
//
// Solidity: function quoteOFT((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam) view returns((uint256,uint256) oftLimit, (int256,string)[] oftFeeDetails, (uint256,uint256) oftReceipt)
func (_MyOFT *MyOFTCallerSession) QuoteOFT(_sendParam SendParam) (struct {
	OftLimit      OFTLimit
	OftFeeDetails []OFTFeeDetail
	OftReceipt    OFTReceipt
}, error) {
	return _MyOFT.Contract.QuoteOFT(&_MyOFT.CallOpts, _sendParam)
}

// QuoteSend is a free data retrieval call binding the contract method 0x3b6f743b.
//
// Solidity: function quoteSend((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam, bool _payInLzToken) view returns((uint256,uint256) msgFee)
func (_MyOFT *MyOFTCaller) QuoteSend(opts *bind.CallOpts, _sendParam SendParam, _payInLzToken bool) (MessagingFee, error) {
	var out []interface{}
	err := _MyOFT.contract.Call(opts, &out, "quoteSend", _sendParam, _payInLzToken)

	if err != nil {
		return *new(MessagingFee), err
	}

	out0 := *abi.ConvertType(out[0], new(MessagingFee)).(*MessagingFee)

	return out0, err

}

// QuoteSend is a free data retrieval call binding the contract method 0x3b6f743b.
//
// Solidity: function quoteSend((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam, bool _payInLzToken) view returns((uint256,uint256) msgFee)
func (_MyOFT *MyOFTSession) QuoteSend(_sendParam SendParam, _payInLzToken bool) (MessagingFee, error) {
	return _MyOFT.Contract.QuoteSend(&_MyOFT.CallOpts, _sendParam, _payInLzToken)
}

// QuoteSend is a free data retrieval call binding the contract method 0x3b6f743b.
//
// Solidity: function quoteSend((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam, bool _payInLzToken) view returns((uint256,uint256) msgFee)
func (_MyOFT *MyOFTCallerSession) QuoteSend(_sendParam SendParam, _payInLzToken bool) (MessagingFee, error) {
	return _MyOFT.Contract.QuoteSend(&_MyOFT.CallOpts, _sendParam, _payInLzToken)
}

// SharedDecimals is a free data retrieval call binding the contract method 0x857749b0.
//
// Solidity: function sharedDecimals() view returns(uint8)
func (_MyOFT *MyOFTCaller) SharedDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _MyOFT.contract.Call(opts, &out, "sharedDecimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// SharedDecimals is a free data retrieval call binding the contract method 0x857749b0.
//
// Solidity: function sharedDecimals() view returns(uint8)
func (_MyOFT *MyOFTSession) SharedDecimals() (uint8, error) {
	return _MyOFT.Contract.SharedDecimals(&_MyOFT.CallOpts)
}

// SharedDecimals is a free data retrieval call binding the contract method 0x857749b0.
//
// Solidity: function sharedDecimals() view returns(uint8)
func (_MyOFT *MyOFTCallerSession) SharedDecimals() (uint8, error) {
	return _MyOFT.Contract.SharedDecimals(&_MyOFT.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MyOFT *MyOFTCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MyOFT.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MyOFT *MyOFTSession) Symbol() (string, error) {
	return _MyOFT.Contract.Symbol(&_MyOFT.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MyOFT *MyOFTCallerSession) Symbol() (string, error) {
	return _MyOFT.Contract.Symbol(&_MyOFT.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_MyOFT *MyOFTCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MyOFT.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_MyOFT *MyOFTSession) Token() (common.Address, error) {
	return _MyOFT.Contract.Token(&_MyOFT.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_MyOFT *MyOFTCallerSession) Token() (common.Address, error) {
	return _MyOFT.Contract.Token(&_MyOFT.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MyOFT *MyOFTCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MyOFT.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MyOFT *MyOFTSession) TotalSupply() (*big.Int, error) {
	return _MyOFT.Contract.TotalSupply(&_MyOFT.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MyOFT *MyOFTCallerSession) TotalSupply() (*big.Int, error) {
	return _MyOFT.Contract.TotalSupply(&_MyOFT.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_MyOFT *MyOFTTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _MyOFT.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_MyOFT *MyOFTSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _MyOFT.Contract.Approve(&_MyOFT.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_MyOFT *MyOFTTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _MyOFT.Contract.Approve(&_MyOFT.TransactOpts, spender, value)
}

// LzReceive is a paid mutator transaction binding the contract method 0x13137d65.
//
// Solidity: function lzReceive((uint32,bytes32,uint64) _origin, bytes32 _guid, bytes _message, address _executor, bytes _extraData) payable returns()
func (_MyOFT *MyOFTTransactor) LzReceive(opts *bind.TransactOpts, _origin Origin, _guid [32]byte, _message []byte, _executor common.Address, _extraData []byte) (*types.Transaction, error) {
	return _MyOFT.contract.Transact(opts, "lzReceive", _origin, _guid, _message, _executor, _extraData)
}

// LzReceive is a paid mutator transaction binding the contract method 0x13137d65.
//
// Solidity: function lzReceive((uint32,bytes32,uint64) _origin, bytes32 _guid, bytes _message, address _executor, bytes _extraData) payable returns()
func (_MyOFT *MyOFTSession) LzReceive(_origin Origin, _guid [32]byte, _message []byte, _executor common.Address, _extraData []byte) (*types.Transaction, error) {
	return _MyOFT.Contract.LzReceive(&_MyOFT.TransactOpts, _origin, _guid, _message, _executor, _extraData)
}

// LzReceive is a paid mutator transaction binding the contract method 0x13137d65.
//
// Solidity: function lzReceive((uint32,bytes32,uint64) _origin, bytes32 _guid, bytes _message, address _executor, bytes _extraData) payable returns()
func (_MyOFT *MyOFTTransactorSession) LzReceive(_origin Origin, _guid [32]byte, _message []byte, _executor common.Address, _extraData []byte) (*types.Transaction, error) {
	return _MyOFT.Contract.LzReceive(&_MyOFT.TransactOpts, _origin, _guid, _message, _executor, _extraData)
}

// LzReceiveAndRevert is a paid mutator transaction binding the contract method 0xbd815db0.
//
// Solidity: function lzReceiveAndRevert(((uint32,bytes32,uint64),uint32,address,bytes32,uint256,address,bytes,bytes)[] _packets) payable returns()
func (_MyOFT *MyOFTTransactor) LzReceiveAndRevert(opts *bind.TransactOpts, _packets []InboundPacket) (*types.Transaction, error) {
	return _MyOFT.contract.Transact(opts, "lzReceiveAndRevert", _packets)
}

// LzReceiveAndRevert is a paid mutator transaction binding the contract method 0xbd815db0.
//
// Solidity: function lzReceiveAndRevert(((uint32,bytes32,uint64),uint32,address,bytes32,uint256,address,bytes,bytes)[] _packets) payable returns()
func (_MyOFT *MyOFTSession) LzReceiveAndRevert(_packets []InboundPacket) (*types.Transaction, error) {
	return _MyOFT.Contract.LzReceiveAndRevert(&_MyOFT.TransactOpts, _packets)
}

// LzReceiveAndRevert is a paid mutator transaction binding the contract method 0xbd815db0.
//
// Solidity: function lzReceiveAndRevert(((uint32,bytes32,uint64),uint32,address,bytes32,uint256,address,bytes,bytes)[] _packets) payable returns()
func (_MyOFT *MyOFTTransactorSession) LzReceiveAndRevert(_packets []InboundPacket) (*types.Transaction, error) {
	return _MyOFT.Contract.LzReceiveAndRevert(&_MyOFT.TransactOpts, _packets)
}

// LzReceiveSimulate is a paid mutator transaction binding the contract method 0xd045a0dc.
//
// Solidity: function lzReceiveSimulate((uint32,bytes32,uint64) _origin, bytes32 _guid, bytes _message, address _executor, bytes _extraData) payable returns()
func (_MyOFT *MyOFTTransactor) LzReceiveSimulate(opts *bind.TransactOpts, _origin Origin, _guid [32]byte, _message []byte, _executor common.Address, _extraData []byte) (*types.Transaction, error) {
	return _MyOFT.contract.Transact(opts, "lzReceiveSimulate", _origin, _guid, _message, _executor, _extraData)
}

// LzReceiveSimulate is a paid mutator transaction binding the contract method 0xd045a0dc.
//
// Solidity: function lzReceiveSimulate((uint32,bytes32,uint64) _origin, bytes32 _guid, bytes _message, address _executor, bytes _extraData) payable returns()
func (_MyOFT *MyOFTSession) LzReceiveSimulate(_origin Origin, _guid [32]byte, _message []byte, _executor common.Address, _extraData []byte) (*types.Transaction, error) {
	return _MyOFT.Contract.LzReceiveSimulate(&_MyOFT.TransactOpts, _origin, _guid, _message, _executor, _extraData)
}

// LzReceiveSimulate is a paid mutator transaction binding the contract method 0xd045a0dc.
//
// Solidity: function lzReceiveSimulate((uint32,bytes32,uint64) _origin, bytes32 _guid, bytes _message, address _executor, bytes _extraData) payable returns()
func (_MyOFT *MyOFTTransactorSession) LzReceiveSimulate(_origin Origin, _guid [32]byte, _message []byte, _executor common.Address, _extraData []byte) (*types.Transaction, error) {
	return _MyOFT.Contract.LzReceiveSimulate(&_MyOFT.TransactOpts, _origin, _guid, _message, _executor, _extraData)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_MyOFT *MyOFTTransactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MyOFT.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_MyOFT *MyOFTSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MyOFT.Contract.Mint(&_MyOFT.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_MyOFT *MyOFTTransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MyOFT.Contract.Mint(&_MyOFT.TransactOpts, to, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MyOFT *MyOFTTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyOFT.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MyOFT *MyOFTSession) RenounceOwnership() (*types.Transaction, error) {
	return _MyOFT.Contract.RenounceOwnership(&_MyOFT.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MyOFT *MyOFTTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MyOFT.Contract.RenounceOwnership(&_MyOFT.TransactOpts)
}

// Send is a paid mutator transaction binding the contract method 0xc7c7f5b3.
//
// Solidity: function send((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam, (uint256,uint256) _fee, address _refundAddress) payable returns((bytes32,uint64,(uint256,uint256)) msgReceipt, (uint256,uint256) oftReceipt)
func (_MyOFT *MyOFTTransactor) Send(opts *bind.TransactOpts, _sendParam SendParam, _fee MessagingFee, _refundAddress common.Address) (*types.Transaction, error) {
	return _MyOFT.contract.Transact(opts, "send", _sendParam, _fee, _refundAddress)
}

// Send is a paid mutator transaction binding the contract method 0xc7c7f5b3.
//
// Solidity: function send((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam, (uint256,uint256) _fee, address _refundAddress) payable returns((bytes32,uint64,(uint256,uint256)) msgReceipt, (uint256,uint256) oftReceipt)
func (_MyOFT *MyOFTSession) Send(_sendParam SendParam, _fee MessagingFee, _refundAddress common.Address) (*types.Transaction, error) {
	return _MyOFT.Contract.Send(&_MyOFT.TransactOpts, _sendParam, _fee, _refundAddress)
}

// Send is a paid mutator transaction binding the contract method 0xc7c7f5b3.
//
// Solidity: function send((uint32,bytes32,uint256,uint256,bytes,bytes,bytes) _sendParam, (uint256,uint256) _fee, address _refundAddress) payable returns((bytes32,uint64,(uint256,uint256)) msgReceipt, (uint256,uint256) oftReceipt)
func (_MyOFT *MyOFTTransactorSession) Send(_sendParam SendParam, _fee MessagingFee, _refundAddress common.Address) (*types.Transaction, error) {
	return _MyOFT.Contract.Send(&_MyOFT.TransactOpts, _sendParam, _fee, _refundAddress)
}

// SetDelegate is a paid mutator transaction binding the contract method 0xca5eb5e1.
//
// Solidity: function setDelegate(address _delegate) returns()
func (_MyOFT *MyOFTTransactor) SetDelegate(opts *bind.TransactOpts, _delegate common.Address) (*types.Transaction, error) {
	return _MyOFT.contract.Transact(opts, "setDelegate", _delegate)
}

// SetDelegate is a paid mutator transaction binding the contract method 0xca5eb5e1.
//
// Solidity: function setDelegate(address _delegate) returns()
func (_MyOFT *MyOFTSession) SetDelegate(_delegate common.Address) (*types.Transaction, error) {
	return _MyOFT.Contract.SetDelegate(&_MyOFT.TransactOpts, _delegate)
}

// SetDelegate is a paid mutator transaction binding the contract method 0xca5eb5e1.
//
// Solidity: function setDelegate(address _delegate) returns()
func (_MyOFT *MyOFTTransactorSession) SetDelegate(_delegate common.Address) (*types.Transaction, error) {
	return _MyOFT.Contract.SetDelegate(&_MyOFT.TransactOpts, _delegate)
}

// SetEnforcedOptions is a paid mutator transaction binding the contract method 0xb98bd070.
//
// Solidity: function setEnforcedOptions((uint32,uint16,bytes)[] _enforcedOptions) returns()
func (_MyOFT *MyOFTTransactor) SetEnforcedOptions(opts *bind.TransactOpts, _enforcedOptions []EnforcedOptionParam) (*types.Transaction, error) {
	return _MyOFT.contract.Transact(opts, "setEnforcedOptions", _enforcedOptions)
}

// SetEnforcedOptions is a paid mutator transaction binding the contract method 0xb98bd070.
//
// Solidity: function setEnforcedOptions((uint32,uint16,bytes)[] _enforcedOptions) returns()
func (_MyOFT *MyOFTSession) SetEnforcedOptions(_enforcedOptions []EnforcedOptionParam) (*types.Transaction, error) {
	return _MyOFT.Contract.SetEnforcedOptions(&_MyOFT.TransactOpts, _enforcedOptions)
}

// SetEnforcedOptions is a paid mutator transaction binding the contract method 0xb98bd070.
//
// Solidity: function setEnforcedOptions((uint32,uint16,bytes)[] _enforcedOptions) returns()
func (_MyOFT *MyOFTTransactorSession) SetEnforcedOptions(_enforcedOptions []EnforcedOptionParam) (*types.Transaction, error) {
	return _MyOFT.Contract.SetEnforcedOptions(&_MyOFT.TransactOpts, _enforcedOptions)
}

// SetMsgInspector is a paid mutator transaction binding the contract method 0x6fc1b31e.
//
// Solidity: function setMsgInspector(address _msgInspector) returns()
func (_MyOFT *MyOFTTransactor) SetMsgInspector(opts *bind.TransactOpts, _msgInspector common.Address) (*types.Transaction, error) {
	return _MyOFT.contract.Transact(opts, "setMsgInspector", _msgInspector)
}

// SetMsgInspector is a paid mutator transaction binding the contract method 0x6fc1b31e.
//
// Solidity: function setMsgInspector(address _msgInspector) returns()
func (_MyOFT *MyOFTSession) SetMsgInspector(_msgInspector common.Address) (*types.Transaction, error) {
	return _MyOFT.Contract.SetMsgInspector(&_MyOFT.TransactOpts, _msgInspector)
}

// SetMsgInspector is a paid mutator transaction binding the contract method 0x6fc1b31e.
//
// Solidity: function setMsgInspector(address _msgInspector) returns()
func (_MyOFT *MyOFTTransactorSession) SetMsgInspector(_msgInspector common.Address) (*types.Transaction, error) {
	return _MyOFT.Contract.SetMsgInspector(&_MyOFT.TransactOpts, _msgInspector)
}

// SetPeer is a paid mutator transaction binding the contract method 0x3400288b.
//
// Solidity: function setPeer(uint32 _eid, bytes32 _peer) returns()
func (_MyOFT *MyOFTTransactor) SetPeer(opts *bind.TransactOpts, _eid uint32, _peer [32]byte) (*types.Transaction, error) {
	return _MyOFT.contract.Transact(opts, "setPeer", _eid, _peer)
}

// SetPeer is a paid mutator transaction binding the contract method 0x3400288b.
//
// Solidity: function setPeer(uint32 _eid, bytes32 _peer) returns()
func (_MyOFT *MyOFTSession) SetPeer(_eid uint32, _peer [32]byte) (*types.Transaction, error) {
	return _MyOFT.Contract.SetPeer(&_MyOFT.TransactOpts, _eid, _peer)
}

// SetPeer is a paid mutator transaction binding the contract method 0x3400288b.
//
// Solidity: function setPeer(uint32 _eid, bytes32 _peer) returns()
func (_MyOFT *MyOFTTransactorSession) SetPeer(_eid uint32, _peer [32]byte) (*types.Transaction, error) {
	return _MyOFT.Contract.SetPeer(&_MyOFT.TransactOpts, _eid, _peer)
}

// SetPreCrime is a paid mutator transaction binding the contract method 0xd4243885.
//
// Solidity: function setPreCrime(address _preCrime) returns()
func (_MyOFT *MyOFTTransactor) SetPreCrime(opts *bind.TransactOpts, _preCrime common.Address) (*types.Transaction, error) {
	return _MyOFT.contract.Transact(opts, "setPreCrime", _preCrime)
}

// SetPreCrime is a paid mutator transaction binding the contract method 0xd4243885.
//
// Solidity: function setPreCrime(address _preCrime) returns()
func (_MyOFT *MyOFTSession) SetPreCrime(_preCrime common.Address) (*types.Transaction, error) {
	return _MyOFT.Contract.SetPreCrime(&_MyOFT.TransactOpts, _preCrime)
}

// SetPreCrime is a paid mutator transaction binding the contract method 0xd4243885.
//
// Solidity: function setPreCrime(address _preCrime) returns()
func (_MyOFT *MyOFTTransactorSession) SetPreCrime(_preCrime common.Address) (*types.Transaction, error) {
	return _MyOFT.Contract.SetPreCrime(&_MyOFT.TransactOpts, _preCrime)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_MyOFT *MyOFTTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MyOFT.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_MyOFT *MyOFTSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MyOFT.Contract.Transfer(&_MyOFT.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_MyOFT *MyOFTTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MyOFT.Contract.Transfer(&_MyOFT.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_MyOFT *MyOFTTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MyOFT.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_MyOFT *MyOFTSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MyOFT.Contract.TransferFrom(&_MyOFT.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_MyOFT *MyOFTTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MyOFT.Contract.TransferFrom(&_MyOFT.TransactOpts, from, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MyOFT *MyOFTTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MyOFT.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MyOFT *MyOFTSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MyOFT.Contract.TransferOwnership(&_MyOFT.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MyOFT *MyOFTTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MyOFT.Contract.TransferOwnership(&_MyOFT.TransactOpts, newOwner)
}

// MyOFTApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the MyOFT contract.
type MyOFTApprovalIterator struct {
	Event *MyOFTApproval // Event containing the contract specifics and raw log

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
func (it *MyOFTApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyOFTApproval)
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
		it.Event = new(MyOFTApproval)
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
func (it *MyOFTApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyOFTApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyOFTApproval represents a Approval event raised by the MyOFT contract.
type MyOFTApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MyOFT *MyOFTFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*MyOFTApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MyOFT.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &MyOFTApprovalIterator{contract: _MyOFT.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MyOFT *MyOFTFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MyOFTApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MyOFT.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyOFTApproval)
				if err := _MyOFT.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_MyOFT *MyOFTFilterer) ParseApproval(log types.Log) (*MyOFTApproval, error) {
	event := new(MyOFTApproval)
	if err := _MyOFT.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MyOFTEnforcedOptionSetIterator is returned from FilterEnforcedOptionSet and is used to iterate over the raw logs and unpacked data for EnforcedOptionSet events raised by the MyOFT contract.
type MyOFTEnforcedOptionSetIterator struct {
	Event *MyOFTEnforcedOptionSet // Event containing the contract specifics and raw log

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
func (it *MyOFTEnforcedOptionSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyOFTEnforcedOptionSet)
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
		it.Event = new(MyOFTEnforcedOptionSet)
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
func (it *MyOFTEnforcedOptionSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyOFTEnforcedOptionSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyOFTEnforcedOptionSet represents a EnforcedOptionSet event raised by the MyOFT contract.
type MyOFTEnforcedOptionSet struct {
	EnforcedOptions []EnforcedOptionParam
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterEnforcedOptionSet is a free log retrieval operation binding the contract event 0xbe4864a8e820971c0247f5992e2da559595f7bf076a21cb5928d443d2a13b674.
//
// Solidity: event EnforcedOptionSet((uint32,uint16,bytes)[] _enforcedOptions)
func (_MyOFT *MyOFTFilterer) FilterEnforcedOptionSet(opts *bind.FilterOpts) (*MyOFTEnforcedOptionSetIterator, error) {

	logs, sub, err := _MyOFT.contract.FilterLogs(opts, "EnforcedOptionSet")
	if err != nil {
		return nil, err
	}
	return &MyOFTEnforcedOptionSetIterator{contract: _MyOFT.contract, event: "EnforcedOptionSet", logs: logs, sub: sub}, nil
}

// WatchEnforcedOptionSet is a free log subscription operation binding the contract event 0xbe4864a8e820971c0247f5992e2da559595f7bf076a21cb5928d443d2a13b674.
//
// Solidity: event EnforcedOptionSet((uint32,uint16,bytes)[] _enforcedOptions)
func (_MyOFT *MyOFTFilterer) WatchEnforcedOptionSet(opts *bind.WatchOpts, sink chan<- *MyOFTEnforcedOptionSet) (event.Subscription, error) {

	logs, sub, err := _MyOFT.contract.WatchLogs(opts, "EnforcedOptionSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyOFTEnforcedOptionSet)
				if err := _MyOFT.contract.UnpackLog(event, "EnforcedOptionSet", log); err != nil {
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

// ParseEnforcedOptionSet is a log parse operation binding the contract event 0xbe4864a8e820971c0247f5992e2da559595f7bf076a21cb5928d443d2a13b674.
//
// Solidity: event EnforcedOptionSet((uint32,uint16,bytes)[] _enforcedOptions)
func (_MyOFT *MyOFTFilterer) ParseEnforcedOptionSet(log types.Log) (*MyOFTEnforcedOptionSet, error) {
	event := new(MyOFTEnforcedOptionSet)
	if err := _MyOFT.contract.UnpackLog(event, "EnforcedOptionSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MyOFTMsgInspectorSetIterator is returned from FilterMsgInspectorSet and is used to iterate over the raw logs and unpacked data for MsgInspectorSet events raised by the MyOFT contract.
type MyOFTMsgInspectorSetIterator struct {
	Event *MyOFTMsgInspectorSet // Event containing the contract specifics and raw log

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
func (it *MyOFTMsgInspectorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyOFTMsgInspectorSet)
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
		it.Event = new(MyOFTMsgInspectorSet)
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
func (it *MyOFTMsgInspectorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyOFTMsgInspectorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyOFTMsgInspectorSet represents a MsgInspectorSet event raised by the MyOFT contract.
type MyOFTMsgInspectorSet struct {
	Inspector common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMsgInspectorSet is a free log retrieval operation binding the contract event 0xf0be4f1e87349231d80c36b33f9e8639658eeaf474014dee15a3e6a4d4414197.
//
// Solidity: event MsgInspectorSet(address inspector)
func (_MyOFT *MyOFTFilterer) FilterMsgInspectorSet(opts *bind.FilterOpts) (*MyOFTMsgInspectorSetIterator, error) {

	logs, sub, err := _MyOFT.contract.FilterLogs(opts, "MsgInspectorSet")
	if err != nil {
		return nil, err
	}
	return &MyOFTMsgInspectorSetIterator{contract: _MyOFT.contract, event: "MsgInspectorSet", logs: logs, sub: sub}, nil
}

// WatchMsgInspectorSet is a free log subscription operation binding the contract event 0xf0be4f1e87349231d80c36b33f9e8639658eeaf474014dee15a3e6a4d4414197.
//
// Solidity: event MsgInspectorSet(address inspector)
func (_MyOFT *MyOFTFilterer) WatchMsgInspectorSet(opts *bind.WatchOpts, sink chan<- *MyOFTMsgInspectorSet) (event.Subscription, error) {

	logs, sub, err := _MyOFT.contract.WatchLogs(opts, "MsgInspectorSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyOFTMsgInspectorSet)
				if err := _MyOFT.contract.UnpackLog(event, "MsgInspectorSet", log); err != nil {
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

// ParseMsgInspectorSet is a log parse operation binding the contract event 0xf0be4f1e87349231d80c36b33f9e8639658eeaf474014dee15a3e6a4d4414197.
//
// Solidity: event MsgInspectorSet(address inspector)
func (_MyOFT *MyOFTFilterer) ParseMsgInspectorSet(log types.Log) (*MyOFTMsgInspectorSet, error) {
	event := new(MyOFTMsgInspectorSet)
	if err := _MyOFT.contract.UnpackLog(event, "MsgInspectorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MyOFTOFTReceivedIterator is returned from FilterOFTReceived and is used to iterate over the raw logs and unpacked data for OFTReceived events raised by the MyOFT contract.
type MyOFTOFTReceivedIterator struct {
	Event *MyOFTOFTReceived // Event containing the contract specifics and raw log

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
func (it *MyOFTOFTReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyOFTOFTReceived)
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
		it.Event = new(MyOFTOFTReceived)
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
func (it *MyOFTOFTReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyOFTOFTReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyOFTOFTReceived represents a OFTReceived event raised by the MyOFT contract.
type MyOFTOFTReceived struct {
	Guid             [32]byte
	SrcEid           uint32
	ToAddress        common.Address
	AmountReceivedLD *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterOFTReceived is a free log retrieval operation binding the contract event 0xefed6d3500546b29533b128a29e3a94d70788727f0507505ac12eaf2e578fd9c.
//
// Solidity: event OFTReceived(bytes32 indexed guid, uint32 srcEid, address indexed toAddress, uint256 amountReceivedLD)
func (_MyOFT *MyOFTFilterer) FilterOFTReceived(opts *bind.FilterOpts, guid [][32]byte, toAddress []common.Address) (*MyOFTOFTReceivedIterator, error) {

	var guidRule []interface{}
	for _, guidItem := range guid {
		guidRule = append(guidRule, guidItem)
	}

	var toAddressRule []interface{}
	for _, toAddressItem := range toAddress {
		toAddressRule = append(toAddressRule, toAddressItem)
	}

	logs, sub, err := _MyOFT.contract.FilterLogs(opts, "OFTReceived", guidRule, toAddressRule)
	if err != nil {
		return nil, err
	}
	return &MyOFTOFTReceivedIterator{contract: _MyOFT.contract, event: "OFTReceived", logs: logs, sub: sub}, nil
}

// WatchOFTReceived is a free log subscription operation binding the contract event 0xefed6d3500546b29533b128a29e3a94d70788727f0507505ac12eaf2e578fd9c.
//
// Solidity: event OFTReceived(bytes32 indexed guid, uint32 srcEid, address indexed toAddress, uint256 amountReceivedLD)
func (_MyOFT *MyOFTFilterer) WatchOFTReceived(opts *bind.WatchOpts, sink chan<- *MyOFTOFTReceived, guid [][32]byte, toAddress []common.Address) (event.Subscription, error) {

	var guidRule []interface{}
	for _, guidItem := range guid {
		guidRule = append(guidRule, guidItem)
	}

	var toAddressRule []interface{}
	for _, toAddressItem := range toAddress {
		toAddressRule = append(toAddressRule, toAddressItem)
	}

	logs, sub, err := _MyOFT.contract.WatchLogs(opts, "OFTReceived", guidRule, toAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyOFTOFTReceived)
				if err := _MyOFT.contract.UnpackLog(event, "OFTReceived", log); err != nil {
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

// ParseOFTReceived is a log parse operation binding the contract event 0xefed6d3500546b29533b128a29e3a94d70788727f0507505ac12eaf2e578fd9c.
//
// Solidity: event OFTReceived(bytes32 indexed guid, uint32 srcEid, address indexed toAddress, uint256 amountReceivedLD)
func (_MyOFT *MyOFTFilterer) ParseOFTReceived(log types.Log) (*MyOFTOFTReceived, error) {
	event := new(MyOFTOFTReceived)
	if err := _MyOFT.contract.UnpackLog(event, "OFTReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MyOFTOFTSentIterator is returned from FilterOFTSent and is used to iterate over the raw logs and unpacked data for OFTSent events raised by the MyOFT contract.
type MyOFTOFTSentIterator struct {
	Event *MyOFTOFTSent // Event containing the contract specifics and raw log

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
func (it *MyOFTOFTSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyOFTOFTSent)
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
		it.Event = new(MyOFTOFTSent)
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
func (it *MyOFTOFTSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyOFTOFTSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyOFTOFTSent represents a OFTSent event raised by the MyOFT contract.
type MyOFTOFTSent struct {
	Guid             [32]byte
	DstEid           uint32
	FromAddress      common.Address
	AmountSentLD     *big.Int
	AmountReceivedLD *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterOFTSent is a free log retrieval operation binding the contract event 0x85496b760a4b7f8d66384b9df21b381f5d1b1e79f229a47aaf4c232edc2fe59a.
//
// Solidity: event OFTSent(bytes32 indexed guid, uint32 dstEid, address indexed fromAddress, uint256 amountSentLD, uint256 amountReceivedLD)
func (_MyOFT *MyOFTFilterer) FilterOFTSent(opts *bind.FilterOpts, guid [][32]byte, fromAddress []common.Address) (*MyOFTOFTSentIterator, error) {

	var guidRule []interface{}
	for _, guidItem := range guid {
		guidRule = append(guidRule, guidItem)
	}

	var fromAddressRule []interface{}
	for _, fromAddressItem := range fromAddress {
		fromAddressRule = append(fromAddressRule, fromAddressItem)
	}

	logs, sub, err := _MyOFT.contract.FilterLogs(opts, "OFTSent", guidRule, fromAddressRule)
	if err != nil {
		return nil, err
	}
	return &MyOFTOFTSentIterator{contract: _MyOFT.contract, event: "OFTSent", logs: logs, sub: sub}, nil
}

// WatchOFTSent is a free log subscription operation binding the contract event 0x85496b760a4b7f8d66384b9df21b381f5d1b1e79f229a47aaf4c232edc2fe59a.
//
// Solidity: event OFTSent(bytes32 indexed guid, uint32 dstEid, address indexed fromAddress, uint256 amountSentLD, uint256 amountReceivedLD)
func (_MyOFT *MyOFTFilterer) WatchOFTSent(opts *bind.WatchOpts, sink chan<- *MyOFTOFTSent, guid [][32]byte, fromAddress []common.Address) (event.Subscription, error) {

	var guidRule []interface{}
	for _, guidItem := range guid {
		guidRule = append(guidRule, guidItem)
	}

	var fromAddressRule []interface{}
	for _, fromAddressItem := range fromAddress {
		fromAddressRule = append(fromAddressRule, fromAddressItem)
	}

	logs, sub, err := _MyOFT.contract.WatchLogs(opts, "OFTSent", guidRule, fromAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyOFTOFTSent)
				if err := _MyOFT.contract.UnpackLog(event, "OFTSent", log); err != nil {
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

// ParseOFTSent is a log parse operation binding the contract event 0x85496b760a4b7f8d66384b9df21b381f5d1b1e79f229a47aaf4c232edc2fe59a.
//
// Solidity: event OFTSent(bytes32 indexed guid, uint32 dstEid, address indexed fromAddress, uint256 amountSentLD, uint256 amountReceivedLD)
func (_MyOFT *MyOFTFilterer) ParseOFTSent(log types.Log) (*MyOFTOFTSent, error) {
	event := new(MyOFTOFTSent)
	if err := _MyOFT.contract.UnpackLog(event, "OFTSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MyOFTOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MyOFT contract.
type MyOFTOwnershipTransferredIterator struct {
	Event *MyOFTOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MyOFTOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyOFTOwnershipTransferred)
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
		it.Event = new(MyOFTOwnershipTransferred)
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
func (it *MyOFTOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyOFTOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyOFTOwnershipTransferred represents a OwnershipTransferred event raised by the MyOFT contract.
type MyOFTOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MyOFT *MyOFTFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MyOFTOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MyOFT.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MyOFTOwnershipTransferredIterator{contract: _MyOFT.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MyOFT *MyOFTFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MyOFTOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MyOFT.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyOFTOwnershipTransferred)
				if err := _MyOFT.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_MyOFT *MyOFTFilterer) ParseOwnershipTransferred(log types.Log) (*MyOFTOwnershipTransferred, error) {
	event := new(MyOFTOwnershipTransferred)
	if err := _MyOFT.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MyOFTPeerSetIterator is returned from FilterPeerSet and is used to iterate over the raw logs and unpacked data for PeerSet events raised by the MyOFT contract.
type MyOFTPeerSetIterator struct {
	Event *MyOFTPeerSet // Event containing the contract specifics and raw log

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
func (it *MyOFTPeerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyOFTPeerSet)
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
		it.Event = new(MyOFTPeerSet)
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
func (it *MyOFTPeerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyOFTPeerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyOFTPeerSet represents a PeerSet event raised by the MyOFT contract.
type MyOFTPeerSet struct {
	Eid  uint32
	Peer [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPeerSet is a free log retrieval operation binding the contract event 0x238399d427b947898edb290f5ff0f9109849b1c3ba196a42e35f00c50a54b98b.
//
// Solidity: event PeerSet(uint32 eid, bytes32 peer)
func (_MyOFT *MyOFTFilterer) FilterPeerSet(opts *bind.FilterOpts) (*MyOFTPeerSetIterator, error) {

	logs, sub, err := _MyOFT.contract.FilterLogs(opts, "PeerSet")
	if err != nil {
		return nil, err
	}
	return &MyOFTPeerSetIterator{contract: _MyOFT.contract, event: "PeerSet", logs: logs, sub: sub}, nil
}

// WatchPeerSet is a free log subscription operation binding the contract event 0x238399d427b947898edb290f5ff0f9109849b1c3ba196a42e35f00c50a54b98b.
//
// Solidity: event PeerSet(uint32 eid, bytes32 peer)
func (_MyOFT *MyOFTFilterer) WatchPeerSet(opts *bind.WatchOpts, sink chan<- *MyOFTPeerSet) (event.Subscription, error) {

	logs, sub, err := _MyOFT.contract.WatchLogs(opts, "PeerSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyOFTPeerSet)
				if err := _MyOFT.contract.UnpackLog(event, "PeerSet", log); err != nil {
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

// ParsePeerSet is a log parse operation binding the contract event 0x238399d427b947898edb290f5ff0f9109849b1c3ba196a42e35f00c50a54b98b.
//
// Solidity: event PeerSet(uint32 eid, bytes32 peer)
func (_MyOFT *MyOFTFilterer) ParsePeerSet(log types.Log) (*MyOFTPeerSet, error) {
	event := new(MyOFTPeerSet)
	if err := _MyOFT.contract.UnpackLog(event, "PeerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MyOFTPreCrimeSetIterator is returned from FilterPreCrimeSet and is used to iterate over the raw logs and unpacked data for PreCrimeSet events raised by the MyOFT contract.
type MyOFTPreCrimeSetIterator struct {
	Event *MyOFTPreCrimeSet // Event containing the contract specifics and raw log

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
func (it *MyOFTPreCrimeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyOFTPreCrimeSet)
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
		it.Event = new(MyOFTPreCrimeSet)
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
func (it *MyOFTPreCrimeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyOFTPreCrimeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyOFTPreCrimeSet represents a PreCrimeSet event raised by the MyOFT contract.
type MyOFTPreCrimeSet struct {
	PreCrimeAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPreCrimeSet is a free log retrieval operation binding the contract event 0xd48d879cef83a1c0bdda516f27b13ddb1b3f8bbac1c9e1511bb2a659c2427760.
//
// Solidity: event PreCrimeSet(address preCrimeAddress)
func (_MyOFT *MyOFTFilterer) FilterPreCrimeSet(opts *bind.FilterOpts) (*MyOFTPreCrimeSetIterator, error) {

	logs, sub, err := _MyOFT.contract.FilterLogs(opts, "PreCrimeSet")
	if err != nil {
		return nil, err
	}
	return &MyOFTPreCrimeSetIterator{contract: _MyOFT.contract, event: "PreCrimeSet", logs: logs, sub: sub}, nil
}

// WatchPreCrimeSet is a free log subscription operation binding the contract event 0xd48d879cef83a1c0bdda516f27b13ddb1b3f8bbac1c9e1511bb2a659c2427760.
//
// Solidity: event PreCrimeSet(address preCrimeAddress)
func (_MyOFT *MyOFTFilterer) WatchPreCrimeSet(opts *bind.WatchOpts, sink chan<- *MyOFTPreCrimeSet) (event.Subscription, error) {

	logs, sub, err := _MyOFT.contract.WatchLogs(opts, "PreCrimeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyOFTPreCrimeSet)
				if err := _MyOFT.contract.UnpackLog(event, "PreCrimeSet", log); err != nil {
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

// ParsePreCrimeSet is a log parse operation binding the contract event 0xd48d879cef83a1c0bdda516f27b13ddb1b3f8bbac1c9e1511bb2a659c2427760.
//
// Solidity: event PreCrimeSet(address preCrimeAddress)
func (_MyOFT *MyOFTFilterer) ParsePreCrimeSet(log types.Log) (*MyOFTPreCrimeSet, error) {
	event := new(MyOFTPreCrimeSet)
	if err := _MyOFT.contract.UnpackLog(event, "PreCrimeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MyOFTTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the MyOFT contract.
type MyOFTTransferIterator struct {
	Event *MyOFTTransfer // Event containing the contract specifics and raw log

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
func (it *MyOFTTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyOFTTransfer)
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
		it.Event = new(MyOFTTransfer)
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
func (it *MyOFTTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyOFTTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyOFTTransfer represents a Transfer event raised by the MyOFT contract.
type MyOFTTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MyOFT *MyOFTFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MyOFTTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MyOFT.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MyOFTTransferIterator{contract: _MyOFT.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MyOFT *MyOFTFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MyOFTTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MyOFT.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyOFTTransfer)
				if err := _MyOFT.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_MyOFT *MyOFTFilterer) ParseTransfer(log types.Log) (*MyOFTTransfer, error) {
	event := new(MyOFTTransfer)
	if err := _MyOFT.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
