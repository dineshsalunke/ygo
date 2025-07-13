package ygo

type ContentFormat struct {
	key   string
	value any
}

func newContentFormat(key string, value any) *ContentFormat {
	return &ContentFormat{key: key, value: value}
}

func (c *ContentFormat) Length() uint64 {
	panic("not implemented")
}

func (c *ContentFormat) Content() []any {
	panic("not implemented")
}

func (c *ContentFormat) IsCountable() bool {
	panic("not implemented")
}

func (c *ContentFormat) Copy() SharedContent {
	panic("not implemented")
}

func (c *ContentFormat) Splice(offset uint64) SharedContent {
	panic("not implemented")
}

func (c *ContentFormat) MergeWith(right SharedContent) bool {
	panic("not implemented")
}

func (c *ContentFormat) Integrate(tx *Transaction, item SharedStruct) {
	panic("not implemented")
}

func (c *ContentFormat) Delete(tx *Transaction) {
	panic("not implemented")
}

func (c *ContentFormat) GC(store *StructStore) {
	panic("not implemented")
}

func (c *ContentFormat) Write(encoder UpdateEncoder, offset uint64) error {
	panic("not implemented")
}

func (c *ContentFormat) GetRef() uint64 {
	panic("not implemented")
}
