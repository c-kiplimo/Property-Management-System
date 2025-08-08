package http

import (
	"net/http"
	"strconv"

	"tenant-management/internal/domains/property/application/commands"
	"tenant-management/internal/domains/property/application/handlers"
	"tenant-management/internal/domains/property/application/queries"
	"tenant-management/internal/shared/application/dto"
	"tenant-management/internal/shared/infrastructure/middleware"

	"github.com/gin-gonic/gin"
)

type PropertyHandler struct {
	commandHandler *handlers.PropertyCommandHandler
	queryHandler   *handlers.PropertyQueryHandler
}

func NewPropertyHandler(
	commandHandler *handlers.PropertyCommandHandler,
	queryHandler *handlers.PropertyQueryHandler,
) *PropertyHandler {
	return &PropertyHandler{
		commandHandler: commandHandler,
		queryHandler:   queryHandler,
	}
}

func (h *PropertyHandler) CreateProperty(c *gin.Context) {
	claims := c.MustGet("claims").(*middleware.Claims)

	var req dto.CreatePropertyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cmd := commands.NewCreatePropertyCommand(claims.LandlordID, req)
	property, err := h.commandHandler.HandleCreateProperty(cmd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, property)
}

func (h *PropertyHandler) GetProperties(c *gin.Context) {
	claims := c.MustGet("claims").(*middleware.Claims)

	query := queries.NewGetPropertiesByLandlordQuery(claims.LandlordID)
	properties, err := h.queryHandler.HandleGetPropertiesByLandlord(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, properties)
}

func (h *PropertyHandler) CreateUnit(c *gin.Context) {
	propertyID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid property ID"})
		return
	}

	var req dto.CreateUnitRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cmd := commands.NewAddUnitCommand(uint(propertyID), req)
	unit, err := h.commandHandler.HandleAddUnit(cmd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, unit)
}
