package usecases

import (
	"gofiber-boilerplatev3/internal/domain/models"
	"gofiber-boilerplatev3/internal/domain/services"
)

// UserUsecase defines the interface for user-related use cases
type UserUsecase interface {
	CreateUser(username, email string) (*models.User, error)
	GetUserByID(id uint) (*models.User, error)
	UpdateUser(id uint, username, email string) (*models.User, error)
	DeleteUser(id uint) error
}

// userUsecaseImpl implements the UserUsecase interface
type userUsecaseImpl struct {
	userService services.UserService
}

// NewUserUsecase creates a new instance of userUsecaseImpl
func NewUserUsecase(userService services.UserService) UserUsecase {
	return &userUsecaseImpl{userService: userService}
}

// CreateUser creates a new user and returns the created user
func (u *userUsecaseImpl) CreateUser(username, email string) (*models.User, error) {
	return u.userService.CreateUser(username, email)
}

// GetUserByID retrieves a user by ID
func (u *userUsecaseImpl) GetUserByID(id uint) (*models.User, error) {
	return u.userService.GetUserByID(id)
}

// UpdateUser updates the details of an existing user
func (u *userUsecaseImpl) UpdateUser(id uint, username, email string) (*models.User, error) {
	return u.userService.UpdateUser(id, username, email)
}

// DeleteUser deletes a user by ID
func (u *userUsecaseImpl) DeleteUser(id uint) error {
	return u.userService.DeleteUser(id)
}
