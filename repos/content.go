package repos

import "github.com/lc-tut/club-portal/models"

type ClubContentRepo interface {
	GetContentByID(activityID uint32) (*models.ClubContent, error)
	GetContentByClubUUID(uuid string) (*models.ClubContent, error)

	CreateContent(clubUUID string, content string) error

	UpdateContent(clubUUID string, content string) error
}

func (r *Repository) GetContentByID(activityID uint32) (*models.ClubContent, error) {
	panic("implement me")
}

func (r *Repository) GetContentByClubUUID(uuid string) (*models.ClubContent, error) {
	panic("implement me")
}

func (r *Repository) CreateContent(clubUUID string, content string) error {
	panic("implement me")
}

func (r *Repository) UpdateContent(clubUUID string, content string) error {
	panic("implement me")
}
