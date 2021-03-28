package utils

import "github.com/gin-gonic/gin"

func DeleteCookie(ctx *gin.Context, name string) {
	ctx.SetCookie(name, "", -1, "", "", false, false)
}
