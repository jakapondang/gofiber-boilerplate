package usecases

import (
	"context"
	"gofiber-boilerplatev3/internal/v1/app/dto"
	"gofiber-boilerplatev3/pkg/utils/auth/jwt"
)

// AuthUsecase defines the interface for Authorization-related use cases
type AuthUsecase interface {
	AuthRegister(ctx context.Context, req *dto.RegisterDTO) (*jwt.Token, error)
	AuthLogin(ctx context.Context, req *dto.LoginDTO) (*jwt.Token, error)
	RefreshToken(ctx context.Context, req *dto.RefreshTokenDTO) (*jwt.Token, error)
	//UpdateUser(id uint, username, mail string) (*models.User, error)
	//DeleteUser(id uint) error
}
