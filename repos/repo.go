package repos

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type IRepository interface {
	ClubAchievementRepo
	ClubContentRepo
	ClubImageRepo
	ClubLinkRepo
	ClubPageRepo
	ClubScheduleRepo
	ClubVideoRepo
	ClubTimeRepo
	ClubPlaceRepo
	ClubTimeAndPlaceRepo
}

type Repository struct {
	logger *zap.Logger
	db     *gorm.DB
}

func NewRepository(logger *zap.Logger, db *gorm.DB) *Repository {
	r := &Repository{
		logger: logger,
		db:     db,
	}
	return r
}
