package ygo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodingUint8(t *testing.T) {
	encoder := NewEncoder()
	err := encoder.WriteUint8(128)
	assert.Nil(t, err)
	decoder := NewDecoder(encoder.Bytes())
	val, err := decoder.ReadUint8()
	assert.Nil(t, err)
	assert.Equal(t, uint8(128), val)
}

func TestEncodingUint16(t *testing.T) {
	encoder := NewEncoder()
	err := encoder.WriteUint16(128)
	assert.Nil(t, err)
	decoder := NewDecoder(encoder.Bytes())
	val, err := decoder.ReadUint16()
	assert.Nil(t, err)
	assert.Equal(t, uint16(128), val)
}

func TestEncodingUint32(t *testing.T) {
	encoder := NewEncoder()
	err := encoder.WriteUint32(128)
	assert.Nil(t, err)
	decoder := NewDecoder(encoder.Bytes())
	val, err := decoder.ReadUint32()
	assert.Nil(t, err)
	assert.Equal(t, uint32(128), val)
}

func TestEncodingVarUint(t *testing.T) {
	encoder := NewEncoder()
	err := encoder.WriteVarUint(4096)
	assert.Nil(t, err)
	decoder := NewDecoder(encoder.Bytes())
	val, err := decoder.ReadVarUint()
	assert.Nil(t, err)
	assert.Equal(t, uint64(4096), val)
}

func TestEncodingVarint(t *testing.T) {
	encoder := NewEncoder()
	err := encoder.WriteVarint(65535)
	assert.Nil(t, err)

	err = encoder.WriteVarint(-65535)
	assert.Nil(t, err)

	decoder := NewDecoder(encoder.Bytes())

	val, err := decoder.ReadVarint()
	assert.Nil(t, err)
	assert.Equal(t, int64(65535), val)

	val, err = decoder.ReadVarint()
	assert.Nil(t, err)
	assert.Equal(t, int64(-65535), val)
}

func TestEncodingVarString(t *testing.T) {
	encoder := NewEncoder()
	str := "this is some long long long long long long and again long text with 🌍 emoji"
	err := encoder.WriteVarString(str)
	assert.Nil(t, err)
	decoder := NewDecoder(encoder.Bytes())
	val, err := decoder.ReadVarString()
	assert.Nil(t, err)
	assert.Equal(t, str, val)
}

func TestEncodingVarUint8Array(t *testing.T) {
	encoder := NewEncoder()
	err := encoder.WriteVarUint8Array([]byte{20, 40, 80})
	assert.Nil(t, err)
	decoder := NewDecoder(encoder.Bytes())
	val, err := decoder.ReadVarUint8Array()
	assert.Nil(t, err)
	assert.Equal(t, []byte{20, 40, 80}, val)
}

func TestEncodingFloats(t *testing.T) {
	encoder := NewEncoder()
	err := encoder.WriteFloat32(float32(3.14159))
	assert.Nil(t, err)
	err = encoder.WriteFloat64(float64(3.14159))
	assert.Nil(t, err)

	decoder := NewDecoder(encoder.Bytes())
	val, err := decoder.ReadFloat32()
	assert.Nil(t, err)
	assert.Equal(t, float32(3.14159), val)

	f64, err := decoder.ReadFloat64()
	assert.Nil(t, err)
	assert.Equal(t, float64(3.14159), f64)
}

func TestEncodingAny(t *testing.T) {
	encoder := NewEncoder()
	err := encoder.WriteAny(int64(20))
	assert.Nil(t, err)
	err = encoder.WriteAny([]any{2000, 3000})
	assert.Nil(t, err)
	err = encoder.WriteAny(map[string]any{"boolean": true})
	assert.Nil(t, err)

	decoder := NewDecoder(encoder.Bytes())
	val, err := decoder.ReadAny()
	assert.Nil(t, err)
	assert.Equal(t, int64(20), val)

	val, err = decoder.ReadAny()
	assert.Nil(t, err)
	assert.Equal(t, []any{int64(2000), int64(3000)}, val)

	val, err = decoder.ReadAny()
	assert.Nil(t, err)
	assert.Equal(t, map[string]any{"boolean": true}, val)
}
