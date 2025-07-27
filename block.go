package ygo

type Kind byte

type Block interface {
	Kind() Kind
	Length() uint64
}

type block struct {
	id     *ID
	length uint64
}

func newBlock(id *ID, length uint64) *block {
	return &block{id: id, length: length}
}

func (self *block) Length() uint64 {
	return self.length
}

type GC struct{ *block }

func newGC(id *ID, length uint64) *GC {
	return &GC{
		block: newBlock(id, length),
	}
}

func (self *GC) Kind() Kind {
	return BlockKindGC
}

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
