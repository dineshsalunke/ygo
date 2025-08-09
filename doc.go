package ygo

import (
	"reflect"
)

type Doc struct {
	store *StructStore
	share map[string]SharedType
}

func newDoc() *Doc {
	return &Doc{
		share: make(map[string]SharedType),
		store: newStructStore(),
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

	t = newBaseSharedType()
	if err := t.Integrate(doc, nil); err != nil {
		return nil, err
	}
	doc.share[key] = t
	return t, nil
}

func (doc *Doc) GetItemKey(item any) (string, bool) {
	for key, t := range doc.share {
		if reflect.TypeOf(t) == reflect.TypeOf(item) {
			return key, true
		}
	}
	return "", false
}
