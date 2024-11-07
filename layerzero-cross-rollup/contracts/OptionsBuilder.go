package contracts

import (
	"bytes"
	"encoding/binary"
	"errors"
	"math/big"
)

// Constants for OptionsBuilder
const (
	TYPE_1 uint16 = 1 // legacy options type 1
	TYPE_2 uint16 = 2 // legacy options type 2
	TYPE_3 uint16 = 3
)

// Constants for ExecutorOptions
const (
	WORKER_ID                     uint8 = 1
	OPTION_TYPE_LZRECEIVE         uint8 = 1
	OPTION_TYPE_NATIVE_DROP       uint8 = 2
	OPTION_TYPE_LZCOMPOSE         uint8 = 3
	OPTION_TYPE_ORDERED_EXECUTION uint8 = 4
)

// NewOptions initializes a new options byte slice with TYPE_3.
func NewOptions() []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, TYPE_3)
	return buf.Bytes()
}

// AddExecutorLzReceiveOption adds an LZ Receive option to the existing options.
// It ensures that the existing options are of TYPE_3.
func AddExecutorLzReceiveOption(options []byte, gas, value *big.Int) ([]byte, error) {
	if !isType3(options) {
		return nil, errors.New("options type is not TYPE_3")
	}

	option, err := ExecutorOptionsEncodeLzReceiveOption(gas, value)
	if err != nil {
		return nil, err
	}

	return AddExecutorOption(options, OPTION_TYPE_LZRECEIVE, option)
}

// AddExecutorOption adds a generic executor option to the existing options.
// It ensures that the existing options are of TYPE_3.
func AddExecutorOption(options []byte, optionType uint8, option []byte) ([]byte, error) {
	if !isType3(options) {
		return nil, errors.New("options type is not TYPE_3")
	}

	buf := new(bytes.Buffer)
	buf.Write(options)

	// Append WORKER_ID
	buf.WriteByte(WORKER_ID)

	// Calculate option length: len(option) + 1 for optionType
	optionLength := uint16(len(option) + 1)
	if err := binary.Write(buf, binary.BigEndian, optionLength); err != nil {
		return nil, err
	}

	// Append OPTION_TYPE
	buf.WriteByte(optionType)

	// Append the actual option bytes
	buf.Write(option)

	return buf.Bytes(), nil
}

// isType3 checks if the options start with TYPE_3.
func isType3(options []byte) bool {
	if len(options) < 2 {
		return false
	}
	return binary.BigEndian.Uint16(options[:2]) == TYPE_3
}

// ExecutorOptions functions
// EncodeLzReceiveOption encodes the LZ Receive option based on _gas and _value.
// If _value is zero, it encodes only _gas. Otherwise, it encodes both _gas and _value.
func ExecutorOptionsEncodeLzReceiveOption(gas, value *big.Int) ([]byte, error) {
	if gas == nil {
		return nil, errors.New("gas cannot be nil")
	}

	gasBytes, err := encodeUint128(gas)
	if err != nil {
		return nil, err
	}

	if value == nil {
		return nil, errors.New("value cannot be nil")
	}

	if value.Cmp(big.NewInt(0)) == 0 {
		return gasBytes, nil
	}

	valueBytes, err := encodeUint128(value)
	if err != nil {
		return nil, err
	}

	// Concatenate gasBytes and valueBytes
	return append(gasBytes, valueBytes...), nil
}

// encodeUint128 encodes a big.Int into a 16-byte big-endian representation.
func encodeUint128(n *big.Int) ([]byte, error) {
	if n.Sign() < 0 {
		return nil, errors.New("cannot encode negative numbers as uint128")
	}

	bytes128 := n.Bytes()
	if len(bytes128) > 16 {
		return nil, errors.New("integer overflow: cannot encode numbers larger than 128 bits")
	}

	// Prepend zeros if necessary to make it 16 bytes
	padded := make([]byte, 16)
	copy(padded[16-len(bytes128):], bytes128)
	return padded, nil
}
