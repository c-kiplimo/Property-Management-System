package events

import (
	"fmt"
	"tenant-management/internal/shared/domain/events"
)

type LeaseSignedEvent struct {
	events.BaseEvent
	LeaseID  uint `json:"lease_id"`
	UnitID   uint `json:"unit_id"`
	TenantID uint `json:"tenant_id"`
}

func NewLeaseSignedEvent(leaseID, unitID, tenantID uint) *LeaseSignedEvent {
	return &LeaseSignedEvent{
		BaseEvent: events.NewBaseEvent("lease.signed", fmt.Sprintf("%d", leaseID)),
		LeaseID:   leaseID,
		UnitID:    unitID,
		TenantID:  tenantID,
	}
}
