package controller

import (
	"fakeBank/internal/models"
	"fakeBank/internal/service"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	GetProfile(cxt *gin.Context)
	ChangePassword(cxt *gin.Context)
	DeleteUser(cxt *gin.Context)
	UpdateUser(cxt *gin.Context)
}
type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) UserController {
	return &userHandler{userService: userService}
}

func (u userHandler) GetProfile(cxt *gin.Context) {
	res, err := u.userService.GetProfile(cxt)
	if err != nil {
		cxt.JSON(500, gin.H{"error": err.Error()})
		return
	}
	cxt.JSON(200, res)
}

func (u userHandler) ChangePassword(cxt *gin.Context) {
	var req models.ChangePasswordJson
	if err := cxt.ShouldBindJSON(&req); err != nil {
		cxt.JSON(400, gin.H{"error": "Invalid request"})
		return
	}
	res, err := u.userService.ChangePassword(req.CurrentPassword, req.NewPassword, cxt)
	if err != nil {
		cxt.JSON(500, gin.H{"error": err.Error()})
		return
	}
	cxt.JSON(200, gin.H{"success": res})
}

func (u userHandler) DeleteUser(cxt *gin.Context) {
	res, err := u.userService.DeleteUser(cxt)
	if err != nil {
		cxt.JSON(500, gin.H{"error": err.Error()})
		return
	}
	cxt.JSON(200, gin.H{"success": res})
}

func (u userHandler) UpdateUser(cxt *gin.Context) {
	var req models.UpdateUserJson
	if err := cxt.ShouldBindJSON(&req); err != nil {
		cxt.JSON(400, gin.H{"error": "Invalid request"})
		return
	}
	res, isSuccess, err := u.userService.UpdateUser(req, cxt)
	if err != nil {
		cxt.JSON(500, gin.H{"error": err.Error()})
		return
	}
	cxt.JSON(200, gin.H{"success": isSuccess, "user": res})
}
