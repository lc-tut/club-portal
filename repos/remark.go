package repos

import (
	"github.com/lc-tut/club-portal/models"
	"github.com/lc-tut/club-portal/utils"
)

type ClubRemarkArgs struct {
	TimeID       uint32
	PlaceID      uint32
	TimeRemarks  string
	PlaceRemarks string
}

type ClubRemarkRepo interface {
	GetRemarkByClubUUID(uuid string) ([]models.ClubRemark, error)

	CreateRemark(uuid string, args []ClubRemarkArgs) error
}

func (r *Repository) GetRemarkByClubUUID(uuid string) ([]models.ClubRemark, error) {
	remarks := make([]models.ClubRemark, 0)

	tx := r.db.Where("club_uuid = ?", uuid).Find(&remarks)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return remarks, nil
}

func (r *Repository) CreateRemark(uuid string, args []ClubRemarkArgs) error {
	remarks := make([]models.ClubRemark, len(args))

	for i, arg := range args {
		remark := models.ClubRemark{
			ClubUUID:    uuid,
			TimeID:      arg.TimeID,
			PlaceID:     arg.PlaceID,
			TimeRemark:  utils.ToNullString(arg.TimeRemarks),
			PlaceRemark: utils.ToNullString(arg.PlaceRemarks),
		}
		remarks[i] = remark
	}

	tx := r.db.Create(&remarks)

	if err := tx.Error; err != nil {
		return err
	}

	return nil
}
