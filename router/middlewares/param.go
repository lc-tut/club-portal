package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/lc-tut/club-portal/consts"
)

func (mw *Middleware) SetClubIDKey() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		v := ctx.Param("clubslug")
		ctx.Set(consts.ClubSlugKeyName, v)
		ctx.Next()
	}
}
