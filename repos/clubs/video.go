package clubs

import (
	"errors"
	"github.com/lc-tut/club-portal/models/clubs"
	"gorm.io/gorm"
)

type ClubVideoRepo interface {
	GetVideoByID(videoID uint32) (*clubs.ClubVideo, error)

	GetVideosByClubUUID(uuid string) ([]clubs.ClubVideo, error)

	CreateVideo(clubUUID string, path []string) ([]clubs.ClubVideo, error)
	CreateVideoWithTx(tx *gorm.DB, clubUUID string, path []string) ([]clubs.ClubVideo, error)

	UpdateVideo(clubUUID string, path []string) ([]clubs.ClubVideo, error)
	UpdateVideoWithTx(tx *gorm.DB, clubUUID string, path []string) ([]clubs.ClubVideo, error)
}

func (r *ClubRepository) GetVideoByID(videoID uint32) (*clubs.ClubVideo, error) {
	video := &clubs.ClubVideo{}
	tx := r.db.Where("video_id = ?", videoID).Take(video)

	if err := tx.Error; errors.Is(err, gorm.ErrRecordNotFound) {
		r.logger.Info(err.Error())
		return nil, err
	} else if err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}

	return video, nil
}

func (r *ClubRepository) GetVideosByClubUUID(uuid string) ([]clubs.ClubVideo, error) {
	video := make([]clubs.ClubVideo, 0)
	tx := r.db.Where("club_uuid = ?", uuid).Find(&video)

	if err := tx.Error; errors.Is(err, gorm.ErrRecordNotFound) {
		r.logger.Info(err.Error())
		return nil, err
	} else if err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}

	return video, nil
}

func (r *ClubRepository) CreateVideo(clubUUID string, path []string) ([]clubs.ClubVideo, error) {
	length := len(path)

	if length == 0 {
		return []clubs.ClubVideo{}, nil
	}

	videos := make([]clubs.ClubVideo, length)

	for i, p := range path {
		video := clubs.ClubVideo{
			ClubUUID: clubUUID,
			Path:     p,
		}
		videos[i] = video
	}

	tx := r.db.Create(&videos)

	if err := tx.Error; err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}

	return videos, nil
}

func (r *ClubRepository) CreateVideoWithTx(tx *gorm.DB, clubUUID string, path []string) ([]clubs.ClubVideo, error) {
	length := len(path)

	if length == 0 {
		return []clubs.ClubVideo{}, nil
	}

	videos := make([]clubs.ClubVideo, length)

	for i, p := range path {
		video := clubs.ClubVideo{
			ClubUUID: clubUUID,
			Path:     p,
		}
		videos[i] = video
	}

	if err := tx.Create(&videos).Error; err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}

	return videos, nil
}

func (r *ClubRepository) UpdateVideo(clubUUID string, path []string) ([]clubs.ClubVideo, error) {
	length := len(path)

	if length == 0 {
		return []clubs.ClubVideo{}, nil
	}

	tx := r.db.Where("club_uuid = ?", clubUUID).Delete(&clubs.ClubVideo{})

	if err := tx.Error; errors.Is(err, gorm.ErrRecordNotFound) {
		r.logger.Info(err.Error())
		return nil, err
	} else if err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}

	videos, err := r.CreateVideo(clubUUID, path)

	if err != nil {
		return nil, err
	}

	return videos, nil
}

func (r *ClubRepository) UpdateVideoWithTx(tx *gorm.DB, clubUUID string, path []string) ([]clubs.ClubVideo, error) {
	length := len(path)

	if length == 0 {
		return []clubs.ClubVideo{}, nil
	}

	tx = tx.Where("club_uuid", clubUUID).Delete(&clubs.ClubVideo{})

	if err := tx.Error; errors.Is(err, gorm.ErrRecordNotFound) {
		r.logger.Info(err.Error())
		return nil, err
	} else if err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}

	videos, err := r.CreateVideoWithTx(tx, clubUUID, path)

	if err != nil {
		return nil, err
	}

	return videos, nil
}
