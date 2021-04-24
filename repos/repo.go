package repos

import (
	"github.com/lc-tut/club-portal/repos/clubs"
	"github.com/lc-tut/club-portal/repos/users"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type IRepository interface {
	clubs.IClubRepository
	users.IUserRepository
}

type Repository struct {
	*clubs.ClubRepository
	*users.UserRepository
	logger *zap.Logger
	db     *gorm.DB
}

func NewRepository(logger *zap.Logger, db *gorm.DB) *Repository {
	clubRepository := clubs.NewClubRepository(logger, db)
	userRepository := users.NewUserRepository(logger, db)
	r := &Repository{
		ClubRepository: clubRepository,
		UserRepository: userRepository,
		logger:         logger,
		db:             db,
	}
	return r
}
