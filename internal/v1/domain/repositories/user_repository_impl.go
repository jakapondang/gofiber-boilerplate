package repositories

import (
	"context"
	"gofiber-boilerplatev3/internal/v1/domain/models"
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

// Save persists a user to the database
func (r *UserRepositoryImpl) Create(ctx context.Context, res *models.User) error {
	err := r.DB.WithContext(ctx).Create(&res).Error
	if err != nil {
		return err
	}
	if err := r.DB.WithContext(ctx).First(&res, res.ID).Error; err != nil {
		return err
	}
	return nil
}

// FindByEmail retrieves a user by email from the database
func (r *UserRepositoryImpl) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := r.DB.WithContext(ctx).Unscoped().Where("email = ?", email).Last(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

//
//// FindByID retrieves a user by ID from the database
//func (r *UserRepositoryImpl) FindByID(id uint) (*models.User, error) {
//	var user models.User
//	result := r.DB.First(&user, id)
//	if result.Error != nil {
//		return nil, result.Error
//	}
//	return &user, nil
//}
//

//
//// Delete removes a user from the database
//func (r *UserRepositoryImpl) Delete(id uint) error {
//	result := r.DB.Delete(&models.User{}, id)
//	return result.Error
//}
