package ygo

type GC struct {
	*BaseBlockType
}

func newGC(id *ID, length uint64) *GC {
	return &GC{
		BaseBlockType: &BaseBlockType{
			id:     id,
			length: length,
		},
	}
}

func (self *GC) Splice(diff uint64) (Block, error) {
	gc := newGC(newID(self.id.client, self.id.clock+diff), self.length-diff)
	gc.length = diff
	return gc, nil
}
