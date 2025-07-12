package encoding

import (
	"github.com/dineshsalunke/ygo/pkg/core"
	"github.com/dineshsalunke/ygo/pkg/encoding/binary"
)

type DsEncoder interface {
	Bytes() []byte
	ResetDsCurVal()
	WriteDsClock(clock uint64) error
	WriteDsLength(length uint64) error
	Encoder() *binary.Encoder
}

type UpdateEncoder interface {
	DsEncoder
	WriteLeftID(id *core.ID) error
	WriteRightID(id *core.ID) error
	WriteClient(client uint64) error
	WriteInfo(info uint8) error
	WriteString(str string) error
	WriteParentInfo(isYKey bool) error
	WriteTypeRef(info uint8) error
	WriteLength(length uint64) error
	WriteAny(val any) error
	WriteBuf(buf []byte) error
	WriteJSON(embed any) error
	WriteKey(key string) error
}
