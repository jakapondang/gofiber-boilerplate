package dto

import (
	"gofiber-boilerplatev3/internal/v1/domain/models"
	"time"
)

type ForgotPasswordDTO struct {
	Email     string    `json:"email" validate:"required,email"`
	UserID    string    `json:"userID"`
	ExpiresAt time.Time `json:"expiresAt"`
}

type EmaiForgotPasswordDTO struct {
	Username   string `json:"username"`
	ResetToken string `json:"resetToken"`
	ResetLink  string `json:"resetLink"`
	Year       int    `json:"Year"`
}

// Trasnform User Register DTO to Model User
func NewForgotPasswordUser(req *ForgotPasswordDTO) *models.User {
	return &models.User{
		Email: req.Email,
	}
}

// Trasnform User Register DTO to Model User
func NewPasswordReset(req *ForgotPasswordDTO) *models.PasswordReset {
	return &models.PasswordReset{
		UserID:    req.UserID,
		ExpiresAt: req.ExpiresAt, // Set 1 hour expiration
	}
}
