package service

import (
	"context"
	"goamartha/domain/model"
)

type HelloService interface {
	FindById(ctx context.Context, id string) model.Users
}
