package ygo

type IdRange struct {
	clock  uint64
	length uint64
}

// helper method for more symantic name for `Clock` incase of `IdRange`
func (ir *IdRange) start() uint64 {
	return ir.clock
}

// helper method for more semantic name for `Clock + Length` incase of `IdRange`
func (ir *IdRange) end() uint64 {
	return ir.clock + ir.length
}

func newIdRange(clock, length uint64) *IdRange {
	return &IdRange{
		clock:  clock,
		length: length,
	}
}
