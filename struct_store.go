package ygo

	"slices"
type PendingStructs struct {
	missingState StateVector
	update       []byte
}

type StructStore struct {
	clients            map[uint64][]Block // TODO: extend this into custom type and slice of block so we could add binary search funcs
	skips              *IdSet
	pendingStructs     *PendingStructs
	pendingIdSetUpdate []byte
}

func newStructStore() *StructStore {
	return &StructStore{
		clients:            make(map[uint64][]Block),
		skips:              newIdSet(),
		pendingStructs:     nil,
		pendingIdSetUpdate: nil,
	}
}

func (ss *StructStore) IntegrateStructs(tx *Transaction) (*PendingStructs, error) {
	panic("not implemented")
func (ss *StructStore) GetStateVector() StateVector {
	cl := len(ss.clients)
	sl := len(ss.skips.clients)
	sm := make(StateVector, cl+sl)
	for client, blocks := range ss.clients {
		block := blocks[len(blocks)-1]
		sm[client] = block.ClockEnd()
	}

	for client, block := range ss.skips.clients {
		r := block.getIdRanges()
		if len(r) > 0 {
			sm[client] = r[0].clock
		}
	}
	return sm
}

func (ss *StructStore) ApplyIdSet(decoder UpdateDecoder, tx *Transaction) ([]byte, error) {
	panic("not implemented")
}

// get client clock length
func (ss *StructStore) GetClientClockEnd(client uint64) uint64 {
	blocks, has := ss.clients[client]
	if !has || len(blocks) == 0 {
		return 0
	}
	lastBlock := blocks[len(blocks)-1]
	return lastBlock.ClockEnd() + 1
}

