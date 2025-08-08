package entities

import (
	"time"

	"tenant-management/internal/shared/domain/valueobjects"
)

type PropertyType string

const (
	PropertyTypeApartment  PropertyType = "apartment"
	PropertyTypeHouse      PropertyType = "house"
	PropertyTypeCommercial PropertyType = "commercial"
)

type Property struct {
	ID            uint                 `json:"id" gorm:"primaryKey"`
	LandlordID    uint                 `json:"landlord_id" gorm:"not null"`
	Name          string               `json:"name" gorm:"not null"`
	Address       valueobjects.Address `json:"address" gorm:"embedded;embedPrefix:address_"`
	PropertyType  PropertyType         `json:"property_type" gorm:"not null"`
	Description   string               `json:"description"`
	PurchasePrice *valueobjects.Money  `json:"purchase_price" gorm:"embedded;embedPrefix:purchase_price_"`
	CurrentValue  *valueobjects.Money  `json:"current_value" gorm:"embedded;embedPrefix:current_value_"`
	PurchaseDate  *time.Time           `json:"purchase_date"`
	CreatedAt     time.Time            `json:"created_at"`
	UpdatedAt     time.Time            `json:"updated_at"`
	Units         []Unit               `json:"units,omitempty" gorm:"foreignKey:PropertyID"`
}

func NewProperty(landlordID uint, name string, address valueobjects.Address, propertyType PropertyType) *Property {
	return &Property{
		LandlordID:   landlordID,
		Name:         name,
		Address:      address,
		PropertyType: propertyType,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}

func (p *Property) UpdateValue(newValue *valueobjects.Money) {
	p.CurrentValue = newValue
	p.UpdatedAt = time.Now()
}

func (p *Property) AddUnit(unit *Unit) {
	unit.PropertyID = p.ID
	p.Units = append(p.Units, *unit)
}
