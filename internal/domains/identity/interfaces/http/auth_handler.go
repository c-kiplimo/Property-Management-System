package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"tenant-management/internal/domains/identity/application/commands"
	"tenant-management/internal/domains/identity/application/handlers"
	"tenant-management/internal/domains/identity/application/queries"
	"tenant-management/internal/shared/infrastructure/config"
	"tenant-management/internal/shared/infrastructure/middleware"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"gorm.io/gorm"
)

type AuthHandler struct {
	commandHandler *handlers.IdentityCommandHandler
	queryHandler   *handlers.IdentityQueryHandler
	googleConfig   *oauth2.Config
	authMiddleware *middleware.AuthMiddleware
}

type GoogleUserInfo struct {
	ID      string `json:"id"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

func NewAuthHandler(
	commandHandler *handlers.IdentityCommandHandler,
	queryHandler *handlers.IdentityQueryHandler,
	cfg *config.Config,
) *AuthHandler {
	googleConfig := &oauth2.Config{
		ClientID:     cfg.GoogleClientID,
		ClientSecret: cfg.GoogleClientSecret,
		RedirectURL:  cfg.GoogleRedirectURL,
		Scopes:       []string{"openid", "profile", "email"},
		Endpoint:     google.Endpoint,
	}

	authMiddleware := middleware.NewAuthMiddleware(cfg.JWTSecret)

	return &AuthHandler{
		commandHandler: commandHandler,
		queryHandler:   queryHandler,
		googleConfig:   googleConfig,
		authMiddleware: authMiddleware,
	}
}

func (h *AuthHandler) Login(c *gin.Context) {
	state := "random-state-string" // In production, generate a secure random state
	url := h.googleConfig.AuthCodeURL(state, oauth2.AccessTypeOffline)
	c.JSON(http.StatusOK, gin.H{
		"auth_url": url,
	})
}

func (h *AuthHandler) Callback(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing authorization code"})
		return
	}

	// Exchange code for token
	token, err := h.googleConfig.Exchange(context.Background(), code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to exchange code: %v", err)})
		return
	}

	// Get user info from Google
	client := h.googleConfig.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to get user info: %v", err)})
		return
	}
	defer resp.Body.Close()

	var googleUser GoogleUserInfo
	if err := json.NewDecoder(resp.Body).Decode(&googleUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to decode user info: %v", err)})
		return
	}

	// Find existing landlord or create new one
	query := queries.NewGetLandlordByGoogleIDQuery(googleUser.ID)
	landlord, err := h.queryHandler.HandleGetLandlordByGoogleID(query)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// Create new landlord
			cmd := commands.NewRegisterLandlordCommand(googleUser.ID, googleUser.Email, googleUser.Name, googleUser.Picture)
			landlord, err = h.commandHandler.HandleRegisterLandlord(cmd)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to register landlord: %v", err)})
				return
			}
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Database error: %v", err)})
			return
		}
	}

	// Generate JWT token
	jwtToken, err := h.authMiddleware.GenerateToken(landlord.ID, landlord.Email.Value)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to generate token: %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"landlord": landlord,
		"token":    jwtToken,
	})
}

func (h *AuthHandler) Me(c *gin.Context) {
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userClaims := claims.(*middleware.Claims)
	c.JSON(http.StatusOK, gin.H{
		"landlord_id": userClaims.LandlordID,
		"email":       userClaims.Email,
	})
}
