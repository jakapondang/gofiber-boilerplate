package repositories

import (
	"context"
	"github.com/google/uuid"
	"gofiber-boilerplatev3/internal/v1/domain/models"
	"gorm.io/gorm"
)

type PasswordResetRepository interface {
	Create(ctx context.Context, tx *gorm.DB, reset *models.PasswordReset) (*models.PasswordReset, error)
	FindByToken(ctx context.Context, tx *gorm.DB, token uuid.UUID) (*models.PasswordReset, error)
	Update(ctx context.Context, tx *gorm.DB, reset *models.PasswordReset) error
	DeleteExpired(ctx context.Context, tx *gorm.DB) error
	FindByUserID(ctx context.Context, tx *gorm.DB, userID uuid.UUID) (*models.PasswordReset, error)
}
