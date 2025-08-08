package entities

import (
	"time"

	"tenant-management/internal/shared/domain/valueobjects"
)

type Tenant struct {
	ID               uint                `json:"id" gorm:"primaryKey"`
	FirstName        string              `json:"first_name" gorm:"not null"`
	LastName         string              `json:"last_name" gorm:"not null"`
	Email            valueobjects.Email  `json:"email" gorm:"embedded;embedPrefix:email_"`
	Phone            string              `json:"phone"`
	EmergencyContact string              `json:"emergency_contact"`
	EmergencyPhone   string              `json:"emergency_phone"`
	DateOfBirth      *time.Time          `json:"date_of_birth"`
	SSN              string              `json:"ssn"` // Store encrypted in production
	Employment       string              `json:"employment"`
	MonthlyIncome    *valueobjects.Money `json:"monthly_income" gorm:"embedded;embedPrefix:monthly_income_"`
	CreatedAt        time.Time           `json:"created_at"`
	UpdatedAt        time.Time           `json:"updated_at"`
}

func NewTenant(firstName, lastName, email string) (*Tenant, error) {
	emailVO, err := valueobjects.NewEmail(email)
	if err != nil {
		return nil, err
	}

	return &Tenant{
		FirstName: firstName,
		LastName:  lastName,
		Email:     *emailVO,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func (t *Tenant) UpdatePersonalInfo(firstName, lastName, phone string, dateOfBirth *time.Time) {
	t.FirstName = firstName
	t.LastName = lastName
	t.Phone = phone
	t.DateOfBirth = dateOfBirth
	t.UpdatedAt = time.Now()
}

func (t *Tenant) UpdateEmploymentInfo(employment string, monthlyIncome *valueobjects.Money) {
	t.Employment = employment
	t.MonthlyIncome = monthlyIncome
	t.UpdatedAt = time.Now()
}

func (t *Tenant) FullName() string {
	return t.FirstName + " " + t.LastName
}
