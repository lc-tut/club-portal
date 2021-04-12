package models

import "github.com/lc-tut/club-portal/utils"

type ClubPlace struct {
	PlaceID  uint32     `gorm:"type:int unsigned not null auto_increment;primaryKey"`
	Place    string     `gorm:"type:text;not null;unique"`
	ClubTime []ClubTime `gorm:"many2many:activity_details;foreignKey:PlaceID;joinForeignKey:PlaceID;references:TimeID;joinReferences:TimeID"`
}

func (cp *ClubPlace) GetPlaceID() uint32 {
	return cp.PlaceID
}

func (cp *ClubPlace) GetPlace() string {
	return cp.Place
}

type Places []ClubPlace

func (p Places) ToPlaceResponse(remarks []ClubRemark) *[]PlaceResponse {
	res := make([]PlaceResponse, len(p))

	for i, place := range p {
		placeRes := PlaceResponse{
			Place:   place.Place,
			Remarks: utils.ToNilOrString(remarks[i].PlaceRemarks),
		}
		res[i] = placeRes
	}

	return &res
}

type PlaceRequest struct {
	PlaceID uint32  `json:"place_id"`
	Place   string  `json:"place"`
	Remarks *string `json:"remarks"`
}

type PlaceResponse struct {
	Place   string  `json:"place"`
	Remarks *string `json:"remarks"`
}
