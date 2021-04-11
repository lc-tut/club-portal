package repos

import (
	"errors"
	"github.com/lc-tut/club-portal/models"
	"github.com/lc-tut/club-portal/router/utils"
	"gorm.io/gorm"
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

	for _, arg := range args {
		t := models.ClubTime{
			TimeID: arg.TimeID,
			Date:   arg.Date,
			Time:   arg.Time,
		}
		timeModels = append(timeModels, t)
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

	for _, arg := range args {
		t := models.ClubTime{
			Date: arg.Date,
			Time: arg.Time,
		}
		timeModels = append(timeModels, t)
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

	for _, arg := range args {
		placeModel := models.ClubPlace{
			PlaceID: arg.PlaceID,
			Place:   arg.Place,
		}
		placeModels = append(placeModels, placeModel)
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

	for _, arg := range args {
		placeModel := models.ClubPlace{
			Place: arg.Place,
		}
		placeModels = append(placeModels, placeModel)
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

func (r *Repository) CreateClubActivityDetail(uuid string, args []ClubActivityDetailArgs) error {
	adModels := make([]models.ActivityDetail, len(args))
	crArgs := make([]ClubRemarkArgs, len(args))

	for _, arg := range args {
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
		adModels = append(adModels, model)
		crArgs = append(crArgs, crArg)
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

type ClubTimeAndPlaceRepo interface {
	GetClubTimeAndPlaces(uuid string) ([]models.ClubTimeAndPlace, error)
	CreateClubTimeAndPlaces(uuid string, tps []models.ClubTimeAndPlace) error
}

func (r *Repository) GetClubTimeAndPlaces(uuid string) ([]models.ClubTimeAndPlace, error) {
	tps := make([]models.ClubTimeAndPlace, 0)
	rows, err := r.db.Table("activity_details").Where("club_uuid = ?", uuid).Select("club_times.*, club_places.*").Joins("inner join club_times on activity_details.time_id = club_times.time_id").Joins("inner join club_places on activity_details.place_id = club_places.place_id").Rows()

	if err != nil {
		return nil, err
	}

	var ct models.ClubTime
	var cp models.ClubPlace

	for rows.Next() {
		if err := r.db.ScanRows(rows, &ct); err != nil {
			return nil, err
		}

		if err := r.db.ScanRows(rows, &cp); err != nil {
			return nil, err
		}

		tp := models.ClubTimeAndPlace{
			TimeID:  ct.GetTimeID(),
			Date:    ct.GetDate(),
			Time:    ct.GetTime(),
			PlaceID: cp.GetPlaceID(),
			Place:   cp.GetPlace(),
		}
		tps = append(tps, tp)
	}

	return tps, nil
}

func (r *Repository) CreateClubTimeAndPlaces(uuid string, tps []models.ClubTimeAndPlace) error {
	// FIXME: probably causes error

	ctArgs := make([]ClubTimeArgs, len(tps))
	cpArgs := make([]ClubPlaceArgs, len(tps))
	adArgs := make([]ClubActivityDetailArgs, len(tps))

	for _, tp := range tps {
		ctArg := ClubTimeArgs{
			TimeID: tp.TimeID,
			Date:   tp.Date,
			Time:   tp.Time,
		}
		cpArg := ClubPlaceArgs{
			PlaceID: tp.PlaceID,
			Place:   tp.Place,
		}
		adArg := ClubActivityDetailArgs{
			TimeID:       tp.TimeID,
			TimeRemarks:  utils.NilToEmptyString(tp.TimeRemarks),
			PlaceID:      tp.PlaceID,
			PlaceRemarks: utils.NilToEmptyString(tp.PlaceRemarks),
		}
		ctArgs = append(ctArgs, ctArg)
		cpArgs = append(cpArgs, cpArg)
		adArgs = append(adArgs, adArg)
	}

	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := r.CreateClubTime(ctArgs); err != nil {
			return err
		}

		if err := r.CreateClubPlace(cpArgs); err != nil {
			return err
		}

		if err := r.CreateClubActivityDetail(uuid, adArgs); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
