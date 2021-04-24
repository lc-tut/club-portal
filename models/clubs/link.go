package clubs

type ClubLink struct {
	LinkID   uint32 `gorm:"type:int unsigned not null auto_increment;primaryKey"`
	ClubUUID string `gorm:"type:char(36);not null"`
	Label    string `gorm:"type:varchar(255);not null;unique"`
	URL      string `gorm:"type:varchar(2047);not null;unique"`
}

type Links []ClubLink

func (l Links) ToLinkResponse() *[]LinkResponse {
	res := make([]LinkResponse, len(l))

	for i, link := range l {
		linkRes := LinkResponse{
			Label: link.Label,
			URL:   link.URL,
		}
		res[i] = linkRes
	}

	return &res
}

type LinkRequest struct {
	Label string `json:"label"`
	URL   string `json:"url"`
}

type LinkResponse struct {
	Label string `json:"label"`
	URL   string `json:"url"`
}
