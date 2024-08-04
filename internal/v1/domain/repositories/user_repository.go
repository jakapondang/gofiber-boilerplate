package repositories

import (
	"context"
	"github.com/google/uuid"
	"gofiber-boilerplatev3/internal/v1/domain/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, tx *gorm.DB, user *models.User) error
	FindByEmail(ctx context.Context, tx *gorm.DB, email string) (*models.User, error)
	FindByID(ctx context.Context, tx *gorm.DB, ID uuid.UUID) (*models.User, error)
	Update(ctx context.Context, tx *gorm.DB, user *models.User) error
}
