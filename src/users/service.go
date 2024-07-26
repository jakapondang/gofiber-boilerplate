package users

import (
	"context"
)

type Service interface {
	FindAll(ctx context.Context) (response []ModelResponse, err error)
	Create(ctx context.Context, request ModelRequest) (ModelResponse, error)
}
