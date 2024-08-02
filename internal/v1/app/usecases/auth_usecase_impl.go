package usecases

import (
	"context"
	"fmt"
	"gofiber-boilerplatev3/internal/v1/app/dto"
	"gofiber-boilerplatev3/internal/v1/domain/services"
	"gofiber-boilerplatev3/pkg/utils/auth/jwt"
	"gofiber-boilerplatev3/pkg/utils/msg"
	"time"
)

// authUsecaseImpl implements the AuthUsecase interface
type authUsecaseImpl struct {
	userService services.UserService
}

// NewAuthUsecase creates a new instance of authUsecaseImpl
func NewAuthUsecase(userService services.UserService) AuthUsecase {
	return &authUsecaseImpl{userService: userService}
}

// AuthRegister creates a new user and returns the created user
func (u *authUsecaseImpl) AuthRegister(ctx context.Context, req *dto.RegisterDTO) (*jwt.Token, error) {

	msg.Validate(req)
	// Create a new Model User
	res := dto.NewRegisterUser(req)
	//Create User
	res, err := u.userService.Create(ctx, res)
	if err != nil {
		return nil, err
	}
	//Transform To User
	resp := dto.NewTokenUser(res)
	// Generate JWT
	tokenAccess, err := jwt.GenerateAccessToken(resp)
	if err != nil {
		panic(err.Error())
	}
	tokenRefresh, err := jwt.GenerateRefreshToken(resp.ID)
	if err != nil {
		panic(err.Error())
	}

	response := jwt.Token{
		AccessToken:  tokenAccess,
		RefreshToken: tokenRefresh,
	}
	return &response, nil
}

// AuthLogin retrieves a user Token
func (u *authUsecaseImpl) AuthLogin(ctx context.Context, req *dto.LoginDTO) (*jwt.Token, error) {
	msg.Validate(req)
	// Get a new Model User
	res := dto.NewLoginUser(req)
	//Get User By Email
	res, err := u.userService.GetUserByEmail(ctx, res)
	if err != nil {
		panic(err.Error())
	}
	// Transform User to New Token User
	resp := dto.NewTokenUser(res)
	// Generate JWT
	tokenAccess, err := jwt.GenerateAccessToken(resp)
	if err != nil {
		panic(err.Error())
	}
	tokenRefresh, err := jwt.GenerateRefreshToken(resp.ID)
	if err != nil {
		panic(err.Error())
	}
	response := jwt.Token{
		AccessToken:  tokenAccess,
		RefreshToken: tokenRefresh,
	}
	//Update Login
	go func() {
		now := time.Now()
		res.LastLogin = &now
		if err := u.userService.Update(ctx, res); err != nil {
			// Handle error (log it, etc.)
			fmt.Printf("Error updating user login: %v\n", err)
		}
	}()

	return &response, nil
}

// RefreshToken retrieves a user refresh token
func (u *authUsecaseImpl) RefreshToken(ctx context.Context, req *dto.RefreshTokenDTO) (*jwt.Token, error) {
	// Validation
	msg.Validate(req)

	// Validate Refresh token
	token, err := jwt.ValidateRefreshToken(req.RefreshToken)
	if err != nil {
		return nil, err
	}
	//Get User By ID
	res, err := u.userService.GetUserByID(ctx, token.ID)
	if err != nil {
		return nil, err
	}
	resp := dto.NewTokenUser(res)
	tokenAccess, err := jwt.GenerateAccessToken(resp)
	if err != nil {
		panic(err.Error())
	}
	tokenRefresh, err := jwt.GenerateRefreshToken(resp.ID)
	if err != nil {
		panic(err.Error())
	}
	response := jwt.Token{
		AccessToken:  tokenAccess,
		RefreshToken: tokenRefresh,
	}
	return &response, nil
}
