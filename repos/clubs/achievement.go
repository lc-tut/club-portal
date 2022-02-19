package clubs

import (
	"errors"
	"github.com/lc-tut/club-portal/models/clubs"
	"gorm.io/gorm"
)

type ClubAchievementRepo interface {
	GetAchievementByID(achievementID uint32) (*clubs.ClubAchievement, error)

	GetAchievementsByClubUUID(uuid string) ([]clubs.ClubAchievement, error)

	CreateAchievement(clubUUID string, achievements []string) ([]clubs.ClubAchievement, error)
	CreateAchievementWithTx(tx *gorm.DB, clubUUID string, achievements []string) ([]clubs.ClubAchievement, error)

	UpdateAchievement(clubUUID string, achievements []string) ([]clubs.ClubAchievement, error)
	UpdateAchievementWithTx(tx *gorm.DB, clubUUID string, achievements []string) ([]clubs.ClubAchievement, error)
}

func (r *ClubRepository) GetAchievementByID(achievementID uint32) (*clubs.ClubAchievement, error) {
	achievement := &clubs.ClubAchievement{}
	tx := r.db.Where("achievement_id = ?", achievementID).Take(achievement)

	if err := tx.Error; errors.Is(err, gorm.ErrRecordNotFound) {
		r.logger.Info(err.Error())
		return nil, err
	} else if err != nil {
		return nil, err
	}

	return achievement, nil
}

func (r *ClubRepository) GetAchievementsByClubUUID(uuid string) ([]clubs.ClubAchievement, error) {
	achievement := make([]clubs.ClubAchievement, 0)
	tx := r.db.Where("club_uuid = ?", uuid).Find(achievement)

	if err := tx.Error; errors.Is(err, gorm.ErrRecordNotFound) {
		r.logger.Info(err.Error())
		return nil, err
	} else if err != nil {
		return nil, err
	}

	return achievement, nil
}

func (r *ClubRepository) CreateAchievement(clubUUID string, achievements []string) ([]clubs.ClubAchievement, error) {
	length := len(achievements)

	if length == 0 {
		return []clubs.ClubAchievement{}, nil
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
		return nil, err
	}

	return achieveModels, nil
}

func (r *ClubRepository) CreateAchievementWithTx(tx *gorm.DB, clubUUID string, achievements []string) ([]clubs.ClubAchievement, error) {
	length := len(achievements)

	if length == 0 {
		return []clubs.ClubAchievement{}, nil
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
		return nil, err
	}

	return achieveModels, nil
}

func (r *ClubRepository) UpdateAchievement(clubUUID string, achievements []string) ([]clubs.ClubAchievement, error) {
	length := len(achievements)

	if length == 0 {
		return []clubs.ClubAchievement{}, nil
	}

	tx := r.db.Where("club_uuid = ?", clubUUID).Delete(&clubs.Achievements{})

	if err := tx.Error; errors.Is(err, gorm.ErrRecordNotFound) {
		r.logger.Info(err.Error())
		return nil, err
	} else if err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}

	ach, err := r.CreateAchievement(clubUUID, achievements)

	if err != nil {
		return nil, err
	}

	return ach, nil
}

func (r *ClubRepository) UpdateAchievementWithTx(tx *gorm.DB, clubUUID string, achievements []string) ([]clubs.ClubAchievement, error) {
	length := len(achievements)

	if length == 0 {
		return []clubs.ClubAchievement{}, nil
	}

	tx = tx.Where("club_uuid = ?", clubUUID).Delete(&clubs.Achievements{})

	if err := tx.Error; errors.Is(err, gorm.ErrRecordNotFound) {
		r.logger.Info(err.Error())
		return nil, err
	} else if err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}

	ach, err := r.CreateAchievementWithTx(tx, clubUUID, achievements)

	if err != nil {
		return nil, err
	}

	return ach, nil
}
