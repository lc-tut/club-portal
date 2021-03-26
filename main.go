package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lc-tut/club-portal/router"
)

func main() {
	r := gin.Default()

	if err := loadConfig(); err != nil {
		panic(err)
	}

	router.Init(r)

	_ = r.Run(":8080")
}
