package ygo

type ContentAny struct {
	arr []any
}

func newContentAny(arr []any) *ContentAny {
	return &ContentAny{
		arr: arr,
	}
}

func (c *ContentAny) Length() uint64 {
	panic("not implemented")
}

func (c *ContentAny) Content() []any {
	panic("not implemented")
}

func (c *ContentAny) IsCountable() bool {
	panic("not implemented")
}

func (c *ContentAny) Copy() SharedContent {
	panic("not implemented")
}

func (c *ContentAny) Splice(offset uint64) SharedContent {
	panic("not implemented")
}

func (c *ContentAny) MergeWith(right SharedContent) bool {
	panic("not implemented")
}

func (c *ContentAny) Integrate(tx *Transaction, item SharedStruct) {
	panic("not implemented")
}

func (c *ContentAny) Delete(tx *Transaction) {
	panic("not implemented")
}

func (c *ContentAny) GC(store *StructStore) {
	panic("not implemented")
}

func (c *ContentAny) Write(encoder UpdateEncoder, offset uint64) error {
	panic("not implemented")
}

func (c *ContentAny) GetRef() uint64 {
	panic("not implemented")
}
