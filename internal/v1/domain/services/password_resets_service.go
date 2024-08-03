package services

import (
	"context"
	"gofiber-boilerplatev3/internal/v1/domain/models"
)

type PasswordResetService interface {
	CreatePasswordResets(ctx context.Context, req *models.PasswordReset) (*models.PasswordReset, error)
	//LoginUserByEmail(ctx context.Context, req *models.User) (*models.User, error)
	//GetUserByID(ctx context.Context, ID string) (*models.User, error)
	//UpdateUser(ctx context.Context, res *models.User) error
	//GetUserByEmail(ctx context.Context, req *models.User) (*models.User, error)
}
