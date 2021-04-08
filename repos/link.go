package repos

import "github.com/lc-tut/club-portal/models"

type ClubLinkRepo interface {
	GetAllLinks() ([]models.ClubLinks, error)
	GetLinkByID(linkID uint32) (*models.ClubLinks, error)
	GetLinkByClubUUID(uuid string) (*models.ClubLinks, error)

	CreateLink(clubUUID string, label string, url string) error

	UpdateLink(clubUUID string, label string, url string) error
}

func (r *Repository) GetAllLinks() ([]models.ClubLinks, error) {
	panic("implement me")
}

func (r *Repository) GetLinkByID(linkID uint32) (*models.ClubLinks, error) {
	panic("implement me")
}

func (r *Repository) GetLinkByClubUUID(uuid string) (*models.ClubLinks, error) {
	panic("implement me")
}

func (r *Repository) CreateLink(clubUUID string, label string, url string) error {
	panic("implement me")
}

func (r *Repository) UpdateLink(clubUUID string, label string, url string) error {
	panic("implement me")
}
