package ygo

type Kind byte

type Block interface {
	Kind() Kind
	Length() uint64
	ID() *ID
	Parent() SharedType
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
