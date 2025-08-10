package ygo

import (
	"reflect"
)

type Doc struct {
	store               *StructStore
	share               map[string]SharedType
	activeTransaction   *Transaction
	transactionCleanups []*Transaction
}

func newDoc() *Doc {
	return &Doc{
		share:               make(map[string]SharedType),
		store:               newStructStore(),
		transactionCleanups: make([]*Transaction, 0),
	}
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

func (doc *Doc) GetMap(key string) (*YMap, error) {
	share, has := doc.share[key]
	if has {
		ymap, isymap := share.(*YMap)
		if isymap {
			return ymap, nil
		}

		basetype, isbasetype := share.(*BaseSharedType)
		if isbasetype {
			ymap := &YMap{
				BaseSharedType: basetype,
			}
			ymap.Integrate(doc, nil)
			doc.share[key] = ymap
			return ymap, nil
		}
	}

	ymap := newYMap(make(map[string]any))
	if err := ymap.Integrate(doc, nil); err != nil {
		return nil, err
	}

	return ymap, nil
}

func (doc *Doc) GetItemKey(item any) (string, bool) {
	for key, t := range doc.share {
		if reflect.TypeOf(t) == reflect.TypeOf(item) {
			return key, true
		}
	}
	return "", false
}
