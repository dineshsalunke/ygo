package ygo

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Name      string `json:"name"`
	Type      string `json:"type"`
	Value     any    `json:"value"`
	ValueType string `json:"valuetype"`
	Encoded   []byte `json:"encoded"`
}

type TestData struct {
	Suites []*TestCase `json:"test_cases"`
}

func TestDecoding(t *testing.T) {
	buff, err := os.ReadFile("./testdata/encoding_decoding_test_data.json")
	if err != nil {
		t.Error(err)
	}

	testdata := &TestData{}
	if err := json.Unmarshal(buff, testdata); err != nil {
		t.Error(err)
	}

	for _, suite := range testdata.Suites {
		if suite.Type == "uint" {
			t.Run(suite.Name, func(t *testing.T) {
				decoder := NewDecoder(suite.Encoded)
				val, err := decoder.ReadVarUint()
				assert.Nil(t, err)
				assert.Equal(t, uint64(suite.Value.(float64)), val)
			})
		}
		if suite.Type == "int" {
			t.Run(suite.Name, func(t *testing.T) {
				decoder := NewDecoder(suite.Encoded)
				val, err := decoder.ReadVarint()
				assert.Nil(t, err)
				assert.Equal(t, int64(suite.Value.(float64)), val)
			})
		}
		if suite.Type == "float32" {
			t.Run(suite.Name, func(t *testing.T) {
				decoder := NewDecoder(suite.Encoded)
				val, err := decoder.ReadFloat32()
				assert.Nil(t, err)
				assert.Equal(t, float32(suite.Value.(float64)), val)
			})
		}
		if suite.Type == "float64" {
			t.Run(suite.Name, func(t *testing.T) {
				decoder := NewDecoder(suite.Encoded)
				val, err := decoder.ReadFloat64()
				assert.Nil(t, err)
				assert.Equal(t, float64(suite.Value.(float64)), val)
			})
		}
		if suite.Type == "string" {
			t.Run(suite.Name, func(t *testing.T) {
				decoder := NewDecoder(suite.Encoded)
				val, err := decoder.ReadVarString()
				assert.Nil(t, err)
				assert.Equal(t, string(suite.Value.(string)), val)
			})
		}
		if suite.Type == "any" {
			t.Run(suite.Name, func(t *testing.T) {
				decoder := NewDecoder(suite.Encoded)
				val, err := decoder.ReadAny()
				assert.Nil(t, err)
				switch suite.ValueType {
				case "integer":
					assert.Equal(t, int64(suite.Value.(float64)), val)
				case "float":
					assert.Equal(t, float64(suite.Value.(float64)), val)
				default:
					assert.Equal(t, suite.Value, val)
				}
				t.Logf("%T", suite.Value)
			})
		}
	}
}
