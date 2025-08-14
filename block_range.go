package ygo

type BlockRange struct {
	i    int
	refs BlockList
}

func (self *BlockRange) FirstBlock() Block {
	return self.refs[0]
}

func (self *BlockRange) LastBlock() Block {
	return self.refs[0]
}
