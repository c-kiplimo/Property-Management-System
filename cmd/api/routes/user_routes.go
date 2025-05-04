package routes

import (
	"github.com/gin-gonic/gin"
	"tenant-management/cmd/internal/user/handler/http"
	"tenant-management/cmd/internal/user/usecase"
)

func RegisterUserRoutes(r *gin.Engine, userUC *usecase.AuthUsecase) { // Change to accept a pointer
	userHandler := http.NewAuthHandler(userUC)
	userGroup := r.Group("/users")
	{
		userGroup.POST("/register", userHandler.Register)
		userGroup.POST("/login", userHandler.Login)
	}
}
