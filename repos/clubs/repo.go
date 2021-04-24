package clubs

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type IClubRepository interface {
	ClubAchievementRepo
	ClubContentRepo
	ClubImageRepo
	ClubLinkRepo
	ClubPageRepo
	ClubScheduleRepo
	ClubVideoRepo
	ClubTimeRepo
	ClubPlaceRepo
	ClubActivityDetailRepo
	ClubRemarkRepo
}

type ClubRepository struct {
	logger *zap.Logger
	db     *gorm.DB
}

func NewClubRepository(logger *zap.Logger, db *gorm.DB) *ClubRepository {
	r := &ClubRepository{
		logger: logger,
		db:     db,
	}
	return r
}