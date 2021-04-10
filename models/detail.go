package models

import "database/sql"

type ClubTime struct {
	TimeID    uint32         `gorm:"type:int unsigned not null auto_increment;primaryKey"`
	Date      string         `gorm:"type:varchar(3);not null;unique"`
	Time      string         `gorm:"type:varchar(255);not null;unique"`
	Remarks   sql.NullString `gorm:"type:text"`
	ClubPlace []ClubPlace    `gorm:"many2many:activity_details;foreignKey:TimeID;joinForeignKey:TimeID;references:PlaceID;joinReferences:PlaceID"`
}

func (ct *ClubTime) GetDate() string {
	return ct.Date
}

func (ct *ClubTime) GetTime() string {
	return ct.Time
}

func (ct *ClubTime) GetRemarks() string {
	return ct.Remarks.String
}

type ClubPlace struct {
	PlaceID  uint32         `gorm:"type:int unsigned not null auto_increment;primaryKey"`
	Place    string         `gorm:"type:text;not null;unique"`
	Remarks  sql.NullString `gorm:"type:text"`
	ClubTime []ClubTime     `gorm:"many2many:activity_details;foreignKey:PlaceID;joinForeignKey:PlaceID;references:TimeID;joinReferences:TimeID"`
}

func (cp *ClubPlace) GetPlace() string {
	return cp.Place
}

func (cp *ClubPlace) GetRemarks() string {
	return cp.Remarks.String
}

type ActivityDetail struct {
	TimeID   uint32 `gorm:"type:int unsigned;not null;primaryKey"`
	PlaceID  uint32 `gorm:"type:int unsigned;not null;primaryKey"`
	ClubUUID string `gorm:"type:char(36);not null"`
}

type ClubTimeAndPlace struct {
	ClubTime
	ClubPlace
}
