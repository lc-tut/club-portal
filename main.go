package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/lc-tut/club-portal/consts"
	"github.com/lc-tut/club-portal/router"
	"github.com/lc-tut/club-portal/utils"
)

func main() {
	r := gin.Default()
	// TODO: keyPairs must be secured
	store, err := redis.NewStore(10, "tcp", "redis:6379", "", []byte("secret"))
	if err != nil {
		panic(err)
	}
	store.Options(utils.CookieOption)
	r.Use(sessions.Sessions(consts.SessionCookieName, store))

	router.Init(r)

	_ = r.Run(":8080")
}
