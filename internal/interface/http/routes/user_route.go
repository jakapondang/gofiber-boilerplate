package routes

import (
	"github.com/gofiber/fiber/v3"
	"gofiber-boilerplatev3/internal/application/handlers"
	"gofiber-boilerplatev3/internal/interface/http/middlewares"
)

// SetupUserRoutes sets up the user routes
func SetupUserRoutes(app *fiber.App, userHandler *handlers.UserHandler) {
	api := app.Group("/api")

	// Apply the authentication middleware to the user routes
	user := api.Group("/users", middlewares.AuthMiddleware)

	user.Post("/", userHandler.CreateUser)
	user.Get("/:id", userHandler.GetUserByID)
}
