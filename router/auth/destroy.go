package auth

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/lc-tut/club-portal/consts"
	"net/http"
)

func (h *Handler) Destroy() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer deleteCookie(ctx, consts.SessionCookieName)

		sess := sessions.Default(ctx)
		sess.Set(consts.SessionKey, nil)

		if err := sess.Save(); err != nil {
			ctx.Status(http.StatusInternalServerError)
		} else {
			ctx.Status(http.StatusCreated)
		}
	}
}
