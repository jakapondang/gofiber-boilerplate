package services

import (
	"context"
	"github.com/google/uuid"
	"gofiber-boilerplatev3/internal/v1/domain/models"
	"gorm.io/gorm"
)

type UserService interface {
	CreateUser(ctx context.Context, tx *gorm.DB, res *models.User) (*models.User, error)
	LoginUserByEmail(ctx context.Context, tx *gorm.DB, req *models.User) (*models.User, error)
	GetUserByID(ctx context.Context, tx *gorm.DB, ID uuid.UUID) (*models.User, error)
	UpdateUser(ctx context.Context, tx *gorm.DB, res *models.User) error
	UpdatePasswordUser(ctx context.Context, tx *gorm.DB, res *models.User) error
	GetUserByEmail(ctx context.Context, tx *gorm.DB, req *models.User) (*models.User, error)
}
