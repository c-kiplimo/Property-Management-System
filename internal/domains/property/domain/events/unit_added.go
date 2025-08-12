package events

import (
	"fmt"
	"tenant-management/internal/shared/domain/events"
)

// UnitAddedEvent is triggered when a unit is added to a property.
type UnitAddedEvent struct {
	events.BaseEvent
	UnitID     uint `json:"unit_id"`
	PropertyID uint `json:"property_id"`
}

// NewUnitAddedEvent creates a new UnitAddedEvent instance.
func NewUnitAddedEvent(unitID, propertyID uint) UnitAddedEvent {
	return UnitAddedEvent{
		BaseEvent:  events.NewBaseEvent("UnitAdded", fmt.Sprintf("%d", unitID)),
		UnitID:     unitID,
		PropertyID: propertyID,
	}
}
