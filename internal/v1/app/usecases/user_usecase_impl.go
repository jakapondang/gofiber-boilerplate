package usecases

import (
	"context"
	"github.com/google/uuid"
	"gofiber-boilerplatev3/internal/v1/app/dto"
	"gofiber-boilerplatev3/internal/v1/domain"
	"gofiber-boilerplatev3/internal/v1/domain/services"
)

// userUsecaseImpl implements the UserUsecase interface
type userUsecaseImpl struct {
	TrxDomain   domain.TrxDomain
	UserService services.UserService
}

// NewUserUsecase creates a new instance of userUsecaseImpl
func NewUserUsecase(trxDomain domain.TrxDomain, userService services.UserService) UserUsecase {
	return &userUsecaseImpl{
		TrxDomain:   trxDomain,
		UserService: userService}
}

// UserFindByID retrieves a user by ID
func (u *userUsecaseImpl) UserFindByID(ctx context.Context, ID string) (*dto.UserDTO, error) {
	// Begin Trx
	tx, err := u.TrxDomain.BeginTx(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		if r := recover(); r != nil {
			u.TrxDomain.RollbackTx(tx)
		}
	}()
	//Get User
	userId, err := uuid.Parse(ID)
	res, err := u.UserService.GetUserByID(ctx, tx, userId)
	if err != nil {
		u.TrxDomain.RollbackTx(tx)
		return nil, err
	}

	resp := dto.NewUser(res)
	// Commit
	if err := u.TrxDomain.CommitTx(tx); err != nil {
		return nil, err
	}
	return resp, nil
}
