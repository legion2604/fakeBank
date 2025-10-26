package models

type TransactionTransfer struct {
	Id          string  `json:"id"`
	AccountId   string  `json:"accountId"`
	Type        string  `json:"type"`
	Amount      float64 `json:"amount"`
	Currency    string  `json:"currency"`
	Description string  `json:"description"`
	FromAccount string  `json:"fromAccount"`
	ToAccount   string  `json:"toAccount"`
	Status      string  `json:"status"`
	CreatedAt   string  `json:"createdAt"`
}

type TransactionDeposit struct {
	Id          string  `json:"id"`
	AccountId   string  `json:"accountId"`
	Type        string  `json:"type"`
	Amount      float64 `json:"amount"`
	Currency    string  `json:"currency"`
	Description string  `json:"description"`
	Status      string  `json:"status"`
	CreatedAt   string  `json:"createdAt"`
}

type TransactionWithdrawal struct {
	Id          string  `json:"id"`
	AccountId   string  `json:"accountId"`
	Type        string  `json:"type"`
	Amount      float64 `json:"amount"`
	Currency    string  `json:"currency"`
	Description string  `json:"description"`
	Status      string  `json:"status"`
	CreatedAt   string  `json:"createdAt"`
}
