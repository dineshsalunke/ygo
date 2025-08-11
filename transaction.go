package ygo

type Transaction struct {
	doc            *Doc
	deleteSet      *IdSet
	cleanups       *IdSet
	insertSet      *IdSet
	beforeState    StateVector
	afterState     StateVector
	origin         any
	local          bool
	mergeBlocks    BlockList
	pendingStructs *BlockSet
	done           bool
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
		mergeBlocks:    make(BlockList, 0),
		pendingStructs: newBlockSet(0),
		local:          false,
		origin:         nil,
		done:           false,
	}
	for _, opt := range opts {
		opt(tx)
	}
	return tx
}

func (tx *Transaction) Store() *BlockStore {
	return tx.doc.store
}

func (tx *Transaction) commitTransaction() error {
	panic("not implemented")
}

type TransactionHandler func(tx *Transaction) (any, error)

type TransactOption func(tx *Transaction)

func cleanUpTransactions(transactions []*Transaction, index int) error {
	if index < len(transactions) {
		transaction := transactions[index]
		panic("not implemented")
		if len(transactions) < index+1 {
			transaction.doc.transactionCleanups = make([]*Transaction, 0)
		} else {
			cleanUpTransactions(transactions, index+1)
		}
	}
	return nil
}

func Transact(doc *Doc, handler TransactionHandler, origin any, local bool) (any, error) {
	initialCall := false
	if doc.activeTransaction == nil {
		initialCall = true
		doc.activeTransaction = newTransaction(doc)
		doc.activeTransaction.local = local
		doc.activeTransaction.origin = origin
		doc.transactionCleanups = append(doc.transactionCleanups, doc.activeTransaction)
	}

	result, err := handler(doc.activeTransaction)

	if initialCall {
		finishCleanup := doc.activeTransaction == doc.transactionCleanups[0]
		doc.activeTransaction = nil
		if finishCleanup {
			err = cleanUpTransactions(doc.transactionCleanups, 0)
		}
	}

	return result, err
}
