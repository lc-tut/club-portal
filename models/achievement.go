package models

type ClubAchievement struct {
	AchievementID uint8  `gorm:"type:int unsigned;not null;primaryKey;autoIncrement"`
	UUID          string `gorm:"type:char(36);not null"`
	Achievement   string `gorm:"type:text;not null"`
}
