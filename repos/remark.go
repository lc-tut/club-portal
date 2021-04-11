package repos

import (
	"github.com/lc-tut/club-portal/models"
	"github.com/lc-tut/club-portal/utils"
)

type ClubRemarkArgs struct {
	ClubUUID     string
	TimeID       uint32
	PlaceID      uint32
	TimeRemarks  string
	PlaceRemarks string
}

type ClubRemarkRepo interface {
	GetRemarkByClubUUID(uuid string) (*models.ClubRemark, error)

	CreateRemark(args []ClubRemarkArgs) error
}

func (r *Repository) GetRemarkByClubUUID(uuid string) (*models.ClubRemark, error) {
	remark := &models.ClubRemark{}

	tx := r.db.Where("club_uuid = ?", uuid).Take(&remark)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return remark, nil
}

func (r *Repository) CreateRemark(args []ClubRemarkArgs) error {
	remarks := make([]models.ClubRemark, len(args))

	for _, arg := range args {
		remark := models.ClubRemark{
			ClubUUID:     arg.ClubUUID,
			TimeID:       arg.TimeID,
			PlaceID:      arg.PlaceID,
			TimeRemarks:  utils.ToNullString(arg.TimeRemarks),
			PlaceRemarks: utils.ToNullString(arg.PlaceRemarks),
		}
		remarks = append(remarks, remark)
	}

	tx := r.db.Create(&remarks)

	if err := tx.Error; err != nil {
		return err
	}

	return nil
}
