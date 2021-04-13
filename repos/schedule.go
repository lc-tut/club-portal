package repos

import (
	"github.com/lc-tut/club-portal/models"
	"github.com/lc-tut/club-portal/utils"
	"gorm.io/gorm"
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
	CreateScheduleWithTx(tx *gorm.DB, clubUUID string, args []ClubScheduleArgs) error

	UpdateSchedule(clubUUID string, args []ClubScheduleArgs) error
	UpdateScheduleWithTx(tx *gorm.DB, clubUUID string, args []ClubScheduleArgs) error
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

	for i, arg := range args {
		sch := models.ClubSchedule{
			ClubUUID: clubUUID,
			Month:    arg.Month,
			Schedule: arg.Schedule,
			Remarks:  utils.ToNullString(arg.Remarks),
		}
		schedules[i] = sch
	}

	tx := r.db.Create(&schedules)

	if err := tx.Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) CreateScheduleWithTx(tx *gorm.DB, clubUUID string, args []ClubScheduleArgs) error {
	schedules := make([]models.ClubSchedule, len(args))

	for i, arg := range args {
		sch := models.ClubSchedule{
			ClubUUID: clubUUID,
			Month:    arg.Month,
			Schedule: arg.Schedule,
			Remarks:  utils.ToNullString(arg.Remarks),
		}
		schedules[i] = sch
	}

	if err := tx.Create(&schedules).Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateSchedule(clubUUID string, args []ClubScheduleArgs) error {
	length := len(args)

	if length == 0 {
		return nil
	}

	schedules := make([]models.ClubSchedule, length)

	for i, arg := range args {
		sch := models.ClubSchedule{
			ClubUUID: clubUUID,
			Month:    arg.Month,
			Schedule: arg.Schedule,
			Remarks:  utils.ToNullString(arg.Remarks),
		}
		schedules[i] = sch
	}

	tx := r.db.Model(&models.ClubSchedule{}).Where("club_uuid = ?", clubUUID).Updates(schedules)

	if err := tx.Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateScheduleWithTx(tx *gorm.DB, clubUUID string, args []ClubScheduleArgs) error {
	length := len(args)

	if length == 0 {
		return nil
	}

	if err := tx.Where("club_uuid", clubUUID).Delete(&models.ClubSchedule{}).Error; err != nil {
		return err
	}

	if err := r.CreateScheduleWithTx(tx, clubUUID, args); err != nil {
		return err
	}

	return nil
}
