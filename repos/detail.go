package repos

import "github.com/lc-tut/club-portal/models"

type ClubTimeRepo interface {
	GetClubTimeByID(timeID uint32) (*models.ClubTime, error)

	GetClubTimesByClubUUID(uuid string) ([]models.ClubTime, error)

	CreateClubTime(clubUUID string, date string, time string, remarks string) error

	UpdateClubTime(clubUUID string, date string, time string, remarks string) error
}

func (r *Repository) GetClubTimeByID(timeID uint32) (*models.ClubTime, error) {
	panic("implement me")
}

func (r *Repository) GetClubTimesByClubUUID(uuid string) ([]models.ClubTime, error) {
	panic("implement me")
}

func (r *Repository) CreateClubTime(clubUUID string, date string, time string, remarks string) error {
	panic("implement me")
}

func (r *Repository) UpdateClubTime(clubUUID string, date string, time string, remarks string) error {
	panic("implement me")
}

type ClubPlaceRepo interface {
	GetClubPlaceByID(placeID uint32) (*models.ClubPlace, error)

	GetClubPlacesByClubUUID(uuid string) ([]models.ClubPlace, error)

	CreateClubPlace(clubUUID string, place string, remarks string) error

	UpdateClubPlace(clubUUID string, place string, remarks string) error
}

func (r *Repository) GetClubPlaceByID(placeID uint32) (*models.ClubPlace, error) {
	panic("implement me")
}

func (r *Repository) GetClubPlacesByClubUUID(uuid string) ([]models.ClubPlace, error) {
	panic("implement me")
}

func (r *Repository) CreateClubPlace(clubUUID string, place string, remarks string) error {
	panic("implement me")
}

func (r *Repository) UpdateClubPlace(clubUUID string, place string, remarks string) error {
	panic("implement me")
}

type ClubDetailRepo interface {
	GetClubDetailsByClubUUID(uuid string) ([]models.ClubDetail, error)
}

func (r *Repository) GetClubDetailsByClubUUID(uuid string) ([]models.ClubDetail, error) {
	panic("implement me")
}
