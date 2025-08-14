package ygo

import (
	"cmp"
	"slices"
)

type IdSet struct {
	clients map[uint64]*IdRanges
}

func newIdSet() *IdSet {
	return &IdSet{
		clients: make(map[uint64]*IdRanges),
	}
}

func (set *IdSet) GetClientIds(client uint64) []*IdRange {
	idRanges, found := set.clients[client]
	if found {
		return idRanges.getIdRanges()
	}
	return nil
}

func (set *IdSet) add(client, clock, length uint64) error {
	if length > 0 {
		idRanges, has := set.clients[client]
		if !has {
			set.clients[client] = newIdRanges([]*IdRange{newIdRange(clock, length)})
		} else {
			idRanges.addIdRange(clock, length)
		}
	}
	return nil
}

func (set *IdSet) delete(client, clock, length uint64) error {
	idRanges, has := set.clients[client]
	if has && length > 0 {
		ids := idRanges.getIdRanges()
		index, has := slices.BinarySearchFunc(ids, clock, func(idRange *IdRange, c uint64) int {
			return int(idRange.clock - c)
		})
		if has {
			r := ids[index]
			for index < len(ids) && r.clock < clock+length {
				if r.clock < clock {
					ids[index] = newIdRange(r.clock, clock-r.clock)
					if clock+length < r.clock+r.length {
						ids = slices.Insert(ids, index+1, newIdRange(clock+length, r.clock+r.length-clock-length))
					}
				} else if clock+length < r.clock+r.length {
					ids[index] = newIdRange(clock+length, r.clock+r.length-clock-length)
				} else if len(ids) == 1 {
					delete(set.clients, client)
				} else {
					index = index - 1
					ids = slices.Delete(ids, index, 1)
				}
				index = index + 1
				r = ids[index]
			}
		}
	}
	return nil
}

func findIndexInIdRanges(idRange []*IdRange, clock uint64) (int, bool) {
	return slices.BinarySearchFunc(idRange, clock, func(b *IdRange, c uint64) int {
		if b.clock <= c && c < b.end() {
			return 0
		}
		return cmp.Compare(b.end(), c)
	})
}

func (idSet *IdSet) Has(client, clock uint64) bool {
	idrange, ok := idSet.clients[client]
	if ok {
		_, has := findIndexInIdRanges(idrange.getIdRanges(), clock)
		return has
	}
	return false
}

func (idSet *IdSet) HasID(id *ID) bool {
	return idSet.Has(id.client, id.clock)
}

func writeIdSet(idSet *IdSet, encoder IdSetEncoder) error {
	clientIds := make([]uint64, len(idSet.clients))
	slices.Sort(clientIds)
	slices.Reverse(clientIds)

	for _, client := range clientIds {
		clientIdRanges := idSet.clients[client]
		idRanges := clientIdRanges.getIdRanges()
		encoder.ResetDsCurrVal()
		if err := encoder.WriteVarUint(client); err != nil {
			return err
		}
		if err := encoder.WriteVarUint(uint64(len(idRanges))); err != nil {
			return err
		}
		for _, idRange := range idRanges {
			if err := encoder.WriteDsClock(idRange.clock); err != nil {
				return err
			}
			if err := encoder.WriteDsLength(idRange.length); err != nil {
				return err
			}
		}
	}
	return nil
}

func readIdSet(decoder IdSetDecoder) (*IdSet, error) {
	idSet := newIdSet()
	numClients, err := decoder.ReadVarUint()
	if err != nil {
		return nil, err
	}
	for range numClients {
		decoder.ResetDsCurrVal()
		client, err := decoder.ReadVarUint()
		if err != nil {
			return nil, err
		}
		numOfDeletes, err := decoder.ReadVarUint()
		if err != nil {
			return nil, err
		}
		if numOfDeletes > 0 {
			dsRanges := make([]*IdRange, numOfDeletes)
			for range numOfDeletes {
				clock, err := decoder.ReadDsClock()
				if err != nil {
					return nil, err
				}
				length, err := decoder.ReadDsLength()
				if err != nil {
					return nil, err
				}
				dsRanges = append(dsRanges, newIdRange(clock, length))
			}
			idSet.clients[client] = newIdRanges(dsRanges)
		}
	}

	return idSet, nil
}
