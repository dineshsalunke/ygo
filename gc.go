package ygo

type GC struct {
	*BaseBlockType
}

func newGC(id *ID, length uint64) *GC {
	return &GC{
		BaseBlockType: &BaseBlockType{
			id:     id,
			length: length,
		},
	}
}
