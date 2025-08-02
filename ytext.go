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
