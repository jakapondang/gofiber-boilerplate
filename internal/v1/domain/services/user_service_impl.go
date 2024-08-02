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

	// Check if a user with the same mail already exists
	existingUser, _ := s.userRepo.FindByEmail(ctx, res.Email)
	if existingUser != nil {
		return nil, errors.New("user with this mail already exists")
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

// GetUserByEmail retrieves a user by Email from the repository
func (s *UserServiceImpl) GetUserByEmail(ctx context.Context, req *models.User) (*models.User, error) {
	// Check if a user with the same mail already exists
	res, err := s.userRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		return nil, errors.New("User mail not found")
	}
	// Check User Active
	if !res.IsActive {
		return nil, errors.New("User is not active , please contact administrator")
	}
	checkPass := auth.CheckPasswordHash(req.PasswordHash, res.PasswordHash)
	if !checkPass {
		return nil, errors.New("User mail and password doesnt match")
	}
	return res, nil
}

// GetUserByID retrieves a user by ID from the repository
func (s *UserServiceImpl) GetUserByID(ctx context.Context, ID string) (*models.User, error) {
	// Check if a user with the same mail already exists
	res, err := s.userRepo.FindByID(ctx, ID)
	if err != nil {
		return nil, errors.New("User not found")
	}

	return res, nil
}

// Update update a user and persists it in the repository
func (s *UserServiceImpl) Update(ctx context.Context, res *models.User) error {
	// Save the user entity in the repository
	if err := s.userRepo.Update(ctx, res); err != nil {
		return err
	}
	return nil
}
