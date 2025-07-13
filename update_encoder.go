package ygo

type DsEncoder interface {
	Encoder() *Encoder
	Bytes() []byte
	ResetDsCurVal() error
	WriteDsClock(clock uint64) error
	WriteDsLength(l uint64) error
}

type UpdateEncoder interface {
	DsEncoder
	WriteLeftID() (*ID, error)
	WriteRightID() (*ID, error)
	WriteClient(client uint64) error
	WriteInfo(info uint64) error
	WriteString(str string) error
	WriteParentInfo(info bool) error
	WriteTypeRef(ref uint64) error
	WriteLength(l uint64) error
	WriteAny(a any) error
	WriteBuf(buf []byte) error
	WriteJSON(buf any) error
	WriteKey(key string) error
}
