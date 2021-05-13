package clubs

type ClubThumbnail struct {
	ThumbnailID uint32 `gorm:"type:int unsigned not null auto_increment;primaryKey"`
	ClubUUID    string `gorm:"type:char(36);not null"`
	Path        string `gorm:"->"`
}

func (ct *ClubThumbnail) ToThumbnailResponse() ThumbnailResponse {
	return ThumbnailResponse{
		ThumbnailID: ct.ThumbnailID,
		Path:        ct.Path,
	}
}

type ThumbnailResponse struct {
	ThumbnailID uint32 `json:"thumbnail_id"`
	Path        string `json:"path"`
}
