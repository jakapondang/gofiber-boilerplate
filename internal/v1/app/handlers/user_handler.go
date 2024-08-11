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
type userHandler struct {
	UserHandler usecases.UserUsecase
	Config      config.Config
}

// NewUserHandler creates a new UserHandler instance
func NewUserHandler(userUsecase usecases.UserUsecase, config config.Config) *userHandler {
	return &userHandler{
		UserHandler: userUsecase,
		Config:      config,
	}
}

// UserProfile handles GET Profile for retrieving a user by ID
func (h *userHandler) UserProfile(c fiber.Ctx) error {
	claims := c.Locals("claims").(*jwt.AccessTokenClaims)
	// Convert the `data` field to JSON
	dataJSON, err := json.Marshal(claims.Data)
	if err != nil {
		panic("Failed to convert data to JSON" + err.Error())

	}
	// Unmarshal JSON bytes into a User Token DTO
	var UserToken dto.UserTokenDTO
	if err := json.Unmarshal(dataJSON, &UserToken); err != nil {
		panic("Failed to unmarshal data" + err.Error())
	}

	//Get User Data User
	resp, err := h.UserHandler.UserFindByID(c.Context(), UserToken.ID)
	if err != nil {
		panic(err.Error())
	}

	return middlewares.Send(c, fiber.StatusOK, resp)

}

// UserProfileUpdate handles Update Profile
func (h *userHandler) UserProfileUpdate(c fiber.Ctx) error {
	//Get Jwt claims
	claims := c.Locals("claims").(*jwt.AccessTokenClaims)
	// Convert the `data` field to JSON
	dataJSON, err := json.Marshal(claims.Data)
	if err != nil {
		panic("Failed to convert data to JSON" + err.Error())

	}
	// Unmarshal JSON bytes into a User Token DTO
	var user dto.UserDTO
	if err := json.Unmarshal(dataJSON, &user); err != nil {
		panic("Failed to unmarshal data" + err.Error())
	}

	//Get Request
	var req *dto.UserProfileUpdateRequestDTO
	if err := json.Unmarshal(c.Body(), &req); err != nil {
		if err != nil {
			panic(msg.BadRequestError{
				Message: err.Error(),
			})
		}
	}
	req.ID = user.ID

	//Get User Data User
	resp, err := h.UserHandler.UserUpdateProfile(c.Context(), req)
	if err != nil {
		panic(err.Error())
	}

	return middlewares.Send(c, fiber.StatusOK, resp)

}
