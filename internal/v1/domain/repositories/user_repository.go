package repositories

import (
	"context"
	"gofiber-boilerplatev3/internal/v1/domain/models"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	//FindByID(id uint) (*models.User, error)
	//Delete(id uint) error
}