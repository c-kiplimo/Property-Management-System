package domain

type Tenant struct {
	ID           string
	Email        string
	MobileNumber string
	Names        string
	UnitID       string
}

type Unit struct {
	Rooms      int
	Floor      int
	UnitNumber string
}

type Landlord struct {
	Names        string
	Idno         string
	Email        string
	MobileNumber string
}

type Caretaker struct {
	Name   string
	Email  string
	Mobile string
	Idno   string
}

type TenantData struct {
	CurrentMeterReading float64
	PreviousReading     float64
	Consumption         float64
	Arrears             float64
	PaidPreviousMonth   float64
	PaidCurrentMonth    float64
}
