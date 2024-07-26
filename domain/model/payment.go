package model

import (
	"goamartha/domain/entity"
)

const PaymentTerms = 50
const PaymentTax = 0.1

type Payment struct {
	Id           uint    `json:"id"`
	BorrowId     uint    `json:"borrowId"`
	TotalPayment float64 `json:"totalPayment"`
	IsPaid       bool    `json:"isPaid"`
	Created      string  `json:"created"`
}
type PaymentSummary struct {
	CountPayment        int     `json:"countPayment"`
	TotalPaid           float64 `json:"totalPaid"`
	OutstandingBallance float64 `json:"outstandingBallance"`
	TotalDelinquent     float64 `json:"totalDeliquent"`
}
type PaymentSchedules struct {
	Week         int     `json:"week"`
	TotalPayment float64 `json:"totalPayment"`
	DueDate      string  `json:"due_date"`
	IsPaid       bool    `json:"isPaid"`
	IsDelinquent bool    `json:"isDelinquent"`
}

func SinglePayment(entity entity.Payment) Payment {
	return Payment{
		Id:           entity.Id,
		BorrowId:     entity.BorrowId,
		TotalPayment: entity.TotalPayment,
		IsPaid:       entity.IsPaid,
		Created:      entity.Created.Format("2006-01-02 15:04:05"),
	}
}
