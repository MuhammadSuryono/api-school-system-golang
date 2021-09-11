package main

import (
	"github.com/MuhammadSuryono/siakad-api-golang/common/db"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()
	r := gin.Default()

	r.GET("ping", func(c *gin.Context) {
		c.JSON(200, "PONG")
	})

	r.Run()
}
