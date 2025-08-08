package entities

import (
	"fmt"
	"time"

	"tenant-management/internal/shared/domain/valueobjects"
)

type LeaseStatus string

const (
	LeaseStatusActive     LeaseStatus = "active"
	LeaseStatusExpired    LeaseStatus = "expired"
	LeaseStatusTerminated LeaseStatus = "terminated"
	LeaseStatusPending    LeaseStatus = "pending"
)

type Lease struct {
	ID              uint                `json:"id" gorm:"primaryKey"`
	UnitID          uint                `json:"unit_id" gorm:"not null"`
	TenantID        uint                `json:"tenant_id" gorm:"not null"`
	StartDate       time.Time           `json:"start_date" gorm:"not null"`
	EndDate         time.Time           `json:"end_date" gorm:"not null"`
	MonthlyRent     valueobjects.Money  `json:"monthly_rent" gorm:"embedded;embedPrefix:monthly_rent_"`
	SecurityDeposit *valueobjects.Money `json:"security_deposit" gorm:"embedded;embedPrefix:security_deposit_"`
	Status          LeaseStatus         `json:"status" gorm:"default:'pending'"`
	LeaseTerms      string              `json:"lease_terms" gorm:"type:text"`
	SignedDate      *time.Time          `json:"signed_date"`
	CreatedAt       time.Time           `json:"created_at"`
	UpdatedAt       time.Time           `json:"updated_at"`
}

func NewLease(unitID, tenantID uint, startDate, endDate time.Time, monthlyRent valueobjects.Money) (*Lease, error) {
	if startDate.After(endDate) {
		return nil, fmt.Errorf("start date cannot be after end date")
	}

	return &Lease{
		UnitID:      unitID,
		TenantID:    tenantID,
		StartDate:   startDate,
		EndDate:     endDate,
		MonthlyRent: monthlyRent,
		Status:      LeaseStatusPending,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}

func (l *Lease) Sign() error {
	if l.Status != LeaseStatusPending {
		return fmt.Errorf("can only sign pending leases")
	}

	now := time.Now()
	l.SignedDate = &now
	l.Status = LeaseStatusActive
	l.UpdatedAt = now
	return nil
}

func (l *Lease) Terminate() error {
	if l.Status != LeaseStatusActive {
		return fmt.Errorf("can only terminate active leases")
	}

	l.Status = LeaseStatusTerminated
	l.UpdatedAt = time.Now()
	return nil
}

func (l *Lease) IsActive() bool {
	now := time.Now()
	return l.Status == LeaseStatusActive && l.StartDate.Before(now) && l.EndDate.After(now)
}

func (l *Lease) IsExpiring(days int) bool {
	return time.Now().AddDate(0, 0, days).After(l.EndDate)
}
