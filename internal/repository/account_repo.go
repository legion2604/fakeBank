package repository

import (
	"database/sql"
	"fakeBank/internal/models"
)

type AccountRepository interface {
	GetAccounts(userID int) ([]models.Account, error)
	GetAccountByID(userID int) (models.Account, error)
	CreateAccount(accId int, accountType, accountName string) (models.Account, error)
}

type accountRepo struct {
	db *sql.DB
}

func NewAccountRepository(db *sql.DB) AccountRepository {
	return &accountRepo{db: db}
}

func (r *accountRepo) GetAccounts(userID int) ([]models.Account, error) {
	row, err := r.db.Query("SELECT id, account_number, account_type, balance, currency, created_at FROM accounts WHERE account_user_id = $1", userID)
	if err != nil {
		return []models.Account{}, err
	}
	defer row.Close()
	var account []models.Account
	for row.Next() {
		var acc models.Account
		err := row.Scan(&acc.ID, &acc.AccountNumber, &acc.AccountType, &acc.Balance, &acc.Currency, &acc.CreatedAt)
		if err != nil {
			return []models.Account{}, err
		}
		account = append(account, acc)
	}
	return account, nil
}

func (r *accountRepo) GetAccountByID(accId int) (models.Account, error) {
	var res models.Account
	err := r.db.QueryRow("SELECT id, account_number, account_type, balance, currency, created_at FROM accounts WHERE id = $1", accId).Scan(&res.ID, &res.AccountNumber, &res.AccountType, &res.Balance, &res.Currency, &res.CreatedAt)
	if err != nil {
		return models.Account{}, err
	}
	return res, nil
}

func (r *accountRepo) CreateAccount(userId int, accountType, accountName string) (models.Account, error) {
	var res models.Account
	err := r.db.QueryRow("INSERT into accounts (account_type, account_name, account_user_id) values ($1, $2,$3) RETURNING id,account_number,balance,currency,created_at", accountType, accountName, userId).Scan(&res.ID, &res.AccountNumber, &res.Balance, &res.Currency, &res.CreatedAt)
	res.AccountType = accountType

	if err != nil {
		return models.Account{}, err
	}
	return res, nil
}
