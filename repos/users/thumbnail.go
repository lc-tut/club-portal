package users

import "github.com/lc-tut/club-portal/models/users"

type UploadedThumbnailRepo interface {
	GetThumbnail(thumbnailID uint32) (*users.UploadedThumbnail, error)

	CreateThumbnail(path string) (*users.UploadedThumbnail, error)

	UpdateThumbnail(thumbnailID uint32, path string) error

	DeleteThumbnail(thumbnailID uint32) error
}

func (r *UserRepository) GetThumbnail(thumbnailID uint32) (*users.UploadedThumbnail, error) {
	thumbnail := &users.UploadedThumbnail{}
	tx := r.db.Where("thumbnail_id = ?", thumbnailID).Find(thumbnail)

	if err := tx.Error; err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}

	return thumbnail, nil
}

func (r *UserRepository) CreateThumbnail(path string) (*users.UploadedThumbnail, error) {
	thumbnail := &users.UploadedThumbnail{
		Path: path,
	}
	tx := r.db.Create(thumbnail)

	if err := tx.Error; err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}

	return thumbnail, nil
}

func (r *UserRepository) UpdateThumbnail(thumbnailID uint32, path string) error {
	tx := r.db.Model(&users.UploadedThumbnail{}).Where("thumbnail_id = ?", thumbnailID).Update("path", path)

	if err := tx.Error; err != nil {
		r.logger.Error(err.Error())
		return err
	}

	return nil
}

func (r *UserRepository) DeleteThumbnail(thumbnailID uint32) error {
	tx := r.db.Where("thumbnail_id = ?", thumbnailID).Delete(&users.UploadedThumbnail{})

	if err := tx.Error; err != nil {
		r.logger.Error(err.Error())
		return err
	}

	return nil
}
