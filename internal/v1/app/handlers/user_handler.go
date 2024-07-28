package handlers

import (
	"encoding/json"
	"github.com/gofiber/fiber/v3"
	"gofiber-boilerplatev3/internal/v1/app/dto"
	"gofiber-boilerplatev3/internal/v1/app/usecases"
	"gofiber-boilerplatev3/pkg/msg"
	"net/http"
)

// UserHandler handles HTTP requests related to users
type UserHandler struct {
	userUsecase usecases.UserUsecase
}

// NewUserHandler creates a new UserHandler instance
func NewUserHandler(userUsecase usecases.UserUsecase) *UserHandler {
	return &UserHandler{userUsecase: userUsecase}
}

// CreateUser handles POST requests for creating a new user
func (h *UserHandler) RegisterUser(c fiber.Ctx) error {
	errT := "(CreateUser) "
	var dto dto.UserRegisterDTO
	if err := json.Unmarshal(c.Body(), &dto); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	resp, err := h.userUsecase.RegisterUser(c.Context(), &dto)
	if err != nil {
		panic(errT + err.Error())
	}

	return msg.Send(c, fiber.StatusCreated, resp)
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
