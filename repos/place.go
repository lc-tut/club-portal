package repos

import (
	"errors"
	"github.com/lc-tut/club-portal/models"
)

type ClubPlaceArgs struct {
	PlaceID uint32
	Place   string
}

type ClubPlaceRepo interface {
	GetPlaceByID(placeID uint32) (*models.ClubPlace, error)

	GetPlacesByClubUUID(uuid string) ([]models.ClubPlace, error)

	CreatePlace(args []ClubPlaceArgs) error

	UpdatePlace(placeID []uint32, args []ClubPlaceArgs) error
}

func (r *Repository) GetPlaceByID(placeID uint32) (*models.ClubPlace, error) {
	place := &models.ClubPlace{}
	tx := r.db.Where("place_id = ?", placeID).Take(place)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return place, nil
}

func (r *Repository) GetPlacesByClubUUID(uuid string) ([]models.ClubPlace, error) {
	places := make([]models.ClubPlace, 0)
	tx := r.db.Where("club_uuid = ?", uuid).Find([]models.ActivityDetail{}).Joins("inner join club_places on activity_details.place_id = club_places.place_id").Scan(places)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return places, nil
}

func (r *Repository) CreatePlace(args []ClubPlaceArgs) error {
	placeModels := make([]models.ClubPlace, len(args))

	for i, arg := range args {
		placeModel := models.ClubPlace{
			PlaceID: arg.PlaceID,
			Place:   arg.Place,
		}
		placeModels[i] = placeModel
	}

	tx := r.db.Create(&placeModels)

	if err := tx.Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdatePlace(placeID []uint32, args []ClubPlaceArgs) error {
	if len(placeID) != len(args) {
		return errors.New("invalid argument")
	}

	placeModels := make([]models.ClubPlace, len(placeID))

	for i, arg := range args {
		placeModel := models.ClubPlace{
			Place: arg.Place,
		}
		placeModels[i] = placeModel
	}

	tx := r.db.Model(&models.ClubPlace{}).Where("place_id = ?", placeID).Updates(&placeModels)

	if err := tx.Error; err != nil {
		return err
	}

	return nil
}
