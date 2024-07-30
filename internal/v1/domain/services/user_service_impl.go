package services

import (
	"context"
	"errors"
	"gofiber-boilerplatev3/internal/v1/domain/models"
	"gofiber-boilerplatev3/internal/v1/domain/repositories"
	"gofiber-boilerplatev3/pkg/utils/auth"
)

// UserServiceImpl provides business logic related to users
type UserServiceImpl struct {
	userRepo repositories.UserRepository
}

// NewUserService creates a new instance of UserServiceImpl
func NewUserService(userRepo repositories.UserRepository) UserService {
	return &UserServiceImpl{userRepo: userRepo}
}

// CreateUser creates a new user and persists it in the repository
func (s *UserServiceImpl) Create(ctx context.Context, res *models.User) (*models.User, error) {

	// Check if a user with the same email already exists
	existingUser, _ := s.userRepo.FindByEmail(ctx, res.Email)
	if existingUser != nil {
		return nil, errors.New("user with this email already exists")
	}

	//Create password hash
	passwordHash, err := auth.HashPassword(res.PasswordHash)
	if err != nil {
		return nil, err
	}
	res.PasswordHash = passwordHash

	// Save the user entity in the repository
	if err := s.userRepo.Create(ctx, res); err != nil {
		return nil, err
	}

	return res, nil
}

// GetUserByEmail retrieves a user by ID from the repository
func (s *UserServiceImpl) GetUserByEmail(ctx context.Context, req *models.User) (*models.User, error) {
	// Check if a user with the same email already exists
	res, err := s.userRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		return nil, errors.New("User email not found")
	}
	// Check User Active
	if !res.IsActive {
		return nil, errors.New("User is not active , please contact administrator")
	}
	checkPass := auth.CheckPasswordHash(req.PasswordHash, res.PasswordHash)
	if !checkPass {
		return nil, errors.New("User email and password doesnt match")
	}
	return res, nil
}

//// UpdateUser updates the details of an existing user
//func (s *UserServiceImpl) UpdateUser(id uint, username, email string) (*models.User, error) {
//	if username == "" || email == "" {
//		return nil, errors.New("username and email cannot be empty")
//	}
//
//	// Retrieve the existing user
//	user, err := s.userRepo.FindByID(id)
//	if err != nil {
//		return nil, err
//	}
//	if user == nil {
//		return nil, errors.New("user not found")
//	}
//
//	// Update the user details
//	user.Username = username
//	user.Email = email
//
//	// Save the updated user entity
//	if err := s.userRepo.Save(user); err != nil {
//		return nil, err
//	}
//
//	return user, nil
//}
//
//// DeleteUser deletes a user by ID
//func (s *UserServiceImpl) DeleteUser(id uint) error {
//	// Check if the user exists
//	user, err := s.userRepo.FindByID(id)
//	if err != nil {
//		return err
//	}
//	if user == nil {
//		return errors.New("user not found")
//	}
//
//	// Delete the user from the repository
//	if err := s.userRepo.Delete(id); err != nil {
//		return err
//	}
//
//	return nil
//}
