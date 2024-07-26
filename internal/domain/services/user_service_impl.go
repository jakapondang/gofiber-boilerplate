package services

import (
	"errors"
	"gofiber-boilerplatev3/internal/domain/models"
	"gofiber-boilerplatev3/internal/domain/repositories"
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
func (s *UserServiceImpl) CreateUser(username, email string) (*models.User, error) {
	if username == "" || email == "" {
		return nil, errors.New("username and email cannot be empty")
	}

	// Check if a user with the same email already exists
	existingUser, _ := s.userRepo.FindByEmail(email)
	if existingUser != nil {
		return nil, errors.New("user with this email already exists")
	}

	// Create a new user entity
	user := models.NewUser(username, email)

	// Save the user entity in the repository
	if err := s.userRepo.Save(user); err != nil {
		return nil, err
	}

	return user, nil
}

// GetUserByID retrieves a user by ID from the repository
func (s *UserServiceImpl) GetUserByID(id uint) (*models.User, error) {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}

// UpdateUser updates the details of an existing user
func (s *UserServiceImpl) UpdateUser(id uint, username, email string) (*models.User, error) {
	if username == "" || email == "" {
		return nil, errors.New("username and email cannot be empty")
	}

	// Retrieve the existing user
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	// Update the user details
	user.Username = username
	user.Email = email

	// Save the updated user entity
	if err := s.userRepo.Save(user); err != nil {
		return nil, err
	}

	return user, nil
}

// DeleteUser deletes a user by ID
func (s *UserServiceImpl) DeleteUser(id uint) error {
	// Check if the user exists
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("user not found")
	}

	// Delete the user from the repository
	if err := s.userRepo.Delete(id); err != nil {
		return err
	}

	return nil
}
