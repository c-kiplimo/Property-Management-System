package events

import (
	"fmt"
	"tenant-management/internal/shared/domain/events"
)

// LandlordRegisteredEvent is triggered when a landlord registers in the system.
type LandlordRegisteredEvent struct {
	events.BaseEvent
	LandlordID uint   `json:"landlord_id"`
	Email      string `json:"email"`
}

// NewLandlordRegisteredEvent creates a new LandlordRegisteredEvent instance.
func NewLandlordRegisteredEvent(landlordID uint, email string) LandlordRegisteredEvent {
	return LandlordRegisteredEvent{
		BaseEvent:  events.NewBaseEvent("LandlordRegistered", fmt.Sprintf("%d", landlordID)),
		LandlordID: landlordID,
		Email:      email,
	}
}
