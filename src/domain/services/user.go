package services

import (
	"bm-users/src/domain/repositories/user"
	"bm-users/src/entities"
	"errors"
	"fmt"
	"go.uber.org/zap"
)

// UserService is an implementation of the UserService interface.
type UserService struct {
	userRepo user.UserRepository
	logger   *zap.Logger
}

// NewUserService creates a new instance of UserService.
func NewUserService(userRepo user.UserRepository, logger *zap.Logger) *UserService {
	return &UserService{
		userRepo: userRepo,
		logger:   logger,
	}
}

// CreateUser adds a new user.
func (s *UserService) CreateUser(user *entities.User) error {
	if user == nil {
		return errors.New("user cannot be nil")
	}

	if err := s.userRepo.Create(user); err != nil {
		s.logger.Error("Failed to create user", zap.Error(err))
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

// GetUserByID retrieves a user by its ID.
func (s *UserService) GetUserByID(id int64) (*entities.User, error) {
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		s.logger.Error("Failed to get user by ID", zap.Int64("id", id), zap.Error(err))
		return nil, fmt.Errorf("failed to get user by ID: %w", err)
	}

	return user, nil
}

// GetAllUsers retrieves all users.
func (s *UserService) GetAllUsers() ([]*entities.User, error) {
	users, err := s.userRepo.GetAll()
	if err != nil {
		s.logger.Error("Failed to get all users", zap.Error(err))
		return nil, fmt.Errorf("failed to get all users: %w", err)
	}

	return users, nil
}

// UpdateUser modifies an existing user.
func (s *UserService) UpdateUser(user *entities.User) error {
	if user == nil {
		return errors.New("user cannot be nil")
	}

	if err := s.userRepo.Update(user); err != nil {
		s.logger.Error("Failed to update user", zap.Error(err))
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}

// DeleteUser removes an user by its ID.
func (s *UserService) DeleteUser(id int) error {
	if err := s.userRepo.Delete(id); err != nil {
		s.logger.Error("Failed to delete user", zap.Int("id", id), zap.Error(err))
		return fmt.Errorf("failed to delete user: %w", err)
	}

	return nil
}
