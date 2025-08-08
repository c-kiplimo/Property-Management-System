package valueobjects

import "fmt"

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zip_code"`
}

func NewAddress(street, city, state, zipCode string) (*Address, error) {
	if street == "" || city == "" || state == "" || zipCode == "" {
		return nil, fmt.Errorf("all address fields are required")
	}
	return &Address{
		Street:  street,
		City:    city,
		State:   state,
		ZipCode: zipCode,
	}, nil
}

func (a Address) String() string {
	return fmt.Sprintf("%s, %s, %s %s", a.Street, a.City, a.State, a.ZipCode)
}
