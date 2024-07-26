package handlers

import (
	"encoding/json"
	"github.com/gofiber/fiber/v3"
	"gofiber-boilerplatev3/internal/application/dto"
	"gofiber-boilerplatev3/internal/application/usecases"
	"net/http"
	"strconv"
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
func (h *UserHandler) CreateUser(c fiber.Ctx) error {
	var userDTO dto.UserDTO
	if err := json.Unmarshal(c.Body(), &userDTO); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	user, err := h.userUsecase.CreateUser(userDTO.Username, userDTO.Email)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusCreated).JSON(user)
}

// GetUserByID handles GET requests for retrieving a user by ID
func (h *UserHandler) GetUserByID(c fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	user, err := h.userUsecase.GetUserByID(uint(id))
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	return c.Status(http.StatusOK).JSON(user)
}
