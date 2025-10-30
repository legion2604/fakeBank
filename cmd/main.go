package main

import (
	"fakeBank/internal/controller"
	operationController "fakeBank/internal/controller/operation"
	"fakeBank/internal/middleware"
	"fakeBank/internal/repository"
	operationRepo "fakeBank/internal/repository/operation"
	"fakeBank/internal/routes"
	operationRoutes "fakeBank/internal/routes/operation"
	"fakeBank/internal/service"
	operationService "fakeBank/internal/service/operation"
	"fakeBank/pkg/config"
	"fakeBank/pkg/database"

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

	transferRepo := operationRepo.NewTransferRepository(database.DB)
	transferService := operationService.NewTransferService(transferRepo)
	transferController := operationController.NewTransactionHandler(transferService)

	c.Use(middleware.CORSMiddleware())

	api := c.Group("/api")
	{
		routes.RegisterAuthRoutes(api, authController)
		routes.RegisterUserRoutes(api, userController)
		routes.RegisterAccountRoutes(api, accController)
		routes.RegisterTransactionRoutes(api, transactionController)
		operationRoutes.RegisterTransferRoutes(api, transferController)
	}

	c.Run(":8080")
}
