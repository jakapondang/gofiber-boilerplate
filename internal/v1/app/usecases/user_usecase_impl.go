package usecases

import (
	"context"
	"gofiber-boilerplatev3/internal/v1/app/dto"
	"gofiber-boilerplatev3/internal/v1/domain/services"
)

// userUsecaseImpl implements the UserUsecase interface
type userUsecaseImpl struct {
	UserService services.UserService
}

// NewUserUsecase creates a new instance of userUsecaseImpl
func NewUserUsecase(userService services.UserService) UserUsecase {
	return &userUsecaseImpl{UserService: userService}
}

// UserFindByID retrieves a user by ID
func (u *userUsecaseImpl) UserFindByID(ctx context.Context, ID string) (*dto.UserDTO, error) {
	//Get User
	res, err := u.UserService.GetUserByID(ctx, ID)
	if err != nil {
		return nil, err
	}
	resp := dto.NewUser(res)
	return resp, nil
}
