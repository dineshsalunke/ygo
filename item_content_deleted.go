package ygo

type ItemContentDeleted struct {
	length uint64
}

func (itemcontentdeleted *ItemContentDeleted) Content() any {
	panic("not implemented") // TODO: Implement
}

func (itemcontentdeleted *ItemContentDeleted) Copy() ItemContent {
	panic("not implemented") // TODO: Implement
}

func (itemcontentdeleted *ItemContentDeleted) MergeWith(right ItemContent) bool {
	panic("not implemented") // TODO: Implement
}

func (itemcontentdeleted *ItemContentDeleted) Integrate(tx *Transaction, item Block) error {
	panic("not implemented") // TODO: Implement
}

func (itemcontentdeleted *ItemContentDeleted) Delete(tx *Transaction) error {
	panic("not implemented") // TODO: Implement
}

func (itemcontentdeleted *ItemContentDeleted) GC(tx *Transaction) error {
	panic("not implemented") // TODO: Implement
}

func (content *ItemContentDeleted) Write(encoder UpdateEncoder, offset uint64, offsetEnd uint64) error {
	panic("not implemented") // TODO: Implement
}

func newItemContentDeleted(length uint64) *ItemContentDeleted {
	return &ItemContentDeleted{
		length: length,
	}
}

func (content *ItemContentDeleted) Length() uint64 {
	return content.length
}

func (content *ItemContentDeleted) IsCountable() bool {
	return true
}

func (content *ItemContentDeleted) GetRef() Kind {
	return BlockKindItemDeleted
}

func (content *ItemContentDeleted) Splice(offset uint64) ItemContent {
	right := newItemContentDeleted(content.length - offset)
	content.length = offset
	return right
}

func init() {
	Decoders[BlockKindItemDeleted] = func(decoder UpdateDecoder, info Kind) (ItemContent, error) {
		length, err := decoder.ReadLength()
		if err != nil {
			return nil, err
		}
		return newItemContentDeleted(length), nil
	}
}
