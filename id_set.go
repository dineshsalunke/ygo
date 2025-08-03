package ygo

import (
	"fmt"
	"slices"
	"strings"
)
type IdSet struct {
	clients map[uint64]*IdRanges
}

func newIdSet() *IdSet {
	return &IdSet{
		clients: make(map[uint64]*IdRanges),
	}
}

func (ss *IdSet) GoString() string {
	builder := strings.Builder{}
	builder.WriteString("IdSet {\r\tclients: {")
	for client, ranges := range ss.clients {
		builder.WriteString(fmt.Sprintf("%d:%#v\n", client, ranges))
	}
	builder.WriteString("}\r}")
	return builder.String()
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

