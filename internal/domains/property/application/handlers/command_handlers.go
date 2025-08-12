package handlers

import (
	"fmt"

	"tenant-management/internal/domains/property/application/commands"
	"tenant-management/internal/domains/property/domain/entities"
	domainEvents "tenant-management/internal/domains/property/domain/events" // ✅ Import property domain events
	"tenant-management/internal/domains/property/domain/repositories"
	"tenant-management/internal/shared/domain/events"
)

type PropertyCommandHandler struct {
	propertyRepo    repositories.PropertyRepository
	unitRepo        repositories.UnitRepository
	eventDispatcher *events.EventDispatcher
}

func NewPropertyCommandHandler(
	propertyRepo repositories.PropertyRepository,
	unitRepo repositories.UnitRepository,
	eventDispatcher *events.EventDispatcher,
) *PropertyCommandHandler {
	return &PropertyCommandHandler{
		propertyRepo:    propertyRepo,
		unitRepo:        unitRepo,
		eventDispatcher: eventDispatcher,
	}
}

func (h *PropertyCommandHandler) HandleCreateProperty(cmd *commands.CreatePropertyCommand) (*entities.Property, error) {
	// Convert DTO to value objects
	address, err := cmd.Request.Address.ToValueObject()
	if err != nil {
		return nil, fmt.Errorf("invalid address: %w", err)
	}

	// Convert string to PropertyType enum
	var propertyType entities.PropertyType
	switch cmd.Request.PropertyType {
	case "apartment":
		propertyType = entities.PropertyTypeApartment
	case "house":
		propertyType = entities.PropertyTypeHouse
	case "commercial":
		propertyType = entities.PropertyTypeCommercial
	default:
		return nil, fmt.Errorf("invalid property type: %s", cmd.Request.PropertyType)
	}

	// Create property entity
	property := entities.NewProperty(cmd.LandlordID, cmd.Request.Name, *address, propertyType)
	property.Description = cmd.Request.Description
	property.PurchaseDate = cmd.Request.PurchaseDate

	if cmd.Request.PurchasePrice != nil {
		purchasePrice, err := cmd.Request.PurchasePrice.ToValueObject()
		if err != nil {
			return nil, fmt.Errorf("invalid purchase price: %w", err)
		}
		property.PurchasePrice = purchasePrice
	}

	if cmd.Request.CurrentValue != nil {
		currentValue, err := cmd.Request.CurrentValue.ToValueObject()
		if err != nil {
			return nil, fmt.Errorf("invalid current value: %w", err)
		}
		property.CurrentValue = currentValue
	}

	// Save property
	if err := h.propertyRepo.Save(property); err != nil {
		return nil, fmt.Errorf("failed to save property: %w", err)
	}

	// ✅ Use domain event from property events package
	event := domainEvents.NewPropertyCreatedEvent(property.ID, cmd.LandlordID)
	h.eventDispatcher.Dispatch(event)

	return property, nil
}

func (h *PropertyCommandHandler) HandleAddUnit(cmd *commands.AddUnitCommand) (*entities.Unit, error) {
	// Convert rent to value object
	rent, err := cmd.Request.Rent.ToValueObject()
	if err != nil {
		return nil, fmt.Errorf("invalid rent amount: %w", err)
	}

	// Create unit entity
	unit := entities.NewUnit(cmd.PropertyID, cmd.Request.UnitNumber, *rent)
	unit.Bedrooms = cmd.Request.Bedrooms
	unit.Bathrooms = cmd.Request.Bathrooms
	unit.SquareFeet = cmd.Request.SquareFeet
	unit.Description = cmd.Request.Description

	if cmd.Request.Deposit != nil {
		deposit, err := cmd.Request.Deposit.ToValueObject()
		if err != nil {
			return nil, fmt.Errorf("invalid deposit amount: %w", err)
		}
		unit.Deposit = deposit
	}

	// Save unit
	if err := h.unitRepo.Save(unit); err != nil {
		return nil, fmt.Errorf("failed to save unit: %w", err)
	}

	// ✅ Use domain event from property events package
	event := domainEvents.NewUnitAddedEvent(unit.ID, cmd.PropertyID)
	h.eventDispatcher.Dispatch(event)

	return unit, nil
}
