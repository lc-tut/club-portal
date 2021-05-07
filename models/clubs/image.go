package clubs

type ClubImage struct {
	ImageID  uint32 `gorm:"type:int unsigned not null auto_increment;primaryKey"`
	ClubUUID string `gorm:"type:char(36);not null"`
	Path     string `gorm:"-"`
}

type Images []ClubImage

func (im Images) ToImageResponse() []ImageResponse {
	res := make([]ImageResponse, len(im))

	for i, image := range im {
		imageRes := ImageResponse{
			ImageID: image.ImageID,
			Path:    image.Path,
		}
		res[i] = imageRes
	}

	return res
}

type ImageRequest struct {
	ImageID uint32 `json:"image_id"`
}

type ImageResponse struct {
	ImageID uint32 `json:"image_id"`
	Path    string `json:"path"`
}
