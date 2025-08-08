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
}

type BaseSharedType struct {
	item   Block
	blocks map[string]Block
	start  Block
	doc    *Doc
	length uint64
}

func (t *BaseSharedType) Start() Block {
	return t.start
}

func (t *BaseSharedType) ID() *ID {
	return t.item.ID()
}

func (t *BaseSharedType) Doc() *Doc {
	return t.doc
}

func (t *BaseSharedType) Parent() SharedType {
	if t.item != nil {
		return t.item.Parent()
	}
	return nil
}

func (t *BaseSharedType) GetBlock(key string) Block {
	return t.blocks[key]
}

func (t *BaseSharedType) SetBlock(key string, block Block) {
	t.blocks[key] = block
}

func (t *BaseSharedType) Integrate(doc *Doc, item *Item) error {
	t.item = item
	t.doc = doc
	return nil
}

func (t *BaseSharedType) Write(encoder UpdateEncoder) error {
	panic("not implemented")
}
