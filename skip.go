package ygo

type Skip struct {
	*BaseBlockType
}

func newSkip(id *ID, length uint64) *Skip {
	return &Skip{
		BaseBlockType: &BaseBlockType{
			id:     id,
			length: length,
		},
	}
}

func (self *Skip) Splice(diff uint64) (Block, error) {
	gc := newSkip(newID(self.id.client, self.id.clock+diff), self.length-diff)
	gc.length = diff
	return gc, nil
}
