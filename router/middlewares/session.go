package middlewares

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/lc-tut/club-portal/consts"
	"github.com/lc-tut/club-portal/router/utils"
	"go.uber.org/zap"
	"net/http"
)

func (mw *Middleware) CheckSession() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sess := sessions.Default(ctx)
		sessionData, ok := sess.Get(consts.SessionKey).([]byte)

		if !ok {
			mw.logger.Error("(session) failed type assert")
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		s, err := utils.ByteSliceToSessionData(sessionData)

		if err != nil {
			mw.logger.Error(err.Error())
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		ctx.Set(consts.SessionUserUUID, s.UserUUID)
		ctx.Set(consts.SessionUserEmail, s.Email)
		ctx.Set(consts.SessionUserName, s.Name)
		ctx.Set(consts.SessionUserRole, s.Role)

		mw.logger.Debug("set sessions to context",
			zap.String("user_uuid", s.UserUUID),
			zap.String("email", s.Email),
			zap.String("name", s.Name),
			zap.String("role", s.Role),
		)

		ctx.Next()
	}
}
