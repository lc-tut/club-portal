package models

type ClubLink struct {
	LinkID   uint32 `gorm:"type:int unsigned not null auto_increment;primaryKey"`
	ClubUUID string `gorm:"type:char(36);not null"`
	Label    string `gorm:"type:varchar(255);not null;unique"`
	URL      string `gorm:"type:varchar(2047);not null;unique"`
}
