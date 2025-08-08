package main

import (
	"log"
	"net/http"
	"os"

	"property-management/internal/interfaces/http"
	"property-management/internal/shared/infrastructure/config"
	"property-management/internal/shared/infrastructure/database"

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

	// Initialize HTTP server
	server := http.NewServer(db, cfg)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Property Management System starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, server.Router()))
}
