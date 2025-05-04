package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tenant-management/cmd/internal/user/domain"
	"tenant-management/cmd/internal/user/usecase"
	"tenant-management/pkg/idgen"
	"time"
)

type AuthHandler struct {
	usecase *usecase.AuthUsecase
}

func NewAuthHandler(u *usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{usecase: u}
}

func (h *AuthHandler) Register(c *gin.Context) {
	type req struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Mobile   string `json:"mobile_number"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}
	var r req
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := domain.User{
		ID:           idgen.GenerateID(),
		Name:         r.Name,
		Email:        r.Email,
		MobileNumber: r.Mobile,
		Role:         domain.Role(r.Role),
		CreatedAt:    time.Now().Format(time.RFC3339),
	}
	err := h.usecase.Register(user, r.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "registered successfully"})
}

func (h *AuthHandler) Login(c *gin.Context) {
	type loginReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var req loginReq
	_ = c.ShouldBindJSON(&req)
	token, err := h.usecase.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}
