package repositoryimpl

import (
	"context"
	"goamartha/domain/entity"
	"goamartha/repository"
	"gorm.io/gorm"
)

type PaymentRespositoryImpl struct {
	*gorm.DB
}

func NewPaymentRespositoryImpl(DB *gorm.DB) repository.PaymentRepository {
	return &PaymentRespositoryImpl{DB: DB}
}

func (repository *PaymentRespositoryImpl) Insert(ctx context.Context, request *entity.Payment) error {

	err := repository.DB.WithContext(ctx).Create(&request).Error
	if err != nil {
		return err
	}
	if err := repository.DB.WithContext(ctx).First(&request, request.Id).Error; err != nil {
		return err
	}
	return err
}

func (repository *PaymentRespositoryImpl) CountPayment(ctx context.Context, request *entity.Payment) (int64, error) {
	var count int64
	err := repository.DB.WithContext(ctx).Model(&request).Where("borrow_id = ? AND is_paid = ?", request.BorrowId, true).Count(&count).Error

	if err != nil {
		return count, err
	}

	return count, err
}
