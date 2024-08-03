package services

import (
	"context"
	"errors"
	"gofiber-boilerplatev3/internal/v1/domain/models"
	"gofiber-boilerplatev3/internal/v1/domain/repositories"
)

// UserServiceImpl provides business logic related to users
type passwordResetServiceImpl struct {
	PasswordResetRepo repositories.PasswordResetRepository
}

// NewUserService creates a new instance of UserServiceImpl
func NewPasswordResetService(passwordResetRepo repositories.PasswordResetRepository) PasswordResetService {
	return &passwordResetServiceImpl{PasswordResetRepo: passwordResetRepo}
}

// CreateUser creates a new user and persists it in the repository
func (s *passwordResetServiceImpl) CreatePasswordResets(ctx context.Context, req *models.PasswordReset) (*models.PasswordReset, error) {

	// Check if a user with the same mailpack already exists
	existingPasswordResetRequest, _ := s.PasswordResetRepo.FindByUserID(ctx, req.UserID)
	if existingPasswordResetRequest != nil {
		return nil, errors.New("User already requested password reset")
	}

	res, err := s.PasswordResetRepo.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
