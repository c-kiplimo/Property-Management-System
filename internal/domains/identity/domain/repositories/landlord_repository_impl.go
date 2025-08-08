package repositories

import (
	"gorm.io/gorm"
	"tenant-management/internal/domains/identity/domain/entities"
)

type LandlordRepositoryImpl struct {
	db *gorm.DB
}

func NewLandlordRepositoryImpl(db *gorm.DB) LandlordRepository {
	return &LandlordRepositoryImpl{db: db}
}

func (r *LandlordRepositoryImpl) Save(landlord *entities.Landlord) error {
	return r.db.Save(landlord).Error
}

func (r *LandlordRepositoryImpl) FindByID(id uint) (*entities.Landlord, error) {
	var landlord entities.Landlord
	err := r.db.First(&landlord, id).Error
	if err != nil {
		return nil, err
	}
	return &landlord, nil
}

func (r *LandlordRepositoryImpl) FindByGoogleID(googleID string) (*entities.Landlord, error) {
	var landlord entities.Landlord
	err := r.db.Where("google_id = ?", googleID).First(&landlord).Error
	if err != nil {
		return nil, err
	}
	return &landlord, nil
}

func (r *LandlordRepositoryImpl) FindByEmail(email string) (*entities.Landlord, error) {
	var landlord entities.Landlord
	err := r.db.Where("email_value = ?", email).First(&landlord).Error
	if err != nil {
		return nil, err
	}
	return &landlord, nil
}
