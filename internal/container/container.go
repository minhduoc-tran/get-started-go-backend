package container

import (
	"github.com/minhduoc-tran/get-started-go-backend/internal/handlers"
	"github.com/minhduoc-tran/get-started-go-backend/internal/repository"
	"github.com/minhduoc-tran/get-started-go-backend/internal/services"
	"github.com/minhduoc-tran/get-started-go-backend/pkg/logger"
)

// Container - Holds all application dependencies
type Container struct {
	// Repositories
	UserRepository *repository.UserRepository

	// Services
	UserService *services.UserService

	// Handlers
	UserHandler *handlers.UserHandler
}

// NewContainer - Initialize and return all dependencies
func NewContainer() *Container {
	logger.Info("Container: Initializing dependencies...")

	// Initialize repositories
	userRepository := repository.NewUserRepository()

	// Initialize services
	userService := services.NewUserService(userRepository)

	// Initialize handlers
	userHandler := handlers.NewUserHandler(userService)

	logger.Info("Container: All dependencies initialized successfully")

	return &Container{
		UserRepository: userRepository,
		UserService:    userService,
		UserHandler:    userHandler,
	}
}
