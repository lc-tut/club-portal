package clubs

import (
	"errors"
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

	CreateSchedule(clubUUID string, args []ClubScheduleArgs) ([]clubs.ClubSchedule, error)
	CreateScheduleWithTx(tx *gorm.DB, clubUUID string, args []ClubScheduleArgs) ([]clubs.ClubSchedule, error)

	UpdateSchedule(clubUUID string, args []ClubScheduleArgs) ([]clubs.ClubSchedule, error)
	UpdateScheduleWithTx(tx *gorm.DB, clubUUID string, args []ClubScheduleArgs) ([]clubs.ClubSchedule, error)
}

func (r *ClubRepository) GetScheduleByID(scheduleID uint32) (*clubs.ClubSchedule, error) {
	schedule := &clubs.ClubSchedule{}
	tx := r.db.Where("schedule_id = ?", scheduleID).Take(schedule)

	if err := tx.Error; errors.Is(err, gorm.ErrRecordNotFound) {
		r.logger.Info(err.Error())
		return nil, err
	} else if err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}

	return schedule, nil
}

func (r *ClubRepository) GetSchedulesByClubUUID(uuid string) ([]clubs.ClubSchedule, error) {
	schedule := make([]clubs.ClubSchedule, 0)
	tx := r.db.Where("club_uuid = ?", uuid).Find(&schedule)

	if err := tx.Error; errors.Is(err, gorm.ErrRecordNotFound) {
		r.logger.Info(err.Error())
		return nil, err
	} else if err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}

	return schedule, nil
}

func (r *ClubRepository) CreateSchedule(clubUUID string, args []ClubScheduleArgs) ([]clubs.ClubSchedule, error) {
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
		return nil, err
	}

	return schedules, nil
}

func (r *ClubRepository) CreateScheduleWithTx(tx *gorm.DB, clubUUID string, args []ClubScheduleArgs) ([]clubs.ClubSchedule, error) {
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
		return nil, err
	}

	return schedules, nil
}

func (r *ClubRepository) UpdateSchedule(clubUUID string, args []ClubScheduleArgs) ([]clubs.ClubSchedule, error) {
	length := len(args)

	if length == 0 {
		return []clubs.ClubSchedule{}, nil
	}

	tx := r.db.Where("club_uuid = ?", clubUUID).Delete(&clubs.ClubSchedule{})

	if err := tx.Error; errors.Is(err, gorm.ErrRecordNotFound) {
		r.logger.Info(err.Error())
		return nil, err
	} else if err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}

	sch, err := r.CreateSchedule(clubUUID, args)

	if err != nil {
		return nil, err
	}

	return sch, nil
}

func (r *ClubRepository) UpdateScheduleWithTx(tx *gorm.DB, clubUUID string, args []ClubScheduleArgs) ([]clubs.ClubSchedule, error) {
	length := len(args)

	if length == 0 {
		return []clubs.ClubSchedule{}, nil
	}

	tx = tx.Where("club_uuid", clubUUID).Delete(&clubs.ClubSchedule{})

	if err := tx.Error; errors.Is(err, gorm.ErrRecordNotFound) {
		r.logger.Info(err.Error())
		return nil, err
	} else if err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}

	sch, err := r.CreateScheduleWithTx(tx, clubUUID, args)

	if err != nil {
		return nil, err
	}

	return sch, nil
}
