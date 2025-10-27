package service

import (
	"errors"
	"fakeBank/internal/models"
	"fakeBank/internal/repository"
	"strconv"
)

type TransactionService interface {
	GetTransactions(accountId, limitStr, offsetStr int) ([]interface{}, error)
	GetTransactionById(transactionId int) (interface{}, error)
	CreateTransactionTransfer(json models.TransactionTransferReq) (models.TransactionTransfer, error)
}

type transactionService struct {
	repo repository.TransactionRepository
}

func NewTransactionService(repo repository.TransactionRepository) TransactionService {
	return &transactionService{repo: repo}
}

func (s *transactionService) GetTransactions(accountId, limitStr, offsetStr int) ([]interface{}, error) {
	res, err := s.repo.GetTransactions(accountId, limitStr, offsetStr)
	if err != nil {
		return nil, err
	}
	return res, err
}

func (s *transactionService) GetTransactionById(transactionId int) (interface{}, error) {
	res, err := s.repo.GetTransactionById(transactionId)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *transactionService) CreateTransactionTransfer(json models.TransactionTransferReq) (models.TransactionTransfer, error) {
	aacId, _ := strconv.Atoi(json.FromAccount)
	getAccountById, err := s.repo.GetAccountIdByEmail(json.ToEmail)
	var res models.TransactionTransfer
	if err != nil {
		return models.TransactionTransfer{}, errors.New("Recipient not found")
	}
	getBalanceById, err := s.repo.GetBalanceById(getAccountById)
	if err != nil {
		return models.TransactionTransfer{}, err
	}
	if getBalanceById < json.Amount {
		res, err = s.repo.CreateTransactionTransfer(aacId, aacId, getAccountById, json.Amount, "transfer", "failed", json.Description)
		if err != nil {
			return models.TransactionTransfer{}, err
		}
		return models.TransactionTransfer{}, errors.New("Insufficient funds")
	}

	res, err = s.repo.CreateTransactionTransfer(aacId, aacId, getAccountById, json.Amount, "transfer", "completed", json.Description)
	if err != nil {
		return models.TransactionTransfer{}, err
	}
	return res, nil
}
