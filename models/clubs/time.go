package clubs

import "github.com/lc-tut/club-portal/utils"

type ClubTime struct {
	TimeID    uint32      `gorm:"type:int unsigned not null auto_increment;primaryKey"`
	Date      string      `gorm:"type:varchar(3);not null;unique"`
	Time      string      `gorm:"type:varchar(255);not null;unique"`
	ClubPlace []ClubPlace `gorm:"many2many:activity_details;foreignKey:TimeID;joinForeignKey:TimeID;references:PlaceID;joinReferences:PlaceID"`
}

func (ct *ClubTime) GetTimeID() uint32 {
	return ct.TimeID
}

func (ct *ClubTime) GetDate() string {
	return ct.Date
}

func (ct *ClubTime) GetTime() string {
	return ct.Time
}

type Times []ClubTime

func (t Times) ToTimeResponse(remarks []ClubRemark) *[]TimeResponse {
	res := make([]TimeResponse, len(t))

	for i, time := range t {
		timeRes := TimeResponse{
			Date:    time.Date,
			Time:    time.Time,
			Remarks: utils.ToNilOrString(remarks[i].TimeRemark),
		}
		res[i] = timeRes
	}

	return &res
}

type TimeResponse struct {
	Date    string  `json:"date"`
	Time    string  `json:"time"`
	Remarks *string `json:"remarks"`
}
