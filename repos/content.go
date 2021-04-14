package repos

import (
	"github.com/lc-tut/club-portal/models"
	"gorm.io/gorm"
)

type ClubContentRepo interface {
	GetContentByID(contentID uint32) (*models.ClubContent, error)

	GetContentsByClubUUID(uuid string) ([]models.ClubContent, error)

	CreateContent(clubUUID string, content []string) error
	CreateContentWithTx(tx *gorm.DB, clubUUID string, contents []string) error

	UpdateContent(clubUUID string, content []string) error
	UpdateContentWithTx(tx *gorm.DB, clubUUID string, contents []string) error
}

func (r *Repository) GetContentByID(contentID uint32) (*models.ClubContent, error) {
	content := &models.ClubContent{}
	tx := r.db.Where("content_id = ?", contentID).Take(content)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return content, nil
}

func (r *Repository) GetContentsByClubUUID(uuid string) ([]models.ClubContent, error) {
	content := make([]models.ClubContent, 0)
	tx := r.db.Where("club_uuid = ?", uuid).Find(content)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return content, nil
}

func (r *Repository) CreateContent(clubUUID string, contents []string) error {
	contModels := make([]models.ClubContent, len(contents))

	for i, c := range contents {
		cont := models.ClubContent{
			ClubUUID: clubUUID,
			Content:  c,
		}
		contModels[i] = cont
	}

	tx := r.db.Create(&contModels)

	if err := tx.Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) CreateContentWithTx(tx *gorm.DB, clubUUID string, contents []string) error {
	contModels := make([]models.ClubContent, len(contents))

	for i, c := range contents {
		cont := models.ClubContent{
			ClubUUID: clubUUID,
			Content:  c,
		}
		contModels[i] = cont
	}

	if err := tx.Create(&contModels).Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateContent(clubUUID string, contents []string) error {
	length := len(contents)

	if length == 0 {
		return nil
	}

	contModels := make([]models.ClubContent, length)

	for i, c := range contents {
		cont := models.ClubContent{
			ClubUUID: clubUUID,
			Content:  c,
		}
		contModels[i] = cont
	}

	tx := r.db.Model(&models.ClubContent{}).Where("club_uuid = ?", clubUUID).Updates(contModels)

	if err := tx.Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateContentWithTx(tx *gorm.DB, clubUUID string, contents []string) error {
	length := len(contents)

	if length == 0 {
		return nil
	}

	if err := tx.Where("club_uuid = ?", clubUUID).Delete(&models.ClubContent{}).Error; err != nil {
		return err
	}

	if err := r.CreateContentWithTx(tx, clubUUID, contents); err != nil {
		return err
	}

	return nil
}
