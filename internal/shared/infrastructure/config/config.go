package config

import (
	"os"
)

type Config struct {
	DatabaseURL        string
	GoogleClientID     string
	GoogleClientSecret string
	GoogleRedirectURL  string
	JWTSecret          string
	Port               string
}

func Load() *Config {
	return &Config{
		DatabaseURL:        getEnv("DATABASE_URL", "host=localhost user=postgres password=postgres dbname=property_management port=5432 sslmode=disable"),
		GoogleClientID:     getEnv("GOOGLE_CLIENT_ID", ""),
		GoogleClientSecret: getEnv("GOOGLE_CLIENT_SECRET", ""),
		GoogleRedirectURL:  getEnv("GOOGLE_REDIRECT_URL", "http://localhost:8080/api/auth/callback"),
		JWTSecret:          getEnv("JWT_SECRET", "your-super-secret-jwt-key-change-in-production"),
		Port:               getEnv("PORT", "8080"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
