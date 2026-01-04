package services

import (
	"github.com/minhduoc-tran/get-started-go-backend/internal/repository"
	"github.com/minhduoc-tran/get-started-go-backend/pkg/logger"
)

// UserService - Struct chứa business logic cho User
type UserService struct {
	userRepo *repository.UserRepository
}

// NewUserService - Factory function tạo UserService
func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

// GetAllUsers - Get all users
func (us *UserService) GetAllUsers() (string, error) {
	logger.Info("Service: GetAllUsers called")
	data, err := us.userRepo.GetAllUsers()
	if err != nil {
		logger.Errorf("Service: Error getting all users - %v", err)
		return "", err
	}
	logger.Info("Service: GetAllUsers completed successfully")
	return data, nil
}

// GetUserByID - Get user by ID
func (us *UserService) GetUserByID(id string) (string, error) {
	logger.Infof("Service: GetUserByID called with id=%s", id)
	data, err := us.userRepo.GetUserByID(id)
	if err != nil {
		logger.Errorf("Service: Error getting user %s - %v", id, err)
		return "", err
	}
	logger.Infof("Service: GetUserByID for id=%s completed successfully", id)
	return data, nil
}

// CreateUser - Create user
func (us *UserService) CreateUser() (string, error) {
	logger.Info("Service: CreateUser called")
	data, err := us.userRepo.CreateUser()
	if err != nil {
		logger.Errorf("Service: Error creating user - %v", err)
		return "", err
	}
	logger.Info("Service: CreateUser completed successfully")
	return data, nil
}

// UpdateUser - Update user
func (us *UserService) UpdateUser(id string) (string, error) {
	logger.Infof("Service: UpdateUser called with id=%s", id)
	data, err := us.userRepo.UpdateUser(id)
	if err != nil {
		logger.Errorf("Service: Error updating user %s - %v", id, err)
		return "", err
	}
	logger.Infof("Service: UpdateUser for id=%s completed successfully", id)
	return data, nil
}

// DeleteUser - Delete user
func (us *UserService) DeleteUser(id string) (string, error) {
	logger.Infof("Service: DeleteUser called with id=%s", id)
	data, err := us.userRepo.DeleteUser(id)
	if err != nil {
		logger.Errorf("Service: Error deleting user %s - %v", id, err)
		return "", err
	}
	logger.Infof("Service: DeleteUser for id=%s completed successfully", id)
	return data, nil
}