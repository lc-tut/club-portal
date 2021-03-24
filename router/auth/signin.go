package auth

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/lc-tut/club-portal/consts"
	"github.com/lc-tut/club-portal/utils"
	"net/http"
)

func (h *AuthHandler) SignIn() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sess := sessions.Default(ctx)
		authState := sess.Get(consts.SessionKey)

		if authState != nil {
			ctx.Status(http.StatusNoContent)
			return
		}

		state, err := utils.CreateAuthState()

		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return
		}

		ctx.SetSameSite(http.SameSiteLaxMode)
		ctx.SetCookie(consts.AuthCSRFCookieName, state, 60*15, "/", "localhost", false, true)

		url := utils.AuthConfig.AuthCodeURL(state)

		ctx.Redirect(http.StatusFound, url)
	}
}
