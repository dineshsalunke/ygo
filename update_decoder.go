package ygo

type UpdateDecoder interface {
	ResetDsCurrVal()
	ReadDsClock() (uint32, error)
	ReadDsLength() (uint32, error)
	ReadLeftID() (*ID, error)
	ReadRightID() (*ID, error)
	ReadClient() (ClientID, error)
	ReadInfo() (byte, error)
	ReadParentInfo() (bool, error)
	ReadTypeRef() (byte, error)
	ReadLength() (uint32, error)
	ReadAny() (any, error)
	ReadJson() (any, error)
	ReadKey() (string, error)
}
