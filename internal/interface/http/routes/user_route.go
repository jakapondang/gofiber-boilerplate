package routes

import (
	"github.com/gofiber/fiber/v3"
	"gofiber-boilerplatev3/internal/application/handlers"
	"gofiber-boilerplatev3/internal/application/usecases"
	"gofiber-boilerplatev3/internal/domain/repositories"
	"gofiber-boilerplatev3/internal/domain/services"
	"gorm.io/gorm"
)

// SetupUserRoutes sets up the user routes
func SetupUserRoutes(app *fiber.App, db *gorm.DB) {
	// Initialize repositories
	userRepo := repositories.NewUserRepository(db)

	// Initialize services
	userService := services.NewUserService(userRepo)

	// Initialize use cases
	userUsecase := usecases.NewUserUsecase(userService)

	// Initialize handlers
	userHandler := handlers.NewUserHandler(userUsecase)
	api := app.Group("/api")

	// Apply the authentication middleware to the user routes
	//user := api.Group("/users", middlewares.AuthMiddleware)
	//Without Auth
	user := api.Group("/users")

	user.Post("/", userHandler.CreateUser)
	user.Get("/:id", userHandler.GetUserByID)
}
