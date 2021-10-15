package admins

import (
	"errors"
	"github.com/lc-tut/club-portal/consts"
	models "github.com/lc-tut/club-portal/models/users"
	"github.com/lc-tut/club-portal/utils"
	"gorm.io/gorm"
)

type AdminUserRepo interface {
	GetSpecifiedUser(userUUID string) (models.UserInfo, error)

	UpdateSpecifiedDomainUser(userUUID string, name string) error
}

func (r *AdminRepository) GetSpecifiedUser(userUUID string) (models.UserInfo, error) {
	user := &models.User{}
	tx := r.db.Where("user_uuid = ?", userUUID).Take(user)

	if err := tx.Error; errors.Is(err, gorm.ErrRecordNotFound) {
		r.logger.Info(err.Error())
		return nil, err
	} else if err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}

	role, _ := utils.ToUserType(user.Role)

	switch role {
	case consts.AdminUser:
		return r.getAdminUser(userUUID)
	case consts.GeneralUser:
		return r.getGeneralUser(userUUID)
	case consts.DomainUser:
		return r.getDomainUser(userUUID)
	default:
		return nil, consts.UnreachableError
	}
}

func (r *AdminRepository) getAdminUser(userUUID string) (*models.AdminUser, error) {
	user := &models.AdminUser{}
	tx := r.db.Where("user_uuid = ?", userUUID).Take(user)

	if err := tx.Error; errors.Is(err, gorm.ErrRecordNotFound) {
		r.logger.Info(err.Error())
		return nil, err
	} else if err != nil {
		r.logger.Error(err.Error())
		return nil, err
	} else {
		return user, nil
	}
}

func (r *AdminRepository) getGeneralUser(userUUID string) (*models.GeneralUser, error) {
	user := &models.GeneralUser{}
	tx := r.db.Where("user_uuid = ?", userUUID).Take(user)

	if err := tx.Error; errors.Is(err, gorm.ErrRecordNotFound) {
		r.logger.Info(err.Error())
		return nil, err
	} else if err != nil {
		r.logger.Error(err.Error())
		return nil, err
	} else {
		return user, nil
	}
}

func (r *AdminRepository) getDomainUser(userUUID string) (*models.DomainUser, error) {
	user := &models.DomainUser{}
	tx := r.db.Where("user_uuid = ?", userUUID).Take(user)

	if err := tx.Error; errors.Is(err, gorm.ErrRecordNotFound) {
		r.logger.Info(err.Error())
		return nil, err
	} else if err != nil {
		r.logger.Error(err.Error())
		return nil, err
	} else {
		return user, nil
	}
}

func (r *AdminRepository) UpdateSpecifiedDomainUser(userUUID string, name string) error {
	user := models.DomainUser{
		UserUUID: userUUID,
		Name:     name,
	}
	tx := r.db.Model(&user).Where("user_uuid = ?", userUUID).Updates(user)

	if err := tx.Error; err != nil {
		r.logger.Error(err.Error())
		return err
	} else if tx.RowsAffected == 0 {
		err := gorm.ErrRecordNotFound
		r.logger.Info(err.Error())
		return err
	} else {
		return nil
	}
}
