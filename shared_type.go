package ygo

type SharedType interface {
	Parent() SharedType
	Integrate(doc *Doc, item Block) error
	ID() *ID
	Doc() *Doc
	Write(encoder UpdateEncoder) error
}

type BaseSharedType struct {
	item   Block
	blocks map[string]Block
	start  Block
	doc    *Doc
	length uint64
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

func (t *BaseSharedType) Integrate(doc *Doc, item Block) error {
	t.item = item
	t.doc = doc
	return nil
}

func (t *BaseSharedType) Write(encoder UpdateEncoder) error {
	panic("not implemented")
}
