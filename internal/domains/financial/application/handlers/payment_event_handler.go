package handlers

import (
	"time"

	"tenant-management/internal/domains/financial/domain/entities"
	"tenant-management/internal/domains/financial/domain/repositories"
	leasingEvents "tenant-management/internal/domains/leasing/domain/events"
	"tenant-management/internal/shared/domain/events"
	"tenant-management/internal/shared/domain/valueobjects"
)

type PaymentEventHandler struct {
	paymentRepo repositories.PaymentRepository
	leaseRepo   repositories.LeaseRepository // We'll need this interface
}

func NewPaymentEventHandler(
	paymentRepo repositories.PaymentRepository,
	leaseRepo repositories.LeaseRepository,
) *PaymentEventHandler {
	return &PaymentEventHandler{
		paymentRepo: paymentRepo,
		leaseRepo:   leaseRepo,
	}
}

func (h *PaymentEventHandler) Handle(event events.DomainEvent) error {
	switch e := event.(type) {
	case *leasingEvents.LeaseSignedEvent:
		return h.handleLeaseSignedEvent(e)
	}
	return nil
}

func (h *PaymentEventHandler) handleLeaseSignedEvent(event *leasingEvents.LeaseSignedEvent) error {
	// Get lease details (in a real implementation, you might pass this in the event)
	// lease, err := h.leaseRepo.FindByID(event.LeaseID)
	// if err != nil {
	//     return err
	// }

	// Generate monthly rent payments for the lease duration
	// This is a simplified version - in reality you'd get the actual lease details

	// For now, create the first month's rent payment
	rent, _ := valueobjects.NewMoney(1000.00, "USD") // This should come from the lease

	payment := entities.NewPayment(
		event.LeaseID,
		event.TenantID,
		*rent,
		entities.PaymentTypeRent,
		time.Now().AddDate(0, 1, 0), // Due next month
	)

	return h.paymentRepo.Save(payment)
}
