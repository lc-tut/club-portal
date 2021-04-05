package models

type ClubSchedule struct {
	ScheduleID uint8  `gorm:"type:int unsigned;not null;primaryKey;autoIncrement"`
	UUID       string `gorm:"type:char(36);not null"`
	Month      uint8  `gorm:"type:tinyint unsigned;not null;unique"`
	Schedule   string `gorm:"type:text;not null;unique"`
	Remarks    string `gorm:"type:text"`
}
