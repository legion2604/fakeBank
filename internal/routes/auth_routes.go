package routes

import (
	"fakeBank/internal/controller"
	"fakeBank/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.RouterGroup, handlers controller.AuthController) {
	user := r.Group("/auth")
	{
		user.POST("login", handlers.Login)
		user.POST("signup", handlers.Signup)
		user.GET("me", middleware.JWTAuthMiddleware(), handlers.Me)
		user.POST("logout", handlers.Logout)
	}
}
