package valueobjects

import (
	"fmt"
	"math"
)

type Money struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

func NewMoney(amount float64, currency string) (*Money, error) {
	if amount < 0 {
		return nil, fmt.Errorf("money amount cannot be negative")
	}
	if currency == "" {
		currency = "USD"
	}
	return &Money{
		Amount:   math.Round(amount*100) / 100, // Round to 2 decimal places
		Currency: currency,
	}, nil
}

func (m Money) Add(other Money) (*Money, error) {
	if m.Currency != other.Currency {
		return nil, fmt.Errorf("cannot add different currencies")
	}
	return NewMoney(m.Amount+other.Amount, m.Currency)
}

func (m Money) Subtract(other Money) (*Money, error) {
	if m.Currency != other.Currency {
		return nil, fmt.Errorf("cannot subtract different currencies")
	}
	return NewMoney(m.Amount-other.Amount, m.Currency)
}
