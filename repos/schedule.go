package repos

import (
	"database/sql"
	"github.com/lc-tut/club-portal/models"
)

type ClubScheduleRepo interface {
	GetScheduleByID(scheduleID uint32) (*models.ClubSchedule, error)
	GetScheduleByClubUUID(uuid string) (*models.ClubSchedule, error)

	CreateSchedule(clubUUID string, month uint8, schedule string, remarks string) error

	UpdateSchedule(clubUUID string, month uint8, schedule string, remarks string) error
}

func (r *Repository) GetScheduleByID(scheduleID uint32) (*models.ClubSchedule, error) {
	schedule := &models.ClubSchedule{}
	tx := r.db.Where("schedule_id = ?", scheduleID).Take(schedule)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return schedule, nil
}

func (r *Repository) GetScheduleByClubUUID(uuid string) (*models.ClubSchedule, error) {
	schedule := &models.ClubSchedule{}
	tx := r.db.Where("club_uuid = ?", uuid).Take(schedule)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return schedule, nil
}

func (r *Repository) CreateSchedule(clubUUID string, month uint8, schedule string, remarks string) error {
	var validatedRemarks sql.NullString

	if remarks == "" {
		validatedRemarks = sql.NullString{
			String: "",
			Valid:  false,
		}
	} else {
		validatedRemarks = sql.NullString{
			String: remarks,
			Valid:  true,
		}
	}

	sch := &models.ClubSchedule{
		ClubUUID: clubUUID,
		Month:    month,
		Schedule: schedule,
		Remarks:  validatedRemarks,
	}

	tx := r.db.Create(sch)

	if err := tx.Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateSchedule(clubUUID string, month uint8, schedule string, remarks string) error {
	var validatedRemarks sql.NullString

	if remarks == "" {
		validatedRemarks = sql.NullString{
			String: "",
			Valid:  false,
		}
	} else {
		validatedRemarks = sql.NullString{
			String: remarks,
			Valid:  true,
		}
	}

	sch := &models.ClubSchedule{
		Month:    month,
		Schedule: schedule,
		Remarks:  validatedRemarks,
	}

	tx := r.db.Model(sch).Where("club_uuid = ?", clubUUID).Updates(sch)

	if err := tx.Error; err != nil {
		return err
	}

	return nil
}
