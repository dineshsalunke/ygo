package ygo

type Transaction struct {
	doc            *Doc
	deleteSet      *IdSet
	insertSet      *IdSet
	cleanups       *IdSet
	beforeState    StateVector
	afterState     StateVector
	origin         any
	local          bool
	mergeStructs   []SharedType
	pendingStructs *StructSet
}

type TxOption func(tx *Transaction)

func TransactionWithLocal(local bool) TxOption {
	return func(tx *Transaction) {
		tx.local = local
	}
}

func TransactionWithOrigin(origin any) TxOption {
	return func(tx *Transaction) {
		tx.origin = origin
	}
}

func newTransaction(doc *Doc, opts ...TxOption) *Transaction {
	tx := &Transaction{
		doc:            doc,
		deleteSet:      newIdSet(),
		insertSet:      newIdSet(),
		cleanups:       newIdSet(),
		beforeState:    make(StateVector),
		afterState:     make(StateVector),
		mergeStructs:   make([]SharedType, 0),
		pendingStructs: newStructSet(0),
		local:          false,
		origin:         nil,
	}
	for _, opt := range opts {
		opt(tx)
	}
	return tx
}

func (tx *Transaction) commitTransaction() error {
	panic("not implemented")
}
