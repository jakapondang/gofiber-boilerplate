package domain

import (
	"context"
	"gorm.io/gorm"
)

// Transaction Domain Manager Interface
type TrxDomain interface {
	BeginTx(ctx context.Context) (*gorm.DB, error)
	CommitTx(tx *gorm.DB) error
	RollbackTx(tx *gorm.DB) error
}

func WithTransaction(ctx context.Context, trxDomain TrxDomain, fn func(tx *gorm.DB) error) error {
	tx, err := trxDomain.BeginTx(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if r := recover(); r != nil {
			trxDomain.RollbackTx(tx)
			panic(r) // Re-panic to be caught by the Fiber recover middleware
		}
	}()

	if err := fn(tx); err != nil {
		trxDomain.RollbackTx(tx)
		return err
	}

	if err := trxDomain.CommitTx(tx); err != nil {
		return err
	}

	return nil
}
