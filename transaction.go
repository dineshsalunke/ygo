package ygo

type Transaction struct {
	doc          *Doc
	origin       any
	local        bool
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
		doc:          doc,
		local:        false,
		origin:       nil,
	}
	for _, opt := range opts {
		opt(tx)
	}
	return tx
}
