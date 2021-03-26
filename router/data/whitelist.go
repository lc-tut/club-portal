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
}

type Whitelist struct {
	users         []string
	generalEmails []string
	adminEmails   []string
	emailDomains  []string
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

	ge := viper.GetStringSlice("general_emails")
	ae := viper.GetStringSlice("admin_emails")
	ed := viper.GetStringSlice("email_domains")

	users := append(ge, ae...)

	w := &Whitelist{
		users:         users,
		generalEmails: ge,
		adminEmails:   ae,
		emailDomains:  ed,
	}

	return w, nil
}
