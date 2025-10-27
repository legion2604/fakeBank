package controller

import (
	"fakeBank/internal/models"
	"fakeBank/internal/service"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AccountController interface {
	GetAccounts(cxt *gin.Context)
	GetAccountById(cxt *gin.Context)
	CreateAccount(cxt *gin.Context)
}

type accountHandler struct {
	accountService service.AccountService
}

func NewAccountHandler(accountService service.AccountService) AccountController {
	return &accountHandler{accountService: accountService}
}

// methods

func (h *accountHandler) GetAccounts(cxt *gin.Context) {
	account, err := h.accountService.GetAccounts(cxt)
	if err != nil {
		cxt.JSON(500, gin.H{"error": err.Error()})
		return
	}
	cxt.JSON(200, account)
}

func (h *accountHandler) GetAccountById(cxt *gin.Context) {
	accId, _ := strconv.Atoi(cxt.Param("accountId"))
	res, err := h.accountService.GetAccountById(accId)
	if err != nil {
		cxt.JSON(500, gin.H{
			"success": false,
			"error":   "Account not found",
		})
	}
	fmt.Println(res)
	cxt.JSON(200, res)
}

func (h *accountHandler) CreateAccount(cxt *gin.Context) {
	var req models.CreateAccountReq
	if err := cxt.ShouldBindJSON(&req); err != nil {
		cxt.JSON(500, gin.H{"error": err.Error()})
		return
	}
	res, err := h.accountService.CreateAccount(cxt, req)
	if err != nil {
		cxt.JSON(500, gin.H{"error": err.Error()})
		return
	}
	cxt.JSON(200, gin.H{"success": true, "account": res})
}
