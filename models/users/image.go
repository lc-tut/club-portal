package users

import "time"

type UploadedImage struct {
	ImageID   uint32 `gorm:"type:int unsigned not null auto_increment;primaryKey"`
	UserUUID  string `gorm:"type:char(36);not null;primaryKey"`
	Path      string `gorm:"type:text;not null;unique"`
	CreatedAt time.Time
}
