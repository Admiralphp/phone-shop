// internal/config/config.go
package config

import (
	"os"
	"strconv"
)

// Config holds all configuration for the application
type Config struct {
	// Server configuration
	ServerPort string
	
	// Database configuration
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	
	// Pagination defaults
	DefaultPageSize int
	MaxPageSize     int
}

// NewConfig creates a new Config struct with values from environment variables
func NewConfig() *Config {
	// Set default values
	config := &Config{
		ServerPort:      "8080",
		DBHost:          "localhost",
		DBPort:          "5432",
		DBUser:          "postgres",
		DBPassword:      "postgres",
		DBName:          "product_service",
		DefaultPageSize: 20,
		MaxPageSize:     100,
	}
	
	// Override with environment variables if they exist
	if port := os.Getenv("SERVER_PORT"); port != "" {
		config.ServerPort = port
	}
	
	if dbHost := os.Getenv("DB_HOST"); dbHost != "" {
		config.DBHost = dbHost
	}
	
	if dbPort := os.Getenv("DB_PORT"); dbPort != "" {
		config.DBPort = dbPort
	}
	
	if dbUser := os.Getenv("DB_USER"); dbUser != "" {
		config.DBUser = dbUser
	}
	
	if dbPassword := os.Getenv("DB_PASSWORD"); dbPassword != "" {
		config.DBPassword = dbPassword
	}
	
	if dbName := os.Getenv("DB_NAME"); dbName != "" {
		config.DBName = dbName
	}
	
	if pageSizeStr := os.Getenv("DEFAULT_PAGE_SIZE"); pageSizeStr != "" {
		if pageSize, err := strconv.Atoi(pageSizeStr); err == nil {
			config.DefaultPageSize = pageSize
		}
	}
	
	if maxPageSizeStr := os.Getenv("MAX_PAGE_SIZE"); maxPageSizeStr != "" {
		if maxPageSize, err := strconv.Atoi(maxPageSizeStr); err == nil {
			config.MaxPageSize = maxPageSize
		}
	}
	
	return config
}
