package domain

import (
	"context"
	"gorm.io/gorm"
)

// Transaction Domain Manager Implementation
type TrxDomainImpl struct {
	DB *gorm.DB
}

func NewTrxDomain(db *gorm.DB) *TrxDomainImpl {
	return &TrxDomainImpl{DB: db}
}

func (tm *TrxDomainImpl) BeginTx(ctx context.Context) (*gorm.DB, error) {
	return tm.DB.WithContext(ctx).Begin(), nil
}

func (tm *TrxDomainImpl) CommitTx(tx *gorm.DB) error {
	return tx.Commit().Error
}

func (tm *TrxDomainImpl) RollbackTx(tx *gorm.DB) error {
	return tx.Rollback().Error
}
