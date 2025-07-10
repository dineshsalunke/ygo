package core

type ID struct {
	client uint64
	clock  uint64
}

func newID(client, clock uint64) *ID {
	return &ID{
		client: client,
		clock:  clock,
	}
}

func NewID(client, clock uint64) *ID {
	return newID(client, clock)
}
