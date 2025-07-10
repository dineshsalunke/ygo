package encoding

import "github.com/dineshsalunke/ygo/pkg/core"

type DsEncoder interface {
	Bytes() []byte
	ResetDsCurVal()
	WriteDsClock(clock uint64) error
	WriteDsLength(length uint64) error
}

type UpdateEncoder interface {
	DsEncoder
	WriteLeftID(id *core.ID)
	WriteRightID(id *core.ID)
	WriteClient(client uint64)
	WriteInfo(info uint8)
	WriteString(str string)
	WriteParentInfo(isYKey bool)
	WriteTypeRef(info uint8)
	WriteLength(length uint64)
	WriteAny(val any)
	WriteBuf(buf []byte)
	WriteJSON(embed any)
	WriteKey(key string)
}
