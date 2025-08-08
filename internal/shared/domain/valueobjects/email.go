package valueobjects

import (
	"fmt"
	"regexp"
)

type Email struct {
	Value string `json:"value"`
}

func NewEmail(email string) (*Email, error) {
	if !isValidEmail(email) {
		return nil, fmt.Errorf("invalid email format: %s", email)
	}
	return &Email{Value: email}, nil
}

func (e Email) String() string {
	return e.Value
}

func isValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}
