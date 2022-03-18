package users

import (
	"errors"
	"github.com/lc-tut/club-portal/models/clubs"
	"github.com/lc-tut/club-portal/models/users"
	"gorm.io/gorm"
)

type UserFavoriteRepo interface {
	GetFavorites(userUUID string) ([]clubs.ClubPageExternalInfo, error)

	CreateFavorite(userUUID string, clubUUID string) error

	DeleteFavorite(userUUID string, clubUUID string) error
}

func (r *UserRepository) GetFavorites(userUUID string) ([]clubs.ClubPageExternalInfo, error) {
	clubPage := make([]clubs.ClubPage, 0)
	joinQuery := "inner join club_pages as cp using (club_uuid)"
	tx := r.db.Table("favorite_clubs").Select("cp.*").Where("user_uuid = ? AND cp.visible is true", userUUID).Joins(joinQuery).Preload("Thumbnail", func(db *gorm.DB) *gorm.DB {
		selectQuery := "club_thumbnails.thumbnail_id, club_thumbnails.club_uuid, ut.path"
		joinQuery := "inner join uploaded_thumbnails as ut using (thumbnail_id)"
		return db.Joins(joinQuery).Select(selectQuery)
	}).Find(&clubPage)

	if err := tx.Error; errors.Is(err, gorm.ErrRecordNotFound) {
		r.logger.Info(err.Error())
		return nil, err
	} else if err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}

	typedPage := clubs.Pages(clubPage)

	return typedPage.ToExternalInfo(), nil
}

func (r *UserRepository) CreateFavorite(userUUID string, clubUUID string) error {
	favorite := users.FavoriteClub{
		UserUUID: userUUID,
		ClubUUID: clubUUID,
	}
	tx := r.db.Create(&favorite)

	if err := tx.Error; err != nil {
		r.logger.Error(err.Error())
		return err
	}

	return nil
}

func (r *UserRepository) DeleteFavorite(userUUID string, clubUUID string) error {
	favorite := users.FavoriteClub{
		UserUUID: userUUID,
		ClubUUID: clubUUID,
	}
	tx := r.db.Where("user_uuid = ? AND club_uuid = ?", userUUID, clubUUID).Delete(&favorite)

	if err := tx.Error; errors.Is(err, gorm.ErrRecordNotFound) {
		r.logger.Info(err.Error())
		return err
	} else if err != nil {
		r.logger.Error(err.Error())
		return err
	}

	return nil
}
