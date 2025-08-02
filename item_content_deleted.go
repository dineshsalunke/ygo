package ygo

type ItemContentDeleted struct {
	length uint64
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

func init() {
	Decoders[BlockKindItemDeleted] = func(decoder UpdateDecoder, info Kind) (ItemContent, error) {
		length, err := decoder.ReadLength()
		if err != nil {
			return nil, err
		}
		return newItemContentDeleted(length), nil
	}
}
