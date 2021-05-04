package clubs

import (
	"github.com/lc-tut/club-portal/models/clubs"
	"gorm.io/gorm"
)

type ClubImageArgs struct {
	ImageID uint32
	Path    string
}

type ClubImageRepo interface {
	GetImageByID(imageID uint32) (*clubs.ClubImage, error)

	GetImagesByClubUUID(uuid string) ([]clubs.ClubImage, error)

	CreateImage(clubUUID string, args []ClubImageArgs) error
	CreateImageWithTx(tx *gorm.DB, clubUUID string, args []ClubImageArgs) error

	UpdateImage(clubUUID string, args []ClubImageArgs) error
	UpdateImageWithTx(tx *gorm.DB, clubUUID string, args []ClubImageArgs) error
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

func (r *ClubRepository) CreateImage(clubUUID string, args []ClubImageArgs) error {
	length := len(args)

	if length == 0 {
		return nil
	}

	imageModels := make([]clubs.ClubImage, length)

	for i, arg := range args {
		image := clubs.ClubImage{
			ImageID:  arg.ImageID,
			ClubUUID: clubUUID,
			Path:     arg.Path,
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

func (r *ClubRepository) CreateImageWithTx(tx *gorm.DB, clubUUID string, args []ClubImageArgs) error {
	length := len(args)

	if length == 0 {
		return nil
	}

	imageModels := make([]clubs.ClubImage, length)

	for i, arg := range args {
		image := clubs.ClubImage{
			ImageID:  arg.ImageID,
			ClubUUID: clubUUID,
			Path:     arg.Path,
		}
		imageModels[i] = image
	}

	if err := tx.Create(&imageModels).Error; err != nil {
		r.logger.Error(err.Error())
		return err
	}

	return nil
}

func (r *ClubRepository) UpdateImage(clubUUID string, args []ClubImageArgs) error {
	length := len(args)

	if length == 0 {
		return nil
	}

	tx := r.db.Where("club_uuid = ?", clubUUID).Delete(&clubs.ClubImage{})

	if err := tx.Error; err != nil {
		r.logger.Error(err.Error())
		return err
	}

	if err := r.CreateImage(clubUUID, args); err != nil {
		return err
	}

	return nil
}

func (r *ClubRepository) UpdateImageWithTx(tx *gorm.DB, clubUUID string, args []ClubImageArgs) error {
	length := len(args)

	if length == 0 {
		return nil
	}

	if err := tx.Where("club_uuid = ?", clubUUID).Delete(&clubs.ClubImage{}).Error; err != nil {
		r.logger.Error(err.Error())
		return err
	}

	if err := r.CreateImageWithTx(tx, clubUUID, args); err != nil {
		return err
	}

	return nil
}
