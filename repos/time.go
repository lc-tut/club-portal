package repos

import (
	"github.com/lc-tut/club-portal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	CreateTimeWithTx(tx *gorm.DB, args []ClubTimeArgs) error
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

	tx := r.db.Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&timeModels)

	if err := tx.Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) CreateTimeWithTx(tx *gorm.DB, args []ClubTimeArgs) error {
	timeModels := make([]models.ClubTime, len(args))

	for i, arg := range args {
		t := models.ClubTime{
			TimeID: arg.TimeID,
			Date:   arg.Date,
			Time:   arg.Time,
		}
		timeModels[i] = t
	}

	if err := tx.Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&timeModels).Error; err != nil {
		return err
	}

	return nil
}
