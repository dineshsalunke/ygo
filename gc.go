package ygo

type GC struct{ *block }

func newGC(id *ID, length uint64) *GC {
	return &GC{
		block: newBlock(id, length),
	}
}

func (self *GC) Parent() SharedType {
	return nil
}

func (self *GC) Kind() Kind {
	return BlockKindGC
}

func (self *GC) Length() uint64 {
	return self.length
}

func (self *GC) ID() *ID {
	return self.id
}

func (self *GC) Write(encoder UpdateEncoder, offset uint64, offsetLength uint64) error {
	if err := encoder.WriteInfo(GCTypeRef); err != nil {
		return err
	}
	if err := encoder.WriteLength(self.length - offset - offsetLength); err != nil {
		return err
	}
	return nil
}
