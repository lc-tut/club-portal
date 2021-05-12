package users

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type IUserRepository interface {
	UserRepo
	UserFavoriteRepo
	UploadedImageRepo
	UploadedThumbnailRepo
}

type UserRepository struct {
	logger *zap.Logger
	db     *gorm.DB
}

func NewUserRepository(logger *zap.Logger, db *gorm.DB) *UserRepository {
	r := &UserRepository{
		logger: logger,
		db:     db,
	}
	return r
}
