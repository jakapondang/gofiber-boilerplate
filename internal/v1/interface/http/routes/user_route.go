package v1

import (
	"github.com/gofiber/fiber/v3"
	"gofiber-boilerplatev3/internal/v1/app/handlers"
	"gofiber-boilerplatev3/internal/v1/app/usecases"
	"gofiber-boilerplatev3/internal/v1/domain/repositories"
	"gofiber-boilerplatev3/internal/v1/domain/services"
	"gofiber-boilerplatev3/pkg/infra/config"
	"gorm.io/gorm"
)

// SetupUserRoutes sets up the user routes
func SetupUserRoutes(app *fiber.App, db *gorm.DB, config config.Config) {
	// Initialize repositories
	userRepo := repositories.NewUserRepository(db)

	// Initialize services
	userService := services.NewUserService(userRepo)

	// Initialize use cases
	userUsecase := usecases.NewUserUsecase(userService)

	// Initialize handlers
	userHandler := handlers.NewUserHandler(userUsecase, config)
	api := app.Group("/api/v1")

	// Apply the authentication middleware to the user routes
	//user := api.Group("/users", middleware.AuthMiddleware)
	//Without Auth
	user := api.Group("/users")

	user.Post("/reg", userHandler.RegisterUser)

}