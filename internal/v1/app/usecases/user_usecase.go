package usecases

import (
	"context"
	"gofiber-boilerplatev3/internal/v1/app/dto"
)

// UserUsecase defines the interface for user-related use cases
type UserUsecase interface {
	UserFindByID(ctx context.Context, ID string) (*dto.UserDTO, error)
	//UpdateUser(id uint, username, mail string) (*models.User, error)
	//DeleteUser(id uint) error
}
