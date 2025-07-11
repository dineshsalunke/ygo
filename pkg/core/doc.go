package core

type Doc struct {
	store *StructStore
}

func newDoc() *Doc {
	return &Doc{}
}

func NewDoc() *Doc {
	return newDoc()
}

func (doc *Doc) Store() *StructStore {
	return doc.store
}
