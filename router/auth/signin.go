package auth

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/lc-tut/club-portal/consts"
	"github.com/lc-tut/club-portal/utils"
	"go.uber.org/zap"
	"net/http"
	"net/url"
	"strings"
)

type signInQuery struct {
	RedirectURL string `form:"redirect_url" binding:"required"`
}

func (h *Handler) SignIn() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		query := &signInQuery{}
		if err := ctx.ShouldBindQuery(query); err != nil {
			query.RedirectURL = "/"
		}

		if !strings.HasPrefix(query.RedirectURL, "/") {
			h.logger.Error("redirect_url should be started with '/'")
			ctx.Status(http.StatusBadRequest)
			return
		}

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
		h.setRedirectURLCookie(ctx, query.RedirectURL, csrfCookie.Path, csrfCookie.Domain, csrfCookie.Secure, csrfCookie.HttpOnly, csrfCookie.SameSite)

		redirectUrl := h.config.GoogleOAuthConfig.AuthCodeURL(state)

		ctx.Redirect(http.StatusFound, redirectUrl)
	}
}

func (h *Handler) setRedirectURLCookie(ctx *gin.Context, value, path, domain string, secure, httpOnly bool, sameSite http.SameSite) {
	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:     consts.RedirectURLCookieName,
		Value:    url.QueryEscape(value),
		Path:     path,
		Domain:   domain,
		Secure:   secure,
		HttpOnly: httpOnly,
		SameSite: sameSite,
	})
	h.logger.Info("created cookie",
		zap.String("cookie_name", consts.RedirectURLCookieName),
	)
}
