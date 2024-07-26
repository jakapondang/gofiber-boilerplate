package services

import (
	"gofiber-boilerplatev3/internal/domain/models"
)

type UserService interface {
	CreateUser(username, email string) (*models.User, error)
	GetUserByID(id uint) (*models.User, error)
	UpdateUser(id uint, username, email string) (*models.User, error)
	DeleteUser(id uint) error
}
