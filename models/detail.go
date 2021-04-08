package models

type ClubTime struct {
	TimeID    uint32      `gorm:"type:int unsigned not null auto_increment;primaryKey"`
	Date      string      `gorm:"type:varchar(3);not null;unique"`
	Time      string      `gorm:"type:varchar(255);not null;unique"`
	Remarks   string      `gorm:"type:text"`
	ClubPlace []ClubPlace `gorm:"many2many:activity_details;foreignKey:TimeID;joinForeignKey:TimeID;references:PlaceID;joinReferences:PlaceID"`
}

type ClubPlace struct {
	PlaceID  uint32     `gorm:"type:int unsigned not null auto_increment;primaryKey"`
	Place    string     `gorm:"type:text;not null;unique"`
	Remarks  string     `gorm:"type:text"`
	ClubTime []ClubTime `gorm:"many2many:activity_details;foreignKey:PlaceID;joinForeignKey:PlaceID;references:TimeID;joinReferences:TimeID"`
}

type ActivityDetail struct {
	TimeID   uint32 `gorm:"type:int unsigned;not null;primaryKey"`
	PlaceID  uint32 `gorm:"type:int unsigned;not null;primaryKey"`
	ClubUUID string `gorm:"type:char(36);not null"`
}

type ClubDetail struct {
	ClubTime
	ClubPlace
}
