package handlers

import (
	"encoding/json"
	"github.com/gofiber/fiber/v3"
	"gofiber-boilerplatev3/internal/v1/app/dto"
	"gofiber-boilerplatev3/internal/v1/app/usecases"
	"gofiber-boilerplatev3/pkg/infra/auth"
	"gofiber-boilerplatev3/pkg/infra/config"
	"gofiber-boilerplatev3/pkg/msg"
	"net/http"
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
	errT := "(CreateUser) "
	var dto dto.UserRegisterDTO
	if err := json.Unmarshal(c.Body(), &dto); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	resp, err := h.UserUsecase.RegisterUser(c.Context(), &dto)
	if err != nil {
		panic(errT + err.Error())
	}
	tokenAccess, err := auth.GenerateAccessToken(h.Config, resp)
	if err != nil {
		panic(errT + err.Error())
	}
	tokenRefresh, err := auth.GenerateRefreshToken(h.Config, resp.ID)
	if err != nil {
		panic(errT + err.Error())
	}

	response := fiber.Map{
		"access_token":  tokenAccess,
		"refresh_token": tokenRefresh,
	}

	return msg.Send(c, fiber.StatusCreated, response)
}

//
//// GetUserByID handles GET requests for retrieving a user by ID
//func (h *UserHandler) GetUserByID(c fiber.Ctx) error {
//	errT := "(GetUserByID) "
//	idStr := c.Params("id")
//	id, err := strconv.Atoi(idStr)
//	if err != nil {
//		panic(msg.BadRequestError{
//			Message: errT + err.Error(),
//		})
//	}
//
//	user, err := h.userUsecase.GetUserByID(uint(id))
//	if err != nil {
//		panic(msg.NotFoundError{
//			Message: errT + err.Error(),
//		})
//	}
//	return c.Status(fiber.StatusOK).JSON(msg.Send(c, user))
//
//}
