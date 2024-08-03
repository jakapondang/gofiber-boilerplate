package usecases

import (
	"context"
	"gofiber-boilerplatev3/internal/v1/app/dto"
	"gofiber-boilerplatev3/internal/v1/domain/services"
	"gofiber-boilerplatev3/pkg/utils/auth/jwt"
	"gofiber-boilerplatev3/pkg/utils/mailpack"
	"gofiber-boilerplatev3/pkg/utils/msg"
	"time"
)

// authUsecaseImpl implements the AuthUsecase interface
type authUsecaseImpl struct {
	UserService   services.UserService
	PasswordReset services.PasswordResetService
}

// NewAuthUsecase creates a new instance of authUsecaseImpl
func NewAuthUsecase(userService services.UserService, passwordReset services.PasswordResetService) AuthUsecase {
	return &authUsecaseImpl{
		UserService:   userService,
		PasswordReset: passwordReset,
	}
}

// AuthRegister creates a new user and returns the created user
func (u *authUsecaseImpl) AuthRegister(ctx context.Context, req *dto.RegisterDTO) (*jwt.Token, error) {

	msg.Validate(req)
	// Create a new Model User
	res := dto.NewRegisterUser(req)
	//Create User
	res, err := u.UserService.CreateUser(ctx, res)
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
	res, err := u.UserService.LoginUserByEmail(ctx, res)
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
		if err := u.UserService.UpdateUser(ctx, res); err != nil {
			// Handle error (log it, etc.)
			panic("Error Update User : " + err.Error())
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
		panic(err.Error())
	}
	//Get User By ID
	res, err := u.UserService.GetUserByID(ctx, token.ID)
	if err != nil {
		panic(err.Error())
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

// RefreshToken retrieves a user refresh token
func (u *authUsecaseImpl) ForgotPassword(ctx context.Context, req *dto.ForgotPasswordDTO) error {
	// Validation
	msg.Validate(req)
	// Get a new Model User
	res := dto.NewForgotPasswordUser(req)

	//Get User By Email
	user, err := u.UserService.GetUserByEmail(ctx, res)
	if err != nil {
		panic(msg.BadRequestError{
			Message: "Theres no such user email",
		})
	}

	req.UserID = user.ID.String()                 // Add userID
	req.ExpiresAt = time.Now().Add(1 * time.Hour) // expire at 1 hour

	// Transform to model & Create Password reset
	reset := dto.NewPasswordReset(req)
	resp, err := u.PasswordReset.CreatePasswordResets(ctx, reset)
	if err != nil {
		panic(err.Error())
	}

	mailData := dto.EmaiForgotPasswordDTO{
		Username:   user.Username,
		ResetToken: resp.ResetToken.String(),
		ResetLink:  "http://admin.id/confirm-password/activate?token=" + resp.ResetToken.String(),
		Year:       time.Now().Year(),
	}

	// Setup Email
	go func() {
		// Send Email
		err = mailpack.SendMail(res.Email, "admin Forgot password", mailData, "forgot_password.html")
		if err != nil {
			panic(
				"Failed Sent Email:" + err.Error(),
			)
		}
	}()

	return nil
}
