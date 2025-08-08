package repositories

import (
	"gorm.io/gorm"
	"tenant-management/internal/domains/property/domain/entities"
)

type PropertyRepositoryImpl struct {
	db *gorm.DB
}

func NewPropertyRepository(db *gorm.DB) PropertyRepository {
	return &PropertyRepositoryImpl{db: db}
}

func (r *PropertyRepositoryImpl) Save(property *entities.Property) error {
	return r.db.Save(property).Error
}

func (r *PropertyRepositoryImpl) FindByID(id uint) (*entities.Property, error) {
	var property entities.Property
	err := r.db.Preload("Units").First(&property, id).Error
	if err != nil {
		return nil, err
	}
	return &property, nil
}

func (r *PropertyRepositoryImpl) FindByLandlordID(landlordID uint) ([]*entities.Property, error) {
	var properties []*entities.Property
	err := r.db.Where("landlord_id = ?", landlordID).Find(&properties).Error
	return properties, err
}

func (r *PropertyRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&entities.Property{}, id).Error
}
