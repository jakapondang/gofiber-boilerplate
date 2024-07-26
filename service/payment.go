package service

import (
	"context"
	"goamartha/domain/model"
)

type PaymentService interface {
	GetSchedule(ctx context.Context, borrower *model.Borrower) error
	PaymentSummaries(ctx context.Context, borrower *model.Borrower)
	MakePayment(ctx context.Context, borrower *model.Borrower) (model.Payment, error)
	MakeRePayment(ctx context.Context, borrower *model.Borrower) (model.Payment, error)
}
