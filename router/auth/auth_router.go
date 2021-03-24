package auth

import (
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
}

func newAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func AddAuthRouter(rg *gin.RouterGroup) {
	h := newAuthHandler()

	authGroup := rg.Group("/auth")
	{
		authGroup.GET("/", h.Auth())
		authGroup.GET("/signin", h.SignIn())
		authGroup.GET("/callback", h.Callback())
		authGroup.POST("/destroy", h.Destroy())
	}
}
