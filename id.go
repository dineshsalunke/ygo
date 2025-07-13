package ygo

type ID struct {
	client uint64
	clock  uint64
}

func (id *ID) Clock() uint64 {
	return id.clock
}

func (id *ID) Client() uint64 {
	return id.client
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
