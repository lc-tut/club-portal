package users

import (
	"errors"
	"github.com/lc-tut/club-portal/consts"
	"github.com/lc-tut/club-portal/models/users"
	"github.com/lc-tut/club-portal/utils"
	"gorm.io/gorm"
)

type UpdateUserArgs struct {
	Name     string
	ClubUUID *string
}

type UserRepo interface {
	GetAllGeneralUser() ([]users.GeneralUser, error)

	GetDomainUserByUUID(uuid string) (*users.DomainUser, error)
	GetDomainUserByEmail(email string) (*users.DomainUser, error)
	GetGeneralUserByUUID(uuid string) (*users.GeneralUser, error)
	GetGeneralUserByEmail(email string) (*users.GeneralUser, error)
	GetAdminUserByUUID(uuid string) (*users.AdminUser, error)
	GetAdminUserByEmail(email string) (*users.AdminUser, error)
	GetUserByUUIDFromRole(uuid string, role string) (users.UserInfo, error)
	GetUserByEmailFromRole(email string, role string) (users.UserInfo, error)

	CreateDomainUser(uuid string, email string, name string) (*users.DomainUser, error)
	CreateGeneralUser(uuid string, email string, name string) (*users.GeneralUser, error)
	CreateAdminUser(uuid string, email string, name string) (*users.AdminUser, error)

	UpdateDomainUser(uuid string, name string) error
	UpdateGeneralUser(uuid string, name string, clubUUID string) error
	UpdateAdminUser(uuid string, name string) error
	UpdateUserFromRole(uuid string, role string, args UpdateUserArgs) error
}

func (r *UserRepository) GetAllGeneralUser() ([]users.GeneralUser, error) {
	generalUsers := make([]users.GeneralUser, 0)
	tx := r.db.Find(&generalUsers)

	if err := tx.Error; errors.Is(err, gorm.ErrRecordNotFound) {
		r.logger.Info(err.Error())
		return nil, err
	} else if err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}

	return generalUsers, nil
}

func (r *UserRepository) GetDomainUserByUUID(uuid string) (*users.DomainUser, error) {
	user := &users.DomainUser{}
	tx := r.db.Where("user_uuid = ?", uuid).Take(user)

	if err := tx.Error; errors.Is(err, gorm.ErrRecordNotFound) {
		r.logger.Info(err.Error())
		return nil, err
	} else if err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) GetDomainUserByEmail(email string) (*users.DomainUser, error) {
	user := &users.DomainUser{}
	tx := r.db.Where("email = ?", email).Take(user)

	if err := tx.Error; errors.Is(err, gorm.ErrRecordNotFound) {
		r.logger.Info(err.Error())
		return nil, err
	} else if err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) GetGeneralUserByUUID(uuid string) (*users.GeneralUser, error) {
	user := &users.GeneralUser{}
	tx := r.db.Where("user_uuid = ?", uuid).Take(user)

	if err := tx.Error; errors.Is(err, gorm.ErrRecordNotFound) {
		r.logger.Info(err.Error())
		return nil, err
	} else if err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) GetGeneralUserByEmail(email string) (*users.GeneralUser, error) {
	user := &users.GeneralUser{}
	tx := r.db.Where("email = ?", email).Take(user)

	if err := tx.Error; errors.Is(err, gorm.ErrRecordNotFound) {
		r.logger.Info(err.Error())
		return nil, err
	} else if err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) GetAdminUserByUUID(uuid string) (*users.AdminUser, error) {
	user := &users.AdminUser{}
	tx := r.db.Where("user_uuid = ?", uuid).Take(user)

	if err := tx.Error; errors.Is(err, gorm.ErrRecordNotFound) {
		r.logger.Info(err.Error())
		return nil, err
	} else if err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) GetAdminUserByEmail(email string) (*users.AdminUser, error) {
	user := &users.AdminUser{}
	tx := r.db.Where("email = ?", email).Take(user)

	if err := tx.Error; errors.Is(err, gorm.ErrRecordNotFound) {
		r.logger.Info(err.Error())
		return nil, err
	} else if err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) GetUserByUUIDFromRole(uuid string, role string) (users.UserInfo, error) {
	userType, err := utils.ToUserType(role)

	if err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}

	switch userType {
	case consts.AdminUser:
		return r.GetAdminUserByUUID(uuid)
	case consts.GeneralUser:
		return r.GetGeneralUserByUUID(uuid)
	case consts.DomainUser:
		return r.GetDomainUserByUUID(uuid)
	default:
		return nil, consts.UnreachableError
	}
}

