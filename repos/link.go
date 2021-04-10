package repos

import "github.com/lc-tut/club-portal/models"

type ClubLinkRepo interface {
	GetAllLinks() ([]models.ClubLink, error)
	GetLinkByID(linkID uint32) (*models.ClubLink, error)

	GetLinksByClubUUID(uuid string) ([]models.ClubLink, error)

	CreateLink(clubUUID string, label string, url string) error

	UpdateLink(clubUUID string, label string, url string) error
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

func (r *Repository) CreateLink(clubUUID string, label string, url string) error {
	link := &models.ClubLink{
		ClubUUID: clubUUID,
		Label:    label,
		URL:      url,
	}

	tx := r.db.Create(link)

	if err := tx.Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateLink(clubUUID string, label string, url string) error {
	updateLink := &models.ClubLink{
		Label: label,
		URL:   url,
	}

	tx := r.db.Model(updateLink).Where("club_uuid = ?", clubUUID).Updates(updateLink)

	if err := tx.Error; err == nil {
		return err
	}

	return nil
}
