package entity

import "time"

type Borrower struct {
	Id              uint      `gorm:"primaryKey;autoIncrement;column:id"`
	Username        string    `gorm:"size:50;not null;column:username"`
	Amount          float64   `gorm:"type:numeric(10,2);check:amount >= 0;column:amount"`
	TaxRate         float64   `gorm:"type:numeric(5,4);default:0.0000;check:tax_rate >= 0 AND tax_rate <= 1;column:tax_rate"`
	TaxAmount       float64   `gorm:"type:numeric(10,2);default:0.00;check:tax_amount >= 0;column:tax_amount"`
	FinalAmount     float64   `gorm:"type:numeric(12,2);->;column:final_amount"`
	IsSettled       bool      `gorm:"default:false;column:is_settled"`
	TransactionDate time.Time `gorm:"column:transaction_date"`
	Created         time.Time `gorm:"default:current_timestamp;column:created"`
}

func (Borrower) TableName() string {
	return "borrower"
}
