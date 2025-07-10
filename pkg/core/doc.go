package core

type Doc struct{}

func newDoc() *Doc {
	return &Doc{}
}

func NewDoc() *Doc {
	return newDoc()
}
