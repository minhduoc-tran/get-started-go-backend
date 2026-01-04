package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/minhduoc-tran/get-started-go-backend/internal/config"
	"github.com/minhduoc-tran/get-started-go-backend/internal/container"
	"github.com/minhduoc-tran/get-started-go-backend/internal/routes"
	"github.com/minhduoc-tran/get-started-go-backend/pkg/logger"
)

func main() {
	// 1. Load configuration
	cfg := config.LoadConfig()

	// 2. Initialize logger
	logger.Init()
	logger.Info("=== Application Starting ===")
	logger.Infof("Environment: %s", cfg.Env)

	// 3. Set Gin mode based on environment
	if cfg.IsProduction() {
		gin.SetMode(gin.ReleaseMode)
		logger.Info("Gin mode: Release (Production)")
	} else {
		gin.SetMode(gin.DebugMode)
		logger.Info("Gin mode: Debug (Development)")
	}

	// 4. Initialize dependency container
	logger.Info("Initializing dependencies...")
	c := container.NewContainer()

	// 5. Create Gin router
	r := gin.Default()

	// 6. Setup all routes
	logger.Info("Setting up routes...")
	routes.SetupRoutes(r, c)

	// 7. Start server
	logger.Info("=== Starting Server ===")
	startServer(cfg, r)
}

// startServer - Start HTTP server
func startServer(cfg *config.Config, r *gin.Engine) {
	addr := fmt.Sprintf(":%s", cfg.Port)
	logger.Infof("Server listening on http://localhost:%s", cfg.Port)

	if err := r.Run(addr); err != nil {
		logger.Errorf("Server error: %v", err)
		panic(err)
	}
}
