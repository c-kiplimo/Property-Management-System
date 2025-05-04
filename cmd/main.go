package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"tenant-management/cmd/api/routes"
	"tenant-management/cmd/internal/user/repositories"
	"tenant-management/cmd/internal/user/service"
	_ "tenant-management/cmd/internal/user/service"
	"tenant-management/cmd/internal/user/usecase"
	"tenant-management/config"
	"tenant-management/pkg/logger"
)

func main() {
	// Initialize logger
	logger.InitLogger()

	// Load config
	cfg, err := config.LoadConfigFromEnv()
	if err != nil {
		log.Fatalf("Failed to load DB config: %v", err)
	}

	// Initialize DB client
	dbClient, err := config.NewDatabaseClient(context.Background(), cfg)
	if err != nil {
		log.Fatalf("Failed to initialize DB client: %v", err)
	}
	defer dbClient.Close()

	// Apply migrations
	if err := dbClient.ApplyMigrations(); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	// Initialize repositories and services
	userRepo := repositories.NewUserRepository(dbClient.DB)
	jwtService := service.NewJWTService(cfg.User) // Pass secret from config or env

	// Setup usecase
	authUsecase := usecase.NewAuthUsecase(userRepo, jwtService)

	// Setup Gin
	r := gin.Default()

	// Register routes
	routes.RegisterUserRoutes(r, authUsecase)

	// Run server
	port := cfg.Port
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
