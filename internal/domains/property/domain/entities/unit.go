package entities

import (
	"time"

	"tenant-management/internal/shared/domain/valueobjects"
)

type Unit struct {
	ID          uint                `json:"id" gorm:"primaryKey"`
	PropertyID  uint                `json:"property_id" gorm:"not null"`
	UnitNumber  string              `json:"unit_number" gorm:"not null"`
	Bedrooms    int                 `json:"bedrooms"`
	Bathrooms   float32             `json:"bathrooms"`
	SquareFeet  int                 `json:"square_feet"`
	Rent        valueobjects.Money  `json:"rent" gorm:"embedded;embedPrefix:rent_"`
	Deposit     *valueobjects.Money `json:"deposit" gorm:"embedded;embedPrefix:deposit_"`
	IsOccupied  bool                `json:"is_occupied" gorm:"default:false"`
	Description string              `json:"description"`
	CreatedAt   time.Time           `json:"created_at"`
	UpdatedAt   time.Time           `json:"updated_at"`
}

func NewUnit(propertyID uint, unitNumber string, rent valueobjects.Money) *Unit {
	return &Unit{
		PropertyID: propertyID,
		UnitNumber: unitNumber,
		Rent:       rent,
		IsOccupied: false,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}

func (u *Unit) MarkAsOccupied() {
	u.IsOccupied = true
	u.UpdatedAt = time.Now()
}

func (u *Unit) MarkAsVacant() {
	u.IsOccupied = false
	u.UpdatedAt = time.Now()
}

func (u *Unit) UpdateRent(newRent valueobjects.Money) {
	u.Rent = newRent
	u.UpdatedAt = time.Now()
}
