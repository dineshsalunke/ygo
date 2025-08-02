package ygo

type YArray struct {
	*BaseSharedType

	// Initial content before integration
	premlimContent []any
}

func newYArray() *YArray {
	return &YArray{
		BaseSharedType: &BaseSharedType{},
		premlimContent: make([]any, 0),
	}
}
