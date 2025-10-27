package repository

import (
	"database/sql"
	"fakeBank/internal/models"
	"fmt"
	"strconv"
)

type TransactionRepository interface {
	GetTransactions(accountId, limitStr, offsetStr int) ([]interface{}, error)
	GetTransactionById(AccId int) (interface{}, error)
	CreateTransactionTransfer(accountId, fromAccountId, toAccountId int, amount float64, TType, status, description string) (models.TransactionTransfer, error)
	GetAccountIdByEmail(email string) (int, error)
	GetBalanceById(accId int) (float64, error)
}
type transactionRepo struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) TransactionRepository {
	return &transactionRepo{db: db}
}

func (r *transactionRepo) GetTransactions(accountId, limitStr, offsetStr int) ([]interface{}, error) {
	rows, err := r.db.Query("SELECT id, account_id, type, amount, currency, description, from_account, to_account, status, created_at FROM transactions WHERE account_id = $1 ORDER BY created_at DESC LIMIT $2 OFFSET $3", accountId, limitStr, offsetStr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []interface{}
	for rows.Next() {
		var (
			id, accId, TransactionType, currency, description, fromAccount, toAccount, status, createdAt string
			amount                                                                                       float64
		)
		err := rows.Scan(&id, &accId, &TransactionType, &amount, &currency, &description, &fromAccount, &toAccount, &status, &createdAt)
		if err != nil {
			return nil, err
		}
		switch TransactionType {
		case "transfer":
			transactions = append(transactions, models.TransactionTransfer{
				Id:          id,
				AccountId:   accId,
				Type:        TransactionType,
				Amount:      amount,
				Currency:    currency,
				Description: description,
				FromAccount: fromAccount,
				ToAccount:   toAccount,
				Status:      status,
				CreatedAt:   createdAt,
			})
		case "deposit":
			transactions = append(transactions, models.TransactionDeposit{
				Id:          id,
				AccountId:   accId,
				Type:        TransactionType,
				Amount:      amount,
				Currency:    currency,
				Description: description,
				Status:      status,
				CreatedAt:   createdAt,
			})
		case "withdrawal":
			transactions = append(transactions, models.TransactionWithdrawal{
				Id:          id,
				AccountId:   accId,
				Type:        TransactionType,
				Amount:      amount,
				Currency:    currency,
				Description: description,
				Status:      status,
				CreatedAt:   createdAt,
			})
		}
	}
	return transactions, nil
}

func (r *transactionRepo) GetTransactionById(TransId int) (interface{}, error) {
	var (
		accId           int
		TransactionType string
		currency        string
		description     string
		fromAccount     string
		toAccount       string
		status          string
		createdAt       string
		amount          float64
	)
	fmt.Println(TransId)
	err := r.db.QueryRow("SELECT account_id, type, amount, currency, description, from_account, to_account, status, created_at FROM transactions WHERE id = $1", TransId).Scan(&accId, &TransactionType, &amount, &currency, &description, &fromAccount, &toAccount, &status, &createdAt)

	if err != nil {
		return nil, err
	}

	switch TransactionType {
	case "transfer":
		return models.TransactionTransfer{
			Id:          strconv.Itoa(TransId),
			AccountId:   strconv.Itoa(accId),
			Type:        TransactionType,
			Amount:      amount,
			Currency:    currency,
			Description: description,
			FromAccount: fromAccount,
			ToAccount:   toAccount,
			Status:      status,
			CreatedAt:   createdAt,
		}, nil
	case "deposit":
		return models.TransactionDeposit{
			Id:          strconv.Itoa(TransId),
			AccountId:   strconv.Itoa(accId),
			Type:        TransactionType,
			Amount:      amount,
			Currency:    currency,
			Description: description,
			Status:      status,
			CreatedAt:   createdAt,
		}, nil
	case "withdrawal":
		return models.TransactionWithdrawal{
			Id:          strconv.Itoa(TransId),
			AccountId:   strconv.Itoa(accId),
			Type:        TransactionType,
			Amount:      amount,
			Currency:    currency,
			Description: description,
			Status:      status,
			CreatedAt:   createdAt,
		}, nil
	}

	return nil, nil
}

func (r *transactionRepo) CreateTransactionTransfer(accountId, fromAccountId, toAccountId int, amount float64, TType, status, description string) (models.TransactionTransfer, error) {
	var res models.TransactionTransfer
	err := r.db.QueryRow("INSERT INTO transactions (account_id, type, amount, description, from_account, to_account, status) VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING id,account_id,from_account,to_account,created_at", accountId, TType, amount, description, fromAccountId, toAccountId, status).Scan(&res.Id, &res.AccountId, &res.FromAccount, &res.ToAccount, &res.CreatedAt)
	if err != nil {
		return models.TransactionTransfer{}, err
	}
	res.Amount = amount
	res.Currency = "USD"
	res.Type = TType
	res.Status = status
	res.Description = description

	return res, nil
}

func (r *transactionRepo) GetAccountIdByEmail(email string) (int, error) {
	var accId int
	err := r.db.QueryRow("SELECT a.id FROM accounts a JOIN users u ON a.account_user_id = u.id WHERE u.email = $1 LIMIT 1;", email).Scan(&accId)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return accId, nil
}

func (r *transactionRepo) GetBalanceById(accId int) (float64, error) {
	var balance float64
	err := r.db.QueryRow("SELECT balance FROM accounts WHERE id = $1", accId).Scan(&balance)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return balance, nil
}
