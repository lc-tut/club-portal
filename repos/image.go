package repos

import "github.com/lc-tut/club-portal/models"

type ClubImageRepo interface {
	GetImageByID(imageID uint32) (*models.ClubImage, error)
	GetImageByClubUUID(uuid string) (*models.ClubImage, error)

	CreateImage(clubUUID string, path string) error

	UpdateImage(clubUUID string, path string) error
}

func (r *Repository) GetImageByID(imageID uint32) (*models.ClubImage, error) {
	image := &models.ClubImage{}
	tx := r.db.Where("image_id = ?", imageID).Take(image)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return image, nil
}

func (r *Repository) GetImageByClubUUID(uuid string) (*models.ClubImage, error) {
	image := &models.ClubImage{}
	tx := r.db.Where("club_uuid = ?", uuid).Take(image)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return image, nil
}

func (r *Repository) CreateImage(clubUUID string, path string) error {
	image := &models.ClubImage{
		ClubUUID: clubUUID,
		Path:     path,
	}

	tx := r.db.Create(image)

	if err := tx.Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateImage(clubUUID string, path string) error {
	image := &models.ClubImage{
		Path: path,
	}

	tx := r.db.Model(image).Where("club_uuid = ?", clubUUID).Updates(image)

	if err := tx.Error; err != nil {
		return err
	}

	return nil
}
