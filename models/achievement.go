package models

type ClubAchievement struct {
	AchievementID uint32 `gorm:"type:int unsigned not null auto_increment;primaryKey"`
	ClubUUID      string `gorm:"type:char(36);not null"`
	Achievement   string `gorm:"type:text;not null"`
}

type Achievements []ClubAchievement

func (ac Achievements) ToAchievementResponse() *[]AchievementResponse {
	res := make([]AchievementResponse, len(ac))

	for i, achieve := range ac {
		achieveRes := AchievementResponse{
			Achievement: achieve.Achievement,
		}
		res[i] = achieveRes
	}

	return &res
}

type AchievementRequest struct {
	Achievement string `json:"achievement"`
}

type AchievementResponse struct {
	Achievement string `json:"achievement"`
}
