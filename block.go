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

func (self *GC) Kind() Kind {
	return BlockKindGC
}

type Skip struct {
	*block
}

func (self *Skip) Kind() Kind {
	return BlockKindSkip
}

type Item struct {
	*block
}

func (self *Item) Kind() Kind {
	panic("not implemented")
}
