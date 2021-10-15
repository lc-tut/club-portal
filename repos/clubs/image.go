package clubs

import (
	"errors"
	"github.com/lc-tut/club-portal/models/clubs"
	"gorm.io/gorm"
)

type ClubImageRepo interface {
	GetImageByID(imageID uint32) (*clubs.ClubImage, error)

	GetImagesByClubUUID(uuid string) ([]clubs.ClubImage, error)

	CreateImage(clubUUID string, imageIDs []uint32) error
	CreateImageWithTx(tx *gorm.DB, clubUUID string, imageIDs []uint32) error

	UpdateImage(clubUUID string, imageIDs []uint32) error
	UpdateImageWithTx(tx *gorm.DB, clubUUID string, imageIDs []uint32) error
}

func (r *ClubRepository) GetImageByID(imageID uint32) (*clubs.ClubImage, error) {
	image := &clubs.ClubImage{}
	selectQuery := "ci.image_id, ci.club_uuid, ui.path"
	joinQuery := "inner join uploaded_images as ui using (image_id)"
	tx := r.db.Table("club_images as ci").Select(selectQuery).Joins(joinQuery).Where("image_id = ?", imageID).Find(image)

	if err := tx.Error; errors.Is(err, gorm.ErrRecordNotFound) {
		r.logger.Info(err.Error())
		return nil, err
	} else if err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}

	return image, nil
}

func (r *ClubRepository) GetImagesByClubUUID(uuid string) ([]clubs.ClubImage, error) {
	image := make([]clubs.ClubImage, 0)
	selectQuery := "ci.image_id, ci.club_uuid, ui.path"
	joinQuery := "inner join uploaded_images as ui using (image_id)"
	tx := r.db.Table("club_images as ci").Select(selectQuery).Joins(joinQuery).Where("club_uuid = ?", uuid).Find(&image)

	if err := tx.Error; errors.Is(err, gorm.ErrRecordNotFound) {
		r.logger.Info(err.Error())
		return nil, err
	} else if err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}

	return image, nil
}

func (r *ClubRepository) CreateImage(clubUUID string, imageIDs []uint32) error {
	length := len(imageIDs)

	if length == 0 {
		return nil
	}

	imageModels := make([]clubs.ClubImage, length)

	for i, imID := range imageIDs {
		image := clubs.ClubImage{
			ImageID:  imID,
			ClubUUID: clubUUID,
		}
		imageModels[i] = image
	}

	tx := r.db.Create(&imageModels)

	if err := tx.Error; err != nil {
		r.logger.Error(err.Error())
		return err
	}

	return nil
}

func (r *ClubRepository) CreateImageWithTx(tx *gorm.DB, clubUUID string, imageIDs []uint32) error {
	length := len(imageIDs)

	if length == 0 {
		return nil
	}

	imageModels := make([]clubs.ClubImage, length)

	for i, imID := range imageIDs {
		image := clubs.ClubImage{
			ImageID:  imID,
			ClubUUID: clubUUID,
		}
		imageModels[i] = image
	}

	if err := tx.Create(&imageModels).Error; err != nil {
		r.logger.Error(err.Error())
		return err
	}

	return nil
}

func (r *ClubRepository) UpdateImage(clubUUID string, imageIDs []uint32) error {
	length := len(imageIDs)

	if length == 0 {
		return nil
	}

	tx := r.db.Where("club_uuid = ?", clubUUID).Delete(&clubs.ClubImage{})

	if err := tx.Error; errors.Is(err, gorm.ErrRecordNotFound) {
		r.logger.Info(err.Error())
		return err
	} else if err != nil {
		r.logger.Error(err.Error())
		return err
	}

	if err := r.CreateImage(clubUUID, imageIDs); err != nil {
		return err
	}

	return nil
}

func (r *ClubRepository) UpdateImageWithTx(tx *gorm.DB, clubUUID string, imageIDs []uint32) error {
	length := len(imageIDs)

	if length == 0 {
		return nil
	}

	tx = tx.Where("club_uuid = ?", clubUUID).Delete(&clubs.ClubImage{})

	if err := tx.Error; errors.Is(err, gorm.ErrRecordNotFound) {
		r.logger.Info(err.Error())
		return err
	} else if err != nil {
		r.logger.Error(err.Error())
		return err
	}

	if err := r.CreateImageWithTx(tx, clubUUID, imageIDs); err != nil {
		return err
	}

	return nil
}
