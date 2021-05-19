package users

type UploadedThumbnail struct {
	ThumbnailID uint32 `gorm:"type:int unsigned not null auto_increment;primaryKey"`
	Path        string `gorm:"type:varchar(255);not null;unique"`
}

func (ut *UploadedThumbnail) ToThumbnailResponse() ThumbnailResponse {
	return ThumbnailResponse{
		ThumbnailID: ut.ThumbnailID,
		Path:        ut.Path,
	}
}

type ThumbnailResponse struct {
	ThumbnailID uint32 `json:"thumbnail_id"`
	Path        string `json:"path"`
}
