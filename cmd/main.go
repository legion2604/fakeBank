package main

import (
	"fakeBank/internal/controller"
	"fakeBank/internal/repository"
	"fakeBank/internal/routes"
	"fakeBank/internal/service"
	"fakeBank/pkg/config"
	"fakeBank/pkg/database"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	database.InitDB()

	c := gin.Default()

	authRepo := repository.NewAuthRepository(database.DB)
	authService := service.NewAuthService(authRepo)
	authController := controller.NewAuthController(authService)

	c.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5174"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	api := c.Group("/api")
	{
		routes.RegisterUserRoutes(api, authController)
	}

	c.Run(":8080")
}
