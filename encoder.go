package ygo

import (
	"bytes"
	"encoding/binary"
)

type Encoder struct {
	*bytes.Buffer
}

func newEncoder(buf []byte) *Encoder {
	return &Encoder{
		Buffer: bytes.NewBuffer(buf),
	}
}

func NewEncoder(buf []byte) *Encoder {
	return newEncoder(buf)
}

func (encoder *Encoder) WriteVarUint(value uint64) error {
	buf := binary.AppendUvarint([]byte{}, value)
	return binary.Write(encoder, binary.LittleEndian, buf)
}

func (encoder *Encoder) WriteVarint(i int64) error {
	buf := binary.AppendVarint([]byte{}, i)
	return binary.Write(encoder, binary.LittleEndian, buf)
}

func (encoder *Encoder) WriteVarString(str string) error {
	buff := []byte(str)
	l := len(buff)
	buf := binary.AppendUvarint([]byte{}, uint64(l))
	if err := binary.Write(encoder, binary.LittleEndian, buf); err != nil {
		return err
	}
	return binary.Write(encoder, binary.LittleEndian, buff)
}

func (encoder *Encoder) WriteVarUintBytes(buff []byte) error {
	l := len(buff)
	buf := binary.AppendUvarint([]byte{}, uint64(l))
	if err := binary.Write(encoder, binary.LittleEndian, buf); err != nil {
		return err
	}
	return binary.Write(encoder, binary.LittleEndian, buff)
}
