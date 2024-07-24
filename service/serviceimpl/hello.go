package serviceimpl

import (
	"context"
	"goamartha/domain/model"
	"goamartha/exception"
	"goamartha/repository"
	"goamartha/service"
)

func NewHelloServiceImpl(helloRepository *repository.HelloRepository) service.HelloService {
	return &HelloServiceImpl{HelloRepository: *helloRepository}
}

type HelloServiceImpl struct {
	repository.HelloRepository
}

func (service *HelloServiceImpl) FindById(ctx context.Context, id string) model.Users {
	helloEntity, err := service.HelloRepository.FindById(ctx, id)
	if err != nil {
		panic(exception.NotFoundError{
			Message: err.Error(),
		})
	}

	return model.Users{
		Id:       helloEntity.Id,
		Username: helloEntity.Username,
	}
}
