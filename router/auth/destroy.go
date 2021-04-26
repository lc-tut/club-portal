package auth

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/lc-tut/club-portal/consts"
	"github.com/lc-tut/club-portal/router/utils"
	"go.uber.org/zap"
	"net/http"
)

func (h *Handler) Destroy() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			h.logger.Info("deleted cookie", zap.String("cookie_name", consts.SessionCookieName))
			utils.DeleteCookie(ctx, consts.SessionCookieName)
		}()

		sess := sessions.Default(ctx)
		sess.Set(consts.SessionKey, nil)

		if err := sess.Save(); err != nil {
			h.logger.Error(err.Error())
			ctx.Status(http.StatusInternalServerError)
		} else {
			ctx.Status(http.StatusCreated)
		}
	}
}
