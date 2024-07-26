package users

import (
	"context"
	"goamartha/common"
	"goamartha/exception"
)

type ServiceImpl struct {
	Repository
}

func NewServiceImpl(repository *Repository) Service {
	return &ServiceImpl{Repository: *repository}
}

func (service *ServiceImpl) FindAll(ctx context.Context) (responses []ModelResponse, err error) {
	entities, err := service.Repository.FindAll(ctx)
	for _, entity := range entities {
		responses = append(responses, ModelResponse{
			ID:        entity.ID,
			Username:  entity.Username,
			Email:     entity.Email,
			FirstName: entity.FirstName,
			LastName:  entity.LastName,
			IsActive:  entity.IsActive,
			IsAdmin:   entity.IsAdmin,
			CreatedAt: entity.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: entity.UpdatedAt.Format("2006-01-02 15:04:05"),
			LastLogin: entity.LastLogin,
			//TransactionDate: entity.TransactionDate.Format("2006-01-02 15:04:05"),
			//Created:         entity.Created.Format("2006-01-02 15:04:05"),
		})
	}

	if len(entities) == 0 {
		panic(exception.NotFoundError{
			Message: "Theres no user found",
		})
	}
	return responses, err
}

func (service *ServiceImpl) Create(ctx context.Context, request ModelRequest) (ModelResponse, error) {

	common.Validate(request)
	requestEntity := Entity{
		Username:  request.Username,
		Email:     request.Email,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		IsActive:  request.IsActive,
		IsAdmin:   request.IsAdmin,
	}

	err := service.Repository.Insert(ctx, &requestEntity)
	response := SingleRow(requestEntity)

	return response, err
}
