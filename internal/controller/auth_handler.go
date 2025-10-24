package controller

import (
	"fakeBank/internal/models"
	"fakeBank/internal/service"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(cxt *gin.Context)
	Signup(cxt *gin.Context)
	Me(cxt *gin.Context)
	Logout(cxt *gin.Context)
}
type authController struct {
	authService service.AuthService
}

func NewAuthController(userSvc service.AuthService) AuthController {
	return &authController{authService: userSvc}
}

func (u authController) Login(cxt *gin.Context) {
	var req models.LoginJson
	if err := cxt.ShouldBindJSON(&req); err != nil {
		cxt.JSON(400, gin.H{"error": "Invalid request"})
		return
	}
	var res models.GetUser
	isValid, res, err := u.authService.Login(req.Email, req.Password, cxt)
	if err != nil {
		cxt.JSON(401, gin.H{
			"success": false,
			"error":   "Invalid email or password",
		})
		return
	}

	cxt.JSON(200, gin.H{"success": isValid, "user": res})
}

func (u authController) Signup(cxt *gin.Context) {
	var req models.UserJson
	if err := cxt.ShouldBindJSON(&req); err != nil {
		cxt.JSON(400, gin.H{"error": "Invalid request"})
		return
	}
	isCreated, userId, createdAt, err := u.authService.Signup(req, cxt)
	if err != nil {
		cxt.JSON(500, gin.H{
			"success": false,
			"error":   "Could not create user",
		})
		return
	}
	cxt.JSON(201, gin.H{"success": isCreated, "user": models.GetUser{Id: userId, FirstName: req.FirstName, LastName: req.LastName, Email: req.Email, Phone: req.Phone, CreatedAt: createdAt}})
}

func (u authController) Me(cxt *gin.Context) {
	res, err := u.authService.Me(cxt)
	if err != nil {
		cxt.JSON(401, gin.H{
			"success": false,
			"error":   "Unauthorized",
		})
		return
	}
	cxt.JSON(200, res)
}

func (u authController) Logout(cxt *gin.Context) {
	res := u.authService.Logout(cxt)
	cxt.JSON(200, gin.H{"success": res})
}
