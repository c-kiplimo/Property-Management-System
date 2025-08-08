package queries

type GetPropertiesByLandlordQuery struct {
	LandlordID uint `json:"landlord_id"`
}

func NewGetPropertiesByLandlordQuery(landlordID uint) *GetPropertiesByLandlordQuery {
	return &GetPropertiesByLandlordQuery{
		LandlordID: landlordID,
	}
}
