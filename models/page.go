package models

import (
	"time"
)

type ClubPage struct {
	UUID            string            `gorm:"type:char(36);not null;primaryKey"`
	ClubID          string            `gorm:"type:char(15);not null;unique"`
	Name            string            `gorm:"type:varchar(63);not null"`
	Description     string            `gorm:"type:text;not null"`
	Campus          uint8             `gorm:"type:tinyint;not null"`
	ClubType        uint8             `gorm:"type:tinyint;not null"`
	Visible         bool              `gorm:"type:tinyint(1);not null"`
	UpdatedAt       time.Time         `gorm:"type:datetime;not null"`
	Activities      []ClubActivity    `gorm:"foreignKey:UUID;references:UUID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Links           []ClubLinks       `gorm:"foreignKey:UUID;references:UUID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Schedules       []ClubSchedule    `gorm:"foreignKey:UUID;references:UUID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Achievements    []ClubAchievement `gorm:"foreignKey:UUID;references:UUID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Image           []ClubImage       `gorm:"foreignKey:UUID;references:UUID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Videos          []ClubVideo       `gorm:"foreignKey:UUID;references:UUID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ActivityDetails []ActivityDetail  `gorm:"foreignKey:UUID;references:UUID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
