package clubs

import (
	"errors"
	"github.com/lc-tut/club-portal/models/clubs"
	"gorm.io/gorm"
)

type ClubLinkArgs struct {
	Label string
	URL   string
}

type ClubLinkRepo interface {
	GetAllLinks() ([]clubs.ClubLink, error)
	GetLinkByID(linkID uint32) (*clubs.ClubLink, error)

	GetLinksByClubUUID(uuid string) ([]clubs.ClubLink, error)

	CreateLink(clubUUID string, args []ClubLinkArgs) error
	CreateLinkWithTx(tx *gorm.DB, clubUUID string, args []ClubLinkArgs) error

	UpdateLink(clubUUID string, args []ClubLinkArgs) error
	UpdateLinkWithTx(tx *gorm.DB, clubUUID string, args []ClubLinkArgs) error
}

func (r *ClubRepository) GetAllLinks() ([]clubs.ClubLink, error) {
	links := make([]clubs.ClubLink, 0)
	tx := r.db.Find(&links)

	if err := tx.Error; errors.Is(err, gorm.ErrRecordNotFound) {
		r.logger.Info(err.Error())
		return nil, err
	} else if err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}

	return links, nil
}

func (r *ClubRepository) GetLinkByID(linkID uint32) (*clubs.ClubLink, error) {
	link := &clubs.ClubLink{}
	tx := r.db.Where("link_id = ?", linkID).Take(link)

	if err := tx.Error; errors.Is(err, gorm.ErrRecordNotFound) {
		r.logger.Info(err.Error())
		return nil, err
	} else if err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}

	return link, nil
}

func (r *ClubRepository) GetLinksByClubUUID(uuid string) ([]clubs.ClubLink, error) {
	link := make([]clubs.ClubLink, 0)
	tx := r.db.Where("club_uuid = ?", uuid).Find(&link)

	if err := tx.Error; errors.Is(err, gorm.ErrRecordNotFound) {
		r.logger.Info(err.Error())
		return nil, err
	} else if err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}

	return link, nil
}

func (r *ClubRepository) CreateLink(clubUUID string, args []ClubLinkArgs) error {
	links := make([]clubs.ClubLink, len(args))

	for i, arg := range args {
		link := clubs.ClubLink{
			ClubUUID: clubUUID,
			Label:    arg.Label,
			URL:      arg.URL,
		}
		links[i] = link
	}

	tx := r.db.Create(&links)

	if err := tx.Error; err != nil {
		r.logger.Error(err.Error())
		return err
	}

	return nil
}

func (r *ClubRepository) CreateLinkWithTx(tx *gorm.DB, clubUUID string, args []ClubLinkArgs) error {
	links := make([]clubs.ClubLink, len(args))

	for i, arg := range args {
		link := clubs.ClubLink{
			ClubUUID: clubUUID,
			Label:    arg.Label,
			URL:      arg.URL,
		}
		links[i] = link
	}

	if err := tx.Create(&links).Error; err != nil {
		r.logger.Error(err.Error())
		return err
	}

	return nil
}

func (r *ClubRepository) UpdateLink(clubUUID string, args []ClubLinkArgs) error {
	tx := r.db.Where("club_uuid = ?", clubUUID).Delete(&clubs.ClubLink{})

	if err := tx.Error; errors.Is(err, gorm.ErrRecordNotFound) {
		r.logger.Info(err.Error())
		return err
	} else if err != nil {
		r.logger.Error(err.Error())
		return err
	}

	if err := r.CreateLink(clubUUID, args); err != nil {
		return err
	}

	return nil
}

func (r *ClubRepository) UpdateLinkWithTx(tx *gorm.DB, clubUUID string, args []ClubLinkArgs) error {
	tx = tx.Where("club_uuid", clubUUID).Delete(&clubs.ClubLink{})

	if err := tx.Error; errors.Is(err, gorm.ErrRecordNotFound) {
		r.logger.Info(err.Error())
		return err
	} else if err != nil {
		r.logger.Error(err.Error())
		return err
	}

	if err := r.CreateLinkWithTx(tx, clubUUID, args); err != nil {
		return err
	}

	return nil
}
