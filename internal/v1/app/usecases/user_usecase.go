package usecases

import (
	"context"
	"gofiber-boilerplatev3/internal/v1/app/dto"
)

// UserUsecase defines the interface for user-related use cases
type UserUsecase interface {
	RegisterUser(ctx context.Context, req *dto.UserRegisterDTO) (*dto.UserDTO, error)
	//GetUserByID(id uint) (*models.User, error)
	//UpdateUser(id uint, username, email string) (*models.User, error)
	//DeleteUser(id uint) error
}
