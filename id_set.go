package ygo

type IdSet struct {
	clients map[uint64]*IdRanges
}

func newIdSet() *IdSet {
	return &IdSet{
		clients: make(map[uint64]*IdRanges),
	}
}

