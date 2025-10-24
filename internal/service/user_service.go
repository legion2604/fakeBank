package service

import (
	"errors"
	"fakeBank/internal/middleware"
	"fakeBank/internal/models"
	"fakeBank/internal/repository"
	"fakeBank/pkg/utils"

	"github.com/gin-gonic/gin"
)

type UserService interface {
	GetProfile(cxt *gin.Context) (models.GetUser, error)
	ChangePassword(currentPassword, newPassword string, cxt *gin.Context) (bool, error)
	DeleteUser(cxt *gin.Context) (bool, error)
	UpdateUser(json models.UpdateUserJson, cxt *gin.Context) (models.GetUser, bool, error)
}
type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (u userService) ChangePassword(currentPassword, newPassword string, cxt *gin.Context) (bool, error) {
	userId := middleware.GetUserIDFromContext(cxt)
	userPass, err := u.userRepo.GetUserPassword(userId)
	if err != nil {
		return false, err
	}
	isPassCorrect := utils.ComparePasswords(userPass, currentPassword)
	if isPassCorrect {
		hashedNewPass := utils.GenerateHashFromPassword(newPassword)
		res, err := u.userRepo.ChangePassword(userId, hashedNewPass)
		if err != nil {
			return false, err
		}
		return res, nil
	}
	return false, errors.New("wrong password")

}

func (u userService) GetProfile(cxt *gin.Context) (models.GetUser, error) {
	userId := middleware.GetUserIDFromContext(cxt)
	res, err := u.userRepo.GetProfile(userId)
	if err != nil {
		return models.GetUser{}, err
	}
	return res, nil
}

func (u userService) DeleteUser(cxt *gin.Context) (bool, error) {
	userId := middleware.GetUserIDFromContext(cxt)
	res, err := u.userRepo.DeleteUser(userId)
	if err != nil {
		return false, err
	}
	return res, nil
}

func (u userService) UpdateUser(json models.UpdateUserJson, cxt *gin.Context) (models.GetUser, bool, error) {
	userId := middleware.GetUserIDFromContext(cxt)
	res, err := u.userRepo.UpdateUser(userId, json)
	if err != nil {
		return models.GetUser{}, false, err
	}
	return res, true, nil
}
