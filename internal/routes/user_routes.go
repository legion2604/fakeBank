package routes

import (
	"fakeBank/internal/controller"
	"fakeBank/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.RouterGroup, handlers controller.UserController) {
	user := r.Group("/user", middleware.JWTAuthMiddleware())
	{
		user.GET("profile", handlers.GetProfile)
		user.POST("change-password", handlers.ChangePassword)
		user.DELETE("account", handlers.DeleteUser)
		user.PUT("profile", handlers.UpdateUser)
	}
}
