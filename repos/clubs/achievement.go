package clubs

import (
	"github.com/lc-tut/club-portal/models/clubs"
	"gorm.io/gorm"
)

type ClubAchievementRepo interface {
	GetAchievementByID(achievementID uint32) (*clubs.ClubAchievement, error)

	GetAchievementsByClubUUID(uuid string) ([]clubs.ClubAchievement, error)

	CreateAchievement(clubUUID string, achievements []string) error
	CreateAchievementWithTx(tx *gorm.DB, clubUUID string, achievements []string) error

	UpdateAchievement(clubUUID string, achievements []string) error
	UpdateAchievementWithTx(tx *gorm.DB, clubUUID string, achievements []string) error
}

func (r *ClubRepository) GetAchievementByID(achievementID uint32) (*clubs.ClubAchievement, error) {
	achievement := &clubs.ClubAchievement{}
	tx := r.db.Where("achievement_id = ?", achievementID).Find(achievement)

	if err := tx.Error; err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}

	return achievement, nil
}

func (r *ClubRepository) GetAchievementsByClubUUID(uuid string) ([]clubs.ClubAchievement, error) {
	achievement := make([]clubs.ClubAchievement, 0)
	tx := r.db.Where("club_uuid = ?", uuid).Find(achievement)

	if err := tx.Error; err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}

	return achievement, nil
}

func (r *ClubRepository) CreateAchievement(clubUUID string, achievements []string) error {
	length := len(achievements)

	if length == 0 {
		return nil
	}

	achieveModels := make([]clubs.ClubAchievement, length)

	for i, achieve := range achievements {
		ach := clubs.ClubAchievement{
			ClubUUID:    clubUUID,
			Achievement: achieve,
		}
		achieveModels[i] = ach
	}

	tx := r.db.Create(&achieveModels)

	if err := tx.Error; err != nil {
		r.logger.Error(err.Error())
		return err
	}

	return nil
}

func (r *ClubRepository) CreateAchievementWithTx(tx *gorm.DB, clubUUID string, achievements []string) error {
	length := len(achievements)

	if length == 0 {
		return nil
	}

	achieveModels := make([]clubs.ClubAchievement, length)

	for i, achieve := range achievements {
		ach := clubs.ClubAchievement{
			ClubUUID:    clubUUID,
			Achievement: achieve,
		}
		achieveModels[i] = ach
	}

	if err := tx.Create(&achieveModels).Error; err != nil {
		r.logger.Error(err.Error())
		return err
	}

	return nil
}

func (r *ClubRepository) UpdateAchievement(clubUUID string, achievements []string) error {
	length := len(achievements)

	if length == 0 {
		return nil
	}

	tx := r.db.Where("club_uuid = ?", clubUUID).Delete(&clubs.Achievements{})

	if err := tx.Error; err != nil {
		r.logger.Error(err.Error())
		return err
	}

	if err := r.CreateAchievement(clubUUID, achievements); err != nil {
		return err
	}

	return nil
}

func (r *ClubRepository) UpdateAchievementWithTx(tx *gorm.DB, clubUUID string, achievements []string) error {
	length := len(achievements)

	if length == 0 {
		return nil
	}

	if err := tx.Where("club_uuid = ?", clubUUID).Delete(&clubs.Achievements{}).Error; err != nil {
		r.logger.Error(err.Error())
		return err
	}

	if err := r.CreateAchievementWithTx(tx, clubUUID, achievements); err != nil {
		return err
	}

	return nil
}
