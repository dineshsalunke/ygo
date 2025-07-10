package encoding

import (
	"github.com/dineshsalunke/ygo/pkg/core"
	"github.com/dineshsalunke/ygo/pkg/encoding/binary"
)

type DsDecoderV2 struct {
	decoder *binary.Decoder
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
	*DsDecoderV1
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
