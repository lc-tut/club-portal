package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/lc-tut/club-portal/consts"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

func (mw *Middleware) SetClubSlugKey() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		v := ctx.Param("clubslug")
		ctx.Set(consts.ClubSlugKeyName, v)
		mw.logger.Debug("set club_slug value to context", zap.String("club_slug", v))
		ctx.Next()
	}
}

func (mw *Middleware) SetUserUUIDKey() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		v := ctx.Param("useruuid")
		ctx.Set(consts.UserUUIDKeyName, v)
		mw.logger.Debug("set user_uuid value to context", zap.String("user_uuid", v))
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
		mw.logger.Debug("set image_id value to context", zap.Uint64("image_id", u))
		ctx.Next()
	}
}

func (mw *Middleware) SetClubUUIDKey() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		v := ctx.Param("clubuuid")
		ctx.Set(consts.ClubUUIDKeyName, v)
		mw.logger.Debug("set club_uuid value to context", zap.String("club_uuid", v))
		ctx.Next()
	}
}

func (mw *Middleware) SetThumbnailIDKey() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		v := ctx.Param("thumbnailid")

		u, err := strconv.ParseUint(v, 10, 32)

		if err != nil {
			mw.logger.Error(err.Error())
			ctx.AbortWithStatus(http.StatusInternalServerError)
		}

		ctx.Set(consts.ThumbnailIDKeyName, u)
		mw.logger.Debug("set thumbnail_id value to context", zap.Uint64("thumbnail_id", u))
	}
}
