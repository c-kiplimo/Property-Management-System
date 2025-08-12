package main

import (
	"log"
	"net/http" // keep standard http for ListenAndServe
	"os"

	httpserver "tenant-management/internal/interfaces/http" // alias to avoid conflict
	"tenant-management/internal/shared/infrastructure/config"
	"tenant-management/internal/shared/infrastructure/database"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Initialize configuration
	cfg := config.Load()

	// Initialize database
	db := database.Initialize(cfg.DatabaseURL)

	// Initialize HTTP server (our package)
	server := httpserver.NewServer(db, cfg)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Property Management System starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, server.Router()))
}
