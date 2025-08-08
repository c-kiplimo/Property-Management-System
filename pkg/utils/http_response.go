package utils

import (
	"github.com/gin-gonic/gin"
)

type APIResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
	Message string      `json:"message,omitempty"`
}

func SuccessResponse(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, APIResponse{
		Success: true,
		Data:    data,
	})
}

func ErrorResponse(c *gin.Context, statusCode int, err error) {
	c.JSON(statusCode, APIResponse{
		Success: false,
		Error:   err.Error(),
	})
}

func MessageResponse(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, APIResponse{
		Success: true,
		Message: message,
	})
}
