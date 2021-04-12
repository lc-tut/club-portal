package repos

import "github.com/lc-tut/club-portal/models"

type ClubContentRepo interface {
	GetContentByID(contentID uint32) (*models.ClubContent, error)

	GetContentsByClubUUID(uuid string) ([]models.ClubContent, error)

	CreateContent(clubUUID string, content []string) error

	UpdateContent(clubUUID string, content []string) error
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

func (r *Repository) UpdateContent(clubUUID string, contents []string) error {
	contModels := make([]models.ClubContent, len(contents))

	for i, c := range contents {
		cont := models.ClubContent{
			ClubUUID: clubUUID,
			Content:  c,
		}
		contModels[i] = cont
	}

	tx := r.db.Model(&models.ClubContent{}).Where("club_uuid = ?", clubUUID).Updates(&contModels)

	if err := tx.Error; err != nil {
		return err
	}

	return nil
}
