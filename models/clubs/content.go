package clubs

type ClubContent struct {
	ContentID uint32 `gorm:"type:int unsigned not null auto_increment;primaryKey"`
	ClubUUID  string `gorm:"type:char(36);not null"`
	Content   string `gorm:"type:text;not null;unique"`
}

type Contents []ClubContent

func (c Contents) ToContentResponse() []ContentResponse {
	res := make([]ContentResponse, len(c))

	for i, con := range c {
		contentRes := ContentResponse{Content: con.Content}
		res[i] = contentRes
	}

	return res
}

type ContentRequest struct {
	Content string `json:"content"`
}

type ContentResponse struct {
	Content string `json:"content"`
}
