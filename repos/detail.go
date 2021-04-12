package repos

import (
	"errors"
	"github.com/lc-tut/club-portal/models"
)

type ClubTimeArgs struct {
	TimeID uint32
	Date   string
	Time   string
}

type ClubTimeRepo interface {
	GetClubTimeByID(timeID uint32) (*models.ClubTime, error)

	GetClubTimesByClubUUID(uuid string) ([]models.ClubTime, error)

	CreateClubTime(args []ClubTimeArgs) error

	UpdateClubTime(timeID []uint32, args []ClubTimeArgs) error
}

func (r *Repository) GetClubTimeByID(timeID uint32) (*models.ClubTime, error) {
	time := &models.ClubTime{}
	tx := r.db.Where("time_id = ?", timeID).Take(time)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return time, nil
}

func (r *Repository) GetClubTimesByClubUUID(uuid string) ([]models.ClubTime, error) {
	times := make([]models.ClubTime, 0)
	tx := r.db.Where("club_uuid = ?", uuid).Find([]models.ActivityDetail{}).Joins("inner join club_times on activity_details.time_id = club_times.time_id").Scan(times)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return times, nil
}

func (r *Repository) CreateClubTime(args []ClubTimeArgs) error {
	timeModels := make([]models.ClubTime, len(args))

	for i, arg := range args {
		t := models.ClubTime{
			TimeID: arg.TimeID,
			Date:   arg.Date,
			Time:   arg.Time,
		}
		timeModels[i] = t
	}

	tx := r.db.Create(&timeModels)

	if err := tx.Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateClubTime(timeID []uint32, args []ClubTimeArgs) error {
	if len(timeID) != len(args) {
		return errors.New("invalid argument")
	}

	timeModels := make([]models.ClubTime, len(timeID))

	for i, arg := range args {
		t := models.ClubTime{
			Date: arg.Date,
			Time: arg.Time,
		}
		timeModels[i] = t
	}

	tx := r.db.Model(&models.ClubTime{}).Where("time_id = ?", timeID).Updates(&timeModels)

	if err := tx.Error; err != nil {
		return err
	}

	return nil
}

type ClubPlaceArgs struct {
	PlaceID uint32
	Place   string
}

type ClubPlaceRepo interface {
	GetClubPlaceByID(placeID uint32) (*models.ClubPlace, error)

	GetClubPlacesByClubUUID(uuid string) ([]models.ClubPlace, error)

	CreateClubPlace(args []ClubPlaceArgs) error

	UpdateClubPlace(placeID []uint32, args []ClubPlaceArgs) error
}

func (r *Repository) GetClubPlaceByID(placeID uint32) (*models.ClubPlace, error) {
	place := &models.ClubPlace{}
	tx := r.db.Where("place_id = ?", placeID).Take(place)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return place, nil
}

func (r *Repository) GetClubPlacesByClubUUID(uuid string) ([]models.ClubPlace, error) {
	places := make([]models.ClubPlace, 0)
	tx := r.db.Where("club_uuid = ?", uuid).Find([]models.ActivityDetail{}).Joins("inner join club_places on activity_details.place_id = club_places.place_id").Scan(places)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return places, nil
}

func (r *Repository) CreateClubPlace(args []ClubPlaceArgs) error {
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

func (r *Repository) UpdateClubPlace(placeID []uint32, args []ClubPlaceArgs) error {
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

type ClubActivityDetailArgs struct {
	TimeID       uint32
	TimeRemarks  string
	PlaceID      uint32
	PlaceRemarks string
}

type ClubActivityDetailRepo interface {
	GetClubActivityDetail(uuid string) ([]models.ActivityDetail, error)

	GetAllRelations(uuid string) ([]models.DetailRelations, error)

	CreateClubActivityDetail(uuid string, args []ClubActivityDetailArgs) error
}

func (r *Repository) GetClubActivityDetail(uuid string) ([]models.ActivityDetail, error) {
	details := make([]models.ActivityDetail, 0)
	tx := r.db.Where("club_uuid = ?", uuid).Find(&details)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return details, nil
}

func (r *Repository) GetAllRelations(uuid string) ([]models.DetailRelations, error) {
	joinQuery1 := "inner join club_times on using (time_id)"
	joinQuery2 := "inner join club_places on using (place_id)"
	joinQuery3 := "inner join club_remarks on using (club_uuid, time_id, place_id)"

	rows, err := r.db.Table("activity_details").Where("activity_details.club_uuid = ?", uuid).Joins(joinQuery1).Joins(joinQuery2).Joins(joinQuery3).Rows()

	if err != nil {
		return nil, err
	}

	relations := make([]models.DetailRelations, len(uuid))
	i := 0

	for rows.Next() {
		var relation models.DetailRelations
		if err := r.db.ScanRows(rows, &relation); err != nil {
			return nil, err
		}
		relations[i] = relation
		i++
	}

	return relations, nil
}

func (r *Repository) CreateClubActivityDetail(uuid string, args []ClubActivityDetailArgs) error {
	adModels := make([]models.ActivityDetail, len(args))
	crArgs := make([]ClubRemarkArgs, len(args))

	for i, arg := range args {
		model := models.ActivityDetail{
			TimeID:   arg.TimeID,
			PlaceID:  arg.PlaceID,
			ClubUUID: uuid,
		}
		crArg := ClubRemarkArgs{
			ClubUUID:     uuid,
			TimeID:       arg.TimeID,
			PlaceID:      arg.PlaceID,
			TimeRemarks:  arg.TimeRemarks,
			PlaceRemarks: arg.PlaceRemarks,
		}
		adModels[i] = model
		crArgs[i] = crArg
	}

	tx := r.db.Create(&adModels)

	if err := tx.Error; err != nil {
		return err
	}

	if err := r.CreateRemark(crArgs); err != nil {
		return err
	}

	return nil
}
