package usecases

import (
	"context"
	"gofiber-boilerplatev3/internal/v1/app/dto"
	"gofiber-boilerplatev3/internal/v1/domain/services"
	"gofiber-boilerplatev3/pkg/utils/msg"
)

// userUsecaseImpl implements the UserUsecase interface
type userUsecaseImpl struct {
	userService services.UserService
}

// NewUserUsecase creates a new instance of userUsecaseImpl
func NewUserUsecase(userService services.UserService) UserUsecase {
	return &userUsecaseImpl{userService: userService}
}

// CreateUser creates a new user and returns the created user
func (u *userUsecaseImpl) RegisterUser(ctx context.Context, req *dto.UserRegisterDTO) (*dto.UserTokenDTO, error) {

	msg.Validate(req)
	// Create a new Model User
	res := dto.NewRegisterUser(req)
	//Create User
	res, err := u.userService.Create(ctx, res)
	if err != nil {
		return nil, err
	}

	resp := dto.NewUserTokenDTO(res)
	return resp, nil
}

//
//// GetUserByID retrieves a user by ID
//func (u *userUsecaseImpl) GetUserByID(id uint) (*models.User, error) {
//	return u.userService.GetUserByID(id)
//}
//
//// UpdateUser updates the details of an existing user
//func (u *userUsecaseImpl) UpdateUser(id uint, username, email string) (*models.User, error) {
//	return u.userService.UpdateUser(id, username, email)
//}
//
//// DeleteUser deletes a user by ID
//func (u *userUsecaseImpl) DeleteUser(id uint) error {
//	return u.userService.DeleteUser(id)
//}
