package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/lc-tut/club-portal/consts"
	"net/http"
)

func (mw *Middleware) PersonalOrAdminOnly() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		email := ctx.GetString(consts.SessionUserEmail)

		if mw.config.WhitelistUsers.IsAdminUser(email) {
			ctx.Next()
			return
		}

		sessUUID := ctx.GetString(consts.SessionUserUUID)
		paramUUID := ctx.GetString(consts.UserUUIDKeyName)

		if sessUUID != paramUUID {
			ctx.AbortWithStatus(http.StatusForbidden)
			return
		}

		ctx.Next()
	}
}

func (mw *Middleware) UserOnly() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		email := ctx.GetString(consts.SessionUserEmail)

		if !mw.config.WhitelistUsers.IsUser(email) {
			ctx.AbortWithStatus(http.StatusForbidden)
			return
		}

		ctx.Next()
	}
}

func (mw *Middleware) OverGeneralOnly() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		email := ctx.GetString(consts.SessionUserEmail)

		if !mw.config.WhitelistUsers.IsGeneralUser(email) || !mw.config.WhitelistUsers.IsAdminUser(email) {
			ctx.AbortWithStatus(http.StatusForbidden)
			return
		}

		ctx.Next()
	}
}

func (mw *Middleware) AdminOnly() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		email := ctx.GetString(consts.SessionUserEmail)

		if !mw.config.WhitelistUsers.IsAdminUser(email) {
			ctx.AbortWithStatus(http.StatusForbidden)
			return
		}

		ctx.Next()
	}
}
