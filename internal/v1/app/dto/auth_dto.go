package dto

import (
	"gofiber-boilerplatev3/internal/v1/domain/models"
	"time"
)

type RegisterDTO struct {
	Username string `json:"username" validate:"required,min=3,max=50"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserTokenDTO struct {
	ID                  string     `json:"id"`
	Username            string     `json:"username"`
	Email               string     `json:"email"`
	IsVerifyEmail       bool       `json:"isVerifyEmail"`
	IsVerifyPhoneNumber bool       `json:"isVerifyPhoneNumber"`
	IsActive            bool       `json:"isActive"`
	IsAdmin             bool       `json:"isAdmin"`
	LastLogin           *time.Time `json:"lastLogin,omitempty"`
}

type RefreshTokenDTO struct {
	RefreshToken string `json:"refreshToken" validate:"required"`
}

// Trasnform User Register DTO to Model User
func NewRegisterUser(req *RegisterDTO) *models.User {
	return &models.User{
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: req.Password,
	}
}

// Trasnform User Login DTO to Model User
func NewLoginUser(req *LoginDTO) *models.User {
	return &models.User{
		Email:        req.Email,
		PasswordHash: req.Password,
	}
}

// Trasnform Model User to User DTO
func NewTokenUser(res *models.User) *UserTokenDTO {
	return &UserTokenDTO{
		//Password:  res.PasswordHash,
		ID:                  res.ID.String(),
		Username:            res.Username,
		Email:               res.Email,
		IsVerifyEmail:       res.IsVerifyEmail,
		IsVerifyPhoneNumber: res.IsVerifyPhoneNumber,
		IsActive:            res.IsActive,
		IsAdmin:             res.IsAdmin,
		LastLogin:           res.LastLogin,
	}
}
