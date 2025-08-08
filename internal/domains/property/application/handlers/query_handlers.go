package handlers

import (
	"tenant-management/internal/domains/property/application/queries"
	"tenant-management/internal/domains/property/domain/entities"
	"tenant-management/internal/domains/property/domain/repositories"
)

type PropertyQueryHandler struct {
	propertyRepo repositories.PropertyRepository
	unitRepo     repositories.UnitRepository
}

func NewPropertyQueryHandler(
	propertyRepo repositories.PropertyRepository,
	unitRepo repositories.UnitRepository,
) *PropertyQueryHandler {
	return &PropertyQueryHandler{
		propertyRepo: propertyRepo,
		unitRepo:     unitRepo,
	}
}

func (h *PropertyQueryHandler) HandleGetPropertiesByLandlord(query *queries.GetPropertiesByLandlordQuery) ([]*entities.Property, error) {
	properties, err := h.propertyRepo.FindByLandlordID(query.LandlordID)
	if err != nil {
		return nil, err
	}

	// Load units for each property
	for _, property := range properties {
		units, err := h.unitRepo.FindByPropertyID(property.ID)
		if err != nil {
			continue // Log error in production
		}

		// Convert slice of pointers to slice of values for GORM compatibility
		unitValues := make([]entities.Unit, len(units))
		for i, unit := range units {
			unitValues[i] = *unit
		}
		property.Units = unitValues
	}

	return properties, nil
}
