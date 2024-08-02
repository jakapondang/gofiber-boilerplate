package dto

import (
	"gofiber-boilerplatev3/internal/v1/domain/models"
	"time"
)

// UserDTO represents a user data transfer object
type UserDTO struct {
	//Password  string     `json:"password"`
	ID                  string     `json:"id"`
	Username            string     `json:"username"`
	Email               string     `json:"mail"`
	FirstName           string     `json:"firstName,omitempty"`
	LastName            string     `json:"lastName,omitempty"`
	PhoneNumber         string     `json:"phoneNumber"`
	IsVerifyEmail       bool       `json:"isVerifyEmail"`
	IsVerifyPhoneNumber bool       `json:"isVerifyPhoneNumber"`
	IsActive            bool       `json:"isActive"`
	IsAdmin             bool       `json:"isAdmin"`
	CreatedAt           string     `json:"createdAt"`
	UpdatedAt           string     `json:"updatedAt"`
	LastLogin           *time.Time `json:"lastLogin,omitempty"`
}

// Trasnform Model User to User DTO
func NewUser(res *models.User) *UserDTO {
	return &UserDTO{
		//Password:  res.PasswordHash,
		ID:                  res.ID.String(),
		Username:            res.Username,
		Email:               res.Email,
		FirstName:           res.FirstName,
		LastName:            res.LastName,
		PhoneNumber:         res.PhoneNumber,
		IsVerifyEmail:       res.IsVerifyEmail,
		IsVerifyPhoneNumber: res.IsVerifyPhoneNumber,
		IsActive:            res.IsActive,
		IsAdmin:             res.IsAdmin,
		CreatedAt:           res.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:           res.UpdatedAt.Format("2006-01-02 15:04:05"),
		LastLogin:           res.LastLogin,
	}
}
