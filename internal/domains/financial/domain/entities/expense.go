package entities

import (
	"time"

	"tenant-management/internal/shared/domain/valueobjects"
)

type ExpenseCategory string

const (
	ExpenseCategoryMaintenance ExpenseCategory = "maintenance"
	ExpenseCategoryInsurance   ExpenseCategory = "insurance"
	ExpenseCategoryTax         ExpenseCategory = "tax"
	ExpenseCategoryMortgage    ExpenseCategory = "mortgage"
	ExpenseCategoryUtility     ExpenseCategory = "utility"
	ExpenseCategoryOther       ExpenseCategory = "other"
)

type Expense struct {
	ID              uint               `json:"id" gorm:"primaryKey"`
	PropertyID      uint               `json:"property_id" gorm:"not null"`
	Amount          valueobjects.Money `json:"amount" gorm:"embedded;embedPrefix:amount_"`
	Category        ExpenseCategory    `json:"category" gorm:"not null"`
	Description     string             `json:"description" gorm:"not null"`
	Date            time.Time          `json:"date" gorm:"not null"`
	Vendor          string             `json:"vendor"`
	Receipt         string             `json:"receipt"`
	IsTaxDeductible bool               `json:"is_tax_deductible" gorm:"default:false"`
	CreatedAt       time.Time          `json:"created_at"`
	UpdatedAt       time.Time          `json:"updated_at"`
}

func NewExpense(propertyID uint, amount valueobjects.Money, category ExpenseCategory, description string, date time.Time) *Expense {
	return &Expense{
		PropertyID:  propertyID,
		Amount:      amount,
		Category:    category,
		Description: description,
		Date:        date,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func (e *Expense) MarkAsTaxDeductible() {
	e.IsTaxDeductible = true
	e.UpdatedAt = time.Now()
}
