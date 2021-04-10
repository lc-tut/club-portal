package models

import "database/sql"

type ClubSchedule struct {
	ScheduleID uint32         `gorm:"type:int unsigned not null auto_increment;primaryKey"`
	ClubUUID   string         `gorm:"type:char(36);not null"`
	Month      uint8          `gorm:"type:tinyint unsigned;not null;unique"`
	Schedule   string         `gorm:"type:text;not null;unique"`
	Remarks    sql.NullString `gorm:"type:text"`
}
