package encoding

import (
	"bytes"
	"encoding/binary"
)

type Decoder struct {
	*bytes.Buffer
}

func newDecoder(buf []byte) *Decoder {
	return &Decoder{
		Buffer: bytes.NewBuffer(buf),
	}
}

func NewDecoder(buf []byte) *Decoder {
	return newDecoder(buf)
}

func (decoder *Decoder) ReadVarUint() (uint64, error) {
	return binary.ReadUvarint(decoder)
}

func (decoder *Decoder) ReadVarint() (int64, error) {
	return binary.ReadVarint(decoder)
}

func (decoder *Decoder) ReadVarString() (string, error) {
	l, err := binary.ReadUvarint(decoder)
	if err != nil {
		return "", err
	}
	buf := make([]byte, l)
	err = binary.Read(decoder, binary.LittleEndian, buf)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}
