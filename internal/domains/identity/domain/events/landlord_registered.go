package events

import (
	"fmt"
	"tenant-management/internal/shared/domain/events"
)

type LandlordRegisteredEvent struct {
	events.BaseEvent
	LandlordID uint   `json:"landlord_id"`
	Email      string `json:"email"`
}

func NewLandlordRegisteredEvent(landlordID uint, email string) *LandlordRegisteredEvent {
	return &LandlordRegisteredEvent{
		BaseEvent:  events.NewBaseEvent("landlord.registered", fmt.Sprintf("%d", landlordID)),
		LandlordID: landlordID,
		Email:      email,
	}
}
