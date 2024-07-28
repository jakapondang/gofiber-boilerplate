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
	Email               string     `json:"email"`
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

type UserRegisterDTO struct {
	Username string `json:"username" validate:"required,min=3,max=50"`
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

// Trasnform Model User to User DTO
func NewUserDTO(res *models.User) *UserDTO {
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

// Trasnform User Register DTO to Model User
func NewRegisterUser(req *UserRegisterDTO) *models.User {
	return &models.User{
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: req.Password,
	}
}

// Trasnform Model User to User DTO
func NewUserTokenDTO(res *models.User) *UserTokenDTO {
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
