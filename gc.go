package ygo

type GC struct{ *block }

func newGC(id *ID, length uint64) *GC {
	return &GC{
		block: newBlock(id, length),
	}
}

func (self *GC) Parent() SharedType {
	return nil
}

func (self *GC) Kind() Kind {
	return BlockKindGC
}

func (self *GC) Length() uint64 {
	return self.length
}

func (self *GC) ID() *ID {
	return self.id
}
