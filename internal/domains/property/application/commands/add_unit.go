package commands

import (
	"tenant-management/internal/shared/application/dto"
)

type AddUnitCommand struct {
	PropertyID uint                  `json:"property_id"`
	Request    dto.CreateUnitRequest `json:"request"`
}

func NewAddUnitCommand(propertyID uint, request dto.CreateUnitRequest) *AddUnitCommand {
	return &AddUnitCommand{
		PropertyID: propertyID,
		Request:    request,
	}
}
