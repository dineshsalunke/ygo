package ygo

import (
	"cmp"
	"slices"
)

type BlockList []Block

func (list *BlockList) BinarySearchByClock(clock uint64) (int, bool) {
	return slices.BinarySearchFunc((*list), clock, func(elem Block, clock uint64) int {
		if elem.ContainsClock(clock) {
			return 0
		}
		return cmp.Compare(elem.ClockStart(), clock)
	})
}
