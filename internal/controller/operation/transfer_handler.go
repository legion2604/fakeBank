package operation

import (
	"fakeBank/internal/models"
	"fakeBank/internal/service/operation"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransferController interface {
	CreateTransactionTransfer(cxt *gin.Context)
}
type transferController struct {
	transactionService operation.TransferService
}

func NewTransactionHandler(transactionService operation.TransferService) TransferController {
	return &transferController{transactionService: transactionService}
}

func (h *transferController) CreateTransactionTransfer(cxt *gin.Context) {
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
