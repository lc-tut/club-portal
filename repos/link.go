package repos

import (
	"github.com/lc-tut/club-portal/models"
	"gorm.io/gorm"
)

type ClubLinkArgs struct {
	Label string
	URL   string
}

type ClubLinkRepo interface {
	GetAllLinks() ([]models.ClubLink, error)
	GetLinkByID(linkID uint32) (*models.ClubLink, error)

	GetLinksByClubUUID(uuid string) ([]models.ClubLink, error)

	CreateLink(clubUUID string, args []ClubLinkArgs) error
	CreateLinkWithTx(tx *gorm.DB, clubUUID string, args []ClubLinkArgs) error

	UpdateLink(clubUUID string, args []ClubLinkArgs) error
}

func (r *Repository) GetAllLinks() ([]models.ClubLink, error) {
	links := make([]models.ClubLink, 0)
	tx := r.db.Find(&links)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return links, nil
}

func (r *Repository) GetLinkByID(linkID uint32) (*models.ClubLink, error) {
	link := &models.ClubLink{}
	tx := r.db.Where("link_id = ?", linkID).Take(link)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return link, nil
}

func (r *Repository) GetLinksByClubUUID(uuid string) ([]models.ClubLink, error) {
	link := make([]models.ClubLink, 0)
	tx := r.db.Where("club_uuid = ?", uuid).Find(link)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return link, nil
}

func (r *Repository) CreateLink(clubUUID string, args []ClubLinkArgs) error {
	links := make([]models.ClubLink, len(args))

	for i, arg := range args {
		link := models.ClubLink{
			ClubUUID: clubUUID,
			Label:    arg.Label,
			URL:      arg.URL,
		}
		links[i] = link
	}

	tx := r.db.Create(&links)

	if err := tx.Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) CreateLinkWithTx(tx *gorm.DB, clubUUID string, args []ClubLinkArgs) error {
	links := make([]models.ClubLink, len(args))

	for i, arg := range args {
		link := models.ClubLink{
			ClubUUID: clubUUID,
			Label:    arg.Label,
			URL:      arg.URL,
		}
		links[i] = link
	}

	if err := tx.Create(&links).Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateLink(clubUUID string, args []ClubLinkArgs) error {
	links := make([]models.ClubLink, len(args))

	for i, arg := range args {
		link := models.ClubLink{
			ClubUUID: clubUUID,
			Label:    arg.Label,
			URL:      arg.URL,
		}
		links[i] = link
	}

	tx := r.db.Model(&models.ClubLink{}).Where("club_uuid = ?", clubUUID).Updates(&links)

	if err := tx.Error; err == nil {
		return err
	}

	return nil
}
