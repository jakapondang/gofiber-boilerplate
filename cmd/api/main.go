package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"gofiber-boilerplatev3/internal/application/handlers"
	"gofiber-boilerplatev3/internal/application/usecases"
	"gofiber-boilerplatev3/internal/domain/repositories"
	"gofiber-boilerplatev3/internal/domain/services"
	"gofiber-boilerplatev3/internal/infrastructure/config"
	"gofiber-boilerplatev3/internal/infrastructure/database"
	"gofiber-boilerplatev3/internal/interface/http/routes"
	"gofiber-boilerplatev3/pkg/utils"
)

func main() {
	// Load configurations
	err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	// Initialize Logrus
	utils.NewLogger()

	app := fiber.New(utils.NewFiberError())

	// Initialize the database connection
	database.Connect(config.AppConfig)

	// Initialize repositories
	userRepo := repositories.NewUserRepository(database.DB)

	// Initialize services
	userService := services.NewUserService(userRepo)

	// Initialize use cases
	userUsecase := usecases.NewUserUsecase(userService)

	// Initialize handlers
	userHandler := handlers.NewUserHandler(userUsecase)

	// Setup routes
	routes.SetupUserRoutes(app, userHandler)

	log.Fatal(app.Listen(":" + config.AppConfig.Server.Port))
}
