package repositories

import (
	"context"
	"github.com/google/uuid"
	"gofiber-boilerplatev3/internal/v1/domain/models"
	"gorm.io/gorm"
)

// UserRepositoryImpl is the implementation of the UserRepository interface
type userRepositoryImpl struct {
	DB *gorm.DB
}

// NewUserRepository creates a new instance of UserRepositoryImpl
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{DB: db}
}

// Save persists a user to the database
func (r *userRepositoryImpl) Create(ctx context.Context, tx *gorm.DB, res *models.User) error {
	err := tx.WithContext(ctx).Create(&res).Error
	if err != nil {
		return err
	}
	if err := tx.WithContext(ctx).First(&res, res.ID).Error; err != nil {
		return err
	}
	return nil
}

// FindByEmail retrieves a user by mailpack from the database
func (r *userRepositoryImpl) FindByEmail(ctx context.Context, tx *gorm.DB, email string) (*models.User, error) {
	var user *models.User

	err := tx.WithContext(ctx).Unscoped().Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

// FindByEmail retrieves a user by ID from the database
func (r *userRepositoryImpl) FindByID(ctx context.Context, tx *gorm.DB, ID uuid.UUID) (*models.User, error) {
	var user *models.User
	err := tx.WithContext(ctx).Unscoped().Where("id = ?", ID).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Update persists a user to the database
func (r *userRepositoryImpl) Update(ctx context.Context, tx *gorm.DB, res *models.User) error {
	err := tx.WithContext(ctx).Where("id = ?", res.ID).Updates(&res).Error
	if err != nil {
		return err
	}
	return nil
}
