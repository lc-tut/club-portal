package clubs

type ClubVideo struct {
	VideoID  uint32 `gorm:"type:int unsigned not null auto_increment;primaryKey"`
	ClubUUID string `gorm:"type:char(36);not null"`
	Path     string `gorm:"type:text;not null;unique"`
}

type Videos []ClubVideo

func (v Videos) ToVideoResponse() []VideoResponse {
	res := make([]VideoResponse, len(v))

	for i, video := range v {
		videoRes := VideoResponse{
			Path: video.Path,
		}
		res[i] = videoRes
	}

	return res
}

type VideoRequest struct {
	Path string `json:"path"`
}

type VideoResponse struct {
	Path string `json:"path"`
}
