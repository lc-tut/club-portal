package clubs

import (
	"errors"
	"github.com/lc-tut/club-portal/models/clubs"
	"github.com/lc-tut/club-portal/utils"
	"gorm.io/gorm"
)

type ClubTPRemarkArgs struct {
	TimeID       uint32
	PlaceID      uint32
	TimeRemarks  string
	PlaceRemarks string
}

type ClubRemarkRepo interface {
	GetTPRemarksByClubUUID(uuid string) ([]clubs.ClubRemark, error)

	CreateTPRemark(uuid string, args []ClubTPRemarkArgs) error
	CreateTPRemarkWithTx(tx *gorm.DB, uuid string, args []ClubTPRemarkArgs) error

	UpdateTPRemarkWithTx(tx *gorm.DB, uuid string, args []ClubTPRemarkArgs) error
}

func (r *ClubRepository) GetTPRemarksByClubUUID(uuid string) ([]clubs.ClubRemark, error) {
	remarks := make([]clubs.ClubRemark, 0)

	tx := r.db.Where("club_uuid = ?", uuid).Find(&remarks)

	if err := tx.Error; errors.Is(err, gorm.ErrRecordNotFound) {
		r.logger.Info(err.Error())
		return nil, err
	} else if err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}

	return remarks, nil
}

func (r *ClubRepository) CreateTPRemark(uuid string, args []ClubTPRemarkArgs) error {
	remarks := make([]clubs.ClubRemark, len(args))

	for i, arg := range args {
		remark := clubs.ClubRemark{
			ClubUUID:    uuid,
			TimeID:      arg.TimeID,
			PlaceID:     arg.PlaceID,
			TimeRemark:  utils.StringToNullString(arg.TimeRemarks),
			PlaceRemark: utils.StringToNullString(arg.PlaceRemarks),
		}
		remarks[i] = remark
	}

	tx := r.db.Create(&remarks)

	if err := tx.Error; err != nil {
		r.logger.Error(err.Error())
		return err
	}

	return nil
}

func (r *ClubRepository) CreateTPRemarkWithTx(tx *gorm.DB, uuid string, args []ClubTPRemarkArgs) error {
	remarks := make([]clubs.ClubRemark, len(args))

	for i, arg := range args {
		remark := clubs.ClubRemark{
			ClubUUID:    uuid,
			TimeID:      arg.TimeID,
			PlaceID:     arg.PlaceID,
			TimeRemark:  utils.StringToNullString(arg.TimeRemarks),
			PlaceRemark: utils.StringToNullString(arg.PlaceRemarks),
		}
		remarks[i] = remark
	}

	if err := tx.Create(&remarks).Error; err != nil {
		r.logger.Error(err.Error())
		return err
	}

	return nil
}

func (r *ClubRepository) UpdateTPRemarkWithTx(tx *gorm.DB, uuid string, args []ClubTPRemarkArgs) error {
	if len(args) == 0 {
		return nil
	}

	tx = tx.Where("club_uuid = ?", uuid).Delete(&clubs.ClubRemark{})

	if err := tx.Error; errors.Is(err, gorm.ErrRecordNotFound) {
		r.logger.Info(err.Error())
		return err
	} else if err != nil {
		r.logger.Error(err.Error())
		return err
	}

	if err := r.CreateTPRemarkWithTx(tx, uuid, args); err != nil {
		return err
	}

	return nil
}
