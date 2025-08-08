package dto

import (
	"tenant-management/internal/shared/domain/valueobjects"
	"time"
)

type CreatePropertyRequest struct {
	Name          string               `json:"name" binding:"required"`
	Address       CreateAddressRequest `json:"address" binding:"required"`
	PropertyType  string               `json:"property_type" binding:"required"`
	Description   string               `json:"description"`
	PurchasePrice *CreateMoneyRequest  `json:"purchase_price"`
	CurrentValue  *CreateMoneyRequest  `json:"current_value"`
	PurchaseDate  *time.Time           `json:"purchase_date"`
}

type CreateAddressRequest struct {
	Street  string `json:"street" binding:"required"`
	City    string `json:"city" binding:"required"`
	State   string `json:"state" binding:"required"`
	ZipCode string `json:"zip_code" binding:"required"`
}

type CreateMoneyRequest struct {
	Amount   float64 `json:"amount" binding:"required"`
	Currency string  `json:"currency"`
}

type CreateUnitRequest struct {
	UnitNumber  string              `json:"unit_number" binding:"required"`
	Bedrooms    int                 `json:"bedrooms"`
	Bathrooms   float32             `json:"bathrooms"`
	SquareFeet  int                 `json:"square_feet"`
	Rent        CreateMoneyRequest  `json:"rent" binding:"required"`
	Deposit     *CreateMoneyRequest `json:"deposit"`
	Description string              `json:"description"`
}

func (r *CreateAddressRequest) ToValueObject() (*valueobjects.Address, error) {
	return valueobjects.NewAddress(r.Street, r.City, r.State, r.ZipCode)
}

func (r *CreateMoneyRequest) ToValueObject() (*valueobjects.Money, error) {
	return valueobjects.NewMoney(r.Amount, r.Currency)
}
