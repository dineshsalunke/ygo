package ygo

type ContentDeleted struct {
	length uint64
}

func newContentDeleted(length uint64) *ContentDeleted {
	return &ContentDeleted{length: length}
}

func (c *ContentDeleted) Length() uint64 {
	panic("not implemented")
}

func (c *ContentDeleted) Content() []any {
	panic("not implemented")
}

func (c *ContentDeleted) IsCountable() bool {
	panic("not implemented")
}

func (c *ContentDeleted) Copy() SharedContent {
	panic("not implemented")
}

func (c *ContentDeleted) Splice(offset uint64) SharedContent {
	panic("not implemented")
}

func (c *ContentDeleted) MergeWith(right SharedContent) bool {
	panic("not implemented")
}

func (c *ContentDeleted) Integrate(tx *Transaction, item SharedStruct) {
	panic("not implemented")
}

func (c *ContentDeleted) Delete(tx *Transaction) {
	panic("not implemented")
}

func (c *ContentDeleted) GC(store *StructStore) {
	panic("not implemented")
}

func (c *ContentDeleted) Write(encoder UpdateEncoder, offset uint64) error {
	panic("not implemented")
}

func (c *ContentDeleted) GetRef() uint64 {
	panic("not implemented")
}
