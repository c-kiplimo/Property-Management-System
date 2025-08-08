package events

import (
	"fmt"
	"tenant-management/internal/shared/domain/events"
)

type PropertyCreatedEvent struct {
	events.BaseEvent
	PropertyID uint `json:"property_id"`
	LandlordID uint `json:"landlord_id"`
}

func NewPropertyCreatedEvent(propertyID, landlordID uint) *PropertyCreatedEvent {
	return &PropertyCreatedEvent{
		BaseEvent:  events.NewBaseEvent("property.created", fmt.Sprintf("%d", propertyID)),
		PropertyID: propertyID,
		LandlordID: landlordID,
	}
}
