package routes

import (
	"fakeBank/internal/controller"
	"fakeBank/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterTransactionRoutes(r *gin.RouterGroup, handlers controller.TransactionController) {
	transactions := r.Group("/transactions", middleware.JWTAuthMiddleware())
	{
		transactions.GET("", handlers.GetTransactions)
		transactions.GET("/:transactionId", handlers.GetTransactionById)

	}
}
