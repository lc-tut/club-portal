package utils

import (
	"github.com/gin-contrib/sessions"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"net/http"
	"os"
)

var (
	CookieOption = sessions.Options{
		Path:     "/",
		Domain:   os.Getenv("DOMAIN"),
		MaxAge:   60 * 60 * 24 * 7,
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}

	AuthConfig = &oauth2.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		Endpoint:     google.Endpoint,
		RedirectURL:  "http://localhost:8080/api/auth/callback",
		Scopes:       []string{"profile", "email"},
	}
)
