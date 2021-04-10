package repos

import (
	"github.com/lc-tut/club-portal/models"
	"github.com/lc-tut/club-portal/utils"
	"gorm.io/gorm"
)

type ClubTimeRepo interface {
	GetClubTimeByID(timeID uint32) (*models.ClubTime, error)

	GetClubTimesByClubUUID(uuid string) ([]models.ClubTime, error)

	CreateClubTime(date string, time string, remarks string) error

	UpdateClubTime(timeID uint32, date string, time string, remarks string) error
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

func (r *Repository) UpdateClubTime(timeID uint32, date string, time string, remarks string) error {
	timeModel := &models.ClubTime{
		Date:    date,
		Time:    time,
		Remarks: utils.ToNullString(remarks),
	}

	tx := r.db.Model(timeModel).Where("time_id = ?", timeID).Updates(timeModel)

	if err := tx.Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) CreateClubTime(date string, time string, remarks string) error {
	timeModel := &models.ClubTime{
		Date:    date,
		Time:    time,
		Remarks: utils.ToNullString(remarks),
	}

	tx := r.db.Create(timeModel)

	if err := tx.Error; err != nil {
		return err
	}

	return nil
}

type ClubPlaceRepo interface {
	GetClubPlaceByID(placeID uint32) (*models.ClubPlace, error)

	GetClubPlacesByClubUUID(uuid string) ([]models.ClubPlace, error)

	CreateClubPlace(place string, remarks string) error

	UpdateClubPlace(placeID uint32, place string, remarks string) error
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

func (r *Repository) UpdateClubPlace(placeID uint32, place string, remarks string) error {
	placeModel := &models.ClubPlace{
		Place:   place,
		Remarks: utils.ToNullString(remarks),
	}

	tx := r.db.Model(placeModel).Where("place_id = ?", placeID).Updates(placeModel)

	if err := tx.Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) CreateClubPlace(place string, remarks string) error {
	placeModel := &models.ClubPlace{
		Place:   place,
		Remarks: utils.ToNullString(remarks),
	}

	tx := r.db.Create(placeModel)

	if err := tx.Error; err != nil {
		return err
	}

	return nil
}

type ClubActivityDetailRepo interface {
	GetClubActivityDetailByClubUUID(uuid string) ([]models.ActivityDetail, error)

	CreateClubActivityDetail(uuid string, clubTime models.ClubTime, clubPlace models.ClubPlace) error
}

func (r *Repository) GetClubActivityDetailByClubUUID(uuid string) ([]models.ActivityDetail, error) {
	details := make([]models.ActivityDetail, 0)
	tx := r.db.Where("club_uuid = ?", uuid).Find(details)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return details, nil
}

func (r *Repository) CreateClubActivityDetail(uuid string, clubTime models.ClubTime, clubPlace models.ClubPlace) error {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := r.CreateClubTime(clubTime.GetDate(), clubTime.GetTime(), clubTime.GetRemarks()); err != nil {
			return err
		}

		var timeID uint32

		tx.Select("last_insert_id()").Scan(&timeID)

		if err := r.CreateClubPlace(clubPlace.GetPlace(), clubPlace.GetRemarks()); err != nil {
			return err
		}

		var placeID uint32

		tx.Select("last_insert_id()").Scan(&placeID)

		actDetail := &models.ActivityDetail{
			TimeID:   timeID,
			PlaceID:  placeID,
			ClubUUID: uuid,
		}

		if err := tx.Create(actDetail).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

type ClubTimeAndPlaceRepo interface {
	GetClubDetailsByClubUUID(uuid string) ([]models.ClubTimeAndPlace, error)
}

func (r *Repository) GetClubDetailsByClubUUID(uuid string) ([]models.ClubTimeAndPlace, error) {
	tps := make([]models.ClubTimeAndPlace, 0)
	rows, err := r.db.Table("activity_details").Where("club_uuid = ?", uuid).Select("club_times.*, club_places.*").Joins("inner join club_times on activity_details.time_id = club_times.time_id").Joins("inner join club_places on activity_details.place_id = club_places.place_id").Rows()

	if err != nil {
		return nil, err
	}

	var clubTime models.ClubTime
	var clubPlace models.ClubPlace

	for rows.Next() {
		if err := r.db.ScanRows(rows, &clubTime); err != nil {
			return nil, err
		}

		if err := r.db.ScanRows(rows, &clubPlace); err != nil {
			return nil, err
		}

		tp := models.ClubTimeAndPlace{
			ClubTime:  clubTime,
			ClubPlace: clubPlace,
		}
		tps = append(tps, tp)
	}

	return tps, nil
}
