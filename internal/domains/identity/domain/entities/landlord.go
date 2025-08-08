package entities

import (
	"time"

	"tenant-management/internal/shared/domain/valueobjects"
)

type Landlord struct {
	ID        uint                  `json:"id" gorm:"primaryKey"`
	GoogleID  string                `json:"google_id" gorm:"uniqueIndex;not null"`
	Email     valueobjects.Email    `json:"email" gorm:"embedded;embedPrefix:email_"`
	Name      string                `json:"name" gorm:"not null"`
	Phone     string                `json:"phone"`
	Picture   string                `json:"picture"`
	Company   string                `json:"company"`
	Address   *valueobjects.Address `json:"address" gorm:"embedded;embedPrefix:address_"`
	CreatedAt time.Time             `json:"created_at"`
	UpdatedAt time.Time             `json:"updated_at"`
}

func NewLandlord(googleID, email, name string) (*Landlord, error) {
	emailVO, err := valueobjects.NewEmail(email)
	if err != nil {
		return nil, err
	}

	return &Landlord{
		GoogleID:  googleID,
		Email:     *emailVO,
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func (l *Landlord) UpdateProfile(name, phone, company string, address *valueobjects.Address) {
	l.Name = name
	l.Phone = phone
	l.Company = company
	l.Address = address
	l.UpdatedAt = time.Now()
}
