package ygo

import "fmt"

type BlockRange struct {
	i    int
	refs []Block
}

func (ss BlockRange) GoString() string {
	return fmt.Sprintf("StructRange {\n\t refs: %#v \n}", ss.refs)
}

func (ss BlockRange) String() string {
	return fmt.Sprintf("StructRange {\n\t refs: %#v \n}", ss.refs)
}
