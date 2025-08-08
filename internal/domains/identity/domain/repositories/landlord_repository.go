package repositories

import (
	_ "gorm.io/gorm"
	"tenant-management/internal/domains/identity/domain/entities"
)

type LandlordRepository interface {
	Save(landlord *entities.Landlord) error
	FindByID(id uint) (*entities.Landlord, error)
	FindByGoogleID(googleID string) (*entities.Landlord, error)
	FindByEmail(email string) (*entities.Landlord, error)
}
