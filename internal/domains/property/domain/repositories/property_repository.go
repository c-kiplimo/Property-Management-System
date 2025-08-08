package repositories

import (
	"tenant-management/internal/domains/property/domain/entities"
)

type PropertyRepository interface {
	Save(property *entities.Property) error
	FindByID(id uint) (*entities.Property, error)
	FindByLandlordID(landlordID uint) ([]*entities.Property, error)
	Delete(id uint) error
}

type UnitRepository interface {
	Save(unit *entities.Unit) error
	FindByID(id uint) (*entities.Unit, error)
	FindByPropertyID(propertyID uint) ([]*entities.Unit, error)
	FindVacantUnits(propertyID uint) ([]*entities.Unit, error)
	Delete(id uint) error
}
