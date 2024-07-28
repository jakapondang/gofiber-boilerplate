package middlewares

import (
	"github.com/gofiber/fiber/v3"
	"net/http"
)

// AuthMiddleware is a simple authentication middleware
func AuthMiddleware(c fiber.Ctx) error {
	// Example: Check for a token in the header (this is a placeholder logic)
	token := c.Get("Authorization")
	if token == "" {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	// Continue to next middleware/handler
	return c.Next()
}
