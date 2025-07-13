package ygo

type ContentEmbed struct {
	embed any
}

func newContentEmbed(embed any) *ContentEmbed {
	return &ContentEmbed{embed}
}

func (c *ContentEmbed) Length() uint64 {
	panic("not implemented")
}

func (c *ContentEmbed) Content() []any {
	panic("not implemented")
}

func (c *ContentEmbed) IsCountable() bool {
	panic("not implemented")
}

func (c *ContentEmbed) Copy() SharedContent {
	panic("not implemented")
}

func (c *ContentEmbed) Splice(offset uint64) SharedContent {
	panic("not implemented")
}

func (c *ContentEmbed) MergeWith(right SharedContent) bool {
	panic("not implemented")
}

func (c *ContentEmbed) Integrate(tx *Transaction, item SharedStruct) {
	panic("not implemented")
}

func (c *ContentEmbed) Delete(tx *Transaction) {
	panic("not implemented")
}

func (c *ContentEmbed) GC(store *StructStore) {
	panic("not implemented")
}

func (c *ContentEmbed) Write(encoder UpdateEncoder, offset uint64) error {
	panic("not implemented")
}

func (c *ContentEmbed) GetRef() uint64 {
	panic("not implemented")
}
