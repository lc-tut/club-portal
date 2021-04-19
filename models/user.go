package models

type UserInfo interface {
	GetUserID() string
	GetEmail() string
	GetName() string
	GetRole() string
}

type DomainUser struct {
	UserUUID string `gorm:"type:char(36);not null;primaryKey"`
	Email    string `gorm:"type:varchar(255);not null;unique"`
	Name     string `gorm:"type:varchar(32);not null;unique"`
}

func (u *DomainUser) GetUserID() string {
	id := u.UserUUID
	return id
}

func (u *DomainUser) GetEmail() string {
	email := u.Email
	return email
}

func (u *DomainUser) GetName() string {
	name := u.Name
	return name
}

func (u *DomainUser) GetRole() string {
	role := "domain"
	return role
}

type GeneralUser struct {
	UserUUID string `gorm:"type:char(36);not null;primaryKey"`
	Email    string `gorm:"type:varchar(255);not null;unique"`
	Name     string `gorm:"type:varchar(32);not null;unique"`
	ClubUUID string `gorm:"foreignKey:ClubUUID;references:ClubUUID"`
}

func (u *GeneralUser) GetUserID() string {
	id := u.UserUUID
	return id
}

func (u *GeneralUser) GetEmail() string {
	email := u.Email
	return email
}

func (u *GeneralUser) GetName() string {
	name := u.Name
	return name
}

func (u *GeneralUser) GetRole() string {
	role := "general"
	return role
}

type AdminUser struct {
	UserUUID string `gorm:"type:char(36);not null;primaryKey"`
	Email    string `gorm:"type:varchar(255);not null;unique"`
	Name     string `gorm:"type:varchar(32);not null;unique"`
}

func (u *AdminUser) GetUserID() string {
	id := u.UserUUID
	return id
}

func (u *AdminUser) GetEmail() string {
	email := u.Email
	return email
}

func (u *AdminUser) GetName() string {
	name := u.Name
	return name
}

func (u *AdminUser) GetRole() string {
	role := "admin"
	return role
}
