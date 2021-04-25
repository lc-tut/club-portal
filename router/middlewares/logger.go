package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/lc-tut/club-portal/consts"
	"go.uber.org/zap"
	"time"
)

// LoggerMiddleware アクセスごとにログを記録する.
//
// 基本的には "github.com/gin-contrib/zap" と同じ処理だが,
// 標準時を JST に強制変換させている.
func LoggerMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now().In(consts.JST)

		path := ctx.Request.URL.Path
		query := ctx.Request.URL.RawQuery
		ctx.Next()

		end := time.Now().In(consts.JST)
		latency := end.Sub(start)

		if len(ctx.Errors) > 0 {
			for _, e := range ctx.Errors.Errors() {
				logger.Error(e)
			}
		} else {
			logger.Info(path,
				zap.Int("status", ctx.Writer.Status()),
				zap.String("method", ctx.Request.Method),
				zap.String("path", path),
				zap.String("query", query),
				zap.String("ip", ctx.ClientIP()),
				zap.String("user-agent", ctx.Request.UserAgent()),
				zap.String("time", end.Format(time.RFC3339)),
				zap.Duration("latency", latency),
			)

		}
	}
}
