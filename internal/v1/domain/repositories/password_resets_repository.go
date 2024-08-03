package repositories

import (
	"context"
	"gofiber-boilerplatev3/internal/v1/domain/models"
)

type PasswordResetRepository interface {
	Create(ctx context.Context, reset *models.PasswordReset) (*models.PasswordReset, error)
	FindByToken(ctx context.Context, token string) (*models.PasswordReset, error)
	MarkAsUsed(ctx context.Context, reset *models.PasswordReset) error
	DeleteExpired(ctx context.Context) error
	FindByUserID(ctx context.Context, userID string) (*models.PasswordReset, error)
}
