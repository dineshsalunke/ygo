package ygo

import (
	"cmp"
	"fmt"
	"slices"
)

type BlockList []Block

func (list *BlockList) Clock() uint64 {
	if len(*list) > 0 {
		lastBlock := (*list)[len(*list)-1]
		return lastBlock.ClockEnd() + 1
	}
	return 0
}

func (list *BlockList) BinarySearchByClock(clock uint64) (int, bool) {
	return slices.BinarySearchFunc((*list), clock, func(elem Block, clock uint64) int {
		if elem.ContainsClock(clock) {
			return 0
		}
		return cmp.Compare(elem.ClockStart(), clock)
	})
}

func (list *BlockList) FindIndexCleanStart(tx *Transaction, clock uint64) (uint64, error) {
	index, exists := (*list).BinarySearchByClock(clock)
	if !exists {
		return 0, fmt.Errorf("no block found for clock %d", clock)
	}

	block := (*list)[index]
	if block.ClockStart() < clock {
		nBlock, err := block.Split(tx, clock-block.ClockStart())
		if err != nil {
			return 0, err
		}
		*list = slices.Insert((*list), index+1, nBlock)
		return uint64(index + 1), nil
	}

	return uint64(index), nil
}
