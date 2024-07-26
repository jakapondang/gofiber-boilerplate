package model

import "goamartha/domain/entity"

type Borrower struct {
	Id               uint               `json:"id"`
	Username         string             `json:"username"`
	Amount           float64            `json:"amount"`
	TaxRate          float64            `json:"taxRate"`
	TaxAmount        float64            `json:"taxAmount"`
	FinalAmount      float64            `json:"finalAmount"`
	IsSettled        bool               `json:"isSettled"`
	TransactionDate  string             `json:"transactionDate"`
	Created          string             `json:"created"`
	PaymentSummary   PaymentSummary     `json:"paymentSummary"`
	PaymentSchedules []PaymentSchedules `json:"paymentSchedules"`
}

type BorrowerRequest struct {
	Username        string  `json:"username" validate:"required,min=3,max=32"`
	Amount          float64 `json:"amount" validate:"required,gt=0"`
	TaxRate         float64 `json:"taxRate" validate:"gte=0"`
	TaxAmount       float64 `json:"taxAmount" validate:"gte=0"`
	FinalAmount     float64 `json:"finalAmount"`
	IsSettled       bool    `json:"isSettled"`
	TransactionDate string  `json:"transactionDate"`
}

type BorrowerUpdateDateRequest struct {
	TransactionDate string `json:"transactionDate" validate:"required"`
}

func SingleBorrower(entity entity.Borrower) Borrower {
	return Borrower{
		Id:              entity.Id,
		Username:        entity.Username,
		Amount:          entity.Amount,
		TaxRate:         entity.TaxRate * 100,
		TaxAmount:       entity.TaxAmount,
		FinalAmount:     entity.FinalAmount,
		IsSettled:       entity.IsSettled,
		TransactionDate: entity.TransactionDate.Format("2006-01-02 15:04:05"),
		Created:         entity.Created.Format("2006-01-02 15:04:05"),
	}
}
