package ygo

type ContentString struct {
	str string
}

func newContentString(str string) *ContentString {
	return &ContentString{str: str}
}

func (c *ContentString) Length() uint64 {
	panic("not implemented")
}

func (c *ContentString) Content() []any {
	panic("not implemented")
}

func (c *ContentString) IsCountable() bool {
	panic("not implemented")
}

func (c *ContentString) Copy() SharedContent {
	panic("not implemented")
}

func (c *ContentString) Splice(offset uint64) SharedContent {
	panic("not implemented")
}

func (c *ContentString) MergeWith(right SharedContent) bool {
	panic("not implemented")
}

func (c *ContentString) Integrate(tx *Transaction, item SharedStruct) {
	panic("not implemented")
}

func (c *ContentString) Delete(tx *Transaction) {
	panic("not implemented")
}

func (c *ContentString) GC(store *StructStore) {
	panic("not implemented")
}

func (c *ContentString) Write(encoder UpdateEncoder, offset uint64) error {
	panic("not implemented")
}

func (c *ContentString) GetRef() uint64 {
	panic("not implemented")
}
