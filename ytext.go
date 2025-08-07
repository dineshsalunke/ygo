package ygo

type pendingOperations func() error

type YText struct {
	*BaseSharedType

	pending []pendingOperations
}

func newYText() *YText {
	return &YText{
		BaseSharedType: &BaseSharedType{},
		pending:        make([]pendingOperations, 0),
	}
}

func (ymap *YText) Write(encoder UpdateEncoder) error {
	return encoder.WriteTypeRef(2)
}
