package models

type ClubImage struct {
	ImageID uint8  `gorm:"type:int unsigned;not null;primaryKey;autoIncrement"`
	UUID    string `gorm:"type:char(36);not null"`
	Path    string `gorm:"type:text;unique"`
}
