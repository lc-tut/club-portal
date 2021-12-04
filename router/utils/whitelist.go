package utils

import (
	models "github.com/lc-tut/club-portal/models/users"
	repos "github.com/lc-tut/club-portal/repos/users"
	"github.com/spf13/viper"
	"strings"
)

type WhitelistInfo interface {
	Users() []string
	GeneralEmails() []string
	AdminEmails() []string
	EmailDomains() []string
	IsUser(email string) bool
	IsDomainUser(email string) bool
	IsGeneralUser(email string) bool
	IsAdminUser(email string) bool
	AddGeneralUser(email string)
}

type Whitelist struct {
	users         []string
	emailDomains  []string
	adminEmails   []string
	generalEmails []string
}

func (w *Whitelist) Users() []string {
	return w.users
}

func (w *Whitelist) GeneralEmails() []string {
	return w.generalEmails
}

func (w *Whitelist) AdminEmails() []string {
	return w.adminEmails
}

func (w *Whitelist) EmailDomains() []string {
	return w.emailDomains
}

func (w *Whitelist) IsUser(email string) bool {
	for _, user := range w.users {
		if email == user || w.IsDomainUser(email) {
			return true
		}
	}

	return false
}

func (w *Whitelist) IsDomainUser(email string) bool {
	for _, domain := range w.emailDomains {
		if strings.HasSuffix(email, domain) {
			return true
		}
	}

	return false
}

func (w *Whitelist) IsAdminUser(email string) bool {
	for _, admin := range w.adminEmails {
		if email == admin {
			return true
		}
	}

	return false
}

func (w *Whitelist) IsGeneralUser(email string) bool {
	for _, general := range w.generalEmails {
		if email == general {
			return true
		}
	}

	return false
}

func (w *Whitelist) AddGeneralUser(email string) {
	w.generalEmails = append(w.generalEmails, email)
	w.users = append(w.users, email)
}

func NewWhitelist(userRepo repos.UserRepo) (WhitelistInfo, error) {
	ed := viper.GetStringSlice("email_domains")
	ae := viper.GetStringSlice("admin_emails")

	generalUser, err := userRepo.GetAllGeneralUser()

	if err != nil {
		return nil, err
	}

	typedGeneralUsers := models.GeneralUserSlice(generalUser)
	ge := typedGeneralUsers.GetEmails()

	users := append(ge, ae...)

	w := &Whitelist{
		users:         users,
		emailDomains:  ed,
		adminEmails:   ae,
		generalEmails: ge,
	}

	return w, nil
}
