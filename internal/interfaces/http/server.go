package http

import (
	"tenant-management/internal/domains/identity/application/handlers"
	identityHandlers "tenant-management/internal/domains/identity/application/handlers"
	identityRepos "tenant-management/internal/domains/identity/infrastructure/repositories"
	identityHttp "tenant-management/internal/domains/identity/interfaces/http"

	propertyHandlers "tenant-management/internal/domains/property/application/handlers"
	propertyRepos "tenant-management/internal/domains/property/infrastructure/repositories"
	propertyHttp "tenant-management/internal/domains/property/interfaces/http"

	"tenant-management/internal/shared/domain/events"
	"tenant-management/internal/shared/infrastructure/config"
	"tenant-management/internal/shared/infrastructure/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	router          *gin.Engine
	eventDispatcher *events.EventDispatcher
}

func NewServer(db *gorm.DB, cfg *config.Config) *Server {
	// Initialize event dispatcher
	eventDispatcher := events.NewEventDispatcher()

	// Initialize repositories
	landlordRepo := identityRepos.NewGormLandlordRepository(db)
	propertyRepo := propertyRepos.NewGormPropertyRepository(db)
	unitRepo := propertyRepos.NewGormUnitRepository(db)

	// Initialize handlers
	identityCommandHandler := identityHandlers.NewIdentityCommandHandler(landlordRepo, eventDispatcher)
	identityQueryHandler := identityHandlers.NewIdentityQueryHandler(landlordRepo)

	propertyCommandHandler := propertyHandlers.NewPropertyCommandHandler(propertyRepo, unitRepo, eventDispatcher)
	propertyQueryHandler := propertyHandlers.NewPropertyQueryHandler(propertyRepo, unitRepo)

	// Initialize HTTP handlers
	authHandler := identityHttp.NewAuthHandler(identityCommandHandler, identityQueryHandler, cfg)
	propertyHandler := propertyHttp.NewPropertyHandler(propertyCommandHandler, propertyQueryHandler)

	// Initialize middleware
	authMiddleware := middleware.NewAuthMiddleware(cfg.JWTSecret)
	corsMiddleware := middleware.NewCORSMiddleware()

	// Setup router
	router := gin.Default()
	router.Use(corsMiddleware.Handle())

	server := &Server{
		router:          router,
		eventDispatcher: eventDispatcher,
	}

	server.setupRoutes(authHandler, propertyHandler, authMiddleware)
	return server
}

func (s *Server) Router() *gin.Engine {
	return s.router
}

func (s *Server) setupRoutes(
	authHandler *identityHttp.AuthHandler,
	propertyHandler *propertyHttp.PropertyHandler,
	authMiddleware *middleware.AuthMiddleware,
) {
	// Health check
	s.router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok", "service": "Property Management System"})
	})

	// Auth routes (public)
	auth := s.router.Group("/api/auth")
	{
		auth.GET("/login", authHandler.Login)
		auth.GET("/callback", authHandler.Callback)
	}

	// Protected routes
	api := s.router.Group("/api")
	api.Use(authMiddleware.Authenticate())
	{
		// Identity
		api.GET("/me", authHandler.Me)

		// Property management
		api.POST("/properties", propertyHandler.CreateProperty)
		api.GET("/properties", propertyHandler.GetProperties)
		api.POST("/properties/:id/units", propertyHandler.CreateUnit)
	}
}
