package repository

import (
	"context"
	"goamartha/domain/entity"
)

type BorrowerRepository interface {
	FindAll(ctx context.Context) ([]entity.Borrower, error)
	FindById(ctx context.Context, borrower *entity.Borrower) error
	FindByUsername(ctx context.Context, borrower *entity.Borrower) error
	Insert(ctx context.Context, borrower *entity.Borrower) error
	Update(ctx context.Context, borrower *entity.Borrower) error
}
