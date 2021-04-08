package repos

import "github.com/lc-tut/club-portal/models"

type ClubImageRepo interface {
	GetImageByID(imageID uint32) (*models.ClubImage, error)
	GetImageByClubUUID(uuid string) (*models.ClubImage, error)

	CreateImage(clubUUUID string, path string) error

	UpdateImage(clubUUID string, path string) error
}

func (r *Repository) GetImageByID(imageID uint32) (*models.ClubImage, error) {
	panic("implement me")
}

func (r *Repository) GetImageByClubUUID(uuid string) (*models.ClubImage, error) {
	panic("implement me")
}

func (r *Repository) CreateImage(clubUUUID string, path string) error {
	panic("implement me")
}

func (r *Repository) UpdateImage(clubUUID string, path string) error {
	panic("implement me")
}
