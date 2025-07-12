package ygo

import (
	"fmt"
	"testing"
)

func TestEncodingDecodingVarUint(t *testing.T) {
	encoder := NewEncoder([]byte{})
	encoder.WriteVarUint(2048)
	encoder.WriteVarUint(4096)
	encoder.WriteVarint(2048)
	encoder.WriteVarint(-2048)
	encoder.WriteVarString("hello world")
	encoder.WriteVarUintBytes([]byte{0x80, 0x22})

	decoder := NewDecoder(encoder.Bytes())

	a, err := decoder.ReadVarUint()
	if err != nil {
		t.Errorf("failed to read var uint %v", err)
	}

	if a != 2048 {
		t.Errorf("should read 2048 but read %v", a)
	}

	a, err = decoder.ReadVarUint()
	if err != nil {
		t.Errorf("failed to read var uint %v", err)
	}
	if a != 4096 {
		t.Errorf("should read 4096 but read %v", a)
	}

	b, err := decoder.ReadVarint()
	if err != nil {
		t.Errorf("failed to read var uint %v", err)
	}
	if b != 2048 {
		t.Errorf("should read 2048 but read %v", a)
	}

	b, err = decoder.ReadVarint()
	if err != nil {
		t.Errorf("failed to read var uint %v", err)
	}
	if b != -2048 {
		t.Errorf("should read -2048 but read %v", a)
	}

	str, err := decoder.ReadVarString()
	if err != nil {
		t.Errorf("failed to read var string %v", err)
	}
	if str != "hello world" {
		t.Errorf("should have read 'hello world' but read '%v'", a)
	}

	bytes, err := decoder.ReadVarUintBytes()
	if err != nil {
		t.Errorf("failed to read var string %v", err)
	}
	fmt.Printf("these are the bytes read %+v \n", bytes)
	if bytes[0] != 0x80 && bytes[1] != 0x22 {
		t.Errorf("should have read 'hello world' but read '%v'", a)
	}
}
