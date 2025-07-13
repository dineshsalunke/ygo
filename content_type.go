package ygo

type ContentType struct {
	t SharedType
}

func newContentType(t SharedType) *ContentType {
	return &ContentType{t: t}
}

func (c *ContentType) Length() uint64 {
	panic("not implemented")
}

func (c *ContentType) Content() []any {
	panic("not implemented")
}

func (c *ContentType) IsCountable() bool {
	panic("not implemented")
}

func (c *ContentType) Copy() SharedContent {
	panic("not implemented")
}

func (c *ContentType) Splice(offset uint64) SharedContent {
	panic("not implemented")
}

func (c *ContentType) MergeWith(right SharedContent) bool {
	panic("not implemented")
}

func (c *ContentType) Integrate(tx *Transaction, item SharedStruct) {
	panic("not implemented")
}

func (c *ContentType) Delete(tx *Transaction) {
	panic("not implemented")
}

func (c *ContentType) GC(store *StructStore) {
	panic("not implemented")
}

func (c *ContentType) Write(encoder UpdateEncoder, offset uint64) error {
	panic("not implemented")
}

func (c *ContentType) GetRef() uint64 {
	panic("not implemented")
}
