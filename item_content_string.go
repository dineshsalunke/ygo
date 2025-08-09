package ygo

type ItemContentString struct {
	str string
}

func (itemcontentstring *ItemContentString) Content() any {
	panic("not implemented") // TODO: Implement
}

func (itemcontentstring *ItemContentString) Copy() ItemContent {
	panic("not implemented") // TODO: Implement
}

func (itemcontentstring *ItemContentString) MergeWith(right ItemContent) bool {
	panic("not implemented") // TODO: Implement
}

func (itemcontentstring *ItemContentString) Integrate(tx *Transaction, item Block) error {
	panic("not implemented") // TODO: Implement
}

func (itemcontentstring *ItemContentString) Delete(tx *Transaction) error {
	panic("not implemented") // TODO: Implement
}

func (itemcontentstring *ItemContentString) GC(tx *Transaction) error {
	panic("not implemented") // TODO: Implement
}

func (content *ItemContentString) Write(encoder UpdateEncoder, offset uint64, offsetEnd uint64) error {
	str := content.str[offset : len(content.str)-int(offsetEnd)]
	if offset == 0 && offsetEnd == 0 {
		str = content.str[:]
	}
	if err := encoder.WriteString(str); err != nil {
		return err
	}
	return nil
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

func (content *ItemContentString) Splice(offset uint64) ItemContent {
	right := newItemContentString(content.str[offset:])
	content.str = content.str[:offset]
	return right
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
