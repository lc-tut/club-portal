package repos

import "github.com/lc-tut/club-portal/models"

type ClubImageRepo interface {
	GetImageByID(imageID uint32) (*models.ClubImage, error)

	GetImagesByClubUUID(uuid string) ([]models.ClubImage, error)

	CreateImage(clubUUID string, path []string) error

	UpdateImage(clubUUID string, path []string) error
}

func (r *Repository) GetImageByID(imageID uint32) (*models.ClubImage, error) {
	image := &models.ClubImage{}
	tx := r.db.Where("image_id = ?", imageID).Take(image)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return image, nil
}

func (r *Repository) GetImagesByClubUUID(uuid string) ([]models.ClubImage, error) {
	image := make([]models.ClubImage, 0)
	tx := r.db.Where("club_uuid = ?", uuid).Find(image)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return image, nil
}

func (r *Repository) CreateImage(clubUUID string, path []string) error {
	imageModels := make([]models.ClubImage, len(path))

	for _, p := range path {
		image := models.ClubImage{
			ClubUUID: clubUUID,
			Path:     p,
		}
		imageModels = append(imageModels, image)
	}

	tx := r.db.Create(&imageModels)

	if err := tx.Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateImage(clubUUID string, path []string) error {
	imageModels := make([]models.ClubImage, len(path))

	for _, p := range path {
		image := models.ClubImage{
			ClubUUID: clubUUID,
			Path:     p,
		}
		imageModels = append(imageModels, image)
	}

	tx := r.db.Model(&models.ClubImage{}).Where("club_uuid = ?", clubUUID).Updates(&imageModels)

	if err := tx.Error; err != nil {
		return err
	}

	return nil
}
