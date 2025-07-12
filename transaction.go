package ygo

type Transaction struct{}

type TransactionHandler func() (*Transaction, error)
