package userroles

import (
	"time"
)

type Entity struct {
	ID          uint      `gorm:"primaryKey;autoIncrement"`
	Name        string    `gorm:"type:varchar(50);unique;not null"`
	Description string    `gorm:"type:text"`
	CreatedAt   time.Time `gorm:"default:current_timestamp"`
	UpdatedAt   time.Time `gorm:"default:current_timestamp"`
	//Users       []users.Entity `gorm:"many2many:user_roles"`
}

func (Entity) TableName() string {
	return "roles"
}
