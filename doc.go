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

func (doc *Doc) GetMap(name string) (*YMap, error) {
	var item *YMap
	var check bool

	st, check := doc.share[name]
	if check {
		item, check = st.(*YMap)
		if !check {
			return nil, fmt.Errorf("type with name %s is already defined with different type", name)
		}
	} else {
		item = &YMap{}
		item.Integrate(doc, nil)
		doc.share[name] = item
	}

	return item, nil
}

func (doc *Doc) GetArray(name string) (*YArray, error) {
	var item *YArray
	var check bool

	st, check := doc.share[name]
	if check {
		item, check = st.(*YArray)
		if !check {
			return nil, fmt.Errorf("type with name %s is already defined with different type", name)
		}
	} else {
		item = &YArray{}
		item.Integrate(doc, nil)
		doc.share[name] = item
	}

	return item, nil
}

func (doc *Doc) GetText(name string) (*YText, error) {
	var item *YText
	var check bool

	st, check := doc.share[name]
	if check {
		item, check = st.(*YText)
		if !check {
			return nil, fmt.Errorf("type with name %s is already defined with different type", name)
		}
	} else {
		item = &YText{}
		item.Integrate(doc, nil)
		doc.share[name] = item
	}

	return item, nil
}
