package repos

import (
	"github.com/lc-tut/club-portal/models"
	"gorm.io/gorm"
)

type ClubVideoRepo interface {
	GetVideoByID(videoID uint32) (*models.ClubVideo, error)

	GetVideosByClubUUID(uuid string) ([]models.ClubVideo, error)

	CreateVideo(clubUUID string, path []string) error
	CreateVideoWithTx(tx *gorm.DB, clubUUID string, path []string) error

	UpdateVideo(clubUUID string, path []string) error
	UpdateVideoWithTx(tx *gorm.DB, clubUUID string, path []string) error
}

func (r *Repository) GetVideoByID(videoID uint32) (*models.ClubVideo, error) {
	video := &models.ClubVideo{}
	tx := r.db.Where("video_id = ?", videoID).Take(video)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return video, nil
}

func (r *Repository) GetVideosByClubUUID(uuid string) ([]models.ClubVideo, error) {
	video := make([]models.ClubVideo, 0)
	tx := r.db.Where("club_uuid = ?", uuid).Find(video)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return video, nil
}

func (r *Repository) CreateVideo(clubUUID string, path []string) error {
	length := len(path)

	if length == 0 {
		return nil
	}

	videos := make([]models.ClubVideo, length)

	for i, p := range path {
		video := models.ClubVideo{
			ClubUUID: clubUUID,
			Path:     p,
		}
		videos[i] = video
	}

	tx := r.db.Create(&videos)

	if err := tx.Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) CreateVideoWithTx(tx *gorm.DB, clubUUID string, path []string) error {
	length := len(path)

	if length == 0 {
		return nil
	}

	videos := make([]models.ClubVideo, length)

	for i, p := range path {
		video := models.ClubVideo{
			ClubUUID: clubUUID,
			Path:     p,
		}
		videos[i] = video
	}

	if err := tx.Create(&videos).Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateVideo(clubUUID string, path []string) error {
	length := len(path)

	if length == 0 {
		return nil
	}

	videos := make([]models.ClubVideo, length)

	for i, p := range path {
		video := models.ClubVideo{
			ClubUUID: clubUUID,
			Path:     p,
		}
		videos[i] = video
	}
	tx := r.db.Model(&models.ClubVideo{}).Where("club_uuid = ?", clubUUID).Updates(videos)

	if err := tx.Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateVideoWithTx(tx *gorm.DB, clubUUID string, path []string) error {
	length := len(path)

	if length == 0 {
		return nil
	}

	if err := tx.Where("club_uuid", clubUUID).Delete(&models.ClubVideo{}).Error; err != nil {
		return err
	}

	if err := r.CreateVideoWithTx(tx, clubUUID, path); err != nil {
		return err
	}

	return nil
}
