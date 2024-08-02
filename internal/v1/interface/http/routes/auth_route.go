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

// SetupAuthRoutes sets up the auth routes
func SetupAuthRoutes(app *fiber.App, db *gorm.DB, config config.Config) {
	// Initialize repositories
	repo := repositories.NewUserRepository(db)

	// Initialize services
	service := services.NewUserService(repo)

	// Initialize use cases
	usecase := usecases.NewAuthUsecase(service)

	// Initialize handlers
	handler := handlers.NewAuthHandler(usecase, config)

	api := app.Group("/api")

	v1 := api.Group("/v1")

	// Public routes
	v1.Post("/register", handler.AuthRegister)
	v1.Post("/login", handler.AuthLogin)
	v1.Post("/refresh-token", handler.RefreshToken)

}
