package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/lc-tut/club-portal/consts"
	"net/http"
)

func (mw *Middleware) AdminOnly() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		email := ctx.GetString(consts.UserEmail)

		if !mw.config.WhitelistUsers.IsAdminUser(email) {
			ctx.Status(http.StatusForbidden)
			ctx.Abort()
		}

		ctx.Next()
	}
}
