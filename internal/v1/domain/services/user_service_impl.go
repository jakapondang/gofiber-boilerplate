package services

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"gofiber-boilerplatev3/internal/v1/domain/models"
	"gofiber-boilerplatev3/internal/v1/domain/repositories"
	"gofiber-boilerplatev3/pkg/utils/auth"
	"gorm.io/gorm"
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
func (s *UserServiceImpl) CreateUser(ctx context.Context, tx *gorm.DB, res *models.User) (*models.User, error) {

	// Check if a user with the same mailpack already exists
	existingUser, _ := s.userRepo.FindByEmail(ctx, tx, res.Email)
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
	if err := s.userRepo.Create(ctx, tx, res); err != nil {
		return nil, err
	}

	return res, nil
}

// LoginUserByEmail login a user by Email from the repository
func (s *UserServiceImpl) LoginUserByEmail(ctx context.Context, tx *gorm.DB, req *models.User) (*models.User, error) {
	// Check if a user with the same mailpack already exists
	res, err := s.userRepo.FindByEmail(ctx, tx, req.Email)
	if err != nil {
		return nil, errors.New("User email not found")
	}

	// Check User Active
	if !res.IsActive {
		return nil, errors.New("User is not active , please contact administrator")
	}

	// Check Password
	checkPass := auth.CheckPasswordHash(req.PasswordHash, res.PasswordHash)
	if !checkPass {
		return nil, errors.New("User email and password doesnt match")
	}
	return res, nil
}

// GetUserByID retrieves a user by ID from the repository
func (s *UserServiceImpl) GetUserByID(ctx context.Context, tx *gorm.DB, ID uuid.UUID) (*models.User, error) {
	// Check if a user with the same mailpack already exists
	res, err := s.userRepo.FindByID(ctx, tx, ID)
	if err != nil {
		return nil, errors.New("User not found")
	}

	return res, nil
}

// Update a user and persists it in the repository
func (s *UserServiceImpl) UpdateUser(ctx context.Context, tx *gorm.DB, res *models.User) error {
	// Save the user entity in the repository
	if err := s.userRepo.Update(ctx, tx, res); err != nil {
		return err
	}
	return nil
}

// GetUserByEmail retrieves a user by Email from the repository
func (s *UserServiceImpl) GetUserByEmail(ctx context.Context, tx *gorm.DB, req *models.User) (*models.User, error) {
	// Check if a user with the same mailpack already exists
	res, err := s.userRepo.FindByEmail(ctx, tx, req.Email)
	if err != nil {
		return nil, errors.New("User email not found")
	}

	return res, nil
}

// Update password user and persists it in the repository
func (s *UserServiceImpl) UpdatePasswordUser(ctx context.Context, tx *gorm.DB, res *models.User) error {
	//Create password hash
	passwordHash, err := auth.HashPassword(res.PasswordHash)
	if err != nil {
		return err
	}
	res.PasswordHash = passwordHash
	// Save the user entity in the repository
	if err := s.userRepo.Update(ctx, tx, res); err != nil {
		return err
	}
	return nil
}
