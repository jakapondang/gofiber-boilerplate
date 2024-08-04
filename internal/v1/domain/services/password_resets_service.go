package services

import (
	"context"
	"gofiber-boilerplatev3/internal/v1/domain/models"
	"gorm.io/gorm"
)

type PasswordResetService interface {
	CreatePasswordResets(ctx context.Context, tx *gorm.DB, req *models.PasswordReset) (*models.PasswordReset, error)
	FindByToken(ctx context.Context, tx *gorm.DB, req *models.PasswordReset) (*models.PasswordReset, error)

	MarkAsUsed(ctx context.Context, tx *gorm.DB, req *models.PasswordReset) error
}
