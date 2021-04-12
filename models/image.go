package models

type ClubImage struct {
	ImageID  uint32 `gorm:"type:int unsigned not null auto_increment;primaryKey"`
	ClubUUID string `gorm:"type:char(36);not null"`
	Path     string `gorm:"type:text;not null;unique"`
}

type Images []ClubImage

func (im Images) ToImageResponse() *[]ImageResponse {
	res := make([]ImageResponse, len(im))

	for i, image := range im {
		imageRes := ImageResponse{Path: image.Path}
		res[i] = imageRes
	}

	return &res
}

type ImageRequest struct {
	Path string `json:"path"`
}

type ImageResponse struct {
	Path string `json:"path"`
}
