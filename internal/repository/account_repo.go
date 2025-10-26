package repository

import (
	"database/sql"
	"fakeBank/internal/models"
)

type AccountRepository interface {
	GetAccounts(userID int) ([]models.Account, error)
	GetAccountByID(userID int) (models.Account, error)
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

func (r *accountRepo) GetAccountByID(userID int) (models.Account, error) {
	var res models.Account
	err := r.db.QueryRow("SELECT id, account_number, account_type, balance, currency, created_at FROM accounts WHERE account_user_id = $1", userID).Scan(&res.ID, &res.AccountNumber, &res.AccountType, &res.Balance, &res.Currency, &res.CreatedAt)
	if err != nil {
		return models.Account{}, err
	}
	return res, nil
}
