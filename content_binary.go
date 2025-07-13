package ygo

type ContentBinary struct {
	content []byte
}

func newContentBinary(content []byte) *ContentBinary {
	return &ContentBinary{content: content}
}

func (c *ContentBinary) Length() uint64 {
	panic("not implemented")
}

func (c *ContentBinary) Content() []any {
	panic("not implemented")
}

func (c *ContentBinary) IsCountable() bool {
	panic("not implemented")
}

func (c *ContentBinary) Copy() SharedContent {
	panic("not implemented")
}

func (c *ContentBinary) Splice(offset uint64) SharedContent {
	panic("not implemented")
}

func (c *ContentBinary) MergeWith(right SharedContent) bool {
	panic("not implemented")
}

func (c *ContentBinary) Integrate(tx *Transaction, item SharedStruct) {
	panic("not implemented")
}

func (c *ContentBinary) Delete(tx *Transaction) {
	panic("not implemented")
}

func (c *ContentBinary) GC(store *StructStore) {
	panic("not implemented")
}

func (c *ContentBinary) Write(encoder UpdateEncoder, offset uint64) error {
	panic("not implemented")
}

func (c *ContentBinary) GetRef() uint64 {
	panic("not implemented")
}
