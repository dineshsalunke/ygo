package ygo

type IdSetDecoder interface {
	ResetDsCurrVal()
	ReadDsClock() (uint64, error)
	ReadDsLength() (uint64, error)

	//
	ReadUint8() (byte, error)
	ReadUint16() (uint16, error)
	ReadUint32() (uint32, error)
	ReadVarUint() (uint64, error)
	ReadVarint() (int64, error)
	ReadUint8Array(length uint64) ([]byte, error)
	ReadVarUint8Array() ([]byte, error)
	ReadVarString() (string, error)
	ReadFloat32() (float32, error)
	ReadFloat64() (float64, error)
}

type UpdateDecoder interface {
	IdSetDecoder
	//
	ReadLeftID() (*ID, error)
	ReadRightID() (*ID, error)
	ReadClient() (uint64, error)
	ReadInfo() (byte, error)
	ReadParentInfo() (bool, error)
	ReadTypeRef() (byte, error)
	ReadLength() (uint64, error)
	ReadAny() (any, error)
	ReadJson() (any, error)
	ReadKey() (string, error)
}
