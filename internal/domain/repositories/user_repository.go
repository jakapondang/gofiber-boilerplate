package repositories

import "gofiber-boilerplatev3/internal/domain/models"

type UserRepository interface {
	FindByID(id uint) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	Save(user *models.User) error
	Delete(id uint) error
}
