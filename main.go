package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/lc-tut/club-portal/router"
)

func main() {
	r := gin.Default()

	router.Init(r)

	_ = r.Run(":8080")
}

// TODO: settings will not depend on .env (and will depend on config.toml)
