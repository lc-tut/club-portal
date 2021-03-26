package data

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"net/http"
)

func newCookieOptions(path, domain string, maxage int, secure, httponly bool, samesite http.SameSite) *sessions.Options {
	c := &sessions.Options{
		Path:     path,
		Domain:   domain,
		MaxAge:   maxage,
		Secure:   secure,
		HttpOnly: httponly,
		SameSite: samesite,
	}

	return c
}

type Config struct {
	SessionCookieOptions *sessions.Options
	CSRFCookieOptions    *sessions.Options
	GoogleOAuthConfig    *oauth2.Config
	WhitelistUsers       WhitelistInfo
}

func NewConfig(local bool) *Config {
	whitelist := NewWhitelist()

	conf := &Config{
		WhitelistUsers: whitelist,
	}

	newOAuthConf := func(redirectURL string) (conf *oauth2.Config) {
		conf = &oauth2.Config{
			ClientID:     viper.GetString("client_id"),
			ClientSecret: viper.GetString("client_secret"),
			Endpoint:     google.Endpoint,
			RedirectURL:  redirectURL,
			Scopes:       []string{"profile", "email"},
		}
		return
	}

	if local {
		redirectURL := "http://localhost:8080/api/auth/callback"
		conf.SessionCookieOptions = newCookieOptions("/", "localhost", 60*60*24*7, false, true, http.SameSiteLaxMode)
		conf.CSRFCookieOptions = newCookieOptions("/", "localhost", 60*15, false, true, http.SameSiteLaxMode)
		conf.GoogleOAuthConfig = newOAuthConf(redirectURL)
	} else {
		domain := viper.GetString("domain")
		redirectURL := fmt.Sprintf("http://%s:8080/api/auth/callback", domain)
		conf.SessionCookieOptions = newCookieOptions("/", domain, 60*60*24*7, true, true, http.SameSiteLaxMode)
		conf.CSRFCookieOptions = newCookieOptions("/", domain, 60*15, true, true, http.SameSiteStrictMode)
		conf.GoogleOAuthConfig = newOAuthConf(redirectURL)
	}

	return conf
}

func NewRedisServer() (redis.Store, error) {
	secretKey := viper.GetString("redis_secret")
	store, err := redis.NewStore(10, "tcp", "redis:6379", viper.GetString("redis_password"), []byte(secretKey))

	if err != nil {
		return nil, err
	}

	return store, nil
}
