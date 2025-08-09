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

func writeBlocks(encoder UpdateEncoder, blocks []Block, client uint64, idRanges []*IdRange) error {
	blocksToWrite := 0

	type tempStruct struct {
		start      uint64
		end        uint64
		startClock uint64
		endClock   uint64
	}
	indexRanges := make([]tempStruct, 0)
	firstPossibleClock := blocks[0].ClockStart()
	lastBlock := blocks[len(blocks)-1]
	lastPossibleClock := lastBlock.ClockEnd()

	for _, idRange := range idRanges {
		startClock := max(idRange.clock, firstPossibleClock)
		endClock := min(idRange.clock+idRange.length, lastPossibleClock)
		if startClock < endClock {
			startIndex, _ := binarySearchClockIndex(blocks, startClock)
			endIndex, _ := binarySearchClockIndex(blocks, endClock-1)
			blocksToWrite += endIndex - startIndex
			indexRanges = append(indexRanges, tempStruct{start: uint64(startIndex), end: uint64(endIndex) + 1, startClock: startClock, endClock: endClock})
		}
	}

	blocksToWrite += len(idRanges) - 1
	clock := indexRanges[0].startClock
	if err := encoder.WriteVarUint(uint64(blocksToWrite)); err != nil {
		return err
	}
	if err := encoder.WriteClient(client); err != nil {
		return err
	}
	if err := encoder.WriteVarUint(clock); err != nil {
		return err
	}

	for _, indexRange := range indexRanges {
		skipLength := indexRange.startClock - clock
		if skipLength > 0 {
			newSkip(newID(client, clock), skipLength).Write(encoder, 0, 0)
			clock += skipLength
		}
		for i := indexRange.start; i < indexRange.end; i++ {
			block := blocks[i]
			blockEnd := block.ClockEnd()
			offsetEnd := max(blockEnd-indexRange.endClock, 0)
			if err := block.Write(encoder, clock-block.ClockStart(), byte(offsetEnd)); err != nil {
				return err
			}
		}
	}

	return nil
}

func writeClientsBlocks(encoder UpdateEncoder, ss *StructStore, sv StateVector) error {
	sm := make(StateVector, 0)
	for client, clock := range sv {
		cl := ss.GetClientClockEnd(client)
		if cl > clock {
			sm[client] = clock
		}
	}
	for client := range ss.GetStateVector() {
		_, has := sv[client]
		if !has {
			sm[client] = 0
		}
	}
	clientIds := make([]uint64, len(sm))
	for client := range sm {
		clientIds = append(clientIds, client)
	}
	slices.SortFunc(clientIds, func(a, b uint64) int {
		return cmp.Compare(b, a)
	})

	for _, client := range clientIds {
		clock := sm[client]
		blocks, has := ss.clients[client]
		if has && len(blocks) > 0 {
			lastBlock := blocks[len(blocks)-1]
			if err := writeBlocks(encoder, blocks, client, []*IdRange{
				{clock: clock, length: lastBlock.ClockEnd() - clock},
			}); err != nil {
				return err
			}
		}

	}
	return nil
}

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

func (ss *StructStore) IntegrateStructs(tx *Transaction, remoteBlockSet *StructSet) (*PendingStructs, error) {
	stack := make([]Block, 0)
	clientIds := make([]uint64, len(remoteBlockSet.clients))
	i := 0
	for client := range remoteBlockSet.clients {
		clientIds[i] = client
		i++
	}
	slices.SortFunc(clientIds, func(a, b uint64) int {
		return int(b - a)
	})

	// Helper to get the next available struct batch to process.
	// It iterates backwards through the sorted client IDs.
	getNextBlockTarget := func() *StructRange {
		for len(clientIds) > 0 {
			clientID := clientIds[len(clientIds)-1]
			target := remoteBlockSet.clients[clientID]
			if target.i < len(target.refs) {
				return target
			}
			// This client is done, pop it.
			clientIds = clientIds[:len(clientIds)-1]
		}
		return nil
	}

	curBlockTarget := getNextBlockTarget()
	if curBlockTarget == nil {
		return nil, nil // Nothing to integrate
	}

	restBlocks := newStructStore()
	missingSv := make(map[uint64]uint64, 0)

	updateMissingSv := func(client, clock uint64) {
		missingClock, exists := missingSv[client]
		if !exists || missingClock > clock {
			missingSv[client] = clock
		}
	}

	addStackToSS := func() {
		for _, item := range stack {
			client := item.Client()
			inapplicableItems, has := remoteBlockSet.clients[client]
			if has {
				inapplicableItems.i -= 1
				restBlocks.clients[client] = inapplicableItems.refs[inapplicableItems.i:]
				delete(remoteBlockSet.clients, client)
				inapplicableItems.i = 0
				inapplicableItems.refs = make([]Block, 0)
			} else {
				restBlocks.clients[client] = []Block{item}
			}

			filtered := make([]uint64, 0)
			for _, id := range clientIds {
				if id != client {
					filtered = append(filtered, client)
				}
			}
			clientIds = filtered
		}
		stack = make([]Block, 0)
	}

	curBlockTarget.i += 1
	stackHead := curBlockTarget.refs[curBlockTarget.i]

	state := make(map[uint64]uint64, 0)

	for {
		if _, ok := stackHead.(*Skip); !ok {
			localClock, has := state[stackHead.Client()]
			if !has {
				localClock = ss.GetClientClockEnd(stackHead.Client())
				state[stackHead.Client()] = localClock
			}

			offset := int(localClock - stackHead.ClockStart())
			missing, hasMissing, err := stackHead.GetMissing(tx, ss)
			if err != nil {
				return nil, err
			}
			if hasMissing {
				stack = append(stack, stackHead)

				blockRefs, hasBlockRefs := remoteBlockSet.clients[missing]
				missingFromStack := slices.ContainsFunc(stack, func(b Block) bool {
					return b.Client() == missing
				})
				if !hasBlockRefs ||
					blockRefs.i == len(blockRefs.refs) ||
					missing == stackHead.Client() ||
					missingFromStack {
					updateMissingSv(missing, ss.GetClientClockEnd(missing))
					addStackToSS()
				} else if hasBlockRefs {
					blockRefs.i += blockRefs.i
					stackHead = blockRefs.refs[blockRefs.i]
					continue
				}
				//
			} else {
				if offset < 0 {
					skip := newSkip(newID(stackHead.Client(), localClock), uint64(-offset))
					if err := skip.Integrate(tx, 0); err != nil {
						return nil, err
					}
				}
				if err := stackHead.Integrate(tx, 0); err != nil {
					return nil, err
				}
				state[stackHead.Client()] = max(stackHead.ClockEnd(), localClock)
			}
		}

		if len(stack) > 0 {
			stackHead = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		} else if curBlockTarget != nil && curBlockTarget.i < len(curBlockTarget.refs) {
			stackHead = curBlockTarget.refs[curBlockTarget.i]
			curBlockTarget.i += 1
		} else {
			curBlockTarget = getNextBlockTarget()
			if curBlockTarget == nil {
				break
			} else {
				stackHead = curBlockTarget.refs[curBlockTarget.i]
				curBlockTarget.i += 1
			}
		}
	}

	if len(restBlocks.clients) > 0 {
		encoder := newUpdateEncoderV1()
		if err := writeClientsBlocks(encoder, restBlocks, make(StateVector)); err != nil {
			return nil, err
		}
		if err := encoder.WriteVarUint(0); err != nil {
			return nil, err
		}

		return &PendingStructs{
			update:       encoder.Bytes(),
			missingState: missingSv,
		}, nil
	}
	return nil, nil
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
