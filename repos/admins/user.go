package admins

import (
	"errors"
	"github.com/lc-tut/club-portal/consts"
	models "github.com/lc-tut/club-portal/models/users"
	"github.com/lc-tut/club-portal/utils"
	"gorm.io/gorm"
)

type UserArgs struct {
	Email    string
	Name     string
	ClubUUID string
}

type AdminUserRepo interface {
	GetAllUser() ([]models.UserInfo, error)
	GetSpecifiedUser(userUUID string) (models.UserInfo, error)

	UpdateSpecifiedDomainUser(userUUID string, name string) error
	UpdateSpecifiedGeneralUser(userUUID string, args UserArgs) error
}

func (r *AdminRepository) GetAllUser() ([]models.UserInfo, error) {
	domainUsers := make([]models.DomainUser, 0)
	generalUsers := make([]models.GeneralUser, 0)
	adminUsers := make([]models.AdminUser, 0)

	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Find(&domainUsers).Error; err != nil {
			return err
		}

		if err := tx.Find(&generalUsers).Error; err != nil {
			return err
		}

		if err := tx.Find(&adminUsers).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}

	allUsers := make([]models.UserInfo, len(domainUsers)+len(generalUsers)+len(adminUsers))

	for _, user := range domainUsers {
		allUsers = append(allUsers, &user)
	}

	for _, user := range generalUsers {
		allUsers = append(allUsers, &user)
	}

	for _, user := range adminUsers {
		allUsers = append(allUsers, &user)
	}

	return allUsers, nil
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

func (r *AdminRepository) UpdateSpecifiedGeneralUser(userUUID string, args UserArgs) error {
	user := models.GeneralUser{
		UserUUID: userUUID,
		Email:    args.Email,
		Name:     args.Name,
		ClubUUID: utils.StringToNullString(args.ClubUUID),
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
