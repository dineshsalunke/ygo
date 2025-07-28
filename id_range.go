package ygo

import "fmt"

type IdRange struct {
	clock  uint64
	length uint64
}

func (ir *IdRange) String() string {
	return fmt.Sprintf("{clock:%d,length:%d}", ir.clock, ir.length)
}

func (ir *IdRange) GoString() string {
	return fmt.Sprintf("{clock:%d,length:%d}", ir.clock, ir.length)
}

func (ir *IdRange) start() uint64 {
	return ir.clock
}

func (ir *IdRange) end() uint64 {
	return ir.clock + ir.length
}
