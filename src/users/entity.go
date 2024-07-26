package users

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Entity struct {
	ID           uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;column:id"`
	Username     string     `gorm:"type:varchar(50);unique;not null;column:username"`
	Email        string     `gorm:"type:varchar(255);unique;not null;column:email"`
	PasswordHash string     `gorm:"type:varchar(255);not null;column:password_hash"`
	FirstName    string     `gorm:"type:varchar(50);column:first_name"`
	LastName     string     `gorm:"type:varchar(50);column:last_name"`
	IsActive     bool       `gorm:"default:true;column:is_active"`
	IsAdmin      bool       `gorm:"default:false;column:is_admin"`
	CreatedAt    time.Time  `gorm:"default:current_timestamp;column:created_at"`
	UpdatedAt    time.Time  `gorm:"default:current_timestamp;column:updated_at"`
	LastLogin    *time.Time `gorm:"column:last_login"`
	//Roles        []userroles.Entity `gorm:"many2many:user_roles"`
}

func (Entity) TableName() string {
	return "users"
}
func (u *Entity) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}

type EntityRoles struct {
	UserID uuid.UUID `gorm:"type:uuid"`
	RoleID uint      `gorm:"type:integer"`
}

func (EntityRoles) TableName() string {
	return "user_roles"
}
