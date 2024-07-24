package entity

type Users struct {
	Id       uint   `gorm:"primaryKey;autoIncrement"`
	Username string `gorm:"size:50;not null"`
}

func (Users) TableName() string {
	return "users"
}
