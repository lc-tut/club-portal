package users

type UploadedThumbnail struct {
	ThumbnailID uint32 `gorm:"type:int unsigned not null auto_increment;primaryKey"`
	Path        string `gorm:"type:varchar(255);not null;unique"`
}
