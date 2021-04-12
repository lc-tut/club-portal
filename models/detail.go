package models

type ClubTime struct {
	TimeID    uint32      `gorm:"type:int unsigned not null auto_increment;primaryKey"`
	Date      string      `gorm:"type:varchar(3);not null;unique"`
	Time      string      `gorm:"type:varchar(255);not null;unique"`
	ClubPlace []ClubPlace `gorm:"many2many:activity_details;foreignKey:TimeID;joinForeignKey:TimeID;references:PlaceID;joinReferences:PlaceID"`
}

func (ct *ClubTime) GetTimeID() uint32 {
	return ct.TimeID
}

func (ct *ClubTime) GetDate() string {
	return ct.Date
}

func (ct *ClubTime) GetTime() string {
	return ct.Time
}

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

type ActivityDetail struct {
	TimeID   uint32     `gorm:"type:int unsigned;not null;primaryKey"`
	PlaceID  uint32     `gorm:"type:int unsigned;not null;primaryKey"`
	ClubUUID string     `gorm:"type:char(36);not null"`
	Remarks  ClubRemark `gorm:"foreignKey:ClubUUID;references:ClubUUID"`
}

type ClubTimeAndPlace struct {
	TimeID       uint32  `json:"time_id"`
	Date         string  `json:"date"`
	Time         string  `json:"time"`
	TimeRemarks  *string `json:"time_remarks"`
	PlaceID      uint32  `json:"place_id"`
	Place        string  `json:"place"`
	PlaceRemarks *string `json:"place_remarks"`
}
