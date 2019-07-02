package devto

type Transaction struct {
	CheckNumber      int     `json:"check_number"`
	Category         string  `json:"category"`
	Amount           float64 `json:"amount"`
	RemainingBalance float64 `json:"remaining_balance"`
}

type TransactionCollection []Transaction
type TransactionCollectionMap map[string]TransactionCollection
type TransactionFilterFunc func(transaction Transaction) bool
type TransactionMapFunc func(transaction Transaction) Transaction

func (tc TransactionCollection) Filter(filterFunc TransactionFilterFunc) TransactionCollection {
	var c TransactionCollection

	for _, t := range tc {
		if filterFunc(t) {
			c = append(c, t)
		}
	}

	return c
}

func (tc TransactionCollection) Map(mapFunc TransactionMapFunc) TransactionCollection {
	var c TransactionCollection

	for _, t := range tc {
		c = append(c, mapFunc(t))
	}

	return c
}
