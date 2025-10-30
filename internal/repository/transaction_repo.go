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
