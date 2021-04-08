package repos

import "github.com/lc-tut/club-portal/models"

type ClubScheduleRepo interface {
	GetScheduleByID(scheduleID uint32) (*models.ClubSchedule, error)
	GetScheduleByClubUUID(uuid string) (*models.ClubSchedule, error)

	CreateSchedule(clubUUID string, month uint8, schedule string, remarks *string) error

	UpdateSchedule(clubUUID string, month uint8, schedule string, remarks *string) error
}

func (r *Repository) GetScheduleByID(scheduleID uint32) (*models.ClubSchedule, error) {
	panic("implement me")
}

func (r *Repository) GetScheduleByClubUUID(uuid string) (*models.ClubSchedule, error) {
	panic("implement me")
}

func (r *Repository) CreateSchedule(clubUUID string, month uint8, schedule string, remarks *string) error {
	panic("implement me")
}

func (r *Repository) UpdateSchedule(clubUUID string, month uint8, schedule string, remarks *string) error {
	panic("implement me")
}
