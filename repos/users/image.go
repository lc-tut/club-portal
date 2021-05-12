package users

import (
	"github.com/lc-tut/club-portal/models/users"
)

type UploadedImageRepo interface {
	GetUploadedImageByID(imageID uint32) (*users.UploadedImage, error)

	GetImagesByUserUUID(userUUID string) ([]users.UploadedImage, error)

	CreateUploadedImage(userUUID string, path string) error

	DeleteImageByID(imageID uint32) error
}

func (r *UserRepository) GetUploadedImageByID(imageID uint32) (*users.UploadedImage, error) {
	image := &users.UploadedImage{}
	tx := r.db.Where("image_id = ?", imageID).Take(image)

	if err := tx.Error; err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}

	return image, nil
}

func (r *UserRepository) GetImagesByUserUUID(userUUID string) ([]users.UploadedImage, error) {
	images := make([]users.UploadedImage, 0)
	tx := r.db.Where("owner = ?", userUUID).Find(&images)

	if err := tx.Error; err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}

	return images, nil
}

func (r *UserRepository) CreateUploadedImage(userUUID string, path string) error {
	image := &users.UploadedImage{
		Owner: userUUID,
		Path:  path,
	}
	tx := r.db.Create(image)

	if err := tx.Error; err != nil {
		r.logger.Error(err.Error())
		return err
	}

	return nil
}

func (r *UserRepository) DeleteImageByID(imageID uint32) error {
	image := &users.UploadedImage{}
	tx := r.db.Where("image_id = ?", imageID).Delete(image)

	if err := tx.Error; err != nil {
		r.logger.Error(err.Error())
		return err
	}

	return nil
}
