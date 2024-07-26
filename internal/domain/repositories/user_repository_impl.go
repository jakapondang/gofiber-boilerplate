package repositories

import (
	"gofiber-boilerplatev3/internal/domain/models"
	"gorm.io/gorm"
)

// UserRepositoryImpl is the implementation of the UserRepository interface
type UserRepositoryImpl struct {
	DB *gorm.DB
}

// NewUserRepository creates a new instance of UserRepositoryImpl
func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{DB: db}
}

// FindByID retrieves a user by ID from the database
func (r *UserRepositoryImpl) FindByID(id uint) (*models.User, error) {
	var user models.User
	result := r.DB.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// FindByEmail retrieves a user by email from the database
func (r *UserRepositoryImpl) FindByEmail(email string) (*models.User, error) {
	var user models.User
	result := r.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// Save persists a user to the database
func (r *UserRepositoryImpl) Save(user *models.User) error {
	result := r.DB.Save(user)
	return result.Error
}

// Delete removes a user from the database
func (r *UserRepositoryImpl) Delete(id uint) error {
	result := r.DB.Delete(&models.User{}, id)
	return result.Error
}
