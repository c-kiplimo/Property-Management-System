package commands

type RegisterLandlordCommand struct {
	GoogleID string `json:"google_id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Picture  string `json:"picture"`
}

func NewRegisterLandlordCommand(googleID, email, name, picture string) *RegisterLandlordCommand {
	return &RegisterLandlordCommand{
		GoogleID: googleID,
		Email:    email,
		Name:     name,
		Picture:  picture,
	}
}
