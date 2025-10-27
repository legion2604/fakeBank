package models

type Account struct {
	ID            string  `json:"id"`
	AccountNumber string  `json:"accountNumber"`
	AccountType   string  `json:"accountType"`
	Balance       float64 `json:"balance"`
	Currency      string  `json:"currency"`
	CreatedAt     string  `json:"createdAt"`
}

type CreateAccountReq struct {
	AccountType string `json:"accountType"`
	AccountName string `json:"accountName"`
}
