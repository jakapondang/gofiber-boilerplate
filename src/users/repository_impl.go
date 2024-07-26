package users

import (
	"context"
	"errors"
	"gorm.io/gorm"
)

type RepositoryImpl struct {
	*gorm.DB
}

func NewRepositoryImpl(DB *gorm.DB) Repository {
	return &RepositoryImpl{DB: DB}
}

func (repository *RepositoryImpl) FindAll(ctx context.Context) ([]Entity, error) {
	var query []Entity
	result := repository.DB.WithContext(ctx).Find(&query)
	if result.RowsAffected == 0 {
		return query, errors.New("Theres no user found")
	}
	return query, nil
}

func (repository *RepositoryImpl) Insert(ctx context.Context, result *Entity) error {

	err := repository.DB.WithContext(ctx).Create(&result).Error
	if err != nil {
		return err
	}
	if err := repository.DB.WithContext(ctx).First(&result, result.ID).Error; err != nil {
		return err
	}
	return nil
}
