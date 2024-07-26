package dto

import (
	"github.com/google/uuid"
	"time"
)

// UserDTO represents a user data transfer object
type UserDTO struct {
	ID        uuid.UUID  `json:"id"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	FirstName string     `json:"firstName,omitempty"`
	LastName  string     `json:"lastName,omitempty"`
	IsActive  bool       `json:"isActive"`
	IsAdmin   bool       `json:"isAdmin"`
	CreatedAt string     `json:"createdAt"`
	UpdatedAt string     `json:"updatedAt"`
	LastLogin *time.Time `json:"lastLogin,omitempty"`
}
