package ygo

type StructStore struct {
	clients map[uint64][]Block
	skips   *IdSet
}

func newStructStore() *StructStore {
	return &StructStore{
		clients: make(map[uint64][]Block),
		skips:   newIdSet(),
	}
}
