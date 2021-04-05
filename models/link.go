package models

type ClubLinks struct {
	LinkID uint32 `gorm:"type:int unsigned;not null;primaryKey;autoIncrement"`
	UUID   string `gorm:"type:char(36);not null"`
	Label  string `gorm:"type:varchar(255);not null;unique"`
	URL    string `gorm:"type:varchar(2047);not null;unique"`
}
