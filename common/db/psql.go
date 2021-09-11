package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func (d *databaseHandler) InitDatabasePostgre() {
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASS")
	PORT := os.Getenv("DB_PORT")
	HOST := os.Getenv("DB_HOST")
	DB := os.Getenv("DB_NAME")

	args := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", HOST, USER, PASS, DB, PORT)
	db, err := gorm.Open(postgres.Open(args), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect database with setting: %s", args))
	}

	Connection = connection{Query: db}
}
