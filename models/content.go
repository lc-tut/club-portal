package models

type ClubContent struct {
	ContentID uint32 `gorm:"type:int unsigned not null auto_increment;primaryKey"`
	ClubUUID  string `gorm:"type:char(36);not null"`
	Content   string `gorm:"type:text;not null;unique"`
}
