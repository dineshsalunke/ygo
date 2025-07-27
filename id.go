package ygo

import "fmt"

type ClientID = uint64

type ID struct {
	client ClientID
	clock  uint64
}

func newID(client ClientID, clock uint64) *ID {
	return &ID{
		client: client,
		clock:  clock,
	}
}

func NewID(client ClientID, clock uint64) *ID {
	return newID(client, clock)
}

// Implement the fmt Stringer interface
func (id *ID) String() string {
	return fmt.Sprintf("%d:%d", id.client, id.client)
}
