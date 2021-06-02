package clubs

import "github.com/lc-tut/club-portal/utils"

type ClubPlace struct {
	PlaceID  uint32     `gorm:"type:int unsigned not null auto_increment;primaryKey"`
	Place    string     `gorm:"type:text;not null;unique"`
	ClubTime []ClubTime `gorm:"many2many:activity_details;foreignKey:PlaceID;joinForeignKey:PlaceID;references:TimeID;joinReferences:TimeID"`
}

type Places []ClubPlace

func (p Places) ToPlaceResponse(remarks []ClubRemark) []PlaceResponse {
	res := make([]PlaceResponse, len(p))

	for i, place := range p {
		placeRes := PlaceResponse{
			Place:   place.Place,
			Remarks: utils.NullStringToStringP(remarks[i].PlaceRemark),
		}
		res[i] = placeRes
	}

	return res
}

type PlaceResponse struct {
	Place   string  `json:"place"`
	Remarks *string `json:"remarks"`
}
