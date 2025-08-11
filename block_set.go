package ygo

import (
	"slices"
)

type BlockSet struct {
	clients map[uint64]*BlockRange
}

func newStructSet(length uint64) *BlockSet {
	return &BlockSet{
		clients: make(map[uint64]*BlockRange, length),
	}
}

func (ss *BlockSet) addRange(client uint64, refs []Block) {
	ss.clients[client] = &BlockRange{
		refs: refs,
	}
}

func (ss *BlockSet) excludeIdSet(set *IdSet) error {
	for excludeClientId, excludeRanges := range set.clients {
		structRange, ok := ss.clients[excludeClientId]
		if ok {
			// blocks := structRange.refs
			firstBlock := structRange.refs[0]
			lastBlock := structRange.refs[len(structRange.refs)-1]
			for _, excludeRange := range excludeRanges.getIdRanges() {
				var err error
				startIndex := uint64(0)
				endIndex := uint64(0)

				// No need to exclude range if its clock is greater than highest block in store
				if excludeRange.clock >= lastBlock.ClockEnd() {
					continue
				}
				// find first id range whose clock is greater than excludeRange clock
				if excludeRange.clock > firstBlock.ClockStart() {
					startIndex, err = findIndexCleanStart(nil, structRange.refs, excludeRange.clock)
					if err != nil {
						return err
					}
				}

				endIndex = uint64(len(structRange.refs))
				if excludeRange.clock+excludeRange.length <= firstBlock.ID().clock {
					continue
				}
				if excludeRange.end() < lastBlock.ClockEnd() {
					endIndex, err = findIndexCleanStart(nil, structRange.refs, excludeRange.end())
					if err != nil {
						return err
					}
				}

				if startIndex < endIndex {
					structRange.refs[startIndex] = newSkip(newID(excludeClientId, excludeRange.clock), excludeRange.length)
					diff := endIndex - startIndex
					if diff > 1 {
						i := int(startIndex + 1)
						j := int(i + int(diff) - 1)
						structRange.refs = slices.Delete(structRange.refs, i, j)
					}
				}
			}
		}
	}
	return nil
}
