package auth

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/lc-tut/club-portal/consts"
	"github.com/lc-tut/club-portal/utils"
	"go.uber.org/zap"
	"net/http"
)

func (h *Handler) SignIn() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sess := sessions.Default(ctx)
		sessionData := sess.Get(consts.SessionKey)

		if sessionData != nil {
			h.logger.Info("logged user accessed")
			ctx.Status(http.StatusNoContent)
			return
		}

		state, err := utils.GenerateCSRFState()

		if err != nil {
			h.logger.Error(err.Error())
			ctx.Status(http.StatusInternalServerError)
			return
		}

		csrfCookie := h.config.CSRFCookieOptions
		ctx.SetSameSite(csrfCookie.SameSite)
		ctx.SetCookie(consts.AuthCSRFCookieName, state, csrfCookie.MaxAge, csrfCookie.Path, csrfCookie.Domain, csrfCookie.Secure, csrfCookie.HttpOnly)
		h.logger.Info("created cookie",
			zap.String("cookie_name", consts.AuthCSRFCookieName),
		)

		url := h.config.GoogleOAuthConfig.AuthCodeURL(state)

		ctx.Redirect(http.StatusFound, url)
	}
}
