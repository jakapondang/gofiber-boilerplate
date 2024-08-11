package v1

import (
	"github.com/gofiber/fiber/v3"
	"gofiber-boilerplatev3/internal/v1/app/handlers"
	"gofiber-boilerplatev3/internal/v1/app/usecases"
	"gofiber-boilerplatev3/internal/v1/domain"
	"gofiber-boilerplatev3/internal/v1/domain/repositories"
	"gofiber-boilerplatev3/internal/v1/domain/services"
	"gofiber-boilerplatev3/internal/v1/interface/http/middlewares"
	"gofiber-boilerplatev3/pkg/infra/config"
	"gorm.io/gorm"
)

// SetupUserRoutes sets up the user routes
func SetupUserRoutes(app *fiber.App, db *gorm.DB, config config.Config) {
	// Initialize repositories
	userRepo := repositories.NewUserRepository(db)

	// Initialize services
	userService := services.NewUserService(userRepo)

	// Initialize Transaction Domain Manager
	trxDomain := domain.NewTrxDomain(db)
	// Initialize use cases
	userUsecase := usecases.NewUserUsecase(trxDomain, userService)

	// Initialize handlers
	userHandler := handlers.NewUserHandler(userUsecase, config)

	api := app.Group("/api")

	v1 := api.Group("/v1")

	// Protected routes
	protected := v1.Group("/protected", middlewares.AuthMiddleware())
	user := protected.Group("/user")
	user.Get("/profile", userHandler.UserProfile)
	user.Put("/profile", userHandler.UserProfileUpdate)
}
