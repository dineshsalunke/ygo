package ygo

import "fmt"
type Item struct {
	*block
	origin       *ID
	left         Block
	right_origin *ID
	right        Block
	parent       any
	parent_sub   string
	red_one      Block
	content      ItemContent
	info         byte
}

func (self *Item) GoString() string {
	return fmt.Sprintf("Item{id: %#v,length:%d}", self.id, self.length)
}

func (self *Item) Kind() Kind {
	return Kind(self.info & 31)
}

func (self *Item) Length() uint64 {
	return self.length
}

func (self *Item) ID() *ID {
	return self.id
}

func (self *Item) Parent() SharedType {
	return self.parent.(SharedType)
}

func decodeItem(id *ID, decoder UpdateDecoder, info Kind, doc *Doc) (*Item, error) {
	var err error
	var item *Item = &Item{
		block: &block{
			id: id,
		},
	}
	cant_copy_parent_info := byte(info)&(HasOriginFlag|HasRightOriginFlag) == 0

	if byte(info)&HasOriginFlag != 0 {
		item.origin, err = decoder.ReadLeftID()
		if err != nil {
			return nil, err
		}
	}
	if byte(info)&HasRightOriginFlag != 0 {
		item.right_origin, err = decoder.ReadRightID()
		if err != nil {
			return nil, err
		}
	}
	if cant_copy_parent_info {
		parent_info, err := decoder.ReadParentInfo()
		if err != nil {
			return nil, err
		}
		if parent_info {
			// Parent string
			// TODO: Get the parent item from the doc
			key, err := decoder.ReadVarString()
			if err != nil {
				return nil, err
			}
			item.parent, err = doc.get(key)
			if err != nil {
				return nil, err
			}
		} else {
			item.parent, err = decoder.ReadLeftID()
			if err != nil {
				return nil, err
			}
		}

		if (byte(info) & HasParentSubFlag) != 0 {
			item.parent_sub, err = decoder.ReadVarString()
			if err != nil {
				return nil, err
			}
		}
	}

	item.content, err = decodeItemContent(decoder, info)
	if err != nil {
		return nil, err
	}
	item.block.length = item.content.Length()
	return item, nil
}
