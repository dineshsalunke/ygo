package encoding

import "github.com/dineshsalunke/ygo/pkg/core"

type DsDecoder interface {
	ResetDsCurVal()
	ReadDsClock() (uint64, error)
	ReadDsLen() (uint64, error)
}

type UpdateDecoder interface {
	DsDecoder
	ReadLeftID() (*core.ID, error)
	ReadRightID() (*core.ID, error)
	ReadClient() (uint64, error)
	ReadInfo() (uint64, error)
	ReadString() (string, error)
	ReadParentInfo() (bool, error)
	ReadTypeRef() (uint8, error)
	ReadLength() (uint64, error)
	ReadAny() (any, error)
	ReadBuf() ([]byte, error)
	ReadJSON() (any, error)
	ReadKey() (string, error)
}
