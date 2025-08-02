package ygo

type ItemContentString struct {
	str string
}

func newItemContentString(str string) *ItemContentString {
	return &ItemContentString{
		str: str,
	}
}

func (content *ItemContentString) Length() uint64 {
	return uint64(len(content.str))
}

func (content *ItemContentString) IsCountable() bool {
	return true
}

func (content *ItemContentString) GetRef() Kind {
	return BlockKindItemMove
}

func init() {
	Decoders[BlockKindItemString] = func(decoder UpdateDecoder, info Kind) (ItemContent, error) {
		str, err := decoder.ReadVarString()
		if err != nil {
			return nil, err
		}
		return newItemContentString(str), nil
	}
}
