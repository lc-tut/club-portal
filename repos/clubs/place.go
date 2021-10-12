package clubs

import (
	"errors"
	"github.com/lc-tut/club-portal/models/clubs"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ClubPlaceArgs struct {
	PlaceID uint32
	Place   string
}

type ClubPlaceRepo interface {
	GetPlaceByID(placeID uint32) (*clubs.ClubPlace, error)

	GetPlacesByClubUUID(uuid string) ([]clubs.ClubPlace, error)

	CreatePlace(args []ClubPlaceArgs) error
	CreatePlaceWithTx(tx *gorm.DB, args []ClubPlaceArgs) error
}

func (r *ClubRepository) GetPlaceByID(placeID uint32) (*clubs.ClubPlace, error) {
	place := &clubs.ClubPlace{}
	tx := r.db.Where("place_id = ?", placeID).Take(place)

	if err := tx.Error; errors.Is(err, gorm.ErrRecordNotFound) {
		r.logger.Info(err.Error())
		return nil, err
	} else if err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}

	return place, nil
}

func (r *ClubRepository) GetPlacesByClubUUID(uuid string) ([]clubs.ClubPlace, error) {
	places := make([]clubs.ClubPlace, 0)
	tx := r.db.Where("club_uuid = ?", uuid).Find([]clubs.ActivityDetail{}).Joins("inner join club_places on activity_details.place_id = club_places.place_id").Scan(&places)

	if err := tx.Error; errors.Is(err, gorm.ErrRecordNotFound) {
		r.logger.Info(err.Error())
		return nil, err
	} else if err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}

	return places, nil
}

func (r *ClubRepository) CreatePlace(args []ClubPlaceArgs) error {
	placeModels := make([]clubs.ClubPlace, len(args))

	for i, arg := range args {
		placeModel := clubs.ClubPlace{
			PlaceID: arg.PlaceID,
			Place:   arg.Place,
		}
		placeModels[i] = placeModel
	}

	tx := r.db.Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&placeModels)

	if err := tx.Error; err != nil {
		r.logger.Error(err.Error())
		return err
	}

	return nil
}

func (r *ClubRepository) CreatePlaceWithTx(tx *gorm.DB, args []ClubPlaceArgs) error {
	placeModels := make([]clubs.ClubPlace, len(args))

	for i, arg := range args {
		placeModel := clubs.ClubPlace{
			PlaceID: arg.PlaceID,
			Place:   arg.Place,
		}
		placeModels[i] = placeModel
	}

	if err := tx.Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&placeModels).Error; err != nil {
		r.logger.Error(err.Error())
		return err
	}

	return nil
}
