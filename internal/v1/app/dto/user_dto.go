package dto

import (
	"github.com/google/uuid"
	"gofiber-boilerplatev3/internal/v1/domain/models"
	"time"
)

// UserDTO represents a user data transfer object
type UserDTO struct {
	ID        uuid.UUID  `json:"id"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	FirstName string     `json:"firstName,omitempty"`
	LastName  string     `json:"lastName,omitempty"`
	IsActive  bool       `json:"isActive"`
	IsAdmin   bool       `json:"isAdmin"`
	CreatedAt string     `json:"createdAt"`
	UpdatedAt string     `json:"updatedAt"`
	LastLogin *time.Time `json:"lastLogin,omitempty"`
}

type UserRegisterDTO struct {
	Username string `json:"username" validate:"required,min=3,max=50"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// NewUser creates a new User instance
func NewRegisterUser(req *UserRegisterDTO) *models.User {
	return &models.User{
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: req.Password,
	}
}

func NewUserDTO(res *models.User) *UserDTO {
	return &UserDTO{
		ID:        res.ID,
		Username:  res.Username,
		Email:     res.Email,
		Password:  res.PasswordHash,
		FirstName: res.FirstName,
		LastName:  res.LastName,
		IsActive:  res.IsActive,
		IsAdmin:   res.IsAdmin,
		CreatedAt: res.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: res.UpdatedAt.Format("2006-01-02 15:04:05"),
		LastLogin: res.LastLogin,
	}
}
