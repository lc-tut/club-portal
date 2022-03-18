package clubs

import (
	"errors"
	"github.com/lc-tut/club-portal/models/clubs"
	"gorm.io/gorm"
)

type ClubThumbnailRepo interface {
	GetClubThumbnailByID(thumbnailID uint32) (*clubs.ClubThumbnail, error)
	GetClubThumbnailByUUID(clubUUID string) (*clubs.ClubThumbnail, error)

	CreateClubThumbnail(clubUUID string, thumbnailID uint32) error
	CreateClubThumbnailWithTx(tx *gorm.DB, clubUUID string, thumbnailID uint32) error

	UpdateClubThumbnail(clubUUID string, thumbnailID uint32) error
	UpdateClubThumbnailWithTx(tx *gorm.DB, clubUUID string, thumbnailID uint32) error
}

func (r *ClubRepository) GetClubThumbnailByID(thumbnailID uint32) (*clubs.ClubThumbnail, error) {
	thumbnail := &clubs.ClubThumbnail{}
	selectQuery := "club_thumbnails.thumbnail_id, club_thumbnails.club_uuid, ut.path"
	joinQuery := "inner join uploaded_thumbnails as ut using (thumbnail_id)"
	tx := r.db.Joins(joinQuery).Select(selectQuery).Where("thumbnail_id = ?", thumbnailID).Find(thumbnail)

	if err := tx.Error; errors.Is(err, gorm.ErrRecordNotFound) {
		r.logger.Info(err.Error())
		return nil, err
	} else if err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}

	return thumbnail, nil
}

func (r *ClubRepository) GetClubThumbnailByUUID(clubUUID string) (*clubs.ClubThumbnail, error) {
	thumbnail := &clubs.ClubThumbnail{}
	selectQuery := "ct.thumbnail_id, ct.club_uuid, ut.path"
	joinQuery := "inner join uploaded_thumbnails as ut using (thumbnail_id)"
	tx := r.db.Table("club_thumbnails as ct").Joins(joinQuery).Select(selectQuery).Where("club_uuid = ?", clubUUID).Find(thumbnail)

	if err := tx.Error; errors.Is(err, gorm.ErrRecordNotFound) {
		r.logger.Info(err.Error())
		return nil, err
	} else if err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}

	return thumbnail, nil
}

func (r *ClubRepository) CreateClubThumbnail(clubUUID string, thumbnailID uint32) error {
	thumbnail := &clubs.ClubThumbnail{
		ThumbnailID: thumbnailID,
		ClubUUID:    clubUUID,
	}
	tx := r.db.Create(thumbnail)

	if err := tx.Error; err != nil {
		r.logger.Error(err.Error())
		return err
	}

	return nil
}

func (r *ClubRepository) CreateClubThumbnailWithTx(tx *gorm.DB, clubUUID string, thumbnailID uint32) error {
	thumbnail := &clubs.ClubThumbnail{
		ThumbnailID: thumbnailID,
		ClubUUID:    clubUUID,
	}

	if err := tx.Create(thumbnail).Error; err != nil {
		r.logger.Error(err.Error())
		return err
	}

	return nil
}

func (r *ClubRepository) UpdateClubThumbnail(clubUUID string, thumbnailID uint32) error {
	tx := r.db.Model(&clubs.ClubThumbnail{}).Where("club_uuid = ?", clubUUID).Update("thumbnail_id", thumbnailID)

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

func (r *ClubRepository) UpdateClubThumbnailWithTx(tx *gorm.DB, clubUUID string, thumbnailID uint32) error {
	tx = tx.Model(&clubs.ClubThumbnail{}).Where("club_uuid = ?", clubUUID).Update("thumbnail_id", thumbnailID)

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
