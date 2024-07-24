package repositoryimpl

import (
	"context"
	"errors"
	"goamartha/domain/entity"
	"goamartha/repository"
	"gorm.io/gorm"
)

func NewHelloRepositoryImpl(DB *gorm.DB) repository.HelloRepository {
	return &HelloRespositoryImpl{DB: DB}
}

type HelloRespositoryImpl struct {
	*gorm.DB
}

func (repository *HelloRespositoryImpl) FindById(ctx context.Context, id string) (entity.Users, error) {
	var users entity.Users
	result := repository.DB.WithContext(ctx).Unscoped().Where("id = ?", id).First(&users)
	if result.RowsAffected == 0 {
		return entity.Users{}, errors.New("product Not Found")
	}
	return users, nil
}
