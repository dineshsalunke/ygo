package ygo

import (
	"cmp"
	"slices"
)

type Kind byte

type Block interface {
	Deleted() bool
	Delete(tx *Transaction) error
	MergeWith(right Block) error
	Write(encoder UpdateEncoder, offset uint64, encodingRef byte) error
	Integrate(tx *Transaction, offset uint64) error
	GetMissing(tx *Transaction, store *BlockStore) (uint64, bool, error)
	ClockStart() uint64
	ClockEnd() uint64
	ClockRange() (uint64, uint64)
	Client() uint64
	Length() uint64
	ContainsClock(clock uint64) bool
	Split(tx *Transaction, diff uint64) (Block, error)
	ID() *ID
	Parent() SharedType
	LastID() *ID
	Splice(diff uint64) (Block, error)
}

type BaseBlockType struct {
	id     *ID
	length uint64
}

func (self *BaseBlockType) Delete(tx *Transaction) error {
	panic("not implemented")
}

func (self *BaseBlockType) ClockStart() uint64 {
	return self.id.clock
}

func (self *BaseBlockType) ClockEnd() uint64 {
	return self.id.clock + self.length - 1
}

func (self *BaseBlockType) ClockRange() (uint64, uint64) {
	return self.ClockStart(), self.ClockEnd()
}

func (self *BaseBlockType) Client() uint64 {
	return self.id.client
}

func (self *BaseBlockType) Length() uint64 {
	return self.length
}

func (self *BaseBlockType) ContainsClock(clock uint64) bool {
	return clock >= self.id.clock && clock <= self.ClockEnd()
}

func (self *BaseBlockType) ID() *ID {
	return self.id
}

func (self *BaseBlockType) Deleted() bool {
	panic("not implemented") // TODO: Implement
}

func (self *BaseBlockType) MergeWith(right Block) error {
	panic("not implemented") // TODO: Implement
}

func (self *BaseBlockType) Write(encoder UpdateEncoder, offset uint64, encodingRef byte) error {
	panic("not implemented") // TODO: Implement
}

func (self *BaseBlockType) Integrate(tx *Transaction, offset uint64) error {
	panic("not implemented") // TODO: Implement
}

func (self *BaseBlockType) GetMissing(tx *Transaction, store *BlockStore) (uint64, bool, error) {
	panic("not implemented") // TODO: Implement
}

func (self *BaseBlockType) Split(tx *Transaction, diff uint64) (Block, error) {
	return self.Splice(diff)
}

func (self *BaseBlockType) Splice(diff uint64) (Block, error) {
	panic("not implemented")
}

func (self *BaseBlockType) Parent() SharedType {
	panic("not implemented") // TODO: Implement
}

func (self *BaseBlockType) LastID() *ID {
	panic("not implemented") // TODO: Implement
}

func binarySearchClockIndex(blocks []Block, clock uint64) (int, bool) {
	return slices.BinarySearchFunc(blocks, clock, func(b Block, c uint64) int {
		if b.ContainsClock(c) {
			return 0
		}
		return cmp.Compare(b.ClockEnd(), c)
	})
}
