package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type PasswordReset struct {
	UserID     uuid.UUID `gorm:"type:uuid;not null;column:user_id"`
	ResetToken uuid.UUID `gorm:"type:uuid;primaryKey;column:reset_token"`
	CreatedAt  time.Time `gorm:"default:current_timestamp;column:created_at"`
	ExpiresAt  time.Time `gorm:"not null;column:expires_at"`
	Used       bool      `gorm:"default:false;column:used"`
}

// Table Name
func (PasswordReset) TableName() string {
	return "password_resets"
}

// UUID Unique
func (u *PasswordReset) BeforeCreate(tx *gorm.DB) (err error) {
	u.ResetToken = uuid.New()
	return
}
