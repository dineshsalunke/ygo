package encoding

import (
	"github.com/dineshsalunke/ygo/pkg/core"
	"github.com/dineshsalunke/ygo/pkg/encoding/binary"
)

type DsDecoderV1 struct {
	decoder *binary.Decoder
}

func (dec *DsDecoderV1) ResetDsCurVal() {
	panic("not implemented")
}

func (dec *DsDecoderV1) ReadDsClock() (uint64, error) {
	panic("not implemented")
}

func (dec *DsDecoderV1) ReadDsLen() (uint64, error) {
	panic("not implemented")
}

type UpdateDecoderV1 struct {
	*DsDecoderV1
}

func (dec *UpdateDecoderV1) ReadLeftID() (*core.ID, error) {
	panic("not implemented")
}

func (dec *UpdateDecoderV1) ReadRightID() (*core.ID, error) {
	panic("not implemented")
}

func (dec *UpdateDecoderV1) ReadClient() (uint64, error) {
	panic("not implemented")
}

func (dec *UpdateDecoderV1) ReadInfo() (uint64, error) {
	panic("not implemented")
}

func (dec *UpdateDecoderV1) ReadString() (string, error) {
	panic("not implemented")
}

func (dec *UpdateDecoderV1) ReadParentInfo() (bool, error) {
	panic("not implemented")
}

func (dec *UpdateDecoderV1) ReadTypeRef() (uint8, error) {
	panic("not implemented")
}

func (dec *UpdateDecoderV1) ReadLength() (uint64, error) {
	panic("not implemented")
}

func (dec *UpdateDecoderV1) ReadAny() (any, error) {
	panic("not implemented")
}

func (dec *UpdateDecoderV1) ReadBuf() ([]byte, error) {
	panic("not implemented")
}

func (dec *UpdateDecoderV1) ReadJSON() (any, error) {
	panic("not implemented")
}

func (dec *UpdateDecoderV1) ReadKey() (string, error) {
	panic("not implemented")
}
