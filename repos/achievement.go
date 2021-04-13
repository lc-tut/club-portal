package repos

import (
	"github.com/lc-tut/club-portal/models"
	"gorm.io/gorm"
)

type ClubAchievementRepo interface {
	GetAchievementByID(achievementID uint32) (*models.ClubAchievement, error)

	GetAchievementsByClubUUID(uuid string) ([]models.ClubAchievement, error)

	CreateAchievement(clubUUID string, achievements []string) error
	CreateAchievementWithTx(tx *gorm.DB, clubUUID string, achievements []string) error

	UpdateAchievement(clubUUID string, achievements []string) error
	UpdateAchievementWithTx(tx *gorm.DB, clubUUID string, achievements []string) error
}

func (r *Repository) GetAchievementByID(achievementID uint32) (*models.ClubAchievement, error) {
	achievement := &models.ClubAchievement{}
	tx := r.db.Where("achievement_id = ?", achievementID).Take(achievement)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return achievement, nil
}

func (r *Repository) GetAchievementsByClubUUID(uuid string) ([]models.ClubAchievement, error) {
	achievement := make([]models.ClubAchievement, 0)
	tx := r.db.Where("club_uuid = ?", uuid).Find(achievement)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return achievement, nil
}

func (r *Repository) CreateAchievement(clubUUID string, achievements []string) error {
	length := len(achievements)

	if length == 0 {
		return nil
	}

	achieveModels := make([]models.ClubAchievement, length)

	for i, achieve := range achievements {
		ach := models.ClubAchievement{
			ClubUUID:    clubUUID,
			Achievement: achieve,
		}
		achieveModels[i] = ach
	}

	tx := r.db.Create(&achieveModels)

	if err := tx.Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) CreateAchievementWithTx(tx *gorm.DB, clubUUID string, achievements []string) error {
	length := len(achievements)

	if length == 0 {
		return nil
	}

	achieveModels := make([]models.ClubAchievement, length)

	for i, achieve := range achievements {
		ach := models.ClubAchievement{
			ClubUUID:    clubUUID,
			Achievement: achieve,
		}
		achieveModels[i] = ach
	}

	if err := tx.Create(&achieveModels).Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateAchievement(clubUUID string, achievements []string) error {
	length := len(achievements)

	if length == 0 {
		return nil
	}

	achieveModels := make([]models.ClubAchievement, length)

	for i, achieve := range achievements {
		ach := models.ClubAchievement{
			ClubUUID:    clubUUID,
			Achievement: achieve,
		}
		achieveModels[i] = ach
	}

	tx := r.db.Model(&models.ClubAchievement{}).Where("club_uuid = ?", clubUUID).Updates(achieveModels)

	if err := tx.Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateAchievementWithTx(tx *gorm.DB, clubUUID string, achievements []string) error {
	length := len(achievements)

	if length == 0 {
		return nil
	}

	if err := tx.Where("club_uuid = ?", clubUUID).Delete(&models.Achievements{}).Error; err != nil {
		return err
	}

	if err := r.CreateAchievementWithTx(tx, clubUUID, achievements); err != nil {
		return err
	}

	return nil
}
