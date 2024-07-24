package repository

import (
	"context"
	"goamartha/domain/entity"
)

type HelloRepository interface {
	FindById(ctx context.Context, id string) (entity.Users, error)
}
