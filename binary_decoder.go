package ygo

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
)

type BinaryDecoder struct {
	*bytes.Buffer
}

func newDecoder(buf []byte) *BinaryDecoder {
	return &BinaryDecoder{
		Buffer: bytes.NewBuffer(buf),
	}
}

func NewDecoder(buf []byte) *BinaryDecoder {
	return newDecoder(buf)
}

// HasContent
func (dec *BinaryDecoder) HasContent() bool {
	return dec.Len() > 0
}

// ReadUint8
func (dec *BinaryDecoder) ReadUint8() (byte, error) {
	return dec.ReadByte()
}

// ReadUint16
func (dec *BinaryDecoder) ReadUint16() (uint16, error) {
	buf := make([]byte, 2)
	if _, err := dec.Read(buf); err != nil {
		return 0, nil
	}
	return binary.LittleEndian.Uint16(buf), nil
}

// ReadUint32
func (dec *BinaryDecoder) ReadUint32() (uint32, error) {
	buf := make([]byte, 4)
	if _, err := dec.Read(buf); err != nil {
		return 0, nil
	}
	return binary.LittleEndian.Uint32(buf), nil
}

// ReadVarUint
func (dec *BinaryDecoder) ReadVarUint() (uint64, error) {
	return binary.ReadUvarint(dec)
}

// ReadVarint
func (dec *BinaryDecoder) ReadVarint() (int64, error) {
	var mult int64 = 64
	var sign int64 = 1
	b, err := dec.ReadByte()
	if err != nil {
		return 0, err
	}

	var num int64 = int64(b & 0x3f)
	// 64 - BIT7
	if b&0x40 > 0 {
		sign = -1
	}

	// 128 - BIT8
	if b&0x80 == 0 {
		return sign * num, nil
	}

	for range binary.MaxVarintLen64 {
		b, err := dec.ReadByte()
		if err != nil {
			return 0, err
		}
		num = num + int64(b&0x7f)*mult
		mult *= 128

		if b < 0x80 {
			return sign * num, nil
		}
	}

	return 0, fmt.Errorf("unexpected end of array")
}

// ReadUint8Array
func (dec *BinaryDecoder) ReadUint8Array(length uint64) ([]byte, error) {
	buf := make([]byte, length)
	if err := binary.Read(dec, binary.LittleEndian, buf); err != nil {
		return nil, err
	}
	return buf, nil
}

// ReadVarUint8Array
func (dec *BinaryDecoder) ReadVarUint8Array() ([]byte, error) {
	length, err := dec.ReadVarUint()
	if err != nil {
		return nil, err
	}
	return dec.ReadUint8Array(length)
}

// ReadVarString
func (dec *BinaryDecoder) ReadVarString() (string, error) {
	buff, err := dec.ReadVarUint8Array()
	if err != nil {
		return "", err
	}
	return string(buff), nil
}

// ReadFloat32
func (dec *BinaryDecoder) ReadFloat32() (float32, error) {
	buff, err := dec.ReadUint8Array(4)
	if err != nil {
		return 0, err
	}
	return math.Float32frombits(binary.BigEndian.Uint32(buff)), nil
}

// ReadFloat64
func (dec *BinaryDecoder) ReadFloat64() (float64, error) {
	buff, err := dec.ReadUint8Array(8)
	if err != nil {
		return 0, err
	}
	return math.Float64frombits(binary.BigEndian.Uint64(buff)), nil
}

const (
	UndefinedTypeByte    byte = 127
	NullTypeByte         byte = 126
	IntTypeByte          byte = 125
	Float32TypeByte      byte = 124
	Float64TypeByte      byte = 123
	BigIntTypeByte       byte = 122
	BooleanFalseTypeByte byte = 121
	BooleanTrueTypeByte  byte = 120
	StringTypeByte       byte = 119
	MapStringAnyTypeByte byte = 118
	ArrayAnyTypeByte     byte = 117
	Uint8ArrayTypeByte   byte = 116
)

func (dec *BinaryDecoder) ReadAny() (any, error) {
	datatype, err := dec.ReadUint8()
	if err != nil {
		return nil, err
	}

	switch datatype {
	case UndefinedTypeByte:
		return nil, nil
	case NullTypeByte:
		return nil, nil
	case IntTypeByte:
		value, err := dec.ReadVarint()
		if err != nil {
			return nil, err
		}
		return int64(value), nil
	case Float32TypeByte:
		return dec.ReadFloat32()
	case Float64TypeByte:
		return dec.ReadFloat64()
	case BigIntTypeByte:
		return nil, fmt.Errorf("unsupported type bigint")
	case BooleanTrueTypeByte:
		return true, nil
	case BooleanFalseTypeByte:
		return false, nil
	case StringTypeByte:
		return dec.ReadVarString()
	case MapStringAnyTypeByte:
		keysLength, err := dec.ReadUint8()
		if err != nil {
			return nil, err
		}
		obj := make(map[string]any, keysLength)
		for i := range keysLength {
			key, err := dec.ReadVarString()
			if err != nil {
				return nil, fmt.Errorf("failed reading key at index %d %w", i, err)
			}
			val, err := dec.ReadAny()
			if err != nil {
				return nil, fmt.Errorf("failed reading value at %s %w", key, err)
			}
			obj[key] = val
		}
		return obj, nil
	case ArrayAnyTypeByte:
		length, err := dec.ReadUint8()
		if err != nil {
			return nil, err
		}
		arr := make([]any, length)
		for i := range length {
			val, err := dec.ReadAny()
			if err != nil {
				return nil, fmt.Errorf("failed reading value at index %d %w", i, err)
			}
			arr[i] = val
		}
		return arr, nil
	}

	return nil, fmt.Errorf("unexpected type byte %d", datatype)
}
