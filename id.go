package ygo

import "fmt"

type ID struct {
	client uint64
	clock  uint64
}

func newID(client uint64, clock uint64) *ID {
	return &ID{
		client: client,
		clock:  clock,
	}
}

func NewID(client uint64, clock uint64) *ID {
	return newID(client, clock)
}

// Implement the fmt Stringer interface
func (id *ID) GoString() string {
	return fmt.Sprintf("ID{client:%d,clock:%d}", id.client, id.clock)
}
