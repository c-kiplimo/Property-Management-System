package entities

import (
	"time"

	"tenant-management/internal/shared/domain/valueobjects"
)

type PaymentType string
type PaymentStatus string
type PaymentMethod string

const (
	PaymentTypeRent    PaymentType = "rent"
	PaymentTypeDeposit PaymentType = "deposit"
	PaymentTypeLateFee PaymentType = "late_fee"
	PaymentTypeUtility PaymentType = "utility"

	PaymentStatusPending PaymentStatus = "pending"
	PaymentStatusPaid    PaymentStatus = "paid"
	PaymentStatusOverdue PaymentStatus = "overdue"

	PaymentMethodCash     PaymentMethod = "cash"
	PaymentMethodCheck    PaymentMethod = "check"
	PaymentMethodTransfer PaymentMethod = "transfer"
	PaymentMethodOnline   PaymentMethod = "online"
)

type Payment struct {
	ID          uint               `json:"id" gorm:"primaryKey"`
	LeaseID     uint               `json:"lease_id" gorm:"not null"`
	TenantID    uint               `json:"tenant_id" gorm:"not null"`
	Amount      valueobjects.Money `json:"amount" gorm:"embedded;embedPrefix:amount_"`
	PaymentType PaymentType        `json:"payment_type" gorm:"not null"`
	PaymentDate *time.Time         `json:"payment_date"`
	DueDate     time.Time          `json:"due_date" gorm:"not null"`
	Status      PaymentStatus      `json:"status" gorm:"default:'pending'"`
	Method      PaymentMethod      `json:"method"`
	Reference   string             `json:"reference"`
	Notes       string             `json:"notes"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
}

func NewPayment(leaseID, tenantID uint, amount valueobjects.Money, paymentType PaymentType, dueDate time.Time) *Payment {
	return &Payment{
		LeaseID:     leaseID,
		TenantID:    tenantID,
		Amount:      amount,
		PaymentType: paymentType,
		DueDate:     dueDate,
		Status:      PaymentStatusPending,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func (p *Payment) MarkAsPaid(method PaymentMethod, reference string) {
	now := time.Now()
	p.PaymentDate = &now
	p.Status = PaymentStatusPaid
	p.Method = method
	p.Reference = reference
	p.UpdatedAt = now
}

func (p *Payment) MarkAsOverdue() {
	if p.Status == PaymentStatusPending && time.Now().After(p.DueDate) {
		p.Status = PaymentStatusOverdue
		p.UpdatedAt = time.Now()
	}
}

func (p *Payment) IsOverdue() bool {
	return p.Status == PaymentStatusPending && time.Now().After(p.DueDate)
}
