package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"tenant-management/cmd/api/routes"
	tenantRepo "tenant-management/cmd/internal/tenant/repositories"
	tenantUsecase "tenant-management/cmd/internal/tenant/usecase"
	userRepo "tenant-management/cmd/internal/user/repositories"
	userService "tenant-management/cmd/internal/user/service"
	userUsecase "tenant-management/cmd/internal/user/usecase"
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
	defer func(dbClient *config.DatabaseClient) {
		err := dbClient.Close()
		if err != nil {
			// Handle close error
		}
	}(dbClient)

	// Apply migrations
	if err := dbClient.ApplyMigrations(); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	// Initialize repositories and services for users
	userRepo := userRepo.NewUserRepository(dbClient.DB)
	jwtService := userService.NewJWTService(cfg.User)

	// Setup user usecases
	authUsecase := userUsecase.NewAuthUsecase(userRepo, jwtService)

	// Initialize repositories and services for tenants
	tenantRepo := tenantRepo.NewTenantRepository(dbClient.DB)
	tenantUsecase := tenantUsecase.NewTenantUseCase(tenantRepo)

	// Setup Gin
	r := gin.Default()

	// Register user routes
	routes.RegisterUserRoutes(r, authUsecase)

	// Register tenant routes
	routes.RegisterTenantRoutes(r, tenantUsecase) // Register routes for tenant use case

	// Run server
	port := cfg.Port
	if port == "" {
		port = "8080"
	}
	err = r.Run(":" + port)
	if err != nil {
		return
	}
}
