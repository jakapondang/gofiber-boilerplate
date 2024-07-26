package users

import (
	"github.com/google/uuid"
	"time"
)

type ModelResponse struct {
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

type ModelRequest struct {
	Username  string `json:"username" validate:"required,min=3,max=50"`
	Email     string `json:"email" validate:"required,email"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	IsActive  bool   `json:"isActive"`
	IsAdmin   bool   `json:"isAdmin"`
}

func SingleRow(entity Entity) ModelResponse {
	return ModelResponse{
		ID:        entity.ID,
		Username:  entity.Username,
		Email:     entity.Email,
		FirstName: entity.FirstName,
		LastName:  entity.LastName,
		IsActive:  entity.IsActive,
		IsAdmin:   entity.IsAdmin,
		CreatedAt: entity.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: entity.UpdatedAt.Format("2006-01-02 15:04:05"),
		LastLogin: entity.LastLogin,
	}
}
