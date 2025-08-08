package handlers

import (
	"tenant-management/internal/domains/identity/application/queries"
	"tenant-management/internal/domains/identity/domain/entities"
	"tenant-management/internal/domains/identity/domain/repositories"
)

type IdentityQueryHandler struct {
	landlordRepo repositories.LandlordRepository
}

func NewIdentityQueryHandler(landlordRepo repositories.LandlordRepository) *IdentityQueryHandler {
	return &IdentityQueryHandler{
		landlordRepo: landlordRepo,
	}
}

func (h *IdentityQueryHandler) HandleGetLandlordByGoogleID(query *queries.GetLandlordByGoogleIDQuery) (*entities.Landlord, error) {
	return h.landlordRepo.FindByGoogleID(query.GoogleID)
}
