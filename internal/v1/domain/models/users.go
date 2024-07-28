package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// User represents a user entity
type User struct {
	ID                  uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;column:id"`
	Username            string     `gorm:"type:varchar(50);unique;not null;column:username"`
	Email               string     `gorm:"type:varchar(255);unique;not null;column:email"`
	PasswordHash        string     `gorm:"type:varchar(255);not null;column:password_hash"`
	FirstName           string     `gorm:"type:varchar(50);column:first_name"`
	LastName            string     `gorm:"type:varchar(50);column:last_name"`
	PhoneNumber         string     `gorm:"type:varchar(20);column:phone_number"`
	IsVerifyEmail       bool       `gorm:"type:boolean;default:false;column:is_verify_email"`
	IsVerifyPhoneNumber bool       `gorm:"type:boolean;default:false;column:is_verify_phone_number"`
	IsActive            bool       `gorm:"default:true;column:is_active"`
	IsAdmin             bool       `gorm:"default:false;column:is_admin"`
	CreatedAt           time.Time  `gorm:"default:current_timestamp;column:created_at"`
	UpdatedAt           time.Time  `gorm:"default:current_timestamp;column:updated_at"`
	LastLogin           *time.Time `gorm:"column:last_login"`
	//Roles        []userroles.Entity `gorm:"many2many:user_roles"`
}

// Table Name
func (User) TableName() string {
	return "users"
}

// UUID Unique
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
