package handlers

import (
	"encoding/json"
	"github.com/gofiber/fiber/v3"
	"gofiber-boilerplatev3/internal/v1/app/usecases"
	"gofiber-boilerplatev3/internal/v1/interface/http/middlewares"
	"gofiber-boilerplatev3/pkg/infra/config"
	"gofiber-boilerplatev3/pkg/utils/auth/jwt"
)

// UserHandler handles HTTP requests related to users
type UserHandler struct {
	UserUsecase usecases.UserUsecase
	Config      config.Config
}

// NewUserHandler creates a new UserHandler instance
func NewUserHandler(userUsecase usecases.UserUsecase, config config.Config) *UserHandler {
	return &UserHandler{
		UserUsecase: userUsecase,
		Config:      config,
	}
}

// Login handles GET requests for retrieving a user by ID
func (h *UserHandler) UserProfile(c fiber.Ctx) error {
	claims := c.Locals("claims").(*jwt.AccessTokenClaims)

	// Convert claims to JSON string for easier readability
	claimsJSON, err := json.MarshalIndent(claims, "", "  ")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to parse claims",
		})
	}

	// Print the claims
	//fmt.Println("Claims:", string(claimsJSON))

	return middlewares.Send(c, fiber.StatusOK, claimsJSON)

}
