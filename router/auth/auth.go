package auth

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/lc-tut/club-portal/consts"
	utils2 "github.com/lc-tut/club-portal/router/utils"
	"net/http"
)

func (h *Handler) Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sess := sessions.Default(ctx)
		sessionData := sess.Get(consts.SessionKey)

		if sessionData != nil {
			s, err := utils2.ByteSliceToSessionData(sessionData.([]byte))
			if err != nil {
				ctx.Status(http.StatusInternalServerError)
				return
			}
			ctx.JSON(http.StatusOK, s)
		} else {
			ctx.Status(http.StatusUnauthorized)
		}
	}
}
