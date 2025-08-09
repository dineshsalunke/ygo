package ygo

import (
	"cmp"
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
	panic("not implemented")
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
