package clubs

import (
	"database/sql"
	"github.com/lc-tut/club-portal/utils"
)

type ClubSchedule struct {
	ScheduleID uint32         `gorm:"type:int unsigned not null auto_increment;primaryKey"`
	ClubUUID   string         `gorm:"type:char(36);not null"`
	Month      uint8          `gorm:"type:tinyint unsigned;not null;unique"`
	Schedule   string         `gorm:"type:text;not null;unique"`
	Remarks    sql.NullString `gorm:"type:text"`
}

type Schedules []ClubSchedule

func (s Schedules) ToScheduleResponse() []ScheduleResponse {
	res := make([]ScheduleResponse, len(s))

	for i, schedule := range s {
		scheduleRes := ScheduleResponse{
			Month:    schedule.Month,
			Schedule: schedule.Schedule,
			Remarks:  utils.NullStringToStringP(schedule.Remarks),
		}
		res[i] = scheduleRes
	}

	return res
}

type ScheduleRequest struct {
	Month    uint8   `json:"month"`
	Schedule string  `json:"schedule"`
	Remarks  *string `json:"remarks"`
}

type ScheduleResponse struct {
	Month    uint8   `json:"month"`
	Schedule string  `json:"schedule"`
	Remarks  *string `json:"remarks"`
}
