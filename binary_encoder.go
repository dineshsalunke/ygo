package ygo

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
)

type BinaryEncoder struct {
	*bytes.Buffer
}

func newEncoder() *BinaryEncoder {
	return &BinaryEncoder{
		Buffer: bytes.NewBuffer([]byte{}),
	}
}

func NewEncoder() *BinaryEncoder {
	return newEncoder()
}

// WriteUint8
func (enc *BinaryEncoder) WriteUint8(b byte) error {
	return enc.WriteByte(b)
}

// WriteUint16
func (enc *BinaryEncoder) WriteUint16(b uint16) error {
	buff := binary.LittleEndian.AppendUint16([]byte{}, b)
	_, err := enc.Write(buff)
	return err
}

// WriteUint32
func (enc *BinaryEncoder) WriteUint32(b uint32) error {
	buff := binary.LittleEndian.AppendUint32([]byte{}, b)
	return binary.Write(enc, binary.LittleEndian, buff)
}

// WriteVarUint
func (enc *BinaryEncoder) WriteVarUint(b uint64) error {
	buff := binary.AppendUvarint([]byte{}, b)
	_, err := enc.Write(buff)
	return err
}

// WriteVarint
func (enc *BinaryEncoder) WriteVarint(num int64) error {
	isNegative := num < 0
	if isNegative {
		num = -num
	}
	continuebit := byte(0)
	if num > 63 {
		continuebit = 128
	}
	isNegativebit := byte(0)
	if isNegative {
		isNegativebit = 64
	}

	if err := enc.WriteByte(continuebit | isNegativebit | byte(63&num)); err != nil {
		return err
	}

	num = num >> 6
	for num > 0 {
		b := byte(0)
		if num > 127 {
			b = 128
		}
		if err := enc.WriteByte(b | byte(127&num)); err != nil {
			return err
		}
		num = num >> 7
	}
	return nil
}

// WriteUint8Array
func (enc *BinaryEncoder) WriteUint8Array(b []byte) error {
	_, err := enc.Write(b)
	return err
}

// WriteVarUint8Array
func (enc *BinaryEncoder) WriteVarUint8Array(b []byte) error {
	if err := enc.WriteVarUint(uint64(len(b))); err != nil {
		return err
	}
	return enc.WriteUint8Array(b)
}

// WriteVarString
func (enc *BinaryEncoder) WriteVarString(b string) error {
	buf := []byte(b)
	return enc.WriteVarUint8Array(buf)
}

// WriteFloat32
func (enc *BinaryEncoder) WriteFloat32(b float32) error {
	u32 := math.Float32bits(b)
	_, err := enc.Write(binary.BigEndian.AppendUint32([]byte{}, u32))
	return err
}

// WriteFloat64
func (enc *BinaryEncoder) WriteFloat64(b float64) error {
	u64 := math.Float64bits(b)
	_, err := enc.Write(binary.BigEndian.AppendUint64([]byte{}, u64))
	return err
}

// WriteAny
func (enc *BinaryEncoder) WriteAny(b any) error {
	switch b := b.(type) {
	case int, *int:
		if err := enc.WriteByte(IntTypeByte); err != nil {
			return err
		}
		return enc.WriteVarint(int64(b.(int)))
	case int8, *int8:
		if err := enc.WriteByte(IntTypeByte); err != nil {
			return err
		}
		return enc.WriteVarint(int64(b.(int8)))
	case int16, *int16:
		if err := enc.WriteByte(IntTypeByte); err != nil {
			return err
		}
		return enc.WriteVarint(int64(b.(int16)))
	case int32, *int32:
		if err := enc.WriteByte(IntTypeByte); err != nil {
			return err
		}
		return enc.WriteVarint(int64(b.(int32)))
	case int64, *int64:
		if err := enc.WriteByte(IntTypeByte); err != nil {
			return err
		}
		return enc.WriteVarint(int64(b.(int64)))
	case uint, *uint:
		if err := enc.WriteByte(IntTypeByte); err != nil {
			return err
		}
		return enc.WriteVarUint(uint64(b.(uint)))
	case uint8, *uint8:
		if err := enc.WriteByte(IntTypeByte); err != nil {
			return err
		}
		return enc.WriteVarUint(uint64(b.(uint8)))
	case uint16, *uint16:
		if err := enc.WriteByte(IntTypeByte); err != nil {
			return err
		}
		return enc.WriteVarUint(uint64(b.(uint16)))
	case uint32, *uint32:
		if err := enc.WriteByte(IntTypeByte); err != nil {
			return err
		}
		return enc.WriteVarUint(uint64(b.(uint32)))
	case uint64, *uint64:
		if err := enc.WriteByte(IntTypeByte); err != nil {
			return err
		}
		return enc.WriteVarUint(uint64(b.(uint64)))
	case float32, *float32:
		if err := enc.WriteByte(IntTypeByte); err != nil {
			return err
		}
		return enc.WriteFloat32(float32(b.(float32)))
	case float64, *float64:
		if err := enc.WriteByte(IntTypeByte); err != nil {
			return err
		}
		return enc.WriteFloat64(float64(b.(float64)))
	case string:
		if err := enc.WriteByte(StringTypeByte); err != nil {
			return err
		}
		return enc.WriteVarString(b)
	case bool:
		if b == true {
			return enc.WriteByte(120)
		} else {
			return enc.WriteByte(121)
		}
	case []byte:
		if err := enc.WriteByte(Uint8ArrayTypeByte); err != nil {
			return err
		}
		return enc.WriteVarUint8Array(b)
	case []any:
		if err := enc.WriteByte(ArrayAnyTypeByte); err != nil {
			return err
		}
		err := enc.WriteVarUint(uint64(len(b)))
		if err != nil {
			return err
		}
		for index, value := range b {
			err := enc.WriteAny(value)
			if err != nil {
				return fmt.Errorf("failed to write value %v at index %d", value, index)
			}
		}
		return nil
	case map[string]any:
		if err := enc.WriteByte(MapStringAnyTypeByte); err != nil {
			return err
		}
		err := enc.WriteVarUint(uint64(len(b)))
		if err != nil {
			return err
		}
		for key, value := range b {
			err := enc.WriteVarString(key)
			if err != nil {
				return fmt.Errorf("failed to write key %s", key)
			}
			err = enc.WriteAny(value)
			if err != nil {
				return fmt.Errorf("failed to write value at key %s", key)
			}
		}
		return nil
	}

	return fmt.Errorf("unexpected type %T", b)
}
