package middlewares

import (
	"github.com/gofiber/fiber/v3"
	"gofiber-boilerplatev3/pkg/utils/auth/jwt"
)

func AuthMiddleware() fiber.Handler {
	return func(c fiber.Ctx) error {
		// Get the token from the request header
		tokenString := c.Get("Authorization")
		if tokenString == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Authorization header missing",
			})
		}

		// Remove "Bearer " prefix if it exists
		if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
			tokenString = tokenString[7:]
		}

		// Validate the token
		claims, err := jwt.ValidateAccessToken(tokenString)
		if err != nil {
			//fmt.Println("Token validation error:", err) // Debugging line to print error details
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid or expired token",
			})
		}

		// Optionally, you can set the claims in the context for use in handlers
		c.Locals("claims", claims)

		// Continue with the next handler
		return c.Next()
	}
}
