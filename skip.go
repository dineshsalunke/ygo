package ygo

type Skip struct {
	*block
}

func newSkip(id *ID, length uint64) *GC {
	return &GC{
		block: newBlock(id, length),
	}
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
