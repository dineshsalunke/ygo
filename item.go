package ygo

import (
	"fmt"
	"reflect"
)

type Item struct {
	*BaseBlockType
	origin      *ID
	left        Block
	rightOrigin *ID
	right       Block
	parent      any
	parentSub   string
	redOne      *ID
	content     ItemContent
	info        uint16
}

type ItemOption func(item *Item)

func WithContent(content ItemContent) ItemOption {
	return func(item *Item) {
		item.content = content
		item.length = content.Length()
		if content.IsCountable() {
			item.SetCountable()
		}
	}
}

func WithLeft(left Block) ItemOption {
	return func(item *Item) {
		item.left = left
	}
}

func WithLeftId(origin *ID) ItemOption {
	return func(item *Item) {
		item.origin = origin
	}
}

func WithLeftLastID(left Block) ItemOption {
	return func(item *Item) {
		item.left = left
		if left != nil {
			item.origin = left.LastID()
		}
	}
}

func WithParent(parent any) ItemOption {
	return func(item *Item) {
		item.parent = parent
	}
}

func WithParentSub(parentSub string) ItemOption {
	return func(item *Item) {
		item.parentSub = parentSub
	}
}

func newItem(id *ID, opts ...ItemOption) *Item {
	item := &Item{
		BaseBlockType: &BaseBlockType{
			id:     id,
			length: 0,
		},
		info:   0,
		parent: nil,
	}
	for _, opt := range opts {
		opt(item)
	}
	return item
}

func (self *Item) Parent() any {
	return self.parent
}

func (self *Item) Deleted() bool {
	return self.IsDeleted()
}

func (self *Item) MergeWith(right Block) error {
	panic("not implemented") // TODO: Implement
}

func (self *Item) LastID() *ID {
	if self.length == 1 {
		return self.id
	}
	return newID(self.id.client, self.ClockEnd())
}

func (self *Item) Left() Block {
	return self.left
}

func (self *Item) Right() Block {
	return self.right
}

func (self *Item) Delete(tx *Transaction) error {
	if !self.Deleted() {
		p, ok := self.parent.(SharedType)
		if ok && self.IsCountable() && self.parentSub == "" {
			p.SetLength(p.Length() - self.length)
		}
		self.SetMarked()
		// TODO:   addToIdSet(transaction.deleteSet, this.id.client, this.id.clock, this.length)
		// TODO:   addChangedTypeToTransaction(transaction, parent, this.parentSub)
		return self.content.Delete(tx)
	}
	return nil
}

func (self *Item) Integrate(tx *Transaction, offset uint64) error {
	if offset > 0 {
		self.id.clock += offset
		b, err := tx.Store().GetItemCleanEnd(tx, newID(self.id.client, self.id.clock-1))
		if err != nil {
			return err
		}
		self.left = b
		self.origin = self.left.LastID()
		self.content = self.content.Splice(offset)
		self.length -= offset
	}
	if self.parent != nil {
		if self.left != nil {
			right := self.left.(*Item).right
			self.right = right
			self.left.(*Item).right = self
		} else {
			var r *Item
			if self.parentSub != "" {
				for r != nil && r.left != nil {
					r = r.left.(*Item)
				}
			} else {
				r, _ = self.Parent().Start().(*Item)
				self.Parent().SetStart(self)
			}
			self.right = r
		}

		if self.right.(*Item) != nil {
			self.right.(*Item).left = self
		} else if self.parentSub != "" {
			self.Parent().SetBlock(self.parentSub, self)
			if self.left != nil {
				if err := self.left.Delete(tx); err != nil {
					return err
				}
			}
		}
		p, ok := self.parent.(*BaseSharedType)
		if self.parentSub != "" && self.IsCountable() && !self.Deleted() && ok {
			p.length += self.length
		}

		if err := self.content.Integrate(tx, self); err != nil {
			return err
		}
	} else {
		return newGC(self.id, self.length).Integrate(tx, 0)
	}
	return nil
}

func (self *Item) GetMissing(tx *Transaction, store *BlockStore) (uint64, bool, error) {
	if self.origin != nil && (self.origin.clock >= store.GetClientClockEnd(self.origin.client) || store.skips.HasID(self.origin)) {
		return self.origin.client, true, nil
	}
	if self.rightOrigin != nil && (self.rightOrigin.clock >= store.GetClientClockEnd(self.rightOrigin.client) || store.skips.HasID(self.rightOrigin)) {
		return self.rightOrigin.client, true, nil
	}
	if p, ok := self.parent.(*ID); ok && (p.clock >= store.GetClientClockEnd(p.client) || store.skips.HasID(p)) {
		return p.client, true, nil
	}

	if self.origin != nil {
		left, err := store.GetItemCleanEnd(tx, self.origin)
		if err != nil {
			return 0, false, err
		}
		self.left = left
		self.origin = self.left.LastID()
	}
	if self.rightOrigin != nil {
		right, err := store.GetItemCleanStart(tx, self.rightOrigin)
		if err != nil {
			return 0, false, err
		}
		self.right = right
		self.rightOrigin = self.right.ID()
	}

	_, leftIsGc := self.left.(*GC)
	_, rightIsGc := self.right.(*GC)
	if leftIsGc || rightIsGc {
		self.parent = nil
	} else if self.parent != nil {
		if l, _ := self.left.(*Item); !reflect.ValueOf(l).IsNil() {
			self.parent = l.parent
			self.parentSub = l.parentSub
		}
		if r, _ := self.right.(*Item); !reflect.ValueOf(r).IsNil() {
			self.parent = r.parent
			self.parentSub = r.parentSub
		}
	} else if pid, _ := self.parent.(*ID); !reflect.ValueOf(pid).IsNil() {
		parentItem := store.GetBlock(pid)
		if parentItem != nil {
			if _, ok := parentItem.(*GC); ok {
				self.parent = nil
			} else if item, ok := parentItem.(*Item); ok {
				if ic, ok := item.content.(*ItemContentType); ok {
					self.parent = ic.sharedType
				}
			}
		}
	}
	return 0, false, nil
}

