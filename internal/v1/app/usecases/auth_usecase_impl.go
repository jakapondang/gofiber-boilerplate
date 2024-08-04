package usecases

import (
	"context"
	"github.com/google/uuid"
	"gofiber-boilerplatev3/internal/v1/app/dto"
	"gofiber-boilerplatev3/internal/v1/domain"
	"gofiber-boilerplatev3/internal/v1/domain/services"
	"gofiber-boilerplatev3/pkg/utils/auth/jwt"
	"gofiber-boilerplatev3/pkg/utils/mailpack"
	"gofiber-boilerplatev3/pkg/utils/msg"
	"gorm.io/gorm"
	"time"
)

// authUsecaseImpl implements the AuthUsecase interface
type authUsecaseImpl struct {
	TrxDomain            domain.TrxDomain
	UserService          services.UserService
	PasswordResetService services.PasswordResetService
}

// NewAuthUsecase creates a new instance of authUsecaseImpl
func NewAuthUsecase(trxDomain domain.TrxDomain, userService services.UserService, passwordResetService services.PasswordResetService) AuthUsecase {
	return &authUsecaseImpl{
		TrxDomain:            trxDomain,
		UserService:          userService,
		PasswordResetService: passwordResetService,
	}
}

func (u *authUsecaseImpl) AuthRegister(ctx context.Context, req *dto.RegisterDTO) (*jwt.Token, error) {
	msg.Validate(req)
	// Transform For register user
	res := dto.NewRegisterUser(req)
	var response *jwt.Token

	err := domain.WithTransaction(ctx, u.TrxDomain, func(tx *gorm.DB) error {
		var err error

		res, err = u.UserService.CreateUser(ctx, tx, res)
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

		response = &jwt.Token{
			AccessToken:  tokenAccess,
			RefreshToken: tokenRefresh,
		}

		return nil
	})

	if err != nil {
		panic(err.Error())
	}

	return response, nil
}

// AuthLogin retrieves a user Token
func (u *authUsecaseImpl) AuthLogin(ctx context.Context, req *dto.LoginDTO) (*jwt.Token, error) {
	msg.Validate(req)
	// Get a new Model User
	res := dto.NewLoginUser(req)
	var response *jwt.Token

	// Start Transaction
	err := domain.WithTransaction(ctx, u.TrxDomain, func(tx *gorm.DB) error {
		var err error
		//Get User By Email
		res, err = u.UserService.LoginUserByEmail(ctx, tx, res)
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
		response = &jwt.Token{
			AccessToken:  tokenAccess,
			RefreshToken: tokenRefresh,
		}
		// Update user last login
		now := time.Now()
		res.LastLogin = &now
		err = u.UserService.UpdateUser(ctx, tx, res)
		if err != nil {
			panic(err.Error())
		}
		return nil
	})
	if err != nil {
		panic(err.Error())
	}

	return response, nil
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
	var response *jwt.Token
	// Start Transaction
	err = domain.WithTransaction(ctx, u.TrxDomain, func(tx *gorm.DB) error {
		//Transform User ID to UUID
		userId, err := uuid.Parse(token.ID)
		if err != nil {
			panic(err.Error())
		}
		//Get User ID
		res, err := u.UserService.GetUserByID(ctx, tx, userId)
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
		response = &jwt.Token{
			AccessToken:  tokenAccess,
			RefreshToken: tokenRefresh,
		}
		return nil
	})
	if err != nil {
		panic(err.Error())
	}

	return response, nil
}

// PasswordResetRequest Request New Password Reset for user
func (u *authUsecaseImpl) PasswordResetRequest(ctx context.Context, req *dto.PasswordResetRequestDTO) error {
	// Validation
	msg.Validate(req)
	// Get a new Model Forgot Password
	res := dto.NewForgotPasswordUser(req)

	var mailData *dto.PasswordResetRequestEmailDTO
	// Start Transaction
	err := domain.WithTransaction(ctx, u.TrxDomain, func(tx *gorm.DB) error {
		//Get User By Email
		user, err := u.UserService.GetUserByEmail(ctx, tx, res)
		if err != nil {
			panic(msg.BadRequestError{
				Message: "Theres no such user email",
			})
		}
		req.UserID = user.ID.String()                 // Add userID
		req.ExpiresAt = time.Now().Add(1 * time.Hour) // expire at 1 hour

		// Transform to model & Create Password reset
		reset := dto.NewPasswordReset(req)
		resp, err := u.PasswordResetService.CreatePasswordResets(ctx, tx, reset)
		if err != nil {
			panic(err.Error())
		}

		// Setup Email Data
		mailData = &dto.PasswordResetRequestEmailDTO{
			Username:   user.Username,
			ResetToken: resp.ResetToken.String(),
			ResetLink:  "http://iinvite.id/confirm-password/activate?token=" + resp.ResetToken.String(),
			Year:       time.Now().Year(),
		}
		return nil
	})
	// Send Email Async
	go func() {
		err = mailpack.SendMail(res.Email, "IInvite Forgot password", mailData, "forgot_password.html")
		if err != nil {
			panic(
				"Failed Sent Email:" + err.Error(),
			)
		}
	}()

	return nil
}

// PasswordResetUpdate Request New Password Reset for user
func (u *authUsecaseImpl) PasswordResetUpdate(ctx context.Context, req *dto.PasswordResetUpdateRequestDTO) error {
	// Validation
	msg.Validate(req)
	// Get a new Model Forgot Password
	res := dto.NewPasswordResetUpdate(req)

	// Start Transaction
	err := domain.WithTransaction(ctx, u.TrxDomain, func(tx *gorm.DB) error {
		// Find User ID by token Reset
		res, err := u.PasswordResetService.FindByToken(ctx, tx, res)
		if err != nil {
			panic(err.Error())
		}

		// Find User By User ID
		user, err := u.UserService.GetUserByID(ctx, tx, res.UserID)
		if err != nil {
			panic(err.Error())
		}

		// Update password User
		user.PasswordHash = req.Password
		err = u.UserService.UpdatePasswordUser(ctx, tx, user)
		if err != nil {
			panic(err.Error())
		}

		// Mark Already used
		err = u.PasswordResetService.MarkAsUsed(ctx, tx, res)
		if err != nil {
			panic(err.Error())
		}
		// return success
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
