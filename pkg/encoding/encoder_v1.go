package encoding

import (
	"github.com/dineshsalunke/ygo/pkg/core"
	"github.com/dineshsalunke/ygo/pkg/encoding/binary"
)

type DsEncoderV1 struct {
	encoder *binary.Encoder
}

func (encoder *DsEncoderV1) Bytes() []byte {
	panic("not implemented")
}

func (encoder *DsEncoderV1) ResetDsCurVal() {
	// Noop
}

func (encoder *DsEncoderV1) WriteDsClock(clock uint64) error {
	panic("not implemented")
}

func (encoder *DsEncoderV1) WriteDsLength(length uint64) error {
	panic("not implemented")
}

type UpdateEncoderV1 struct {
	*DsEncoderV1
}

func (enc *UpdateEncoderV1) WriteLeftID(id *core.ID) {
	panic("not implemented")
}

func (enc *UpdateEncoderV1) WriteRightID(id *core.ID) {
	panic("not implemented")
}

func (enc *UpdateEncoderV1) WriteClient(client uint64) {
	panic("not implemented")
}

func (enc *UpdateEncoderV1) WriteInfo(info uint8) {
	panic("not implemented")
}

func (enc *UpdateEncoderV1) WriteString(str string) {
	panic("not implemented")
}

func (enc *UpdateEncoderV1) WriteParentInfo(isYKey bool) {
	panic("not implemented")
}

func (enc *UpdateEncoderV1) WriteTypeRef(info uint8) {
	panic("not implemented")
}

func (enc *UpdateEncoderV1) WriteLength(length uint64) {
	panic("not implemented")
}

func (enc *UpdateEncoderV1) WriteAny(val any) {
	panic("not implemented")
}

func (enc *UpdateEncoderV1) WriteBuf(buf []byte) {
	panic("not implemented")
}

func (enc *UpdateEncoderV1) WriteJSON(embed any) {
	panic("not implemented")
}

func (enc *UpdateEncoderV1) WriteKey(key string) {
	panic("not implemented")
}
