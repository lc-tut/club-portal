package repos

import "github.com/lc-tut/club-portal/models"

type UserRepository interface {
	GetAllGeneralUser() (*models.GeneralUser, error)

	GetDomainUserByUUID(uuid string) (*models.DomainUser, error)
	GetDomainUserByEmail(email string) (*models.DomainUser, error)
	GetGeneralUserByUUID(uuid string) (*models.GeneralUser, error)
	GetGeneralUserByEmail(email string) (*models.GeneralUser, error)
	GetAdminUserByUUID(uuid string) (*models.AdminUser, error)
	GetAdminUserByEmail(email string) (*models.AdminUser, error)
	GetUserByUUIDFromRole(uuid string, role string) (models.UserInfo, error)
	GetUserByEmailFromRole(email string, role string) (models.UserInfo, error)

	CreateDomainUser(uuid string, email string, name string) error
	CreateGeneralUser(uuid string, email string, name string, clubUUID string) error

	UpdateDomainUser(uuid string, name string) error
}
