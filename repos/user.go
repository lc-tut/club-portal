package repos

import (
	"github.com/lc-tut/club-portal/consts"
	"github.com/lc-tut/club-portal/models"
	"github.com/lc-tut/club-portal/utils"
)

type UpdateUserArgs struct {
	Name     string
	ClubUUID *string
}

type UserRepo interface {
	GetAllGeneralUser() ([]models.GeneralUser, error)

	GetDomainUserByUUID(uuid string) (*models.DomainUser, error)
	GetDomainUserByEmail(email string) (*models.DomainUser, error)
	GetGeneralUserByUUID(uuid string) (*models.GeneralUser, error)
	GetGeneralUserByEmail(email string) (*models.GeneralUser, error)
	GetAdminUserByUUID(uuid string) (*models.AdminUser, error)
	GetAdminUserByEmail(email string) (*models.AdminUser, error)
	GetUserByUUIDFromRole(uuid string, role string) (models.UserInfo, error)
	GetUserByEmailFromRole(email string, role string) (models.UserInfo, error)

	CreateDomainUser(uuid string, email string, name string) (*models.DomainUser, error)
	CreateGeneralUser(uuid string, email string, name string) (*models.GeneralUser, error)
	CreateAdminUser(uuid string, email string, name string) (*models.AdminUser, error)

	UpdateDomainUser(uuid string, name string) error
	UpdateGeneralUser(uuid string, name string, clubUUID string) error
	UpdateAdminUser(uuid string, name string) error
	UpdateUserFromRole(uuid string, role string, args UpdateUserArgs) error
}

func (r *Repository) GetAllGeneralUser() ([]models.GeneralUser, error) {
	users := make([]models.GeneralUser, 0)
	tx := r.db.Find(&users)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (r *Repository) GetDomainUserByUUID(uuid string) (*models.DomainUser, error) {
	user := &models.DomainUser{}
	tx := r.db.Where("user_uuid = ?", uuid).Take(user)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *Repository) GetDomainUserByEmail(email string) (*models.DomainUser, error) {
	user := &models.DomainUser{}
	tx := r.db.Where("email = ?", email).Take(user)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *Repository) GetGeneralUserByUUID(uuid string) (*models.GeneralUser, error) {
	user := &models.GeneralUser{}
	tx := r.db.Where("user_uuid = ?", uuid).Take(user)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *Repository) GetGeneralUserByEmail(email string) (*models.GeneralUser, error) {
	user := &models.GeneralUser{}
	tx := r.db.Where("email = ?", email).Take(user)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *Repository) GetAdminUserByUUID(uuid string) (*models.AdminUser, error) {
	user := &models.AdminUser{}
	tx := r.db.Where("user_uuid", uuid).Take(user)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *Repository) GetAdminUserByEmail(email string) (*models.AdminUser, error) {
	user := &models.AdminUser{}
	tx := r.db.Where("email = ?", email).Take(user)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *Repository) GetUserByUUIDFromRole(uuid string, role string) (models.UserInfo, error) {
	userType, err := utils.ToUserType(role)

	if err != nil {
		return nil, err
	}

	if userType == consts.AdminUser {
		return r.GetAdminUserByUUID(uuid)
	} else if userType == consts.GeneralUser {
		return r.GetGeneralUserByUUID(uuid)
	} else {
		return r.GetAdminUserByUUID(uuid)
	}
}

func (r *Repository) GetUserByEmailFromRole(email string, role string) (models.UserInfo, error) {
	userType, err := utils.ToUserType(role)

	if err != nil {
		return nil, err
	}

	if userType == consts.AdminUser {
		return r.GetAdminUserByEmail(email)
	} else if userType == consts.GeneralUser {
		return r.GetGeneralUserByEmail(email)
	} else {
		return r.GetAdminUserByEmail(email)
	}
}

func (r *Repository) CreateDomainUser(uuid string, email string, name string) (*models.DomainUser, error) {
	user := &models.DomainUser{
		UserUUID: uuid,
		Email:    email,
		Name:     name,
	}

	tx := r.db.Create(user)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *Repository) CreateGeneralUser(uuid string, email string, name string) (*models.GeneralUser, error) {
	user := &models.GeneralUser{
		UserUUID: uuid,
		Email:    email,
		Name:     name,
		ClubUUID: utils.ToNullString(""),
	}

	tx := r.db.Create(user)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *Repository) CreateAdminUser(uuid string, email string, name string) (*models.AdminUser, error) {
	user := &models.AdminUser{
		UserUUID: uuid,
		Email:    email,
		Name:     name,
	}

	tx := r.db.Create(user)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *Repository) UpdateDomainUser(uuid string, name string) error {
	user := models.DomainUser{
		Name: name,
	}

	tx := r.db.Model(&user).Where("user_uuid = ?", uuid).Updates(user)

	if err := tx.Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateGeneralUser(uuid string, name string, clubUUID string) error {
	user := models.GeneralUser{
		Name:     name,
		ClubUUID: utils.ToNullString(clubUUID),
	}

	tx := r.db.Model(&user).Where("user_uuid = ?", uuid).Updates(user)

	if err := tx.Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateAdminUser(uuid string, name string) error {
	user := models.AdminUser{
		Name: name,
	}

	tx := r.db.Model(&user).Where("user_uuid = ?", uuid).Updates(user)

	if err := tx.Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateUserFromRole(uuid string, role string, args UpdateUserArgs) error {
	userType, err := utils.ToUserType(role)

	if err != nil {
		return err
	}

	if userType == consts.AdminUser {
		err = r.UpdateAdminUser(uuid, args.Name)
	} else if userType == consts.GeneralUser {
		err = r.UpdateGeneralUser(uuid, args.Name, utils.NilToEmptyString(args.ClubUUID))
	} else {
		err = r.UpdateDomainUser(uuid, args.Name)
	}

	if err != nil {
		return err
	}

	return nil
}
