package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/lc-tut/club-portal/consts"
	"github.com/lc-tut/club-portal/router/utils"
	"net/http"
)

func (mw *Middleware) CheckSession() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sess := sessions.Default(ctx)
		sessionData, ok := sess.Get(consts.SessionKey).([]byte)

		if !ok {
			ctx.Status(http.StatusUnauthorized)
			ctx.Abort()
			return
		}

		s, err := utils.ByteSliceToSessionData(sessionData)

		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			ctx.Abort()
			return
		}

		ctx.Set(consts.UserEmail, s.Email)

		ctx.Next()
	}
}
