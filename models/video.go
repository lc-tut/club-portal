package models

type ClubVideo struct {
	VideoID uint8  `gorm:"type:int unsigned;not null;primaryKey;autoIncrement"`
	UUID    string `gorm:"type:char(36);not null"`
	Path    string `gorm:"type:text;not null;unique"`
}
