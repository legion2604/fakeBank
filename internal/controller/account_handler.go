package controller

import (
	"fakeBank/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AccountController interface {
	GetAccounts(cxt *gin.Context)
	GetAccountById(cxt *gin.Context)
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
	cxt.JSON(200, res)
}
