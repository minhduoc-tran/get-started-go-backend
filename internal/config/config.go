package config

import (
	"os"
)

// Config - Application configuration
type Config struct {
	// Server
	Port string
	Env  string // "development", "production"

	// Database
	DBHost string
	DBPort string
	DBName string
	DBUser string
	DBPassword string
}

// LoadConfig - Load configuration from environment variables
func LoadConfig() *Config {
	return &Config{
		Port: getEnv("PORT", "8080"),
		Env:  getEnv("APP_ENV", "development"),

		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBName:     getEnv("DB_NAME", "myapp"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", ""),
	}
}

// getEnv - Get environment variable with default value
func getEnv(key string, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// IsDevelopment - Check if running in development mode
func (c *Config) IsDevelopment() bool {
	return c.Env == "development"
}

// IsProduction - Check if running in production mode
func (c *Config) IsProduction() bool {
	return c.Env == "production"
}
