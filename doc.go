package ygo

import (
	"math/rand"
	"reflect"
)

type Doc struct {
	store               *BlockStore
	share               map[string]SharedType
	activeTransaction   *Transaction
	transactionCleanups []*Transaction
	clientID            uint64
}

func newDoc() *Doc {
	return &Doc{
		share:               make(map[string]SharedType),
		store:               newStructStore(),
		transactionCleanups: make([]*Transaction, 0),
		clientID:            rand.Uint64(),
	}
}

func NewDoc() *Doc {
	return newDoc()
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
			if err := ymap.Integrate(doc, nil); err != nil {
				return nil, err
			}
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

func (doc *Doc) ApplyUpdate(update []byte, origin any) error {
	_, err := Transact(doc, func(tx *Transaction) (any, error) {
		tx.local = false
		decoder := newUpdateDecoderV1(update)

		if err := applyUpdate(decoder, tx, origin); err != nil {
			return nil, err
		}
		return nil, nil
	}, origin, false)
	return err
}
