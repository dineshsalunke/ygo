package ygo

type SharedType interface {
	Parent() SharedType
	Integrate(doc *Doc, item *Item) error
	ID() *ID
	Doc() *Doc
	Write(encoder UpdateEncoder) error
	GetBlock(key string) Block
	SetBlock(key string, block Block)
	Start() Block
	SetStart(block Block)
	Block() Block
	Length() uint64
	SetLength(length uint64)
}

type BaseSharedType struct {
	block  Block
	blocks map[string]Block
	start  Block
	doc    *Doc
	length uint64
}

func newBaseSharedType() *BaseSharedType {
	return &BaseSharedType{
		block:  nil,
		start:  nil,
		blocks: make(map[string]Block),
	}
}

func (self *BaseSharedType) SetLength(length uint64) {
	self.length = length
}

func (self *BaseSharedType) Length() uint64 {
	return self.length
}

func (self *BaseSharedType) Block() Block {
	return self.block.(Block)
}

func (self *BaseSharedType) SetStart(block Block) {
	self.start = block
}

func (self *BaseSharedType) Start() Block {
	return self.start
}

func (self *BaseSharedType) ID() *ID {
	return self.block.ID()
}

func (self *BaseSharedType) Doc() *Doc {
	return self.doc
}

func (self *BaseSharedType) Parent() SharedType {
	return self.block.Parent().(SharedType)
}

func (self *BaseSharedType) GetBlock(key string) Block {
	return self.blocks[key]
}

func (self *BaseSharedType) SetBlock(key string, block Block) {
	self.blocks[key] = block
}

func (self *BaseSharedType) Integrate(doc *Doc, item *Item) error {
	self.block = item
	self.doc = doc
	return nil
}

func (self *BaseSharedType) Write(encoder UpdateEncoder) error {
	panic("not implemented")
}
