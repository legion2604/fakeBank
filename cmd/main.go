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

	userRepo := repository.NewUserRepo(database.DB)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserHandler(userService)

	accRepo := repository.NewAccountRepository(database.DB)
	accService := service.NewAccountService(accRepo)
	accController := controller.NewAccountHandler(accService)

	transactionRepo := repository.NewTransactionRepository(database.DB)
	transactionService := service.NewTransactionService(transactionRepo)
	transactionController := controller.NewTransactionHandler(transactionService)

	c.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5174"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	api := c.Group("/api")
	{
		routes.RegisterAuthRoutes(api, authController)
		routes.RegisterUserRoutes(api, userController)
		routes.RegisterAccountRoutes(api, accController)
		routes.RegisterTransactionRoutes(api, transactionController)

	}

	c.Run(":8080")
}
