package models

type ClubImage struct {
	ImageID  uint32 `gorm:"type:int unsigned not null auto_increment;primaryKey"`
	ClubUUID string `gorm:"type:char(36);not null"`
	Path     string `gorm:"type:text;not null;unique"`
}

type ClubReqImage struct {
	Path string `json:"path"`
}
