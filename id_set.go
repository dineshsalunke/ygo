package ygo

type IdRange struct{}

type IdRanges struct{}

type IdSet struct {
	clients map[uint64]*IdRanges
}

func newIdSet() *IdSet {
	return &IdSet{}
}
