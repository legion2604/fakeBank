package service

import (
	"fakeBank/internal/middleware"
	"fakeBank/internal/models"
	"fakeBank/internal/repository"

	"github.com/gin-gonic/gin"
)

type AccountService interface {
	GetAccounts(cxt *gin.Context) ([]models.Account, error)
	GetAccountById(accId int) (models.Account, error)
	CreateAccount(cxt *gin.Context, user models.CreateAccountReq) (models.Account, error)
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

func (s *accountService) GetAccountById(accId int) (models.Account, error) {
	res, err := s.accountRepo.GetAccountByID(accId)
	if err != nil {
		return models.Account{}, err
	}
	return res, nil
}

func (s *accountService) CreateAccount(cxt *gin.Context, user models.CreateAccountReq) (models.Account, error) {
	userId := middleware.GetUserIDFromContext(cxt)
	res, err := s.accountRepo.CreateAccount(userId, user.AccountName, user.AccountType)
	if err != nil {
		return models.Account{}, err
	}
	return res, nil
}
