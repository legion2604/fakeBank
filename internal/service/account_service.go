package service

import (
	"fakeBank/internal/middleware"
	"fakeBank/internal/models"
	"fakeBank/internal/repository"

	"github.com/gin-gonic/gin"
)

type AccountService interface {
	GetAccounts(cxt *gin.Context) ([]models.Account, error)
	GetAccountById(userId int) (models.Account, error)
}

type accountService struct {
	accountRepo repository.AccountRepository
}

func NewAccountService(accountRepo repository.AccountRepository) AccountService {
	return &accountService{accountRepo: accountRepo}
}

func (s *accountService) GetAccounts(cxt *gin.Context) ([]models.Account, error) {
	userId := middleware.GetUserIDFromContext(cxt)
	acc, err := s.accountRepo.GetAccounts(userId)
	if err != nil {
		return []models.Account{}, err
	}
	return acc, nil
}

func (s *accountService) GetAccountById(userId int) (models.Account, error) {
	res, err := s.accountRepo.GetAccountByID(userId)
	if err != nil {
		return models.Account{}, err
	}
	return res, nil
}
