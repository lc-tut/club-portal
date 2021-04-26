package clubs

import (
	"github.com/lc-tut/club-portal/models/clubs"
	"gorm.io/gorm"
)

type ClubContentRepo interface {
	GetContentByID(contentID uint32) (*clubs.ClubContent, error)

	GetContentsByClubUUID(uuid string) ([]clubs.ClubContent, error)

	CreateContent(clubUUID string, content []string) error
	CreateContentWithTx(tx *gorm.DB, clubUUID string, contents []string) error

	UpdateContent(clubUUID string, content []string) error
	UpdateContentWithTx(tx *gorm.DB, clubUUID string, contents []string) error
}

func (r *ClubRepository) GetContentByID(contentID uint32) (*clubs.ClubContent, error) {
	content := &clubs.ClubContent{}
	tx := r.db.Where("content_id = ?", contentID).Take(content)

	if err := tx.Error; err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}

	return content, nil
}

func (r *ClubRepository) GetContentsByClubUUID(uuid string) ([]clubs.ClubContent, error) {
	content := make([]clubs.ClubContent, 0)
	tx := r.db.Where("club_uuid = ?", uuid).Find(&content)

	if err := tx.Error; err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}

	return content, nil
}

func (r *ClubRepository) CreateContent(clubUUID string, contents []string) error {
	contModels := make([]clubs.ClubContent, len(contents))

	for i, c := range contents {
		cont := clubs.ClubContent{
			ClubUUID: clubUUID,
			Content:  c,
		}
		contModels[i] = cont
	}

	tx := r.db.Create(&contModels)

	if err := tx.Error; err != nil {
		r.logger.Error(err.Error())
		return err
	}

	return nil
}

func (r *ClubRepository) CreateContentWithTx(tx *gorm.DB, clubUUID string, contents []string) error {
	contModels := make([]clubs.ClubContent, len(contents))

	for i, c := range contents {
		cont := clubs.ClubContent{
			ClubUUID: clubUUID,
			Content:  c,
		}
		contModels[i] = cont
	}

	if err := tx.Create(&contModels).Error; err != nil {
		r.logger.Error(err.Error())
		return err
	}

	return nil
}

// FIXME: use delete -> create to update contents.

func (r *ClubRepository) UpdateContent(clubUUID string, contents []string) error {
	length := len(contents)

	if length == 0 {
		return nil
	}

	contModels := make([]clubs.ClubContent, length)

	for i, c := range contents {
		cont := clubs.ClubContent{
			ClubUUID: clubUUID,
			Content:  c,
		}
		contModels[i] = cont
	}

	tx := r.db.Model(&clubs.ClubContent{}).Where("club_uuid = ?", clubUUID).Updates(contModels)

	if err := tx.Error; err != nil {
		r.logger.Error(err.Error())
		return err
	}

	return nil
}

func (r *ClubRepository) UpdateContentWithTx(tx *gorm.DB, clubUUID string, contents []string) error {
	length := len(contents)

	if length == 0 {
		return nil
	}

	if err := tx.Where("club_uuid = ?", clubUUID).Delete(&clubs.ClubContent{}).Error; err != nil {
		r.logger.Error(err.Error())
		return err
	}

	if err := r.CreateContentWithTx(tx, clubUUID, contents); err != nil {
		return err
	}

	return nil
}
