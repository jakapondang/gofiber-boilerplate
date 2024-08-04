package repositories

import (
	"context"
	"github.com/google/uuid"
	"gofiber-boilerplatev3/internal/v1/domain/models"
	"gorm.io/gorm"
	"time"
)

type passwordResetRepository struct {
	DB *gorm.DB
}

func NewPasswordResetRepository(db *gorm.DB) PasswordResetRepository {
	return &passwordResetRepository{DB: db}
}

func (r *passwordResetRepository) Create(ctx context.Context, tx *gorm.DB, res *models.PasswordReset) (*models.PasswordReset, error) {
	result := tx.WithContext(ctx).Create(&res)
	if result.Error != nil {
		return nil, result.Error
	}
	return res, nil
}

func (r *passwordResetRepository) FindByToken(ctx context.Context, tx *gorm.DB, token uuid.UUID) (*models.PasswordReset, error) {
	var reset *models.PasswordReset
	err := tx.WithContext(ctx).Where("reset_token = ? AND expires_at > ? AND used = false", token, time.Now()).First(&reset).Error
	if err != nil {
		return nil, err
	}
	return reset, err
}
func (r *passwordResetRepository) FindByUserID(ctx context.Context, tx *gorm.DB, userID uuid.UUID) (*models.PasswordReset, error) {
	var res *models.PasswordReset
	err := tx.WithContext(ctx).Where("user_id = ? AND expires_at > ? AND used = false", userID, time.Now()).First(&res).Error
	if err != nil {
		return nil, err
	}
	return res, err
}

func (r *passwordResetRepository) Update(ctx context.Context, tx *gorm.DB, reset *models.PasswordReset) error {
	err := tx.WithContext(ctx).Save(&reset).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *passwordResetRepository) DeleteExpired(ctx context.Context, tx *gorm.DB) error {
	err := tx.WithContext(ctx).Where("expires_at < ?", time.Now()).Delete(&models.PasswordReset{}).Error
	if err != nil {
		return err
	}
	return nil
}
