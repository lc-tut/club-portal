package data

import (
	"fmt"
	"github.com/lc-tut/club-portal/consts"
	"github.com/spf13/viper"
	"os"
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

func NewWhitelist() (WhitelistInfo, error) {
	configFn := os.Getenv("CONFIG_FILE")

	if !strings.HasSuffix(configFn, consts.ConfigFileName) {
		return nil, fmt.Errorf("filename must be `%s`", consts.ConfigFileName)
	}

	path := strings.ReplaceAll(configFn, consts.ConfigFileName, "")

	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(path)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	ed := viper.GetStringSlice("email_domains")
	ae := viper.GetStringSlice("admin_emails")
	ge := viper.GetStringSlice("general_emails")

	users := append(ge, ae...)

	w := &Whitelist{
		users:         users,
		emailDomains:  ed,
		adminEmails:   ae,
		generalEmails: ge,
	}

	return w, nil
}
