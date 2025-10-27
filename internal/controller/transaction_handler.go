package controller

import (
	"fakeBank/internal/models"
	"fakeBank/internal/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TransactionController interface {
	GetTransactions(cxt *gin.Context)
	GetTransactionById(cxt *gin.Context)
	CreateTransactionTransfer(cxt *gin.Context)
}
type transactionController struct {
	transactionService service.TransactionService
}

func NewTransactionHandler(transactionService service.TransactionService) TransactionController {
	return &transactionController{transactionService: transactionService}
}

func (h *transactionController) GetTransactions(cxt *gin.Context) {
	accountId, _ := strconv.Atoi(cxt.Query("accountId"))
	limitStr, _ := strconv.Atoi(cxt.DefaultQuery("limit", "50"))
	offsetStr, _ := strconv.Atoi(cxt.DefaultQuery("offset", "0"))
	res, err := h.transactionService.GetTransactions(accountId, limitStr, offsetStr)
	if err != nil {
		cxt.JSON(http.StatusInternalServerError, err)
		return
	}
	fmt.Println(res)
	cxt.JSON(http.StatusOK, res)
}

func (h *transactionController) GetTransactionById(cxt *gin.Context) {
	transactionId, _ := strconv.Atoi(cxt.Param("transactionId"))
	res, err := h.transactionService.GetTransactionById(transactionId)
	if err != nil {
		cxt.JSON(http.StatusInternalServerError, err)
		fmt.Println(err)
		return
	}
	cxt.JSON(http.StatusOK, res)
}

func (h *transactionController) CreateTransactionTransfer(cxt *gin.Context) {
	var req models.TransactionTransferReq
	if err := cxt.ShouldBindJSON(&req); err != nil {
		cxt.JSON(http.StatusBadRequest, err)
		return
	}
	res, err := h.transactionService.CreateTransactionTransfer(req)
	if err != nil {
		cxt.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err})
		fmt.Println(err)

		return
	}
	fmt.Println(err)

	cxt.JSON(http.StatusOK, gin.H{"transactionId": res})
}
