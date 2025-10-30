package operation

import (
	"fakeBank/internal/controller/operation"
	"fakeBank/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterTransferRoutes(r *gin.RouterGroup, handlers operation.TransferController) {
	transfer := r.Group("/transfers", middleware.JWTAuthMiddleware())
	{
		transfer.POST("", middleware.JWTAuthMiddleware(), handlers.CreateTransactionTransfer)
	}
}
