package repos

import "github.com/lc-tut/club-portal/models"

type ClubAchievementRepo interface {
	GetAchievementByID(achievementID uint32) (*models.ClubAchievement, error)
	GetAchievementByClubUUID(uuid string) (*models.ClubAchievement, error)

	CreateAchievement(clubUUID string, achievement string) error

	UpdateAchievement(clubUUID string, achievement string) error
}

func (r *Repository) GetAchievementByID(achievementID uint32) (*models.ClubAchievement, error) {
	panic("implement me")
}

func (r *Repository) GetAchievementByClubUUID(uuid string) (*models.ClubAchievement, error) {
	panic("implement me")
}

func (r *Repository) CreateAchievement(clubUUID string, achievement string) error {
	panic("implement me")
}

func (r *Repository) UpdateAchievement(clubUUID string, achievement string) error {
	panic("implement me")
}
