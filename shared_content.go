package ygo

type SharedContent interface {
	Length() uint64
	Content() []any
	IsCountable() bool
	Copy() SharedContent
	Splice(offset uint64) SharedContent
	MergeWith(right SharedContent) bool
	Integrate(*Transaction, SharedStruct)
	Delete(*Transaction)
	GC(store *StructStore)
	Write(encoder UpdateEncoder, offset uint64) error
	GetRef() uint64
}
