package ygo

import "fmt"

type Doc struct {
	share               map[string]SharedType
	store               *StructStore
	transaction         *Transaction
	transactionCleanups []*Transaction
}

func newDoc() *Doc {
	return &Doc{
		store: newStructStore(),
	}
}

func NewDoc() *Doc {
	return newDoc()
}

