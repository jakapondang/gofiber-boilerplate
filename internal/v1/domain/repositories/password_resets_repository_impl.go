package repositories

import (
	"context"
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

func (r *passwordResetRepository) Create(ctx context.Context, res *models.PasswordReset) (*models.PasswordReset, error) {
	// Assume 'res' is an instance of the struct you want to insert
	result := r.DB.WithContext(ctx).Create(&res)
	if result.Error != nil {
		return nil, result.Error
	}

	// 'res' now contains the created record, including any default values set by the database
	return res, nil
}

func (r *passwordResetRepository) FindByToken(ctx context.Context, token string) (*models.PasswordReset, error) {
	var reset models.PasswordReset
	err := r.DB.WithContext(ctx).Where("reset_token = ? AND expires_at > ? AND used = false", token, time.Now()).First(&reset).Error
	return &reset, err
}
func (r *passwordResetRepository) FindByUserID(ctx context.Context, userID string) (*models.PasswordReset, error) {
	var res models.PasswordReset
	err := r.DB.WithContext(ctx).Where("user_id = ? AND expires_at > ? AND used = false", userID, time.Now()).First(&res).Error
	if err != nil {
		return nil, err
	}
	return &res, err
}

func (r *passwordResetRepository) MarkAsUsed(ctx context.Context, reset *models.PasswordReset) error {
	reset.Used = true
	return r.DB.WithContext(ctx).Save(reset).Error
}

func (r *passwordResetRepository) DeleteExpired(ctx context.Context) error {
	return r.DB.WithContext(ctx).Where("expires_at < ?", time.Now()).Delete(&models.PasswordReset{}).Error
}
