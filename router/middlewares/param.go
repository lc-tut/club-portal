package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/lc-tut/club-portal/consts"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

func (mw *Middleware) SetClubIDKey() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		v := ctx.Param("clubslug")
		ctx.Set(consts.ClubSlugKeyName, v)
		mw.logger.Debug("set clubslug value to context", zap.String("club_slug", v))
		ctx.Next()
	}
}

func (mw *Middleware) SetUserUUIDKey() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		v := ctx.Param("uuid")
		ctx.Set(consts.UserUUIDKeyName, v)
		mw.logger.Debug("set uuid value to context", zap.String("user_uuid", v))
		ctx.Next()
	}
}

func (mw *Middleware) SetImageIDKey() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		v := ctx.Param("imageid")
		u, err := strconv.ParseUint(v, 10, 32)

		if err != nil {
			mw.logger.Error(err.Error())
			ctx.AbortWithStatus(http.StatusInternalServerError)
		}

		ctx.Set(consts.ImageIDKeyName, u)
		mw.logger.Debug("set imageid value to context", zap.Uint64("image_id", u))
		ctx.Next()
	}
}
