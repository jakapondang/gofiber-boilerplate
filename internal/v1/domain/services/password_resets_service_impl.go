package services

import (
	"context"
	"errors"
	"gofiber-boilerplatev3/internal/v1/domain/models"
	"gofiber-boilerplatev3/internal/v1/domain/repositories"
	"gorm.io/gorm"
)

// UserServiceImpl provides business logic related to users
type passwordResetServiceImpl struct {
	PasswordResetRepo repositories.PasswordResetRepository
}

// NewUserService creates a new instance of UserServiceImpl
func NewPasswordResetService(passwordResetRepo repositories.PasswordResetRepository) PasswordResetService {
	return &passwordResetServiceImpl{PasswordResetRepo: passwordResetRepo}
}

// CreatePasswordResets creates a new password reset and persists it in the repository
func (s *passwordResetServiceImpl) CreatePasswordResets(ctx context.Context, tx *gorm.DB, req *models.PasswordReset) (*models.PasswordReset, error) {
	// If exist set error already request
	existingPasswordResetRequest, _ := s.PasswordResetRepo.FindByUserID(ctx, tx, req.UserID)
	if existingPasswordResetRequest != nil {
		return nil, errors.New("User already requested password reset")
	}

	res, err := s.PasswordResetRepo.Create(ctx, tx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// FindByToken Find data password reset by Token
func (s *passwordResetServiceImpl) FindByToken(ctx context.Context, tx *gorm.DB, req *models.PasswordReset) (*models.PasswordReset, error) {
	res, err := s.PasswordResetRepo.FindByToken(ctx, tx, req.ResetToken)
	if err != nil {
		return nil, errors.New("User token for password reset hase expired")
	}

	return res, nil
}

// FindByToken Find data password reset by Token
func (s *passwordResetServiceImpl) MarkAsUsed(ctx context.Context, tx *gorm.DB, req *models.PasswordReset) error {
	req.Used = true // mark as used
	err := s.PasswordResetRepo.Update(ctx, tx, req)
	if err != nil {
		return errors.New("User token for password reset hase expired")
	}
	return nil
}
