package ygo

type pendingOperations func() error

type YText struct {
	*BaseSharedType

	pending []pendingOperations
}

func newYText() *YText {
	return &YText{
		BaseSharedType: newBaseSharedType(),
		pending:        make([]pendingOperations, 0),
	}
}

func NewYText() *YText {
	return newYText()
}

func (ymap *YText) Write(encoder UpdateEncoder) error {
	return encoder.WriteTypeRef(2)
}
