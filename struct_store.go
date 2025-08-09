package ygo

import (
	"cmp"
	"fmt"
	"slices"
)
type PendingStructs struct {
	missingState StateVector
	update       []byte
}

type StructStore struct {
	clients            map[uint64][]Block // TODO: extend this into custom type and slice of block so we could add binary search funcs
	skips              *IdSet
	pendingStructs     *PendingStructs
	pendingIdSetUpdate []byte
}

func newStructStore() *StructStore {
	return &StructStore{
		clients:            make(map[uint64][]Block),
		skips:              newIdSet(),
		pendingStructs:     nil,
		pendingIdSetUpdate: nil,
	}
}

func (ss *StructStore) IntegrateStructs(tx *Transaction) (*PendingStructs, error) {
	panic("not implemented")
func (ss *StructStore) GetStateVector() StateVector {
	cl := len(ss.clients)
	sl := len(ss.skips.clients)
	sm := make(StateVector, cl+sl)
	for client, blocks := range ss.clients {
		block := blocks[len(blocks)-1]
		sm[client] = block.ClockEnd()
	}

	for client, block := range ss.skips.clients {
		r := block.getIdRanges()
		if len(r) > 0 {
			sm[client] = r[0].clock
		}
	}
	return sm
}


func (ss *StructStore) ApplyIdSet(decoder UpdateDecoder, tx *Transaction) ([]byte, error) {
	unappliedIdSet := newIdSet()
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
		blocks, ok := ss.clients[client]
		if !ok {
			blocks = make([]Block, 0)
		}
		state := ss.GetClientClockEnd(client)

		for range numOfDeletes {
			clock, err := decoder.ReadDsClock()
			if err != nil {
				return nil, err
			}
			dsLength, err := decoder.ReadDsLength()
			if err != nil {
				return nil, err
			}
			clockEnd := clock + dsLength
			if clock < state {
				if state < clockEnd {
					if err := unappliedIdSet.add(client, state, clockEnd-state); err != nil {
						return nil, err
					}
				}
				index, found := binarySearchClockIndex(blocks, clock)
				if found {
					item, ok := blocks[index].(*Item)
					if ok && !item.Deleted() && item.ClockStart() < clock {
						rightItem, err := item.Split(tx, clock-item.ClockEnd())
						if err != nil {
							return nil, err
						}
						index += index
						ss.clients[client] = slices.Insert(blocks, index, rightItem)
					}

					for index < len(blocks) {
						block := blocks[index]
						if block.ClockStart() < clockEnd {
							if !block.Deleted() {
								if item, ok := block.(*Item); ok {
									if clockEnd < block.ClockEnd()+1 {
										rightItem, err := item.Split(tx, clockEnd-block.ClockStart())
										if err != nil {
											return nil, err
										}
										ss.clients[client] = slices.Insert(blocks, index, rightItem)
									}
								} else {
									c := max(block.ClockStart(), clock)
									unappliedIdSet.add(client, c, min(block.Length(), clockEnd-c))
								}
							}
						}

						index += 1
					}
				}
			} else {
				unappliedIdSet.add(client, clock, clockEnd-clock)
			}
		}
	}
	if len(unappliedIdSet.clients) > 0 {
		encoder := newUpdateEncoderV1()
		if err := encoder.WriteVarUint(0); err != nil {
			return nil, err
		}
		if err := writeIdSet(unappliedIdSet, encoder); err != nil {
			return nil, err
		}
		return encoder.Bytes(), nil
	}
	return nil, nil
}

// get client clock length
func (ss *StructStore) GetClientClockEnd(client uint64) uint64 {
	blocks, has := ss.clients[client]
	if !has || len(blocks) == 0 {
		return 0
	}
	lastBlock := blocks[len(blocks)-1]
	return lastBlock.ClockEnd() + 1
}

func (ss *StructStore) BinarySearchBlock(id *ID) (Block, int, bool) {
	blocks, has := ss.clients[id.client]
	if has {
		index, present := slices.BinarySearchFunc(blocks, id.clock, func(b Block, c uint64) int {
			if b.ContainsClock(c) {
				return 0
			}
			return cmp.Compare(b.ClockEnd(), c)
		})
		if present {
			return blocks[index], index, true
		}
		return nil, 0, false
	}
	return nil, 0, false
}

func findIndexCleanStart(tx *Transaction, blocks []Block, clock uint64) (uint64, error) {
	index, exists := slices.BinarySearchFunc(blocks, clock, func(b Block, c uint64) int {
		if b.ContainsClock(c) {
			return 0
		}
		return cmp.Compare(b.ClockEnd(), c)
	})
	if !exists {
		return 0, fmt.Errorf("no block found for clock %d", clock)
	}

	block := blocks[index]
	if block.ClockStart() < clock {
		nBlock, err := block.Split(tx, clock-block.ClockStart())
		if err != nil {
			return 0, err
		}
		blocks = slices.Insert(blocks, index+1, nBlock)
		return uint64(index + 1), nil
	}

	return uint64(index), nil
}

func (ss *StructStore) GetItemCleanStart(tx *Transaction, id *ID) (Block, error) {
	blocks, has := ss.clients[id.client]
	if has {
		index, err := findIndexCleanStart(tx, blocks, id.clock)
		if err != nil {
			return nil, err
		}
		return blocks[index], nil
	}
	return nil, fmt.Errorf("failed to find blocks for client : %d", id.client)
}

func (ss *StructStore) GetItemCleanEnd(tx *Transaction, id *ID) (Block, error) {
	block, index, has := ss.BinarySearchBlock(id)
	if has && id.clock != block.ClockEnd() {
		if _, isGC := block.(*GC); !isGC {
			right, err := block.Split(tx, id.clock-block.ClockStart()+1)
			if err != nil {
				return nil, err
			}
			ss.clients[id.client] = slices.Insert(ss.clients[id.client], index, right)
		}
	}
	return block, nil
}

func (ss *StructStore) GetBlock(id *ID) Block {
	block, _, _ := ss.BinarySearchBlock(id)
	return block
}

func (ss *StructStore) GetItem(id *ID) *Item {
	block := ss.GetBlock(id)
	if block != nil {
		return block.(*Item)
	}
	return nil
}