func (self *Item) Split(tx *Transaction, diff uint64) (Block, error) {
	rightItem := newItem(newID(self.id.client, self.id.clock+diff))
	rightItem.left = self
	rightItem.origin = newID(self.id.client, self.id.clock+diff-1)
	rightItem.right = self.right
	rightItem.rightOrigin = self.rightOrigin
	rightItem.parent = self.parent
	rightItem.parentSub = self.parentSub
	rightItem.content = self.content.Splice(diff)

	if self.IsDeleted() {
		rightItem.SetDeleted()
	}

	if self.IsKeep() {
		rightItem.SetKeep()
	}

	if self.redOne != nil {
		rightItem.redOne = newID(self.id.client, self.id.clock+diff)
	}

	if tx != nil {
		self.right = rightItem
		if right := rightItem.right.(*Item); right != nil {
			right.left = rightItem
		}
		tx.mergeBlocks = append(tx.mergeBlocks, rightItem)
		if rightItem.parentSub != "" && rightItem.right != nil {
			sharedType, has := rightItem.parent.(SharedType)
			if has && sharedType != nil {
				sharedType.SetBlock(rightItem.parentSub, rightItem)
			}
		}
	} else {
		rightItem.left = nil
		rightItem.right = nil
	}
	self.length = diff
	return rightItem, nil
}

func (self *Item) Write(encoder UpdateEncoder, offset uint64, offsetEnd byte) error {
	origin := self.origin
	if offset > 0 {
		origin = newID(self.id.client, self.id.clock+offset-1)
	}
	// const rightOrigin = this.rightOrigin
	rightOrigin := self.rightOrigin
	parentSub := self.parentSub
	info := self.content.GetRef() & 31
	if origin != nil {
		info = info & 128
	}
	if rightOrigin != nil {
		info = info & 64
	}
	if parentSub != "" {
		info = info & 32
	}

	if err := encoder.WriteInfo(byte(info)); err != nil {
		return err
	}

	if origin != nil {
		if err := encoder.WriteLeftID(origin); err != nil {
			return err
		}
	}
	if rightOrigin != nil {
		if err := encoder.WriteRightID(rightOrigin); err != nil {
			return err
		}
	}

	if origin == nil && rightOrigin == nil {
		switch parent := self.parent.(type) {
		case string:
			if err := encoder.WriteParentInfo(true); err != nil {
				return err
			}
			if err := encoder.WriteVarString(parent); err != nil {
				return err
			}
		case *ID:
			if err := encoder.WriteParentInfo(false); err != nil {
				return err
			}
			if err := encoder.WriteLeftID(parent); err != nil {
				return err
			}
		case SharedType:
			parentItem := self.Parent()
			if parentItem == nil {
				// FIXME: should we return error if the key for item is not found
				// Ideally it should never happen, but better safe then sorry right !
				ykey, _ := parent.Doc().GetItemKey(parent)
				if err := encoder.WriteParentInfo(true); err != nil {
					return err
				}
				if err := encoder.WriteVarString(ykey); err != nil {
					return err
				}
			} else {
				if err := encoder.WriteParentInfo(false); err != nil {
					return err
				}
				if err := encoder.WriteLeftID(parentItem.ID()); err != nil {
					return err
				}
			}
		default:
			return fmt.Errorf("unexpected case")
		}
	}

	if err := self.content.Write(encoder, offset, uint64(offsetEnd)); err != nil {
		return err
	}
	return nil
}

func decodeItem(id *ID, decoder UpdateDecoder, info Kind, doc *Doc) (*Item, error) {
	var err error
	var item *Item = newItem(id)
	cant_copy_parent_info := byte(info)&(HasOriginFlag|HasRightOriginFlag) == 0

	if byte(info)&HasOriginFlag != 0 {
		item.origin, err = decoder.ReadLeftID()
		if err != nil {
			return nil, fmt.Errorf("decodeItem: failed to read left id %w", err)
		}
	}
	if byte(info)&HasRightOriginFlag != 0 {
		item.rightOrigin, err = decoder.ReadRightID()
		if err != nil {
			return nil, fmt.Errorf("decodeItem: failed to read right id %w", err)
		}
	}
	if cant_copy_parent_info {
		parent_info, err := decoder.ReadParentInfo()
		if err != nil {
			return nil, fmt.Errorf("decodeItem: failed to read parent info %w", err)
		}
		if parent_info {
			// Parent string
			item.parent, err = decoder.ReadVarString()
			if err != nil {
				return nil, fmt.Errorf("decodeItem: failed to read parent string %w", err)
			}
			if doc != nil {
				item.parent, err = doc.get(item.parent.(string))
				if err != nil {
					return nil, fmt.Errorf("decodeItem: failed to get parent from doc %w", err)
				}
			}
		} else {
			item.parent, err = decoder.ReadLeftID()
			if err != nil {
				return nil, fmt.Errorf("decodeItem: failed to read left id parent %w", err)
			}
		}

		if (byte(info) & HasParentSubFlag) != 0 {
			item.parentSub, err = decoder.ReadVarString()
			if err != nil {
				return nil, fmt.Errorf("decodeItem: failed to read parent sub %w", err)
			}
		}
	}

	content, err := decodeItemContent(decoder, info)
	if err != nil {
		return nil, fmt.Errorf("decodeItem: failed to decode item content with info %d %w", info&31, err)
	}
	WithContent(content)(item)

	item.length = item.content.Length()
	return item, nil
}
