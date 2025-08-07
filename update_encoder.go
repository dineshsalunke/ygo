package ygo

type IdSetEncoder interface {
	ResetDsCurrVal()
	WriteDsClock(length uint64) error
	WriteDsLength(length uint64) error

	//
	WriteUint8(val byte) error
	WriteUint16(val uint16) error
	WriteUint32(val uint32) error
	WriteVarUint(val uint64) error
	WriteVarint(val int64) error
	WriteUint8Array(val []byte) error
	WriteVarUint8Array(val []byte) error
	WriteVarString(val string) error
	WriteFloat32(val float32) error
	WriteFloat64(val float64) error
	Bytes() []byte
}

type UpdateEncoder interface {
	IdSetEncoder
	WriteLeftID(id *ID) error
	WriteRightID(id *ID) error
	WriteClient(clientId uint64) error
	WriteInfo(info byte) error
	WriteString(value string) error
	WriteParentInfo(hasParentInfo bool) error
	WriteTypeRef(ref byte) error
	WriteLength(length uint64) error
	WriteAny(val any) error
	WriteJson(val any) error
	WriteKey(key string) error
}
