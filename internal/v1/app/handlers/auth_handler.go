package handlers

import (
	"encoding/json"
	"github.com/gofiber/fiber/v3"
	"gofiber-boilerplatev3/internal/v1/app/dto"
	"gofiber-boilerplatev3/internal/v1/app/usecases"
	"gofiber-boilerplatev3/internal/v1/interface/http/middlewares"
	"gofiber-boilerplatev3/pkg/infra/config"
	"gofiber-boilerplatev3/pkg/utils/auth/jwt"
	"gofiber-boilerplatev3/pkg/utils/msg"
)

// UserHandler handles HTTP requests related to users
type authHandler struct {
	AuthUsecase usecases.AuthUsecase
	Config      config.Config
}

// NewUserHandler creates a new UserHandler instance
func NewAuthHandler(authUsecase usecases.AuthUsecase, config config.Config) *authHandler {
	return &authHandler{
		AuthUsecase: authUsecase,
		Config:      config,
	}
}

// CreateUser handles POST requests for creating a new user
func (h *authHandler) AuthRegister(c fiber.Ctx) error {

	var dto dto.UserRegisterDTO
	if err := json.Unmarshal(c.Body(), &dto); err != nil {
		if err != nil {
			panic(msg.BadRequestError{
				Message: err.Error(),
			})
		}
	}
	// Register User
	resp, err := h.AuthUsecase.AuthRegister(c.Context(), &dto)
	if err != nil {
		panic(err.Error())
	}
	// Generate JWT
	tokenAccess, err := jwt.GenerateAccessToken(resp)
	if err != nil {
		panic(err.Error())
	}
	tokenRefresh, err := jwt.GenerateRefreshToken(resp.ID)
	if err != nil {
		panic(err.Error())
	}

	response := jwt.Token{
		AccessToken:  tokenAccess,
		RefreshToken: tokenRefresh,
	}

	return middlewares.Send(c, fiber.StatusCreated, response)
}

// Login handles GET requests for retrieving a user by ID
func (h *authHandler) AuthLogin(c fiber.Ctx) error {
	var dto dto.UserLoginDTO
	if err := json.Unmarshal(c.Body(), &dto); err != nil {
		if err != nil {
			panic(msg.BadRequestError{
				Message: err.Error(),
			})
		}
	}
	// Login User
	resp, err := h.AuthUsecase.AuthLogin(c.Context(), &dto)
	if err != nil {
		panic(msg.NotFoundError{
			Message: err.Error(),
		})
	}
	// Generate JWT
	tokenAccess, err := jwt.GenerateAccessToken(resp)
	if err != nil {
		panic(err.Error())
	}
	tokenRefresh, err := jwt.GenerateRefreshToken(resp.ID)
	if err != nil {
		panic(err.Error())
	}
	response := jwt.Token{
		AccessToken:  tokenAccess,
		RefreshToken: tokenRefresh,
	}
	return middlewares.Send(c, fiber.StatusOK, response)

}
