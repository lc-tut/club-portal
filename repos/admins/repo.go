package admins

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type IAdminRepository interface {
	AdminUserRepo
}

type AdminRepository struct {
	logger *zap.Logger
	db     *gorm.DB
}

func NewAdminRepository(logger *zap.Logger, db *gorm.DB) *AdminRepository {
	r := &AdminRepository{
		logger: logger,
		db:     db,
	}
	return r
}
