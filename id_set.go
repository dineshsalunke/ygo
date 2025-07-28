package ygo

type IdSet struct {
	clients map[ClientID]*IdRanges
}

func newIdSet() *IdSet {
	return &IdSet{}
}

func (set *IdSet) addRanges(client ClientID, clock, length uint64) {
	if client > 0 {
		id_ranges, ok := set.clients[client]
		if ok {
			id_ranges.addIdRange(clock, length)
		} else {
			ids := make([]*IdRange, 0)
			ids = append(ids, newIdRange(clock, length))
			set.clients[client] = newIdRanges(ids)
		}
	}
}
