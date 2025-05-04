package domain

type Role string

const (
	RoleSuperAdmin Role = "SUPER_ADMIN"
	RoleLandlord   Role = "LANDLORD"
	RoleCaretaker  Role = "CARETAKER"
)

type User struct {
	ID           string
	Name         string
	Email        string
	MobileNumber string
	PasswordHash string
	Role         Role
	CreatedAt    string
}
