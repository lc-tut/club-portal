package clubs

import (
	"errors"
	"github.com/lc-tut/club-portal/models/clubs"
	"gorm.io/gorm"
)

type ClubSearchRepo interface {
	DoSearch(content string) ([]clubs.ClubPageExternalInfo, error)
}

func (r *ClubRepository) DoSearch(content string) ([]clubs.ClubPageExternalInfo, error) {
	pages := make([]clubs.ClubPage, 0)
	tx := r.db.Where("visible is true").Where("description LIKE ? OR short_description LIKE ?", "%"+content+"%", "%"+content+"%").Preload("Thumbnail", func(db *gorm.DB) *gorm.DB {
		selectQuery := "club_thumbnails.thumbnail_id, club_thumbnails.club_uuid, ut.path"
		joinQuery := "inner join uploaded_thumbnails as ut using (thumbnail_id)"
		return db.Joins(joinQuery).Select(selectQuery)
	}).Find(&pages)

	if err := tx.Error; errors.Is(err, gorm.ErrRecordNotFound) {
		r.logger.Info(err.Error())
		return make([]clubs.ClubPageExternalInfo, 0), nil
	} else if err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}

	return clubs.Pages(pages).ToExternalInfo(), nil
}
