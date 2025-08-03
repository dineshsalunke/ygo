package ygo

import (
	"fmt"
	"strings"
)
type IdSet struct {
	clients map[uint64]*IdRanges
}

func newIdSet() *IdSet {
	return &IdSet{
		clients: make(map[uint64]*IdRanges),
	}
}

func (ss *IdSet) GoString() string {
	builder := strings.Builder{}
	builder.WriteString("IdSet {\r\tclients: {")
	for client, ranges := range ss.clients {
		builder.WriteString(fmt.Sprintf("%d:%#v\n", client, ranges))
	}
	builder.WriteString("}\r}")
	return builder.String()
}

