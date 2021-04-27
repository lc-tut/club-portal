package auth

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/lc-tut/club-portal/consts"
	"github.com/lc-tut/club-portal/router/utils"
	"go.uber.org/zap"
	"net/http"
)

func (h *Handler) Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sess := sessions.Default(ctx)
		sessionData, ok := sess.Get(consts.SessionKey).([]byte)

		if ok {
			s, err := utils.ByteSliceToSessionData(sessionData)
			if err != nil {
				ctx.Status(http.StatusInternalServerError)
				return
			}
			h.logger.Info("session user accessed",
				zap.String("session_uuid", s.SessionUUID),
				zap.String("user_uuid", s.UserUUID),
				zap.String("email", s.Email),
				zap.String("name", s.Name),
				zap.String("role", s.Role),
			)
			ctx.JSON(http.StatusOK, s)
		} else {
			h.logger.Info("no session user accessed")
			ctx.JSON(http.StatusOK, nil)
		}
	}
}
