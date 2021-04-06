package models

type ClubVideo struct {
	VideoID uint32 `gorm:"type:int unsigned not null auto_increment;primaryKey"`
	UUID    string `gorm:"type:char(36);not null"`
	Path    string `gorm:"type:text;not null;unique"`
}
