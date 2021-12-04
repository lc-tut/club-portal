package clubs

import "database/sql"

type ActivityDetail struct {
	TimeID   uint32     `gorm:"type:int unsigned;not null;primaryKey"`
	PlaceID  uint32     `gorm:"type:int unsigned;not null;primaryKey"`
	ClubUUID string     `gorm:"type:char(36);not null"`
	Remarks  ClubRemark `gorm:"foreignKey:ClubUUID;references:ClubUUID"`
}

type DetailRelations struct {
	ClubUUID    string
	TimeID      uint32
	Date        string
	Time        string
	PlaceID     uint32
	Place       string
	RemarkID    uint32
	PlaceRemark sql.NullString
	TimeRemark  sql.NullString
}

type Relations []DetailRelations

func (r Relations) ToClubTime() []ClubTime {
	times := make([]ClubTime, len(r))

	for i, rel := range r {
		ct := ClubTime{
			TimeID: rel.TimeID,
			Date:   rel.Date,
			Time:   rel.Time,
		}
		times[i] = ct
	}

	return times
}

func (r Relations) ToClubPlace() []ClubPlace {
	places := make([]ClubPlace, len(r))

	for i, rel := range r {
		cp := ClubPlace{
			PlaceID: rel.PlaceID,
			Place:   rel.Place,
		}
		places[i] = cp
	}

	return places
}

func (r Relations) ToClubRemark() []ClubRemark {
	remarks := make([]ClubRemark, len(r))

	for i, rel := range r {
		cr := ClubRemark{
			RemarkID:    rel.RemarkID,
			ClubUUID:    rel.ClubUUID,
			PlaceID:     rel.PlaceID,
			TimeID:      rel.TimeID,
			PlaceRemark: rel.PlaceRemark,
			TimeRemark:  rel.TimeRemark,
		}
		remarks[i] = cr
	}

	return remarks
}

type ActivityDetailRequest struct {
	TimeID      uint32  `json:"time_id"`
	Date        string  `json:"date"`
	Time        string  `json:"time"`
	TimeRemark  *string `json:"time_remark"`
	PlaceID     uint32  `json:"place_id"`
	Place       string  `json:"place"`
	PlaceRemark *string `json:"place_remark"`
}
