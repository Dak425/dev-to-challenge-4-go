package devto

type CheckBook interface {
	InitialBalance() float64
	CurrentBalance() float64
	Transactions() TransactionCollection
	TransactionsForCategory(category string) TransactionCollection
	AddTransaction(transaction Transaction)
	FullReport() string
}
