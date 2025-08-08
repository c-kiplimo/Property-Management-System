package entities

import (
	"time"

	"tenant-management/internal/shared/domain/valueobjects"
)

type Priority string
type RequestStatus string
type MaintenanceCategory string

const (
	PriorityLow       Priority = "low"
	PriorityMedium    Priority = "medium"
	PriorityHigh      Priority = "high"
	PriorityEmergency Priority = "emergency"

	RequestStatusOpen       RequestStatus = "open"
	RequestStatusInProgress RequestStatus = "in_progress"
	RequestStatusCompleted  RequestStatus = "completed"
	RequestStatusCancelled  RequestStatus = "cancelled"

	CategoryPlumbing   MaintenanceCategory = "plumbing"
	CategoryElectrical MaintenanceCategory = "electrical"
	CategoryHVAC       MaintenanceCategory = "hvac"
	CategoryGeneral    MaintenanceCategory = "general"
)

type MaintenanceRequest struct {
	ID          uint                `json:"id" gorm:"primaryKey"`
	PropertyID  uint                `json:"property_id" gorm:"not null"`
	UnitID      *uint               `json:"unit_id"`
	TenantID    *uint               `json:"tenant_id"`
	Title       string              `json:"title" gorm:"not null"`
	Description string              `json:"description" gorm:"type:text;not null"`
	Priority    Priority            `json:"priority" gorm:"default:'medium'"`
	Status      RequestStatus       `json:"status" gorm:"default:'open'"`
	Category    MaintenanceCategory `json:"category"`
	Cost        *valueobjects.Money `json:"cost" gorm:"embedded;embedPrefix:cost_"`
	AssignedTo  string              `json:"assigned_to"`
	RequestDate time.Time           `json:"request_date" gorm:"not null"`
	CompletedAt *time.Time          `json:"completed_at"`
	CreatedAt   time.Time           `json:"created_at"`
	UpdatedAt   time.Time           `json:"updated_at"`
}

func NewMaintenanceRequest(propertyID uint, title, description string) *MaintenanceRequest {
	return &MaintenanceRequest{
		PropertyID:  propertyID,
		Title:       title,
		Description: description,
		Priority:    PriorityMedium,
		Status:      RequestStatusOpen,
		RequestDate: time.Now(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func (m *MaintenanceRequest) AssignTo(contractor string) {
	m.AssignedTo = contractor
	m.Status = RequestStatusInProgress
	m.UpdatedAt = time.Now()
}

func (m *MaintenanceRequest) Complete(cost *valueobjects.Money) {
	now := time.Now()
	m.Status = RequestStatusCompleted
	m.CompletedAt = &now
	m.Cost = cost
	m.UpdatedAt = now
}

func (m *MaintenanceRequest) Cancel() {
	m.Status = RequestStatusCancelled
	m.UpdatedAt = time.Now()
}

func (m *MaintenanceRequest) SetPriority(priority Priority) {
	m.Priority = priority
	m.UpdatedAt = time.Now()
}
