package users

import (
	"github.com/lc-tut/club-portal/models/clubs"
	"github.com/lc-tut/club-portal/models/users"
)

type UserFavoriteRepo interface {
	GetFavorites(userUUID string) ([]clubs.ClubPageExternalInfo, error)

	CreateFavorite(userUUID string, clubUUID string) error

	DeleteFavorite(userUUID string, clubUUID string) error
}

func (r *UserRepository) GetFavorites(userUUID string) ([]clubs.ClubPageExternalInfo, error) {
	clubPage := make([]clubs.ClubPage, 0)
	joinQuery := "inner join club_pages as cp using (club_uuid)"
	tx := r.db.Table("favorite_clubs").Select("cp.*").Where("user_uuid = ? AND cp.visible is true", userUUID).Joins(joinQuery).Preload("Images").Find(&clubPage)

	if err := tx.Error; err != nil {
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
	tx := r.db.Delete(&favorite)

	if err := tx.Error; err != nil {
		r.logger.Error(err.Error())
		return err
	}

	return nil
}
