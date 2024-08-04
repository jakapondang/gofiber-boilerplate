package dto

import (
	"github.com/google/uuid"
	"gofiber-boilerplatev3/internal/v1/domain/models"
	"time"
)

type PasswordResetRequestDTO struct {
	Email     string    `json:"email" validate:"required,email"`
	UserID    string    `json:"userID"`
	ExpiresAt time.Time `json:"expiresAt"`
}

type PasswordResetRequestEmailDTO struct {
	Username   string `json:"username"`
	ResetToken string `json:"resetToken"`
	ResetLink  string `json:"resetLink"`
	Year       int    `json:"Year"`
}

type PasswordResetUpdateRequestDTO struct {
	ResetToken string `json:"resetToken" validate:"required"`
	Password   string `json:"password" validate:"required"`
}

// Trasnform User Register DTO to Model User
func NewForgotPasswordUser(req *PasswordResetRequestDTO) *models.User {
	return &models.User{
		Email: req.Email,
	}
}

// Trasnform User Register DTO to Model PasswordReset
func NewPasswordReset(req *PasswordResetRequestDTO) *models.PasswordReset {
	userId, _ := uuid.Parse(req.UserID)
	return &models.PasswordReset{
		UserID:    userId,
		ExpiresAt: req.ExpiresAt, // Set 1 hour expiration
	}
}

func NewPasswordResetUpdate(req *PasswordResetUpdateRequestDTO) *models.PasswordReset {
	resetToken, _ := uuid.Parse(req.ResetToken)
	return &models.PasswordReset{
		ResetToken: resetToken,
		Used:       true,
	}
}
