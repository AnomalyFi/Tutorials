package contracts

import (
	"bytes"
	"encoding/binary"
	"errors"
)

// Constants representing option types
const (
	TYPE_1 uint16 = 1
	TYPE_2 uint16 = 2
	TYPE_3 uint16 = 3

	OPTION_TYPE_LZRECEIVE uint8 = 1
	WORKER_ID             uint8 = 0x01
)

// Helper function to convert uint16 to bytes
func toBytesUint16(value uint16) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, value)
	return buf.Bytes()
}

// Helper function to convert uint128 to bytes
func toBytesUint128(value uint64) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, value)
	return buf.Bytes()
}

// Equivalent to Solidity's newOptions()
func NewOptions() []byte {
	return append([]byte{}, byte(TYPE_3))
}

// Equivalent to Solidity's addExecutorLzReceiveOption function
func AddExecutorLzReceiveOption(options []byte, gas uint64, value uint64) ([]byte, error) {
	if !isType3(options) {
		return nil, errors.New("only type 3 options are supported")
	}
	option := EncodeLzReceiveOption(gas, value)
	return AddExecutorOption(options, OPTION_TYPE_LZRECEIVE, option)
}

// Equivalent to Solidity's addExecutorOption function
func AddExecutorOption(options []byte, optionType uint8, option []byte) ([]byte, error) {
	if !isType3(options) {
		return nil, errors.New("only type 3 options are supported")
	}

	optionLength := uint16(len(option) + 1) // +1 for optionType
	optionLengthBytes := toBytesUint16(optionLength)

	encoded := append(options, WORKER_ID)
	encoded = append(encoded, optionLengthBytes...)
	encoded = append(encoded, optionType)
	encoded = append(encoded, option...)

	return encoded, nil
}

// Equivalent to Solidity's ExecutorOptions.encodeLzReceiveOption function
func EncodeLzReceiveOption(gas uint64, value uint64) []byte {
	if value == 0 {
		return toBytesUint128(gas)
	}
	return append(toBytesUint128(gas), toBytesUint128(value)...)
}

// Helper function to check if options are of TYPE_3
func isType3(options []byte) bool {
	return len(options) > 0 && options[0] == byte(TYPE_3)
}
