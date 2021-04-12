package repos

import (
	"errors"
	"github.com/lc-tut/club-portal/models"
)

type ClubTimeArgs struct {
	TimeID uint32
	Date   string
	Time   string
}

type ClubTimeRepo interface {
	GetTimeByID(timeID uint32) (*models.ClubTime, error)

	GetTimesByClubUUID(uuid string) ([]models.ClubTime, error)

	CreateTime(args []ClubTimeArgs) error

	UpdateTime(timeID []uint32, args []ClubTimeArgs) error
}

func (r *Repository) GetTimeByID(timeID uint32) (*models.ClubTime, error) {
	time := &models.ClubTime{}
	tx := r.db.Where("time_id = ?", timeID).Take(time)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return time, nil
}

func (r *Repository) GetTimesByClubUUID(uuid string) ([]models.ClubTime, error) {
	times := make([]models.ClubTime, 0)
	tx := r.db.Where("club_uuid = ?", uuid).Find([]models.ActivityDetail{}).Joins("inner join club_times on activity_details.time_id = club_times.time_id").Scan(times)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return times, nil
}

func (r *Repository) CreateTime(args []ClubTimeArgs) error {
	timeModels := make([]models.ClubTime, len(args))

	for i, arg := range args {
		t := models.ClubTime{
			TimeID: arg.TimeID,
			Date:   arg.Date,
			Time:   arg.Time,
		}
		timeModels[i] = t
	}

	tx := r.db.Create(&timeModels)

	if err := tx.Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateTime(timeID []uint32, args []ClubTimeArgs) error {
	if len(timeID) != len(args) {
		return errors.New("invalid argument")
	}

	timeModels := make([]models.ClubTime, len(timeID))

	for i, arg := range args {
		t := models.ClubTime{

			Date: arg.Date,
			Time: arg.Time,
		}
		timeModels[i] = t
	}

	tx := r.db.Model(&models.ClubTime{}).Where("time_id = ?", timeID).Updates(&timeModels)

	if err := tx.Error; err != nil {
		return err
	}

	return nil
}
