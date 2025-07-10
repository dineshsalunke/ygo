package encoding

import (
	"github.com/dineshsalunke/ygo/pkg/core"
	"github.com/dineshsalunke/ygo/pkg/encoding/binary"
)

type DsEncoderV2 struct {
	encoder *binary.Encoder
}

func (encoder *DsEncoderV2) Bytes() []byte {
	panic("not implemented")
}

func (encoder *DsEncoderV2) ResetDsCurVal() {
	// Noop
}

func (encoder *DsEncoderV2) WriteDsClock(clock uint64) error {
	panic("not implemented")
}

func (encoder *DsEncoderV2) WriteDsLength(length uint64) error {
	panic("not implemented")
}

type UpdateEncoderV2 struct {
	*DsEncoderV2
}

func (enc *UpdateEncoderV2) WriteLeftID(id *core.ID) {
	panic("not implemented")
}

func (enc *UpdateEncoderV2) WriteRightID(id *core.ID) {
	panic("not implemented")
}

func (enc *UpdateEncoderV2) WriteClient(client uint64) {
	panic("not implemented")
}

func (enc *UpdateEncoderV2) WriteInfo(info uint8) {
	panic("not implemented")
}

func (enc *UpdateEncoderV2) WriteString(str string) {
	panic("not implemented")
}

func (enc *UpdateEncoderV2) WriteParentInfo(isYKey bool) {
	panic("not implemented")
}

func (enc *UpdateEncoderV2) WriteTypeRef(info uint8) {
	panic("not implemented")
}

func (enc *UpdateEncoderV2) WriteLength(length uint64) {
	panic("not implemented")
}

func (enc *UpdateEncoderV2) WriteAny(val any) {
	panic("not implemented")
}

func (enc *UpdateEncoderV2) WriteBuf(buf []byte) {
	panic("not implemented")
}

func (enc *UpdateEncoderV2) WriteJSON(embed any) {
	panic("not implemented")
}

func (enc *UpdateEncoderV2) WriteKey(key string) {
	panic("not implemented")
}
