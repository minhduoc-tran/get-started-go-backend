package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/minhduoc-tran/get-started-go-backend/internal/container"
	"github.com/minhduoc-tran/get-started-go-backend/pkg/logger"
)

// SetupRoutes - Khai báo tất cả routes (chỉ focus vào routing)
func SetupRoutes(r *gin.Engine, c *container.Container) {
	logger.Info("Routes: Setting up all routes...")
	
	// User routes - Group API v1
	setupUserRoutes(r, c)

	logger.Info("Routes: All routes setup completed")
}

// setupUserRoutes - User resource routes
func setupUserRoutes(r *gin.Engine, c *container.Container) {
	userRoutes := r.Group("/api/users")
	{
		userRoutes.GET("", c.UserHandler.GetAllUsers)
		userRoutes.GET("/:id", c.UserHandler.GetUserByID)
		userRoutes.POST("", c.UserHandler.CreateUser)
		userRoutes.PUT("/:id", c.UserHandler.UpdateUser)
		userRoutes.DELETE("/:id", c.UserHandler.DeleteUser)
	}
}
