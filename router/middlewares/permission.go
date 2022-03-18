package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/lc-tut/club-portal/consts"
	"github.com/lc-tut/club-portal/utils"
	"go.uber.org/zap"
	"net/http"
)

func (mw *Middleware) GeneralOnly() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		email := ctx.GetString(consts.SessionUserEmail)

		if !mw.config.WhitelistUsers.IsGeneralUser(email) {
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

func (mw *Middleware) IdentifyUUID(key string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sessUUID := ctx.GetString(consts.SessionUserUUID)
		paramUUID := ctx.GetString(key)
		if key == consts.ClubUUIDKeyName {
			res, err := mw.repo.GetGeneralUserByUUID(sessUUID)
			if err != nil {
				ctx.AbortWithStatus(http.StatusInternalServerError)
				return
			}
			clubUUID := utils.StringPToString(utils.NullStringToStringP(res.ClubUUID))
			if clubUUID != paramUUID {
				mw.logger.Warn("invalid user", zap.String("club_uuid", clubUUID), zap.String("param_uuid", paramUUID))
				ctx.AbortWithStatus(http.StatusForbidden)
				return
			}
		} else {
			if sessUUID != paramUUID {
				mw.logger.Warn("invalid user", zap.String("session_user_uuid", sessUUID), zap.String("param_uuid", paramUUID))
				ctx.AbortWithStatus(http.StatusForbidden)
				return
			}
		}

		ctx.Next()
	}
}
