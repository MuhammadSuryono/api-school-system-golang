package main

import (
	"github.com/MuhammadSuryono/siakad-api-golang/common/db"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	// Create connection database
	db.CreateConnection()
	// Migration Database
	db.MigrationDatabase()

	r := gin.Default()
	r.GET("ping", func(c *gin.Context) {
		c.JSON(200, "PONG")
	})

	r.Run()
}
