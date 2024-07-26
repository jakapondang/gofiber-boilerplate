package service

import (
	"context"
	"goamartha/domain/model"
)

type BorrowerService interface {
	FindAll(ctx context.Context) (response []model.Borrower, err error)
	FindById(ctx context.Context, id string) (response model.Borrower, err error)
	FindByUsername(ctx context.Context, username string) (response model.Borrower, err error)
	Create(ctx context.Context, model model.BorrowerRequest) (model.Borrower, error)
	Updates(ctx context.Context, request *model.Borrower) error
}
