package repository

import (
	"fmt"

	"github.com/minhduoc-tran/get-started-go-backend/pkg/logger"
)

// UserRepository - Struct chứa data access logic cho User
type UserRepository struct {
	// Sau này có thể thêm DB connection ở đây
	// db *sql.DB
}

// NewUserRepository - Factory function tạo UserRepository
func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

// GetAllUsers - Get all users from database
func (ur *UserRepository) GetAllUsers() (string, error) {
	logger.Info("Repository: GetAllUsers called")
	msg := "[Repository] Get all users from database"
	fmt.Println(msg)
	logger.Debug("Repository: Returning all users")
	return msg, nil
}

// GetUserByID - Get user by ID from database
func (ur *UserRepository) GetUserByID(id string) (string, error) {
	logger.Infof("Repository: GetUserByID called with id=%s", id)
	msg := fmt.Sprintf("[Repository] Get user by id = %s from database", id)
	fmt.Println(msg)
	logger.Debug("Repository: User found")
	return msg, nil
}

// CreateUser - Create user in database
func (ur *UserRepository) CreateUser() (string, error) {
	logger.Info("Repository: CreateUser called")
	msg := "[Repository] Create user in database"
	fmt.Println(msg)
	logger.Debug("Repository: User created successfully")
	return msg, nil
}

// UpdateUser - Update user in database
func (ur *UserRepository) UpdateUser(id string) (string, error) {
	logger.Infof("Repository: UpdateUser called with id=%s", id)
	msg := fmt.Sprintf("[Repository] Update user by id = %s in database", id)
	fmt.Println(msg)
	logger.Debug("Repository: User updated successfully")
	return msg, nil
}

// DeleteUser - Delete user from database
func (ur *UserRepository) DeleteUser(id string) (string, error) {
	logger.Infof("Repository: DeleteUser called with id=%s", id)
	msg := fmt.Sprintf("[Repository] Delete user by id = %s in database", id)
	fmt.Println(msg)
	logger.Debug("Repository: User deleted successfully")
	return msg, nil
}