package ygo

type ItemContentAny struct {
	contents []any
}

func newItemContentAny(contents []any) *ItemContentAny {
	return &ItemContentAny{
		contents: contents,
	}
}

func (content *ItemContentAny) Length() uint64 {
	return uint64(len(content.contents))
}

func (content *ItemContentAny) IsCountable() bool {
	return true
}

func (content *ItemContentAny) GetRef() Kind {
	return BlockKindItemAny
}

func init() {
	Decoders[BlockKindItemAny] = func(decoder UpdateDecoder, info Kind) (ItemContent, error) {
		length, err := decoder.ReadLength()
		if err != nil {
			return nil, err
		}
		values := make([]any, 0, length)
		for range length {
			val, err := decoder.ReadAny()
			if err != nil {
				return nil, err
			}
			values = append(values, val)
		}
		return newItemContentAny(values), nil
	}
}
