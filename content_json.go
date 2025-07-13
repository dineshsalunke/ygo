package ygo

type ContentJSON struct {
	json any
}

func newContentJSON(json any) *ContentJSON {
	return &ContentJSON{json: json}
}

func (c *ContentJSON) Length() uint64 {
	panic("not implemented")
}

func (c *ContentJSON) Content() []any {
	panic("not implemented")
}

func (c *ContentJSON) IsCountable() bool {
	panic("not implemented")
}

func (c *ContentJSON) Copy() SharedContent {
	panic("not implemented")
}

func (c *ContentJSON) Splice(offset uint64) SharedContent {
	panic("not implemented")
}

func (c *ContentJSON) MergeWith(right SharedContent) bool {
	panic("not implemented")
}

func (c *ContentJSON) Integrate(tx *Transaction, item SharedStruct) {
	panic("not implemented")
}

func (c *ContentJSON) Delete(tx *Transaction) {
	panic("not implemented")
}

func (c *ContentJSON) GC(store *StructStore) {
	panic("not implemented")
}

func (c *ContentJSON) Write(encoder UpdateEncoder, offset uint64) error {
	panic("not implemented")
}

func (c *ContentJSON) GetRef() uint64 {
	panic("not implemented")
}
