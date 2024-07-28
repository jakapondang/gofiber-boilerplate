package models

import "github.com/google/uuid"

type EntityRoles struct {
	UserID uuid.UUID `gorm:"type:uuid"`
	RoleID uint      `gorm:"type:integer"`
}

func (EntityRoles) TableName() string {
	return "user_roles"
}
