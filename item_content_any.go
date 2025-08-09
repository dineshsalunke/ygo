package ygo

type ItemContentAny struct {
	contents []any
}

func (self *ItemContentAny) Content() any {
	panic("not implemented") // TODO: Implement
}

func (self *ItemContentAny) Copy() ItemContent {
	panic("not implemented") // TODO: Implement
}

func (self *ItemContentAny) MergeWith(right ItemContent) bool {
	if cany, ok := right.(*ItemContentAny); ok {
		self.contents = append(self.contents, cany.contents...)
		return true
	}
	return false
}

func (self *ItemContentAny) Integrate(tx *Transaction, item Block) error {
	return nil
}

func (self *ItemContentAny) Delete(tx *Transaction) error {
	return nil
}

func (self *ItemContentAny) GC(tx *Transaction) error {
	return nil
}

func (content *ItemContentAny) Write(encoder UpdateEncoder, offset uint64, offsetEnd uint64) error {
	end := uint64(len(content.contents)) - offset
	if err := encoder.WriteLength(end - offset); err != nil {
		return err
	}
	for i := offset; i < end; i++ {
		c := content.contents[i]
		if err := encoder.WriteAny(c); err != nil {
			return err
		}

	}
	return nil
}

func (self *ItemContentAny) Length() uint64 {
	return uint64(len(self.contents))
}

func (self *ItemContentAny) IsCountable() bool {
	return true
}

func (self *ItemContentAny) GetRef() Kind {
	return BlockKindItemAny
}

func (self *ItemContentAny) Splice(offset uint64) ItemContent {
	c := make([]any, uint64(len(self.contents)-int(offset)))
	copy(c, self.contents[offset:])
	self.contents = self.contents[:offset]
	right := newItemContentAny(c)
	return right
}

func newItemContentAny(contents []any) *ItemContentAny {
	return &ItemContentAny{
		contents: contents,
	}
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
