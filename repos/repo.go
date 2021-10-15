package repos

import (
	"github.com/lc-tut/club-portal/repos/admins"
	"github.com/lc-tut/club-portal/repos/clubs"
	"github.com/lc-tut/club-portal/repos/users"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type IRepository interface {
	clubs.IClubRepository
	users.IUserRepository
	admins.IAdminRepository
}

type Repository struct {
	*clubs.ClubRepository
	*users.UserRepository
	*admins.AdminRepository
	logger *zap.Logger
	db     *gorm.DB
}

func NewRepository(logger *zap.Logger, db *gorm.DB) *Repository {
	clubRepository := clubs.NewClubRepository(logger, db)
	userRepository := users.NewUserRepository(logger, db)
	adminRepository := admins.NewAdminRepository(logger, db)
	r := &Repository{
		ClubRepository:  clubRepository,
		UserRepository:  userRepository,
		AdminRepository: adminRepository,
		logger:          logger,
		db:              db,
	}
	return r
}
