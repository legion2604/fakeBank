package operation

import (
	"database/sql"
	"fakeBank/internal/models"
	"fmt"
)

type TransferRepo interface {
	CreateTransactionTransfer(accountId, fromAccountId, toAccountId int, amount float64, TType, status, description string) (models.TransactionTransfer, error)
	GetAccountIdByEmail(email string) (int, error)
	GetBalanceById(accId int) (float64, error)
	MoneyTransfer(fromAccountId, toAccountId int, amount float64) error
}

type transferRepo struct {
	db *sql.DB
}

func NewTransferRepository(db *sql.DB) TransferRepo {
	return &transferRepo{db: db}
}

func (r *transferRepo) CreateTransactionTransfer(accountId, fromAccountId, toAccountId int, amount float64, TType, status, description string) (models.TransactionTransfer, error) {
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

func (r *transferRepo) GetAccountIdByEmail(email string) (int, error) {
	var accId int
	err := r.db.QueryRow("SELECT a.id FROM accounts a JOIN users u ON a.account_user_id = u.id WHERE u.email = $1 LIMIT 1;", email).Scan(&accId)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return accId, nil
}

func (r *transferRepo) GetBalanceById(accId int) (float64, error) {
	var balance float64
	err := r.db.QueryRow("SELECT balance FROM accounts WHERE id = $1", accId).Scan(&balance)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return balance, nil
}

func (r *transferRepo) MoneyTransfer(fromAccountId, toAccountId int, amount float64) error {
	_, err := r.db.Exec("UPDATE accounts SET balance= balance - $1 WHERE id=$2", amount, fromAccountId)
	if err != nil {
		return err
	}
	_, err = r.db.Exec("UPDATE accounts SET balance= balance + $1 WHERE id=$2", amount, toAccountId)

	if err != nil {
		return err
	}

	return nil
}
