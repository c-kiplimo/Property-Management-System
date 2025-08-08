package commands

import (
	"tenant-management/internal/shared/application/dto"
)

type CreatePropertyCommand struct {
	LandlordID uint                      `json:"landlord_id"`
	Request    dto.CreatePropertyRequest `json:"request"`
}

func NewCreatePropertyCommand(landlordID uint, request dto.CreatePropertyRequest) *CreatePropertyCommand {
	return &CreatePropertyCommand{
		LandlordID: landlordID,
		Request:    request,
	}
}
