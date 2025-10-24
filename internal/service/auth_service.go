package service

import (
	"fakeBank/internal/middleware"
	"fakeBank/internal/models"
	"fakeBank/internal/repository"
	"fakeBank/pkg/utils"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type AuthService interface {
	Login(email, password string, cxt *gin.Context) (bool, models.GetUser, error)
	Signup(json models.UserJson, cxt *gin.Context) (bool, int, time.Time, error)
	Me(cxt *gin.Context) (models.GetUser, error)
	Logout(cxt *gin.Context) bool
}

type authService struct {
	authRepository repository.AuthRepository
}

func NewAuthService(authRepo repository.AuthRepository) AuthService {
	return &authService{authRepository: authRepo}
}

func (u authService) Login(email, password string, cxt *gin.Context) (bool, models.GetUser, error) {
	res, hashedPassword, err := u.authRepository.Login(email)
	if err != nil {
		return false, models.GetUser{}, err
	}
	result := utils.ComparePasswords(hashedPassword, password)
	accessToken, _ := utils.GenerateJWT(res.Id)
	cxt.SetCookie("access_token", accessToken, int((time.Hour * 24).Seconds()), "/", "", false, true) // 24 hours
	if result {
		return true, res, nil
	} else {
		return false, models.GetUser{}, nil
	}
}

func (u authService) Signup(json models.UserJson, cxt *gin.Context) (bool, int, time.Time, error) {
	hashedPassword := utils.GenerateHashFromPassword(json.Password)
	json.Password = hashedPassword
	res, userId, createdAt, err := u.authRepository.Signup(json.FirstName, json.LastName, json.Email, json.Password, json.Phone)
	if err != nil {
		return res, 0, time.Time{}, err
	}
	accessToken, _ := utils.GenerateJWT(userId)
	cxt.SetCookie("access_token", accessToken, int((time.Hour * 24).Seconds()), "/", "", false, true) // 24 hours
	return res, userId, createdAt, nil
}

func (u authService) Me(cxt *gin.Context) (models.GetUser, error) {
	userId := middleware.GetUserIDFromContext(cxt)
	fmt.Println(userId)
	res, err := u.authRepository.Me(userId)
	if err != nil {
		return models.GetUser{}, err
	}
	return res, nil
}

func (u authService) Logout(cxt *gin.Context) bool {
	cxt.SetCookie("access_token", "", -1, "/", "", false, true)
	return true
}
