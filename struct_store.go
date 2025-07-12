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

func (ss *StructStore) GetStateVector() ClientClockMap {
	sv := make(ClientClockMap, len(ss.clients))
	for client, structs := range ss.clients {
		s := structs[len(structs)-1]
		sv[client] = s.ID().clock + s.Length()
	}
	return sv
}

func (ss *StructStore) GetState(client uint64) uint64 {
	structs, ok := ss.clients[client]
	if !ok {
		return 0
	}
	lastStruct := structs[len(structs)-1]
	if !ok {
		return 0
	}

	return lastStruct.ID().clock + lastStruct.Length()
}
