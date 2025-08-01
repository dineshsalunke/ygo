package ygo

import (
	"fmt"
)

type Doc struct {
	share map[string]SharedType
}

func newDoc() *Doc {
	return &Doc{
		share: make(map[string]SharedType),
	}
}

func (doc *Doc) NewTransaction(opts ...TxOption) *Transaction {
	return newTransaction(doc, opts...)
}

func (doc *Doc) get(key string) (SharedType, error) {
	t, has := doc.share[key]
	if has {
		return t, nil
	}
	return nil, fmt.Errorf("no type registered for key %s", key)
}
