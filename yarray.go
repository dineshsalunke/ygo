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

func (ymap *YArray) Write(encoder UpdateEncoder) error {
	return encoder.WriteTypeRef(0)
}
