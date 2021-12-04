package clubs

import (
	"errors"
	"github.com/lc-tut/club-portal/models/clubs"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ClubTimeArgs struct {
	TimeID uint32
	Date   string
	Time   string
}

type ClubTimeRepo interface {
	GetTimeByID(timeID uint32) (*clubs.ClubTime, error)

	GetTimesByClubUUID(uuid string) ([]clubs.ClubTime, error)

	CreateTime(args []ClubTimeArgs) error
	CreateTimeWithTx(tx *gorm.DB, args []ClubTimeArgs) error
}

func (r *ClubRepository) GetTimeByID(timeID uint32) (*clubs.ClubTime, error) {
	time := &clubs.ClubTime{}
	tx := r.db.Where("time_id = ?", timeID).Take(time)

	if err := tx.Error; errors.Is(err, gorm.ErrRecordNotFound) {
		r.logger.Info(err.Error())
		return nil, err
	} else if err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}

	return time, nil
}

func (r *ClubRepository) GetTimesByClubUUID(uuid string) ([]clubs.ClubTime, error) {
	times := make([]clubs.ClubTime, 0)
	tx := r.db.Where("club_uuid = ?", uuid).Find([]clubs.ActivityDetail{}).Joins("inner join club_times on activity_details.time_id = club_times.time_id").Scan(&times)

	if err := tx.Error; errors.Is(err, gorm.ErrRecordNotFound) {
		r.logger.Info(err.Error())
		return nil, err
	} else if err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}

	return times, nil
}

func (r *ClubRepository) CreateTime(args []ClubTimeArgs) error {
	timeModels := make([]clubs.ClubTime, len(args))

	for i, arg := range args {
		t := clubs.ClubTime{
			TimeID: arg.TimeID,
			Date:   arg.Date,
			Time:   arg.Time,
		}
		timeModels[i] = t
	}

	tx := r.db.Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&timeModels)

	if err := tx.Error; err != nil {
		r.logger.Error(err.Error())
		return err
	}

	return nil
}

func (r *ClubRepository) CreateTimeWithTx(tx *gorm.DB, args []ClubTimeArgs) error {
	timeModels := make([]clubs.ClubTime, len(args))

	for i, arg := range args {
		t := clubs.ClubTime{
			TimeID: arg.TimeID,
			Date:   arg.Date,
			Time:   arg.Time,
		}
		timeModels[i] = t
	}

	if err := tx.Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&timeModels).Error; err != nil {
		r.logger.Error(err.Error())
		return err
	}

	return nil
}
