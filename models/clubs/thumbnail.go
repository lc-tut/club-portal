package clubs

type ClubThumbnail struct {
	ThumbnailID uint32 `gorm:"type:int unsigned not null auto_increment;primaryKey"`
	ClubUUID    string `gorm:"type:char(36);not null;primaryKey"`
	Path        string `gorm:"type:varchar(255);not null;unique"`
}
