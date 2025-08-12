package repositories

import (
	"tenant-management/internal/domains/property/domain/entities"

	"gorm.io/gorm"
)

type UnitRepositoryImpl struct {
	db *gorm.DB
}

func NewUnitRepositoryImpl(db *gorm.DB) UnitRepository {
	return &UnitRepositoryImpl{db: db}
}

func (r *UnitRepositoryImpl) Save(unit *entities.Unit) error {
	return r.db.Save(unit).Error
}

func (r *UnitRepositoryImpl) FindByID(id uint) (*entities.Unit, error) {
	var unit entities.Unit
	err := r.db.First(&unit, id).Error
	if err != nil {
		return nil, err
	}
	return &unit, nil
}

func (r *UnitRepositoryImpl) FindByPropertyID(propertyID uint) ([]*entities.Unit, error) {
	var units []*entities.Unit
	err := r.db.Where("property_id = ?", propertyID).Find(&units).Error
	return units, err
}

func (r *UnitRepositoryImpl) FindVacantUnits(propertyID uint) ([]*entities.Unit, error) {
	var units []*entities.Unit
	err := r.db.Where("property_id = ? AND is_occupied = ?", propertyID, false).Find(&units).Error
	return units, err
}

func (r *UnitRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&entities.Unit{}, id).Error
}
