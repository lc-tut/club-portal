package clubs

import (
	"errors"
	"github.com/lc-tut/club-portal/models/clubs"
	"gorm.io/gorm"
)

type ActivityDetailArgs struct {
	TimeID  uint32
	PlaceID uint32
}

type ClubActivityDetailRepo interface {
	GetActivityDetail(uuid string) ([]clubs.ActivityDetail, error)

	GetAllRelations(uuid string) ([]clubs.DetailRelations, error)

	CreateActivityDetail(uuid string, args []ActivityDetailArgs) error
	CreateActivityDetailWithTx(tx *gorm.DB, uuid string, args []ActivityDetailArgs) error

	UpdateActivityDetailWithTx(tx *gorm.DB, uuid string, args []ActivityDetailArgs) error
}

func (r *ClubRepository) GetActivityDetail(uuid string) ([]clubs.ActivityDetail, error) {
	details := make([]clubs.ActivityDetail, 0)
	tx := r.db.Where("club_uuid = ?", uuid).Find(&details)

	if err := tx.Error; errors.Is(err, gorm.ErrRecordNotFound) {
		r.logger.Info(err.Error())
		return nil, err
	} else if err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}

	return details, nil
}

func (r *ClubRepository) GetAllRelations(uuid string) ([]clubs.DetailRelations, error) {
	selectQuery := "ad.club_uuid, ct.time_id, ct.date, ct.time, cp.place_id, cp.place, cr.remark_id, cr.place_remark, cr.time_remark"
	joinQuery1 := "inner join club_times as ct using (time_id)"
	joinQuery2 := "inner join club_places as cp using (place_id)"
	joinQuery3 := "inner join club_remarks as cr using (club_uuid, time_id, place_id)"

	rows, err := r.db.Table("activity_details as ad").Select(selectQuery).Where("ad.club_uuid = ?", uuid).Joins(joinQuery1).Joins(joinQuery2).Joins(joinQuery3).Rows()

	if err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}

	relations := make([]clubs.DetailRelations, 0)

	for rows.Next() {
		var relation clubs.DetailRelations
		if err := r.db.ScanRows(rows, &relation); err != nil {
			r.logger.Error(err.Error())
			return nil, err
		}
		relations = append(relations, relation)
	}

	return relations, nil
}

func (r *ClubRepository) CreateActivityDetail(uuid string, args []ActivityDetailArgs) error {
	adModels := make([]clubs.ActivityDetail, len(args))

	for i, arg := range args {
		model := clubs.ActivityDetail{
			TimeID:   arg.TimeID,
			PlaceID:  arg.PlaceID,
			ClubUUID: uuid,
		}
		adModels[i] = model
	}

	tx := r.db.Create(&adModels)

	if err := tx.Error; err != nil {
		r.logger.Error(err.Error())
		return err
	}

	return nil
}

func (r *ClubRepository) CreateActivityDetailWithTx(tx *gorm.DB, uuid string, args []ActivityDetailArgs) error {
	adModels := make([]clubs.ActivityDetail, len(args))

	for i, arg := range args {
		model := clubs.ActivityDetail{
			TimeID:   arg.TimeID,
			PlaceID:  arg.PlaceID,
			ClubUUID: uuid,
		}
		adModels[i] = model
	}

	if err := tx.Create(&adModels).Error; err != nil {
		r.logger.Error(err.Error())
		return err
	}

	return nil
}

func (r *ClubRepository) UpdateActivityDetailWithTx(tx *gorm.DB, uuid string, args []ActivityDetailArgs) error {
	if len(args) == 0 {
		return nil
	}

	tx = tx.Where("club_uuid = ?", uuid).Delete(&clubs.ActivityDetail{})

	if err := tx.Error; errors.Is(err, gorm.ErrRecordNotFound) {
		r.logger.Info(err.Error())
		return err
	} else if err != nil {
		r.logger.Error(err.Error())
		return err
	}

	if err := r.CreateActivityDetailWithTx(tx, uuid, args); err != nil {
		return err
	}

	return nil
}
