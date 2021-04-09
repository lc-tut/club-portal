package repos

import "github.com/lc-tut/club-portal/models"

type ClubContentRepo interface {
	GetContentByID(contentID uint32) (*models.ClubContent, error)
	GetContentByClubUUID(uuid string) (*models.ClubContent, error)

	CreateContent(clubUUID string, content string) error

	UpdateContent(clubUUID string, content string) error
}

func (r *Repository) GetContentByID(contentID uint32) (*models.ClubContent, error) {
	content := &models.ClubContent{}
	tx := r.db.Where("content_id = ?", contentID).Take(content)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return content, nil
}

func (r *Repository) GetContentByClubUUID(uuid string) (*models.ClubContent, error) {
	content := &models.ClubContent{}
	tx := r.db.Where("club_uuid = ?", uuid).Take(content)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return content, nil
}

func (r *Repository) CreateContent(clubUUID string, content string) error {
	cont := &models.ClubContent{
		ClubUUID: clubUUID,
		Content:  content,
	}

	tx := r.db.Create(cont)

	if err := tx.Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateContent(clubUUID string, content string) error {
	cont := &models.ClubContent{
		Content: content,
	}

	tx := r.db.Model(cont).Where("club_uuid = ?", clubUUID).Updates(cont)

	if err := tx.Error; err != nil {
		return err
	}

	return nil
}
