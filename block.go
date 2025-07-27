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

func (self *block) Length() uint64 {
	return self.length
}

type GC struct{ *block }

func newGC(id *ID, length uint64) *GC {
	return &GC{
		block: &block{id: id, length: length},
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
		block: &block{id: id, length: length},
	}
}

func (self *Skip) Kind() Kind {
	return BlockKindSkip
}
