package usecases

import (
	"context"
	"gofiber-boilerplatev3/internal/v1/app/dto"
	"gofiber-boilerplatev3/internal/v1/domain/services"
)

// userUsecaseImpl implements the UserUsecase interface
type userUsecaseImpl struct {
	userService services.UserService
}

// NewUserUsecase creates a new instance of userUsecaseImpl
func NewUserUsecase(userService services.UserService) UserUsecase {
	return &userUsecaseImpl{userService: userService}
}

// GetUserByID retrieves a user by ID
func (u *userUsecaseImpl) UserFindByID(ctx context.Context, ID string) (*dto.UserTokenDTO, error) {

	//Create User
	res, err := u.userService.GetUserByID(ctx, ID)
	if err != nil {
		return nil, err
	}
	resp := dto.NewUserTokenDTO(res)
	return resp, nil
}
