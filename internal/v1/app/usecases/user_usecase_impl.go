package usecases

import (
	"context"
	"github.com/google/uuid"
	"gofiber-boilerplatev3/internal/v1/app/dto"
	"gofiber-boilerplatev3/internal/v1/domain"
	"gofiber-boilerplatev3/internal/v1/domain/services"
	"gofiber-boilerplatev3/pkg/utils/msg"
	"gorm.io/gorm"
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
	var resp *dto.UserDTO
	// Start Transaction
	err := domain.WithTransaction(ctx, u.TrxDomain, func(tx *gorm.DB) error {
		//Get User
		userId, err := uuid.Parse(ID)
		res, err := u.UserService.GetUserByID(ctx, tx, userId)
		if err != nil {
			return err
		}

		resp = dto.NewUser(res)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// UserFindByID retrieves a user by ID
func (u *userUsecaseImpl) UserUpdateProfile(ctx context.Context, req *dto.UserProfileUpdateRequestDTO) (*dto.UserDTO, error) {

	msg.Validate(req)
	var resp *dto.UserDTO

	// Start Transaction
	err := domain.WithTransaction(ctx, u.TrxDomain, func(tx *gorm.DB) error {
		//Get User
		userId, err := uuid.Parse(req.ID)
		user, err := u.UserService.GetUserByID(ctx, tx, userId)
		if err != nil {
			return err
		}
		user.FirstName = req.FirstName
		user.LastName = req.LastName
		user.PhoneNumber = req.PhoneNumber

		err = u.UserService.UpdateUser(ctx, tx, user)
		if err != nil {
			return err
		}

		resp = dto.NewUser(user)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}
