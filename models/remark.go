package models

import "database/sql"

type ClubRemark struct {
	RemarkID    uint32         `gorm:"type:int unsigned not null auto_increment;primaryKey"`
	ClubUUID    string         `gorm:"type:char(36);not null"`
	PlaceID     uint32         `gorm:"type:int unsigned not null"`
	TimeID      uint32         `gorm:"type:int unsigned not null"`
	PlaceRemark sql.NullString `gorm:"type:text"`
	TimeRemark  sql.NullString `gorm:"type:text"`
}
