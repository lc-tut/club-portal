package auth

import (
	"github.com/gin-gonic/gin"
)

func deleteCookie(ctx *gin.Context, name string) {
	ctx.SetCookie(name, "", -1, "", "", false, false)
}
