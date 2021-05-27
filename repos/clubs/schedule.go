package clubs

import (
	"github.com/lc-tut/club-portal/models/clubs"
	"github.com/lc-tut/club-portal/utils"
	"gorm.io/gorm"
)

type ClubScheduleArgs struct {
	Month    uint8
	Schedule string
	Remarks  string
}

type ClubScheduleRepo interface {
	GetScheduleByID(scheduleID uint32) (*clubs.ClubSchedule, error)

	GetSchedulesByClubUUID(uuid string) ([]clubs.ClubSchedule, error)

	CreateSchedule(clubUUID string, args []ClubScheduleArgs) error
	CreateScheduleWithTx(tx *gorm.DB, clubUUID string, args []ClubScheduleArgs) error

	UpdateSchedule(clubUUID string, args []ClubScheduleArgs) error
	UpdateScheduleWithTx(tx *gorm.DB, clubUUID string, args []ClubScheduleArgs) error
}

func (r *ClubRepository) GetScheduleByID(scheduleID uint32) (*clubs.ClubSchedule, error) {
	schedule := &clubs.ClubSchedule{}
	tx := r.db.Where("schedule_id = ?", scheduleID).Take(schedule)

	if err := tx.Error; err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}

	return schedule, nil
}

func (r *ClubRepository) GetSchedulesByClubUUID(uuid string) ([]clubs.ClubSchedule, error) {
	schedule := make([]clubs.ClubSchedule, 0)
	tx := r.db.Where("club_uuid = ?", uuid).Find(&schedule)

	if err := tx.Error; err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}

	return schedule, nil
}

func (r *ClubRepository) CreateSchedule(clubUUID string, args []ClubScheduleArgs) error {
	schedules := make([]clubs.ClubSchedule, len(args))

	for i, arg := range args {
		sch := clubs.ClubSchedule{
			ClubUUID: clubUUID,
			Month:    arg.Month,
			Schedule: arg.Schedule,
			Remarks:  utils.StringToNullString(arg.Remarks),
		}
		schedules[i] = sch
	}

	tx := r.db.Create(&schedules)

	if err := tx.Error; err != nil {
		r.logger.Error(err.Error())
		return err
	}

	return nil
}

func (r *ClubRepository) CreateScheduleWithTx(tx *gorm.DB, clubUUID string, args []ClubScheduleArgs) error {
	schedules := make([]clubs.ClubSchedule, len(args))

	for i, arg := range args {
		sch := clubs.ClubSchedule{
			ClubUUID: clubUUID,
			Month:    arg.Month,
			Schedule: arg.Schedule,
			Remarks:  utils.StringToNullString(arg.Remarks),
		}
		schedules[i] = sch
	}

	if err := tx.Create(&schedules).Error; err != nil {
		r.logger.Error(err.Error())
		return err
	}

	return nil
}

func (r *ClubRepository) UpdateSchedule(clubUUID string, args []ClubScheduleArgs) error {
	length := len(args)

	if length == 0 {
		return nil
	}

	tx := r.db.Where("club_uuid = ?", clubUUID).Delete(&clubs.ClubSchedule{})

	if err := tx.Error; err != nil {
		r.logger.Error(err.Error())
		return err
	}

	if err := r.CreateSchedule(clubUUID, args); err != nil {
		return err
	}

	return nil
}

func (r *ClubRepository) UpdateScheduleWithTx(tx *gorm.DB, clubUUID string, args []ClubScheduleArgs) error {
	length := len(args)

	if length == 0 {
		return nil
	}

	if err := tx.Where("club_uuid", clubUUID).Delete(&clubs.ClubSchedule{}).Error; err != nil {
		r.logger.Error(err.Error())
		return err
	}

	if err := r.CreateScheduleWithTx(tx, clubUUID, args); err != nil {
		return err
	}

	return nil
}
