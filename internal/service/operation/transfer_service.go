package operation

import (
	"fakeBank/internal/models"
	"fakeBank/internal/repository/operation"
	"fakeBank/pkg/errors"
	"strconv"
)

type TransferService interface {
	CreateTransactionTransfer(json models.TransactionTransferReq) (models.TransactionTransfer, error)
}
type transferService struct {
	repo operation.TransferRepo
}

func NewTransferService(repo operation.TransferRepo) TransferService {
	return &transferService{repo: repo}
}

func (s *transferService) CreateTransactionTransfer(json models.TransactionTransferReq) (models.TransactionTransfer, error) {
	fromAccountId, _ := strconv.Atoi(json.FromAccount)

	toAccountId, err := s.repo.GetAccountIdByEmail(json.ToEmail)

	var res models.TransactionTransfer
	if err != nil {
		return models.TransactionTransfer{}, errors.ErrUserRecipientNotFound
	}
	getBalanceById, err := s.repo.GetBalanceById(toAccountId)
	if err != nil {
		return models.TransactionTransfer{}, err
	}
	if getBalanceById < json.Amount {
		res, err = s.repo.CreateTransactionTransfer(fromAccountId, fromAccountId, toAccountId, json.Amount, "transfer", "failed", json.Description)
		if err != nil {
			return models.TransactionTransfer{}, err
		}
		return models.TransactionTransfer{}, errors.ErrInsufficientFunds
	}

	res, err = s.repo.CreateTransactionTransfer(fromAccountId, fromAccountId, toAccountId, json.Amount, "transfer", "completed", json.Description)
	if err != nil {
		return models.TransactionTransfer{}, err
	}
	_, err = s.repo.CreateTransactionTransfer(toAccountId, fromAccountId, toAccountId, json.Amount, "deposit", "completed", json.Description)

	if err != nil {
		return models.TransactionTransfer{}, err
	}
	err = s.repo.MoneyTransfer(fromAccountId, toAccountId, json.Amount)
	if err != nil {
		return models.TransactionTransfer{}, err
	}
	return res, nil
}
