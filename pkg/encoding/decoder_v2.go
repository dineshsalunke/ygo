package encoding

import (
	"github.com/dineshsalunke/ygo/pkg/core"
	"github.com/dineshsalunke/ygo/pkg/encoding/binary"
)

type DsDecoderV2 struct {
	decoder *binary.Decoder
}

func NewDsDecoderV2(buf []byte) *DsDecoderV2 {
	return &DsDecoderV2{
		decoder: binary.NewDecoder(buf),
	}
}

func (dec *DsDecoderV2) ResetDsCurVal() {
	panic("not implemented")
}

func (dec *DsDecoderV2) ReadDsClock() (uint64, error) {
	panic("not implemented")
}

func (dec *DsDecoderV2) ReadDsLen() (uint64, error) {
	panic("not implemented")
}

type UpdateDecoderV2 struct {
	*DsDecoderV2
}

func NewUpdateDecoderV2(buf []byte) *UpdateDecoderV2 {
	return &UpdateDecoderV2{
		DsDecoderV2: NewDsDecoderV2(buf),
	}
}

func (dec *UpdateDecoderV2) ReadLeftID() (*core.ID, error) {
	panic("not implemented")
}

func (dec *UpdateDecoderV2) ReadRightID() (*core.ID, error) {
	panic("not implemented")
}

func (dec *UpdateDecoderV2) ReadClient() (uint64, error) {
	panic("not implemented")
}

func (dec *UpdateDecoderV2) ReadInfo() (uint64, error) {
	panic("not implemented")
}

func (dec *UpdateDecoderV2) ReadString() (string, error) {
	panic("not implemented")
}

func (dec *UpdateDecoderV2) ReadParentInfo() (bool, error) {
	panic("not implemented")
}

func (dec *UpdateDecoderV2) ReadTypeRef() (uint8, error) {
	panic("not implemented")
}

func (dec *UpdateDecoderV2) ReadLength() (uint64, error) {
	panic("not implemented")
}

func (dec *UpdateDecoderV2) ReadAny() (any, error) {
	panic("not implemented")
}

func (dec *UpdateDecoderV2) ReadBuf() ([]byte, error) {
	panic("not implemented")
}

func (dec *UpdateDecoderV2) ReadJSON() (any, error) {
	panic("not implemented")
}

func (dec *UpdateDecoderV2) ReadKey() (string, error) {
	panic("not implemented")
}
