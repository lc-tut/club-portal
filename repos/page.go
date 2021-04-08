package repos

import (
	"github.com/lc-tut/club-portal/consts"
	"github.com/lc-tut/club-portal/models"
)

type ClubPageCreateArgs struct {
	Name     string
	Desc     string
	Campus   consts.CampusType
	ClubType consts.ClubType
	Visible  consts.Visibility
}

type ClubPageUpdateArgs struct {
	Desc    string
	Visible consts.Visibility
}

type ClubPageRepo interface {
	GetAllPages() ([]models.ClubPage, error)
	GetPageByClubUUID(uuid string) (*models.ClubPage, error)
	GetPageByClubID(clubID string) (*models.ClubPage, error)

	CreatePage(args ClubPageCreateArgs) error

	UpdatePage(args ClubPageUpdateArgs) error
}

func (r *Repository) GetAllPages() ([]models.ClubPage, error) {
	page := make([]models.ClubPage, 0)
	tx := r.db.Find(&page)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return page, nil
}

func (r *Repository) GetPageByClubUUID(uuid string) (*models.ClubPage, error) {
	panic("implement me")
}

func (r *Repository) GetPageByClubID(clubID string) (*models.ClubPage, error) {
	panic("implement me")
}

func (r *Repository) CreatePage(args ClubPageCreateArgs) error {
	panic("implement me")
}

func (r *Repository) UpdatePage(args ClubPageUpdateArgs) error {
	panic("implement me")
}
