package memory

import (
	"fmt"
	"strings"

	"github.com/Dak425/dev-to-challenge-4-go"
)

type Checkbook struct {
	initialBalance float64
	currentBalance float64
	transactions   devto.TransactionCollection
	categoryMap    devto.TransactionCollectionMap
}

func NewInMemoryCheckBook(raw string) *Checkbook {
	cb := &Checkbook{
		transactions: devto.TransactionCollection{},
		categoryMap:  make(devto.TransactionCollectionMap),
	}

	cb.parseRaw(raw)

	return cb
}

func (cb *Checkbook) InitialBalance() float64 {
	return cb.initialBalance
}

func (cb *Checkbook) CurrentBalance() float64 {
	return cb.currentBalance
}

func (cb *Checkbook) Transactions() devto.TransactionCollection {
	return cb.transactions
}

func (cb *Checkbook) TransactionsForCategory(category string) devto.TransactionCollection {
	return cb.categoryMap[category]
}

func (cb *Checkbook) AddTransaction(transaction devto.Transaction) {
	cb.currentBalance -= transaction.Amount
	transaction.RemainingBalance = cb.currentBalance
	cb.transactions = append(cb.transactions, transaction)
	cb.addToCategoryMap(transaction)
}

func (cb *Checkbook) FullReport() string {
	var report strings.Builder
	var cost float64

	// initial balance
	report.WriteString(fmt.Sprintf("Starting Balance: %.2f\n", cb.initialBalance))

	// transaction details
	for i, t := range cb.transactions {
		cost += t.Amount
		report.WriteString(fmt.Sprintf(
			"[%d] -> Check Number: %d, Category: %s, Amount: %.2f, Remaining Balance: %.2f\n",
			i+1,
			t.CheckNumber,
			t.Category,
			t.Amount,
			t.RemainingBalance,
		))
	}

	// total costs
	report.WriteString(fmt.Sprintf(
		"Total Costs: %.2f\n",
		cost,
	))

	// average costs
	report.WriteString(fmt.Sprintf(
		"Average Cost: %.2f",
		cost/float64(len(cb.transactions)),
	))

	return report.String()
}

func (cb *Checkbook) addToCategoryMap(transaction devto.Transaction) {
	if _, ok := cb.categoryMap[transaction.Category]; !ok {
		cb.categoryMap[transaction.Category] = devto.TransactionCollection{}
	}

	cb.categoryMap[transaction.Category] = append(cb.categoryMap[transaction.Category], transaction)
}

func (cb *Checkbook) parseRaw(raw string) {
	p := newParser()
	lines := p.getLines(raw)

	cb.initialBalance, lines = p.initialBalanceFromLines(lines)
	cb.currentBalance = cb.initialBalance

	for _, l := range lines {
		cb.AddTransaction(p.transactionFromLine(l))
	}
}
