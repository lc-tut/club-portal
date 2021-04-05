package models

type ClubActivity struct {
	ActivityID uint32 `gorm:"type:int unsigned;not null;primaryKey;autoIncrement"`
	UUID       string `gorm:"type:char(36);not null"`
	Activity   string `gorm:"type:text;not null;unique"`
}
