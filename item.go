package ygo

type Item struct {
	*block
	origin       *ID
	left         block
	right_origin *ID
	right        Block
	parent       any
	parent_sub   string
	red_one      Block
	content      ItemContent
	info         byte
}

func (self *Item) Kind() Kind {
	panic("not implemented")
}

func decode_item(id *ID, decoder UpdateDecoder, info byte) (*Item, error) {
	panic("not implemented")
}
