package users

import (
	"database/sql"
	"github.com/lc-tut/club-portal/consts"
)

type UserInfo interface {
	GetUserID() string
	GetEmail() string
	GetName() string
	GetRole() consts.UserType
	ToUserResponse() *UserResponse
}

type User struct {
	UserUUID       string          `gorm:"type:char(36);not null;primaryKey"`
	AdminUsers     []AdminUser     `gorm:"foreignKey:UserUUID;references:UserUUID"`
	GeneralUsers   []GeneralUser   `gorm:"foreignKey:UserUUID;references:UserUUID"`
	DomainUsers    []DomainUser    `gorm:"foreignKey:UserUUID;references:UserUUID"`
	UploadedImages []UploadedImage `gorm:"foreignKey:UserUUID;references:UserUUID"`
}

type DomainUser struct {
	UserUUID  string         `gorm:"type:char(36);not null;primaryKey"`
	Email     string         `gorm:"type:varchar(255);not null;unique"`
	Name      string         `gorm:"type:varchar(32);not null;unique"`
	Favorites []FavoriteClub `gorm:"foreignKey:UserUUID;references:UserUUID"`
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

func (u *DomainUser) GetRole() consts.UserType {
	role := consts.DomainUser
	return role
}

func (u *DomainUser) ToUserResponse() *UserResponse {
	res := &UserResponse{
		UserUUID: u.UserUUID,
		Email:    u.Email,
		Name:     u.Name,
		Role:     u.GetRole().ToPrimitive(),
	}

	return res
}

type GeneralUser struct {
	UserUUID string         `gorm:"type:char(36);not null;primaryKey"`
	Email    string         `gorm:"type:varchar(255);not null;unique"`
	Name     string         `gorm:"type:varchar(32);not null;unique"`
	ClubUUID sql.NullString `gorm:"foreignKey:ClubUUID;references:ClubUUID"`
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

func (u *GeneralUser) GetRole() consts.UserType {
	role := consts.GeneralUser
	return role
}

func (u *GeneralUser) ToUserResponse() *UserResponse {
	res := &UserResponse{
		UserUUID: u.UserUUID,
		Email:    u.Email,
		Name:     u.Name,
		Role:     u.GetRole().ToPrimitive(),
	}

	return res
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

func (u *AdminUser) GetRole() consts.UserType {
	role := consts.AdminUser
	return role
}

func (u *AdminUser) ToUserResponse() *UserResponse {
	res := &UserResponse{
		UserUUID: u.UserUUID,
		Email:    u.Email,
		Name:     u.Name,
		Role:     u.GetRole().ToPrimitive(),
	}

	return res
}

type UserResponse struct {
	UserUUID string `json:"user_uuid"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Role     string `json:"role"`
}

type GeneralUserSlice []GeneralUser

func (g GeneralUserSlice) GetEmails() []string {
	emails := make([]string, len(g))

	for i, email := range g {
		emails[i] = email.GetEmail()
	}

	return emails
}
