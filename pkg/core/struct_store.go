package core

type StructStore struct {
	clients map[uint64][]IntegrateWriter
}

func (ss *StructStore) GetStateVector() map[uint64]uint64 {
	sv := make(map[uint64]uint64, len(ss.clients))

	for client, structs := range ss.clients {
		s := structs[len(structs)-1]
		sv[client] = s.Clock() + s.Length()
	}

	return sv
}

func (ss *StructStore) GetState(client uint64) uint64 {
	structs, ok := ss.clients[client]
	if !ok {
		return 0
	}
	lastStruct := structs[len(structs)-1]
	return lastStruct.Clock() + lastStruct.Length()
}
