package clubs

import (
	"github.com/lc-tut/club-portal/models/clubs"
	"gorm.io/gorm"
)

type ClubNameRepo interface {
	UpdateClubName(uuid string, name string) error
}

// クラブ(サークル)名の更新を行います
func (r *ClubRepository) UpdateClubName(uuid string, name string) error {
	tx := r.db.Model(&clubs.ClubPage{}).Where("club_uuid = ?", uuid).Update("name", name)

	if rows := tx.RowsAffected; rows == 0 {
		err := gorm.ErrRecordNotFound
		r.logger.Info(err.Error())
		return err
	} else if err := tx.Error; err != nil {
		r.logger.Error(err.Error())
		return err
	}

	return nil
}
