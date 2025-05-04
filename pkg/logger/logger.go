package logger

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

// Logger is a global instance of the logrus logger
var Logger *logrus.Logger

// InitLogger initializes the logger with custom settings.
func InitLogger() {
	Logger = logrus.New()
	Logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		PrettyPrint:     false,
	})
	Logger.SetOutput(os.Stdout)

	// Set log level based on environment
	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		logLevel = "info"
	}

	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		// Default to info if parsing fails
		level = logrus.InfoLevel
	}
	Logger.SetLevel(level)

	// Log that the logger has been initialized
	Logger.Info("Logger initialized with level: " + logLevel)
}

// LogRequest logs request details (method, path, status)
func LogRequest(method, path string, status int) {
	Logger.WithFields(logrus.Fields{
		"method": method,
		"path":   path,
		"status": status,
	}).Info("Request")
}

// LogError logs an error with custom message
func LogError(message string, err error) {
	Logger.WithFields(logrus.Fields{
		"error": err,
	}).Error(message)
}

// LogInfo logs an informational message
func LogInfo(message string) {
	Logger.Info(message)
}

// LogWarning logs a warning message
func LogWarning(message string) {
	Logger.Warning(message)
}

// LogDebug logs a debug message
func LogDebug(message string) {
	Logger.Debug(message)
}

// LogFatal logs a fatal error message and exits
func LogFatal(message string, err error) {
	Logger.WithFields(logrus.Fields{
		"error": err,
	}).Fatal(message)
}

// GinLoggerMiddleware returns a gin middleware for logging HTTP requests
func GinLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		startTime := time.Now()

		// Process request
		c.Next()

		// Stop timer
		endTime := time.Now()
		latency := endTime.Sub(startTime)

		// Get request details
		clientIP := c.ClientIP()
		method := c.Request.Method
		path := c.Request.URL.Path
		statusCode := c.Writer.Status()

		// Log request
		Logger.WithFields(logrus.Fields{
			"status":    statusCode,
			"latency":   latency,
			"clientIP":  clientIP,
			"method":    method,
			"path":      path,
			"userAgent": c.Request.UserAgent(),
			"errors":    c.Errors.String(),
		}).Info("Request")
	}
}
