package repository

import (
	"context"
	"goamartha/domain/entity"
)

type PaymentRepository interface {
	Insert(ctx context.Context, payment *entity.Payment) error
	CountPayment(ctx context.Context, payment *entity.Payment) (int64, error)
}
