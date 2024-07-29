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

// CreateUser handles POST requests for creating a new user
func (h *UserHandler) RegisterUser(c fiber.Ctx) error {

	var dto dto.UserRegisterDTO
	if err := json.Unmarshal(c.Body(), &dto); err != nil {
		if err != nil {
			panic(msg.BadRequestError{
				Message: err.Error(),
			})
		}
	}
	// Register User
	resp, err := h.UserUsecase.RegisterUser(c.Context(), &dto)
	if err != nil {
		panic(err.Error())
	}
	// Generate JWT
	tokenAccess, err := jwt.GenerateAccessToken(h.Config, resp)
	if err != nil {
		panic(err.Error())
	}
	tokenRefresh, err := jwt.GenerateRefreshToken(h.Config, resp.ID)
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
func (h *UserHandler) LoginUser(c fiber.Ctx) error {
	var dto dto.UserLoginDTO
	if err := json.Unmarshal(c.Body(), &dto); err != nil {
		if err != nil {
			panic(msg.BadRequestError{
				Message: err.Error(),
			})
		}
	}
	// Login User
	resp, err := h.UserUsecase.LoginUser(c.Context(), &dto)
	if err != nil {
		panic(msg.NotFoundError{
			Message: err.Error(),
		})
	}
	// Generate JWT
	tokenAccess, err := jwt.GenerateAccessToken(h.Config, resp)
	if err != nil {
		panic(err.Error())
	}
	tokenRefresh, err := jwt.GenerateRefreshToken(h.Config, resp.ID)
	if err != nil {
		panic(err.Error())
	}
	response := jwt.Token{
		AccessToken:  tokenAccess,
		RefreshToken: tokenRefresh,
	}
	return middlewares.Send(c, fiber.StatusOK, response)

}
