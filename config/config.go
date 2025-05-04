package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	JWTSecret string
	Port      string
)

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found or failed to load it. Falling back to system env.")
	}

	JWTSecret = os.Getenv("JWT_SECRET")
	if JWTSecret == "" {
		log.Println("JWT_SECRET not set in environment, using default (not safe for production)")
		JWTSecret = "default_secret"
	}

	Port = os.Getenv("PORT")
	if Port == "" {
		Port = "9090"
	}

}
