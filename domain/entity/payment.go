package entity

import "time"

// Payment represents the payment table in the database.
type Payment struct {
	Id           uint      `gorm:"primaryKey;autoIncrement;column:id"`
	BorrowId     uint      `gorm:"not null;column:borrow_id"`
	TotalPayment float64   `gorm:"type:numeric(10,2);check:total_payment >= 0;column:total_payment"`
	IsPaid       bool      `gorm:"default:false;column:is_paid"`
	Created      time.Time `gorm:"autoCreateTime;column:created"`
}

func (Payment) TableName() string {
	return "payment"
}
