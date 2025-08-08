package queries

type GetLandlordByGoogleIDQuery struct {
	GoogleID string `json:"google_id"`
}

func NewGetLandlordByGoogleIDQuery(googleID string) *GetLandlordByGoogleIDQuery {
	return &GetLandlordByGoogleIDQuery{
		GoogleID: googleID,
	}
}
