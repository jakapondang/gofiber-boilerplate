package users

import (
	"context"
)

type Repository interface {
	FindAll(ctx context.Context) ([]Entity, error)
	Insert(ctx context.Context, result *Entity) error
}
