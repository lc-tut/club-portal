package users

import "time"

type UploadedImage struct {
	ImageID   uint32 `gorm:"type:int unsigned not null auto_increment;primaryKey"`
	Owner     string `gorm:"type:char(36);not null;primaryKey"`
	Path      string `gorm:"type:varchar(255);not null;unique"`
	CreatedAt time.Time
}

type Images []UploadedImage

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

type ImageResponse struct {
	ImageID uint32 `json:"image_id"`
	Path    string `json:"path"`
}
