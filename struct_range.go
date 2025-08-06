package ygo

import "fmt"

type StructRange struct {
	i    int
	refs []Block
}

func (ss StructRange) GoString() string {
	return fmt.Sprintf("StructRange {\n\t refs: %#v \n}", ss.refs)
}

func (ss StructRange) String() string {
	return fmt.Sprintf("StructRange {\n\t refs: %#v \n}", ss.refs)
}
