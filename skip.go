package ygo

import "fmt"

type Skip struct {
	*block
}

func (self *Skip) GoString() string {
	return fmt.Sprintf("Skip{id:%#v,length:%d}", self.id, self.length)
}

func newSkip(id *ID, length uint64) *Skip {
	return &Skip{
		block: newBlock(id, length),
	}
}

func (self *Skip) Parent() SharedType {
	return nil
}

func (self *Skip) Kind() Kind {
	return BlockKindSkip
}

func (self *Skip) Length() uint64 {
	return self.length
}

func (self *Skip) ID() *ID {
	return self.id
}
