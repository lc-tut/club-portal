package admins

import (
	"github.com/lc-tut/club-portal/repos/clubs"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type IAdminRepository interface {
	AdminUserRepo
	AdminClubRepo
}

type AdminRepository struct {
	logger *zap.Logger
	db     *gorm.DB
	clubs.IClubRepository
}

func NewAdminRepository(logger *zap.Logger, db *gorm.DB, clubRepo clubs.IClubRepository) *AdminRepository {
	r := &AdminRepository{
		logger:          logger,
		db:              db,
		IClubRepository: clubRepo,
	}
	return r
}
