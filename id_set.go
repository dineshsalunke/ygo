package ygo

type IdSet struct {
	clients map[ClientID]*IdRanges
}

func newIdSet() *IdSet {
	return &IdSet{}
}

func decode_id_set(decoder UpdateDecoder) (IdSet, error) {
	panic("not implemented")
}
