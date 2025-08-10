package ygo

import (
	"fmt"
)

type ContentTypeKind byte

type ContentTypeDecoderFactory func(decoder UpdateDecoder) (SharedType, error)

var ContentTypeFactories map[ContentTypeKind]ContentTypeDecoderFactory = make(map[ContentTypeKind]ContentTypeDecoderFactory)

type ItemContentType struct {
	sharedType SharedType
}

func (itemcontenttype *ItemContentType) Content() []any {
	panic("not implemented") // TODO: Implement
}

func (itemcontenttype *ItemContentType) Copy() ItemContent {
	panic("not implemented") // TODO: Implement
}

func (itemcontenttype *ItemContentType) MergeWith(right ItemContent) bool {
	panic("not implemented") // TODO: Implement
}

func (itemcontenttype *ItemContentType) Integrate(tx *Transaction, item Block) error {
	panic("not implemented") // TODO: Implement
}

func (itemcontenttype *ItemContentType) Delete(tx *Transaction) error {
	panic("not implemented") // TODO: Implement
}

func (itemcontenttype *ItemContentType) GC(tx *Transaction) error {
	panic("not implemented") // TODO: Implement
}

func newItemContentType(sharedType SharedType) *ItemContentType {
	return &ItemContentType{
		sharedType: sharedType,
	}
}

func (content *ItemContentType) Write(encoder UpdateEncoder, offset uint64, offsetEnd uint64) error {
	return content.sharedType.Write(encoder)
}

func (content *ItemContentType) Length() uint64 {
	return 1
}

func (content *ItemContentType) IsCountable() bool {
	return true
}

func (content *ItemContentType) GetRef() Kind {
	return BlockKindItemType
}

func (content *ItemContentType) Splice(offset uint64) ItemContent {
	// This method is noop
	return nil
}

func init() {
	Decoders[BlockKindItemType] = func(decoder UpdateDecoder, info Kind) (ItemContent, error) {
		typeRef, err := decoder.ReadTypeRef()
		if err != nil {
			return nil, err
		}
		factory, has := ContentTypeFactories[ContentTypeKind(typeRef)]
		if !has {
			return nil, fmt.Errorf("%v factory not registered", ContentTypeKind(typeRef))
		}
		t, err := factory(decoder)
		if err != nil {
			return nil, err
		}
		return newItemContentType(t), nil
	}
}
