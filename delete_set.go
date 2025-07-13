package ygo

type DeleteItem struct {
	clock  uint64
	length uint64
}

func newDeleteItem(clock, length uint64) *DeleteItem {
	return &DeleteItem{clock: clock, length: length}
}

type DeleteSet struct {
	clients map[uint64][]*DeleteItem
}

func newDeleteSet() *DeleteSet {
	return &DeleteSet{
		clients: make(map[uint64][]*DeleteItem),
	}
}

