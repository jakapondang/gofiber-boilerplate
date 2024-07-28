package usecases

import "gofiber-boilerplatev3/internal/v1/domain/models"

// UserUsecase defines the interface for user-related use cases
type Roles interface {
	CreateRole(username, email string) (*models.User, error)
	GetUserByID(id uint) (*models.User, error)
	UpdateUser(id uint, username, email string) (*models.User, error)
	DeleteUser(id uint) error
}
