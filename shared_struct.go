package ygo

type SharedStruct interface {
	Content() SharedContent
	ID() *ID
	Length() uint64
	Deleted() bool
	Delete(*Transaction) error
	Write(encoder UpdateEncoder, offset uint64) error
}
