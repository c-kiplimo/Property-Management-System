package events

import (
	"fmt"
	"tenant-management/internal/shared/domain/events"
)

// PropertyCreatedEvent is triggered when a property is created.
type PropertyCreatedEvent struct {
	events.BaseEvent
	PropertyID uint `json:"property_id"`
	LandlordID uint `json:"landlord_id"`
}

// NewPropertyCreatedEvent creates a new PropertyCreatedEvent instance.
func NewPropertyCreatedEvent(propertyID, landlordID uint) PropertyCreatedEvent {
	return PropertyCreatedEvent{
		BaseEvent:  events.NewBaseEvent("PropertyCreated", fmt.Sprintf("%d", propertyID)),
		PropertyID: propertyID,
		LandlordID: landlordID,
	}
}
