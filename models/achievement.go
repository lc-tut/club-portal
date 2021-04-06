package models

type ClubAchievement struct {
	AchievementID uint32 `gorm:"type:int unsigned not null auto_increment;primaryKey"`
	UUID          string `gorm:"type:char(36);not null"`
	Achievement   string `gorm:"type:text;not null"`
}
