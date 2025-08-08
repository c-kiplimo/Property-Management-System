package events

import (
	"fmt"
	"tenant-management/internal/shared/domain/events"
)

type UnitAddedEvent struct {
	events.BaseEvent
	UnitID     uint `json:"unit_id"`
	PropertyID uint `json:"property_id"`
}

func NewUnitAddedEvent(unitID, propertyID uint) *UnitAddedEvent {
	return &UnitAddedEvent{
		BaseEvent:  events.NewBaseEvent("unit.added", fmt.Sprintf("%d", unitID)),
		UnitID:     unitID,
		PropertyID: propertyID,
	}
}
