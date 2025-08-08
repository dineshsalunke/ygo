package ygo

type Skip struct {
	*BaseBlockType
}

func newSkip(id *ID, length uint64) *Skip {
	return &Skip{
		BaseBlockType: &BaseBlockType{
			id:     id,
			length: length,
		},
	}
}
