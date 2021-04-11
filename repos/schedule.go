package repos

import (
	"github.com/lc-tut/club-portal/models"
	"github.com/lc-tut/club-portal/utils"
)

type ClubScheduleArgs struct {
	Month    uint8
	Schedule string
	Remarks  string
}

type ClubScheduleRepo interface {
	GetScheduleByID(scheduleID uint32) (*models.ClubSchedule, error)

	GetSchedulesByClubUUID(uuid string) ([]models.ClubSchedule, error)

	CreateSchedule(clubUUID string, args []ClubScheduleArgs) error

	UpdateSchedule(clubUUID string, args []ClubScheduleArgs) error
}

func (r *Repository) GetScheduleByID(scheduleID uint32) (*models.ClubSchedule, error) {
	schedule := &models.ClubSchedule{}
	tx := r.db.Where("schedule_id = ?", scheduleID).Take(schedule)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return schedule, nil
}

func (r *Repository) GetSchedulesByClubUUID(uuid string) ([]models.ClubSchedule, error) {
	schedule := make([]models.ClubSchedule, 0)
	tx := r.db.Where("club_uuid = ?", uuid).Find(schedule)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return schedule, nil
}

func (r *Repository) CreateSchedule(clubUUID string, args []ClubScheduleArgs) error {
	schedules := make([]models.ClubSchedule, len(args))

	for _, arg := range args {
		sch := models.ClubSchedule{
			ClubUUID: clubUUID,
			Month:    arg.Month,
			Schedule: arg.Schedule,
			Remarks:  utils.ToNullString(arg.Remarks),
		}
		schedules = append(schedules, sch)
	}

	tx := r.db.Create(&schedules)

	if err := tx.Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateSchedule(clubUUID string, args []ClubScheduleArgs) error {
	schedules := make([]models.ClubSchedule, len(args))

	for _, arg := range args {
		sch := models.ClubSchedule{
			ClubUUID: clubUUID,
			Month:    arg.Month,
			Schedule: arg.Schedule,
			Remarks:  utils.ToNullString(arg.Remarks),
		}
		schedules = append(schedules, sch)
	}

	tx := r.db.Model(&models.ClubSchedule{}).Where("club_uuid = ?", clubUUID).Updates(&schedules)

	if err := tx.Error; err != nil {
		return err
	}

	return nil
}
