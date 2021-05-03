package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/lc-tut/club-portal/consts"
	"go.uber.org/zap"
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
		ctx.Set(consts.ImageIDKeyName, v)
		mw.logger.Debug("set imageid value to context", zap.String("image_id", v))
		ctx.Next()
	}
}
