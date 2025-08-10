package ygo

import "reflect"
type YMap struct {
	*BaseSharedType

	// Preliminary content before integration
	prelimContent map[string]any
}

func newYMap(prelim map[string]any) *YMap {
	return &YMap{
		BaseSharedType: newBaseSharedType(),
		prelimContent:  prelim,
	}
}

// NewYMap creates a new YMap instance
func NewYMap() *YMap {
	return newYMap(make(map[string]any))
}

func NewYMapWithEntries(entries map[string]any) *YMap {
	return newYMap(entries)
}

func (ymap *YMap) Write(encoder UpdateEncoder) error {
	return encoder.WriteTypeRef(1)
}
func (ymap *YMap) Delete(key string, tx *Transaction) error {
	v, ok := ymap.blocks[key]
	if ok {
		if err := v.Delete(tx); err != nil {
			return err
		}
	}
	return nil
}

func (ymap *YMap) Get(key string) (any, bool) {
	v, ok := ymap.blocks[key]
	if ok {
		if value, ok := v.(*Item); ok && !v.Deleted() {
			return value.content.Content()[value.length-1], true
		}
	}
	return nil, false
}

func (ymap *YMap) Has(key string) bool {
	v, ok := ymap.blocks[key]
	if ok {
		return !reflect.ValueOf(v).IsNil() && !v.Deleted()
	}
	return false
}

