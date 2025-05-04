package config

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

// DBConfig holds database configuration
type DBConfig struct {
	Host            string
	Port            string
	User            string
	Password        string
	DBName          string
	SSLMode         string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
	MigrationsPath  string
}

// DatabaseClient wraps the database connection and provides methods for interaction
type DatabaseClient struct {
	DB     *sql.DB
	Logger *zap.Logger
	Config DBConfig
}

// LoadConfigFromEnv loads database configuration from environment variables
func LoadConfigFromEnv() (DBConfig, error) {
	// Load env variables from .env file if present
	if err := godotenv.Load(); err != nil {
		// This is not a critical error, just log it
	}

	config := DBConfig{
		Host:            getEnv("DB_HOST", "localhost"),
		Port:            getEnv("DB_PORT", "5432"),
		User:            getEnv("DB_USER", ""),
		Password:        getEnv("DB_PASSWORD", ""),
		DBName:          getEnv("DB_NAME", ""),
		SSLMode:         getEnv("DB_SSL_MODE", "disable"),
		MaxOpenConns:    getEnvAsInt("DB_MAX_OPEN_CONNS", 25),
		MaxIdleConns:    getEnvAsInt("DB_MAX_IDLE_CONNS", 25),
		ConnMaxLifetime: time.Duration(getEnvAsInt("DB_CONN_MAX_LIFETIME_SECONDS", 300)) * time.Second,
		MigrationsPath:  getEnv("DB_MIGRATIONS_PATH", "file://migrations"),
	}

	// Validate required fields
	if config.User == "" || config.DBName == "" {
		return config, errors.New("required database configuration missing: DB_USER or DB_NAME")
	}

	return config, nil
}

// NewDatabaseClient creates a new database client with the given configuration
func NewDatabaseClient(ctx context.Context, config DBConfig) (*DatabaseClient, error) {
	logger, _ := zap.NewProduction()
	client := &DatabaseClient{
		Logger: logger,
		Config: config,
	}

	// Connect to the database
	if err := client.Connect(ctx); err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return client, nil
}

// Connect establishes a connection to the database with retry logic
func (c *DatabaseClient) Connect(ctx context.Context) error {
	var db *sql.DB
	var err error

	// Build connection string
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Config.Host, c.Config.Port, c.Config.User, url.QueryEscape(c.Config.Password), c.Config.DBName, c.Config.SSLMode,
	)

	// Retry logic with exponential backoff
	maxRetries := 5
	for i := 0; i < maxRetries; i++ {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			c.Logger.Info("Attempting database connection", zap.Int("attempt", i+1))

			db, err = sql.Open("postgres", dsn)
			if err != nil {
				c.Logger.Error("Failed to open database connection", zap.Error(err))
				time.Sleep(time.Duration(1<<uint(i)) * time.Second) // Exponential backoff
				continue
			}

			// Configure connection pool
			db.SetMaxOpenConns(c.Config.MaxOpenConns)
			db.SetMaxIdleConns(c.Config.MaxIdleConns)
			db.SetConnMaxLifetime(c.Config.ConnMaxLifetime)

			// Test connection
			if err := db.PingContext(ctx); err != nil {
				c.Logger.Error("Failed to ping database", zap.Error(err))
				_ = db.Close()
				time.Sleep(time.Duration(1<<uint(i)) * time.Second) // Exponential backoff
				continue
			}

			c.DB = db
			c.Logger.Info("Successfully connected to database")
			return nil
		}
	}

	return fmt.Errorf("failed to connect to database after %d attempts: %w", maxRetries, err)
}

// Close closes the database connection
func (c *DatabaseClient) Close() error {
	if c.DB != nil {
		c.Logger.Info("Closing database connection")
		return c.DB.Close()
	}
	return nil
}

// ApplyMigrations runs database migrations
func (c *DatabaseClient) ApplyMigrations() error {
	c.Logger.Info("Applying database migrations", zap.String("path", c.Config.MigrationsPath))

	// Acquire a migration lock to prevent concurrent migrations
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Config.Host, c.Config.Port, c.Config.User, url.QueryEscape(c.Config.Password), c.Config.DBName, c.Config.SSLMode,
	)
	lockConn, err := sql.Open("postgres", dsn)
	if err != nil {
		return fmt.Errorf("failed to open lock connection: %w", err)
	}
	defer lockConn.Close()

	_, err = lockConn.Exec("SELECT pg_advisory_lock(42)")
	if err != nil {
		return fmt.Errorf("failed to acquire migration lock: %w", err)
	}
	defer lockConn.Exec("SELECT pg_advisory_unlock(42)")

	// Create migration driver
	driver, err := postgres.WithInstance(c.DB, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("could not create migration driver: %w", err)
	}

	// Initialize migrations
	m, err := migrate.NewWithDatabaseInstance(
		c.Config.MigrationsPath,
		"postgres",
		driver,
	)
	if err != nil {
		return fmt.Errorf("migration initialization failed: %w", err)
	}

	// Apply migrations
	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("migration failed: %w", err)
	}

	c.Logger.Info("Database migrations completed successfully")
	return nil
}

// Helper functions
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func getEnvAsInt(key string, fallback int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intVal, err := strconv.Atoi(value); err == nil {
			return intVal
		}
	}
	return fallback
}
