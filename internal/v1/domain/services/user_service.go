package services

import (
	"context"
	"gofiber-boilerplatev3/internal/v1/domain/models"
)

type UserService interface {
	Create(ctx context.Context, res *models.User) (*models.User, error)
	GetUserByEmail(ctx context.Context, req *models.User) (*models.User, error)
	//GetUserByID(id uint) (*models.User, error)
	//UpdateUser(id uint, username, email string) (*models.User, error)
	//DeleteUser(id uint) error
}
