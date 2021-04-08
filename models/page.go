package models

import (
	"time"
)

type ClubPage struct {
	ClubUUID        string            `gorm:"type:char(36);not null;primaryKey"`
	ClubSlug        string            `gorm:"type:char(15);not null;unique"`
	Name            string            `gorm:"type:varchar(63);not null"`
	Description     string            `gorm:"type:text;not null"`
	Campus          uint8             `gorm:"type:tinyint;not null"`
	ClubType        uint8             `gorm:"type:tinyint;not null"`
	Visible         bool              `gorm:"type:tinyint(1);not null"`
	UpdatedAt       time.Time         `gorm:"type:datetime;not null"`
	Activities      []ClubContent     `gorm:"foreignKey:ClubUUID;references:ClubUUID"`
	Links           []ClubLinks       `gorm:"foreignKey:ClubUUID;references:ClubUUID"`
	Schedules       []ClubSchedule    `gorm:"foreignKey:ClubUUID;references:ClubUUID"`
	Achievements    []ClubAchievement `gorm:"foreignKey:ClubUUID;references:ClubUUID"`
	Image           []ClubImage       `gorm:"foreignKey:ClubUUID;references:ClubUUID"`
	Videos          []ClubVideo       `gorm:"foreignKey:ClubUUID;references:ClubUUID"`
	ActivityDetails []ActivityDetail  `gorm:"foreignKey:ClubUUID;references:ClubUUID"`
}
