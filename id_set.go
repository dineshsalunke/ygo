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
// TODO: we could probably use the BinarySearchFunc from the slices package
// Find Range start in the ranges slice using binary search
func findRangeStartInIdRanges(ids []*IdRange, clock uint64) (int, bool) {
	left := 0
	right := len(ids) - 1
	for left <= right {
		midindex := left + right/2
		mid := ids[midindex]
		midclock := mid.clock
		if midclock <= clock {
			if clock < mid.end() {
				return midindex, true
			}
			left = midindex + 1
		} else {
			right = midindex - 1
		}
	}
	if left < len(ids) {
		return left, true
	}
	return 0, false
}
