package auth

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/lc-tut/club-portal/consts"
	"github.com/lc-tut/club-portal/utils"
	"net/http"
)

func (h *Handler) SignIn() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sess := sessions.Default(ctx)
		sessionData := sess.Get(consts.SessionKey)

		if sessionData != nil {
			ctx.Status(http.StatusNoContent)
			return
		}

		state, err := utils.CreateAuthState()

		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return
		}

		csrfCookie := h.config.CSRFCookieOptions
		ctx.SetSameSite(csrfCookie.SameSite)
		ctx.SetCookie(consts.AuthCSRFCookieName, state, csrfCookie.MaxAge, csrfCookie.Path, csrfCookie.Domain, csrfCookie.Secure, csrfCookie.HttpOnly)

		url := h.config.GoogleOAuthConfig.AuthCodeURL(state)

		ctx.Redirect(http.StatusFound, url)
	}
}
