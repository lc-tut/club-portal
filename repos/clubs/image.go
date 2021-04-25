package clubs

import (
	"github.com/lc-tut/club-portal/models/clubs"
	"gorm.io/gorm"
)

type ClubImageRepo interface {
	GetImageByID(imageID uint32) (*clubs.ClubImage, error)

	GetImagesByClubUUID(uuid string) ([]clubs.ClubImage, error)

	CreateImage(clubUUID string, path []string) error
	CreateImageWithTx(tx *gorm.DB, clubUUID string, path []string) error

	UpdateImage(clubUUID string, path []string) error
	UpdateImageWithTx(tx *gorm.DB, clubUUID string, path []string) error
}

func (r *ClubRepository) GetImageByID(imageID uint32) (*clubs.ClubImage, error) {
	image := &clubs.ClubImage{}
	tx := r.db.Where("image_id = ?", imageID).Take(image)

	if err := tx.Error; err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}

	return image, nil
}

func (r *ClubRepository) GetImagesByClubUUID(uuid string) ([]clubs.ClubImage, error) {
	image := make([]clubs.ClubImage, 0)
	tx := r.db.Where("club_uuid = ?", uuid).Find(&image)

	if err := tx.Error; err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}

	return image, nil
}

func (r *ClubRepository) CreateImage(clubUUID string, path []string) error {
	length := len(path)

	if length == 0 {
		return nil
	}

	imageModels := make([]clubs.ClubImage, length)

	for i, p := range path {
		image := clubs.ClubImage{
			ClubUUID: clubUUID,
			Path:     p,
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

func (r *ClubRepository) CreateImageWithTx(tx *gorm.DB, clubUUID string, path []string) error {
	length := len(path)

	if length == 0 {
		return nil
	}

	imageModels := make([]clubs.ClubImage, length)

	for i, p := range path {
		image := clubs.ClubImage{
			ClubUUID: clubUUID,
			Path:     p,
		}
		imageModels[i] = image
	}

	if err := tx.Create(&imageModels).Error; err != nil {
		r.logger.Error(err.Error())
		return err
	}

	return nil
}

// FIXME: use delete -> create to update images.

func (r *ClubRepository) UpdateImage(clubUUID string, path []string) error {
	length := len(path)

	if length == 0 {
		return nil
	}

	imageModels := make([]clubs.ClubImage, length)

	for i, p := range path {
		image := clubs.ClubImage{
			ClubUUID: clubUUID,
			Path:     p,
		}
		imageModels[i] = image
	}

	tx := r.db.Model(&clubs.ClubImage{}).Where("club_uuid = ?", clubUUID).Updates(imageModels)

	if err := tx.Error; err != nil {
		r.logger.Error(err.Error())
		return err
	}

	return nil
}

func (r *ClubRepository) UpdateImageWithTx(tx *gorm.DB, clubUUID string, path []string) error {
	length := len(path)

	if length == 0 {
		return nil
	}

	if err := tx.Where("club_uuid", clubUUID).Delete(&clubs.ClubImage{}).Error; err != nil {
		return err
	}

	if err := r.CreateImageWithTx(tx, clubUUID, path); err != nil {
		return err
	}

	return nil
}
