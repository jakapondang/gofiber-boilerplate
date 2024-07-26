package repositoryimpl

import (
	"context"
	"errors"
	"goamartha/domain/entity"
	"goamartha/repository"
	"gorm.io/gorm"
)

type BorrowerRespositoryImpl struct {
	*gorm.DB
}

func NewBorrowerRepositoryImpl(DB *gorm.DB) repository.BorrowerRepository {
	return &BorrowerRespositoryImpl{DB: DB}
}

func (repository *BorrowerRespositoryImpl) FindAll(ctx context.Context) ([]entity.Borrower, error) {
	var borrowers []entity.Borrower
	result := repository.DB.WithContext(ctx).Find(&borrowers)
	if result.RowsAffected == 0 {
		return borrowers, errors.New("Error Query Borrowers Not Found")
	}
	return borrowers, nil
}
func (repository *BorrowerRespositoryImpl) FindById(ctx context.Context, borrower *entity.Borrower) error {
	result := repository.DB.WithContext(ctx).Unscoped().Where("id = ?", borrower.Id).First(&borrower)
	if result.RowsAffected == 0 {
		return errors.New("Borrower Not Found")
	}
	return nil
}

func (repository *BorrowerRespositoryImpl) FindByUsername(ctx context.Context, borrower *entity.Borrower) error {

	result := repository.DB.WithContext(ctx).Unscoped().Where("username = ?", &borrower.Username).Last(&borrower)
	if result.RowsAffected == 0 {
		return errors.New("Username Not Found")
	}
	return nil
}

func (repository *BorrowerRespositoryImpl) Insert(ctx context.Context, borrow *entity.Borrower) error {

	err := repository.DB.WithContext(ctx).Create(&borrow).Error
	if err != nil {
		return err
	}
	if err := repository.DB.WithContext(ctx).First(&borrow, borrow.Id).Error; err != nil {
		return err
	}
	return nil
}

func (repository *BorrowerRespositoryImpl) Update(ctx context.Context, borrower *entity.Borrower) error {
	err := repository.DB.WithContext(ctx).Where("id = ?", borrower.Id).Updates(&borrower).Error
	if err != nil {
		return err
	}
	return nil
}
