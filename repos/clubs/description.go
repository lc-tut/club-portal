package clubs

import (
	"errors"
	"github.com/lc-tut/club-portal/models/clubs"
	"gorm.io/gorm"
)

type ClubDescriptionRepo interface {
	GetClubDescription(uuid string) (string, error)

	UpdateClubDescription(uuid string, desc string) (string, error)
}

func (r *ClubRepository) GetClubDescription(uuid string) (string, error) {
	page := &clubs.ClubPage{}
	tx := r.db.Select("description").Where("club_uuid = ?", uuid).Take(page)

	if err := tx.Error; errors.Is(err, gorm.ErrRecordNotFound) {
		r.logger.Info(err.Error())
		return "", err
	} else if err != nil {
		r.logger.Error(err.Error())
		return "", err
	}

	return page.Description, nil
}

func (r *ClubRepository) UpdateClubDescription(uuid string, desc string) (string, error) {
	tx := r.db.Model(&clubs.ClubPage{}).Where("club_uuid = ?", uuid).Update("description", desc)

	if rows := tx.RowsAffected; rows == 0 {
		err := gorm.ErrRecordNotFound
		r.logger.Info(err.Error())
		return "", err
	} else if err := tx.Error; err != nil {
		r.logger.Error(err.Error())
		return "", err
	}

	return desc, nil
}
