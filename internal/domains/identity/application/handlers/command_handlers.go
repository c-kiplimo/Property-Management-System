package handlers

import (
	"tenant-management/internal/domains/identity/application/commands"
	"tenant-management/internal/domains/identity/domain/entities"
	"tenant-management/internal/domains/identity/domain/repositories"
	"tenant-management/internal/shared/domain/events"
)

type IdentityCommandHandler struct {
	landlordRepo    repositories.LandlordRepository
	eventDispatcher *events.EventDispatcher
}

func NewIdentityCommandHandler(
	landlordRepo repositories.LandlordRepository,
	eventDispatcher *events.EventDispatcher,
) *IdentityCommandHandler {
	return &IdentityCommandHandler{
		landlordRepo:    landlordRepo,
		eventDispatcher: eventDispatcher,
	}
}

func (h *IdentityCommandHandler) HandleRegisterLandlord(cmd *commands.RegisterLandlordCommand) (*entities.Landlord, error) {
	landlord, err := entities.NewLandlord(cmd.GoogleID, cmd.Email, cmd.Name)
	if err != nil {
		return nil, err
	}

	landlord.Picture = cmd.Picture

	if err := h.landlordRepo.Save(landlord); err != nil {
		return nil, err
	}

	// Dispatch event
	event := NewLandlordRegisteredEvent(landlord.ID, landlord.Email.Value)
	h.eventDispatcher.Dispatch(event)

	return landlord, nil
}
