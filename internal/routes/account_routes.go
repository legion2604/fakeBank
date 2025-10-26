package routes

import (
	"fakeBank/internal/controller"
	"fakeBank/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterAccountRoutes(r *gin.RouterGroup, handlers controller.AccountController) {
	accounts := r.Group("/accounts", middleware.JWTAuthMiddleware())
	{
		accounts.GET("", handlers.GetAccounts)
		accounts.GET("/:accountId", handlers.GetAccountById)
	}
}
