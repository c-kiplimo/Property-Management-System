package repositories

import (
	"tenant-management/internal/domains/financial/domain/entities"
	"time"
)

type PaymentRepository interface {
	Save(payment *entities.Payment) error
	FindByID(id uint) (*entities.Payment, error)
	FindByLandlordID(landlordID uint) ([]*entities.Payment, error)
	FindOverduePayments(landlordID uint) ([]*entities.Payment, error)
	FindByLeaseID(leaseID uint) ([]*entities.Payment, error)
	FindByTenantID(tenantID uint) ([]*entities.Payment, error)
	FindByDateRange(startDate, endDate time.Time) ([]*entities.Payment, error)
	Delete(id uint) error
}

type ExpenseRepository interface {
	Save(expense *entities.Expense) error
	FindByID(id uint) (*entities.Expense, error)
	FindByPropertyID(propertyID uint) ([]*entities.Expense, error)
	FindByLandlordID(landlordID uint) ([]*entities.Expense, error)
	FindByDateRange(startDate, endDate time.Time) ([]*entities.Expense, error)
	FindTaxDeductibleExpenses(landlordID uint) ([]*entities.Expense, error)
	Delete(id uint) error
}

// We'll also need this for the event handler
type LeaseRepository interface {
	FindByID(id uint) (interface{}, error) // Using interface{} for now
}
