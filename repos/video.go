package repos

import "github.com/lc-tut/club-portal/models"

type ClubVideoRepo interface {
	GetVideoByID(videoID uint32) (*models.ClubVideo, error)
	GetVideoByClubUUID(uuid string) (*models.ClubVideo, error)

	CreateVideo(clubUUID string, path string) error

	UpdateVideo(clubUUID string, path string) error
}

func (r *Repository) GetVideoByID(videoID uint32) (*models.ClubVideo, error) {
	panic("implement me")
}

func (r *Repository) GetVideoByClubUUID(uuid string) (*models.ClubVideo, error) {
	panic("implement me")
}

func (r *Repository) CreateVideo(clubUUID string, path string) error {
	panic("implement me")
}

func (r *Repository) UpdateVideo(clubUUID string, path string) error {
	panic("implement me")
}
