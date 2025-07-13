package ygo

type ContentDoc struct {
	doc *Doc
}

func newContentDoc(doc *Doc) *ContentDoc {
	return &ContentDoc{doc: doc}
}

func (c *ContentDoc) Length() uint64 {
	panic("not implemented")
}

func (c *ContentDoc) Content() []any {
	panic("not implemented")
}

func (c *ContentDoc) IsCountable() bool {
	panic("not implemented")
}

func (c *ContentDoc) Copy() SharedContent {
	panic("not implemented")
}

func (c *ContentDoc) Splice(offset uint64) SharedContent {
	panic("not implemented")
}

func (c *ContentDoc) MergeWith(right SharedContent) bool {
	panic("not implemented")
}

func (c *ContentDoc) Integrate(tx *Transaction, item SharedStruct) {
	panic("not implemented")
}

func (c *ContentDoc) Delete(tx *Transaction) {
	panic("not implemented")
}

func (c *ContentDoc) GC(store *StructStore) {
	panic("not implemented")
}

func (c *ContentDoc) Write(encoder UpdateEncoder, offset uint64) error {
	panic("not implemented")
}

func (c *ContentDoc) GetRef() uint64 {
	panic("not implemented")
}
