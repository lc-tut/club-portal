package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/lc-tut/club-portal/consts"
	"go.uber.org/zap"
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
			mw.logger.Warn("invalid user", zap.String("session_user_uuid", sessUUID), zap.String("param_uuid", paramUUID), zap.String("email", email))
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
			mw.logger.Warn("invalid user", zap.String("email", email))
			ctx.AbortWithStatus(http.StatusForbidden)
			return
		}

		ctx.Next()
	}
}

func (mw *Middleware) OverGeneralOnly() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		email := ctx.GetString(consts.SessionUserEmail)

		if !mw.config.WhitelistUsers.IsGeneralUser(email) && !mw.config.WhitelistUsers.IsAdminUser(email) {
			mw.logger.Warn("invalid user", zap.String("email", email))
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
			mw.logger.Warn("invalid user", zap.String("email", email))
			ctx.AbortWithStatus(http.StatusForbidden)
			return
		}

		ctx.Next()
	}
}

func (mw *Middleware) IdentifyClubUUID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sessUUID := ctx.GetString(consts.SessionUserUUID)
		paramUUID := ctx.GetString(consts.ClubUUIDKeyName)

		if sessUUID != paramUUID {
			mw.logger.Warn("invalid user", zap.String("session_user_uuid", sessUUID), zap.String("param_uuid", paramUUID))
			ctx.AbortWithStatus(http.StatusForbidden)
			return
		}

		ctx.Next()
	}
}
