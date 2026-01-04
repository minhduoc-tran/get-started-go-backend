package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/minhduoc-tran/get-started-go-backend/internal/services"
	"github.com/minhduoc-tran/get-started-go-backend/pkg/logger"
)

// UserHandler - Struct chứa handler functions cho User
type UserHandler struct {
	userService *services.UserService
}

// NewUserHandler - Factory function tạo UserHandler
func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// GetAllUsers - GET /api/users
func (uh *UserHandler) GetAllUsers(c *gin.Context) {
	logger.Info("Handler: GetAllUsers called")
	
	message, err := uh.userService.GetAllUsers()
	if err != nil {
		logger.Errorf("Handler: GetAllUsers error - %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  err.Error(),
			"status": "error",
		})
		return
	}

	logger.Info("Handler: GetAllUsers returning success")
	c.JSON(http.StatusOK, gin.H{
		"message": message,
		"status":  "success",
	})
}

// GetUserByID - GET /api/users/:id
func (uh *UserHandler) GetUserByID(c *gin.Context) {
	idParam := c.Param("id")
	logger.Infof("Handler: GetUserByID called with id=%s", idParam)
	
	message, err := uh.userService.GetUserByID(idParam)
	if err != nil {
		logger.Errorf("Handler: GetUserByID error - %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  err.Error(),
			"status": "error",
		})
		return
	}

	logger.Infof("Handler: GetUserByID returning success for id=%s", idParam)
	c.JSON(http.StatusOK, gin.H{
		"message": message,
		"status":  "success",
	})
}

// CreateUser - POST /api/users
func (uh *UserHandler) CreateUser(c *gin.Context) {
	logger.Info("Handler: CreateUser called")
	
	message, err := uh.userService.CreateUser()
	if err != nil {
		logger.Errorf("Handler: CreateUser error - %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  err.Error(),
			"status": "error",
		})
		return
	}

	logger.Info("Handler: CreateUser returning success")
	c.JSON(http.StatusCreated, gin.H{
		"message": message,
		"status":  "success",
	})
}

// UpdateUser - PUT /api/users/:id
func (uh *UserHandler) UpdateUser(c *gin.Context) {
	idParam := c.Param("id")
	logger.Infof("Handler: UpdateUser called with id=%s", idParam)
	
	message, err := uh.userService.UpdateUser(idParam)
	if err != nil {
		logger.Errorf("Handler: UpdateUser error - %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  err.Error(),
			"status": "error",
		})
		return
	}

	logger.Infof("Handler: UpdateUser returning success for id=%s", idParam)
	c.JSON(http.StatusOK, gin.H{
		"message": message,
		"status":  "success",
	})
}

// DeleteUser - DELETE /api/users/:id
func (uh *UserHandler) DeleteUser(c *gin.Context) {
	idParam := c.Param("id")
	logger.Infof("Handler: DeleteUser called with id=%s", idParam)
	
	message, err := uh.userService.DeleteUser(idParam)
	if err != nil {
		logger.Errorf("Handler: DeleteUser error - %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  err.Error(),
			"status": "error",
		})
		return
	}

	logger.Infof("Handler: DeleteUser returning success for id=%s", idParam)
	c.JSON(http.StatusOK, gin.H{
		"message": message,
		"status":  "success",
	})
}