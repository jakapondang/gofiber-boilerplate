package handlers

import (
	"encoding/json"
	"github.com/gofiber/fiber/v3"
	"gofiber-boilerplatev3/internal/v1/app/dto"
	"gofiber-boilerplatev3/internal/v1/app/usecases"
	"gofiber-boilerplatev3/internal/v1/interface/http/middlewares"
	"gofiber-boilerplatev3/pkg/infra/config"
	"gofiber-boilerplatev3/pkg/utils/msg"
)

// authHandler handles HTTP requests related to users
type authHandler struct {
	AuthUsecase usecases.AuthUsecase
	Config      config.Config
}

// New Handler creates a new instance
func NewAuthHandler(authUsecase usecases.AuthUsecase, config config.Config) *authHandler {
	return &authHandler{
		AuthUsecase: authUsecase,
		Config:      config,
	}
}

// Auth handles POST requests for creating a new user
func (h *authHandler) AuthRegister(c fiber.Ctx) error {

	var req dto.RegisterDTO
	if err := json.Unmarshal(c.Body(), &req); err != nil {
		if err != nil {
			panic(msg.BadRequestError{
				Message: err.Error(),
			})
		}
	}
	// Register User
	resp, err := h.AuthUsecase.AuthRegister(c.Context(), &req)
	if err != nil {
		panic(err.Error())
	}

	return middlewares.Send(c, fiber.StatusCreated, resp)
}

// Auth handles GET requests for retrieving a user by ID
func (h *authHandler) AuthLogin(c fiber.Ctx) error {
	var req dto.LoginDTO
	if err := json.Unmarshal(c.Body(), &req); err != nil {
		if err != nil {
			panic(msg.BadRequestError{
				Message: err.Error(),
			})
		}
	}
	// Login User
	resp, err := h.AuthUsecase.AuthLogin(c.Context(), &req)
	if err != nil {
		panic(msg.NotFoundError{
			Message: err.Error(),
		})
	}

	return middlewares.Send(c, fiber.StatusOK, resp)

}

// Auth handles GET requests for Refresh Token
func (h *authHandler) RefreshToken(c fiber.Ctx) error {
	var req dto.RefreshTokenDTO
	if err := json.Unmarshal(c.Body(), &req); err != nil {
		if err != nil {
			panic(msg.BadRequestError{
				Message: err.Error(),
			})
		}
	}

	resp, err := h.AuthUsecase.RefreshToken(c.Context(), &req)
	if err != nil {
		panic(msg.NotFoundError{
			Message: err.Error(),
		})
	}

	return middlewares.Send(c, fiber.StatusOK, resp)
}

// Auth handles Password Reset Request
func (h *authHandler) PasswordResetRequest(c fiber.Ctx) error {
	var req dto.PasswordResetRequestDTO
	if err := json.Unmarshal(c.Body(), &req); err != nil {
		if err != nil {
			panic(msg.BadRequestError{
				Message: err.Error(),
			})
		}
	}

	err := h.AuthUsecase.PasswordResetRequest(c.Context(), &req)
	if err != nil {
		panic(msg.NotFoundError{
			Message: err.Error(),
		})
	}

	return middlewares.Send(c, fiber.StatusOK, "Password reset requested")
}

// Auth handles Password Reset using token with new password
func (h *authHandler) PasswordResetUpdate(c fiber.Ctx) error {
	var req dto.PasswordResetUpdateRequestDTO
	if err := json.Unmarshal(c.Body(), &req); err != nil {
		if err != nil {
			panic(msg.BadRequestError{
				Message: err.Error(),
			})
		}
	}

	err := h.AuthUsecase.PasswordResetUpdate(c.Context(), &req)
	if err != nil {
		panic(msg.NotFoundError{
			Message: err.Error(),
		})
	}

	return middlewares.Send(c, fiber.StatusOK, "Password reset requested")
}
