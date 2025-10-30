package service

import (
	"fakeBank/internal/repository"
)

type TransactionService interface {
	GetTransactions(accountId, limitStr, offsetStr int) ([]interface{}, error)
	GetTransactionById(transactionId int) (interface{}, error)
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
