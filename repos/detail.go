package repos

import (
	"github.com/lc-tut/club-portal/models"
	"gorm.io/gorm"
)

type ClubActivityDetailArgs struct {
	TimeID  uint32
	PlaceID uint32
}

type ClubActivityDetailRepo interface {
	GetClubActivityDetail(uuid string) ([]models.ActivityDetail, error)

	GetAllRelations(uuid string) ([]models.DetailRelations, error)

	CreateClubActivityDetail(uuid string, args []ClubActivityDetailArgs) error
	CreateClubActivityDetailWithTx(tx *gorm.DB, uuid string, args []ClubActivityDetailArgs) error
}

func (r *Repository) GetClubActivityDetail(uuid string) ([]models.ActivityDetail, error) {
	details := make([]models.ActivityDetail, 0)
	tx := r.db.Where("club_uuid = ?", uuid).Find(&details)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return details, nil
}

func (r *Repository) GetAllRelations(uuid string) ([]models.DetailRelations, error) {
	selectQuery := "ad.club_uuid, ct.time_id, ct.date, ct.time, cp.place_id, cp.place, cr.remark_id, cr.place_remark, cr.time_remark"
	joinQuery1 := "inner join club_times as ct using (time_id)"
	joinQuery2 := "inner join club_places as cp using (place_id)"
	joinQuery3 := "inner join club_remarks as cr using (club_uuid, time_id, place_id)"

	rows, err := r.db.Table("activity_details as ad").Select(selectQuery).Where("ad.club_uuid = ?", uuid).Joins(joinQuery1).Joins(joinQuery2).Joins(joinQuery3).Rows()

	if err != nil {
		return nil, err
	}

	relations := make([]models.DetailRelations, 0)
	i := 0

	for rows.Next() {
		var relation models.DetailRelations
		if err := r.db.ScanRows(rows, &relation); err != nil {
			return nil, err
		}
		relations = append(relations, relation)
		i++
	}

	return relations, nil
}

func (r *Repository) CreateClubActivityDetail(uuid string, args []ClubActivityDetailArgs) error {
	adModels := make([]models.ActivityDetail, len(args))

	for i, arg := range args {
		model := models.ActivityDetail{
			TimeID:   arg.TimeID,
			PlaceID:  arg.PlaceID,
			ClubUUID: uuid,
		}
		adModels[i] = model
	}

	tx := r.db.Create(&adModels)

	if err := tx.Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) CreateClubActivityDetailWithTx(tx *gorm.DB, uuid string, args []ClubActivityDetailArgs) error {
	adModels := make([]models.ActivityDetail, len(args))

	for i, arg := range args {
		model := models.ActivityDetail{
			TimeID:   arg.TimeID,
			PlaceID:  arg.PlaceID,
			ClubUUID: uuid,
		}
		adModels[i] = model
	}

	if err := tx.Create(&adModels).Error; err != nil {
		return err
	}

	return nil
}
