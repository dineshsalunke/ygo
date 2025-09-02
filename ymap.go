package ygo

import (
	"fmt"
	"reflect"
)

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

func (ymap *YMap) Set(key string, value any) (any, error) {
	if ymap.doc != nil {
		return Transact(ymap.doc, func(tx *Transaction) (any, error) {
			left, _ := ymap.blocks[key]
			doc := tx.doc
			ownClientId := doc.clientID
			var content ItemContent = nil
			if value == nil {
				content = newItemContentAny([]any{value})
			} else {
				switch value.(type) {
				case int, *int, int8, *int8, int16, *int16, int32, *int32, int64, *int64, uint8, *uint8, uint16, *uint16, uint32, *uint32, uint64, *uint64, string, *string, bool, *bool:
					content = newItemContentAny([]any{value})
				case []byte, []*byte:
					panic("not implemented")
				case *Doc, Doc:
					panic("not implemented")
				default:
					if st, ok := value.(SharedType); ok {
						content = newItemContentType(st)
					} else {
						return nil, fmt.Errorf("failed to integrate item for key %s, unexpected value type %T", key, value)
					}
				}
			}
			nextClock := doc.store.GetClientClockEnd(ownClientId)
			item := newItem(newID(ownClientId, nextClock), WithLeftLastID(left), WithContent(content), WithParent(ymap), WithParentSub(key))
			if err := item.Integrate(tx, 0); err != nil {
				return value, fmt.Errorf("failed to integrate item for key %s value %#v: %w", key, value, err)
			}
			return value, nil
		}, nil, true)
	} else {
		ymap.prelimContent[key] = value
	}
	return value, nil
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

func (ymap *YMap) Clear() {
	panic("not implemented")
}
