package commands

import (
	"tenant-management/internal/shared/domain/valueobjects"
	"time"
)

type CreateLeaseCommand struct {
	UnitID          uint                `json:"unit_id"`
	TenantID        uint                `json:"tenant_id"`
	StartDate       time.Time           `json:"start_date"`
	EndDate         time.Time           `json:"end_date"`
	MonthlyRent     valueobjects.Money  `json:"monthly_rent"`
	SecurityDeposit *valueobjects.Money `json:"security_deposit"`
	LeaseTerms      string              `json:"lease_terms"`
}

func NewCreateLeaseCommand(
	unitID, tenantID uint,
	startDate, endDate time.Time,
	monthlyRent valueobjects.Money,
	securityDeposit *valueobjects.Money,
	leaseTerms string,
) *CreateLeaseCommand {
	return &CreateLeaseCommand{
		UnitID:          unitID,
		TenantID:        tenantID,
		StartDate:       startDate,
		EndDate:         endDate,
		MonthlyRent:     monthlyRent,
		SecurityDeposit: securityDeposit,
		LeaseTerms:      leaseTerms,
	}
}
