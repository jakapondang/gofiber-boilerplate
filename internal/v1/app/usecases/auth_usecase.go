package usecases

import (
	"context"
	"gofiber-boilerplatev3/internal/v1/app/dto"
)

// UserUsecase defines the interface for user-related use cases
type AuthUsecase interface {
	AuthRegister(ctx context.Context, req *dto.UserRegisterDTO) (*dto.UserTokenDTO, error)
	AuthLogin(ctx context.Context, req *dto.UserLoginDTO) (*dto.UserTokenDTO, error)
	//UpdateUser(id uint, username, email string) (*models.User, error)
	//DeleteUser(id uint) error
}
