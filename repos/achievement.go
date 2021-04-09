package repos

import "github.com/lc-tut/club-portal/models"

type ClubAchievementRepo interface {
	GetAchievementByID(achievementID uint32) (*models.ClubAchievement, error)
	GetAchievementByClubUUID(uuid string) (*models.ClubAchievement, error)

	CreateAchievement(clubUUID string, achievement string) error

	UpdateAchievement(clubUUID string, achievement string) error
}

func (r *Repository) GetAchievementByID(achievementID uint32) (*models.ClubAchievement, error) {
	achievement := &models.ClubAchievement{}
	tx := r.db.Where("achievement_id = ?", achievementID).Take(achievement)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return achievement, nil
}

func (r *Repository) GetAchievementByClubUUID(uuid string) (*models.ClubAchievement, error) {
	achievement := &models.ClubAchievement{}
	tx := r.db.Where("club_uuid = ?", uuid).Take(achievement)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return achievement, nil
}

func (r *Repository) CreateAchievement(clubUUID string, achievement string) error {
	ach := &models.ClubAchievement{
		ClubUUID:    clubUUID,
		Achievement: achievement,
	}

	tx := r.db.Create(ach)

	if err := tx.Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateAchievement(clubUUID string, achievement string) error {
	ach := &models.ClubAchievement{
		Achievement: achievement,
	}

	tx := r.db.Model(ach).Where("club_uuid = ?", clubUUID).Updates(ach)

	if err := tx.Error; err != nil {
		return err
	}

	return nil
}
