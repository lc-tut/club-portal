package repos

import (
	"errors"
	"github.com/google/uuid"
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
	GetPageByClubID(clubID string) (*models.ClubPage, error)

	CreatePage(args ClubPageCreateArgs) error

	UpdatePage(uuid string, clubID string, args ClubPageUpdateArgs) error
}

// TODO: Get records on foreign key.

func (r *Repository) GetAllPages() ([]models.ClubPage, error) {
	page := make([]models.ClubPage, 0)
	tx := r.db.Find(&page)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return page, nil
}

func (r *Repository) GetPageByClubUUID(uuid string) (*models.ClubPage, error) {
	page := &models.ClubPage{}
	tx := r.db.Where("club_uuid = ?", uuid).First(page)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return page, nil
}

func (r *Repository) GetPageByClubID(clubID string) (*models.ClubPage, error) {
	page := &models.ClubPage{}
	tx := r.db.Where("club_id = ?", clubID).First(page)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return page, nil
}

func (r *Repository) CreatePage(args ClubPageCreateArgs) error {
	clubUUID, err := uuid.NewUUID()

	if err != nil {
		return err
	}

	slug, err := utils.GenerateRand15()

	if err != nil {
		return err
	}

	page := &models.ClubPage{
		ClubUUID:    clubUUID.String(),
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

func (r *Repository) UpdatePage(uuid string, clubSlug string, args ClubPageUpdateArgs) error {
	if uuid == "" && clubSlug == "" {
		return errors.New("no uuid or clubSlug")
	}

	var searchID string
	var whereSQL string

	if clubSlug != "" {
		searchID = clubSlug
		whereSQL = "club_slug = ?"
	} else {
		searchID = uuid
		whereSQL = "club_uuid = ?"
	}

	tx := r.db.Model(&models.ClubPage{}).Where(whereSQL, searchID).Updates(args)

	if err := tx.Error; err != nil {
		return err
	}

	return nil
}
