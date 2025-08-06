package ygo

type Kind byte

type Block interface {
	Kind() Kind
	Length() uint64
	ID() *ID
	Parent() SharedType
	ClockLength() uint64
	Splice(offset uint64, tx *Transaction) Block
	GetMissing(tx *Transaction, store *StructStore) (uint64, bool, error)
	Integrate(tx *Transaction, offset uint64) error
	Write(encoder UpdateEncoder, offset uint64, offsetEnd uint64) error
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

func (self *block) ClockLength() uint64 {
	return self.id.clock + self.length
}

func (self *block) Splice(offset uint64, tx *Transaction) Block {
	panic("not implemented")
}

func (self *block) Integrate(tx *Transaction, offset uint64) error {
	panic("not implemented")
}

func (self *block) GetMissing(tx *Transaction, store *StructStore) (uint64, bool, error) {
	panic("not implemented")
}

