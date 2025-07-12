package ygo

type ClientClockMap map[uint64]uint64

type pendingStructs struct {
	missing ClientClockMap
	update  []byte
}

type StructStore struct {
	clients        map[uint64][]SharedStruct
	pendingStructs *pendingStructs
	pendingDs      []byte
}

func newStructStore() *StructStore {
	return &StructStore{
		clients: make(map[uint64][]SharedStruct),
	}
}

