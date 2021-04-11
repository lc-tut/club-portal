package repos

import (
	"github.com/lc-tut/club-portal/consts"
	"github.com/lc-tut/club-portal/models"
	"github.com/lc-tut/club-portal/utils"
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
	GetPageByClubSlug(clubSlug string) (*models.ClubPage, error)

	CreatePage(uuid string, args ClubPageCreateArgs) error

	UpdatePageByClubUUID(uuid string, args ClubPageUpdateArgs) error
	UpdatePageByClubSlug(clubSlug string, args ClubPageUpdateArgs) error
}

func (r *Repository) GetAllPages() ([]models.ClubPage, error) {
	page := make([]models.ClubPage, 0)
	tx := r.db.Preload("Contents").Preload("Links").Preload("Schedules").Preload("Achievements").Preload("Images").Preload("Videos").Preload("ActivityDetails").Find(&page)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return page, nil
}

func (r *Repository) GetPageByClubUUID(uuid string) (*models.ClubPage, error) {
	page := &models.ClubPage{}
	tx := r.db.Where("club_uuid = ?", uuid).Preload("Contents").Preload("Links").Preload("Schedules").Preload("Achievements").Preload("Images").Preload("Videos").Preload("ActivityDetails").Take(page)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return page, nil
}

func (r *Repository) GetPageByClubSlug(clubSlug string) (*models.ClubPage, error) {
	page := &models.ClubPage{}
	tx := r.db.Where("club_id = ?", clubSlug).Preload("Contents").Preload("Links").Preload("Schedules").Preload("Achievements").Preload("Images").Preload("Videos").Preload("ActivityDetails").Take(page)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return page, nil
}

func (r *Repository) CreatePage(uuid string, args ClubPageCreateArgs) error {
	slug, err := utils.GenerateRand15()

	if err != nil {
		return err
	}

	page := &models.ClubPage{
		ClubUUID:    uuid,
		ClubSlug:    slug,
		Name:        args.Name,
		Description: args.Desc,
		Campus:      uint8(args.Campus),
		ClubType:    uint8(args.ClubType),
		Visible:     uint8(args.Visible),
	}

	tx := r.db.Create(page)

	if err := tx.Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdatePageByClubUUID(uuid string, args ClubPageUpdateArgs) error {
	tx := r.db.Model(&models.ClubPage{}).Where("club_uuid = ?", uuid).Updates(args)

	if err := tx.Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdatePageByClubSlug(clubSlug string, args ClubPageUpdateArgs) error {
	tx := r.db.Model(&models.ClubPage{}).Where("club_slug = ?", clubSlug).Updates(args)

	if err := tx.Error; err != nil {
		return err
	}

	return nil
}