func (r *UserRepository) GetUserByEmailFromRole(email string, role string) (users.UserInfo, error) {
	userType, err := utils.ToUserType(role)

	if err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}

	switch userType {
	case consts.AdminUser:
		return r.GetAdminUserByEmail(email)
	case consts.GeneralUser:
		return r.GetGeneralUserByEmail(email)
	case consts.DomainUser:
		return r.GetDomainUserByEmail(email)
	default:
		return nil, consts.UnreachableError
	}
}

func (r *UserRepository) createUser(tx *gorm.DB, uuid string, role string) error {
	user := &users.User{
		UserUUID: uuid,
		Role:     role,
	}

	if err := tx.Create(user).Error; err != nil {
		r.logger.Error(err.Error())
		return err
	}

	return nil
}

func (r *UserRepository) CreateDomainUser(uuid string, email string, name string) (*users.DomainUser, error) {
	user := &users.DomainUser{
		UserUUID: uuid,
		Email:    email,
		Name:     name,
	}

	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := r.createUser(tx, uuid, "domain"); err != nil {
			return err
		}

		if err := tx.Create(user).Error; err != nil {
			r.logger.Error(err.Error())
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) CreateGeneralUser(uuid string, email string, name string) (*users.GeneralUser, error) {
	user := &users.GeneralUser{
		UserUUID: uuid,
		Email:    email,
		Name:     name,
		ClubUUID: utils.StringToNullString(""),
	}

	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := r.createUser(tx, uuid, "general"); err != nil {
			return err
		}

		if err := tx.Create(user).Error; err != nil {
			r.logger.Error(err.Error())
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) CreateAdminUser(uuid string, email string, name string) (*users.AdminUser, error) {
	user := &users.AdminUser{
		UserUUID: uuid,
		Email:    email,
		Name:     name,
	}

	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := r.createUser(tx, uuid, "admin"); err != nil {
			return err
		}

		if err := tx.Create(user).Error; err != nil {
			r.logger.Error(err.Error())
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) UpdateDomainUser(uuid string, name string) error {
	user := users.DomainUser{
		Name: name,
	}

	tx := r.db.Model(&user).Where("user_uuid = ?", uuid).Updates(user)

	if rows := tx.RowsAffected; rows == 0 {
		err := gorm.ErrRecordNotFound
		r.logger.Info(err.Error())
		return err
	} else if err := tx.Error; err != nil {
		r.logger.Error(err.Error())
		return err
	}

	return nil
}

func (r *UserRepository) UpdateGeneralUser(uuid string, name string, clubUUID string) error {
	user := users.GeneralUser{
		Name:     name,
		ClubUUID: utils.StringToNullString(clubUUID),
	}

	tx := r.db.Model(&user).Where("user_uuid = ?", uuid).Updates(user)

	if rows := tx.RowsAffected; rows == 0 {
		err := gorm.ErrRecordNotFound
		r.logger.Info(err.Error())
		return err
	} else if err := tx.Error; err != nil {
		r.logger.Error(err.Error())
		return err
	}

	return nil
}

func (r *UserRepository) UpdateAdminUser(uuid string, name string) error {
	user := users.AdminUser{
		Name: name,
	}

	tx := r.db.Model(&user).Where("user_uuid = ?", uuid).Updates(user)

	if rows := tx.RowsAffected; rows == 0 {
		err := gorm.ErrRecordNotFound
		r.logger.Info(err.Error())
		return err
	} else if err := tx.Error; err != nil {
		r.logger.Error(err.Error())
		return err
	}

	return nil
}

func (r *UserRepository) UpdateUserFromRole(uuid string, role string, args UpdateUserArgs) error {
	userType, err := utils.ToUserType(role)

	if err != nil {
		r.logger.Error(err.Error())
		return err
	}

	switch userType {
	case consts.AdminUser:
		err = r.UpdateAdminUser(uuid, args.Name)
	case consts.GeneralUser:
		err = r.UpdateGeneralUser(uuid, args.Name, utils.StringPToString(args.ClubUUID))
	case consts.DomainUser:
		err = r.UpdateDomainUser(uuid, args.Name)
	default:
		err = consts.UnreachableError
	}

	if err != nil {
		return err
	}

	return nil
}
