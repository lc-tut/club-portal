package auth

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/lc-tut/club-portal/consts"
	"net/http"
)

func (h *AuthHandler) Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sess := sessions.Default(ctx)
		authState := sess.Get(consts.SessionKey)

		if authState != nil {
			ctx.Status(http.StatusOK)
		} else {
			ctx.Status(http.StatusUnauthorized)
		}
	}
}
