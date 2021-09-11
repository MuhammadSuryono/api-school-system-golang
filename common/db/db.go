package db

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

type connection struct {
	Query *gorm.DB
}

var Connect *connection

func Init() {
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASS")
	PORT := os.Getenv("DB_PORT")
	HOST := os.Getenv("DB_HOST")
	DB := os.Getenv("DB_NAME")
	args := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		USER, PASS, HOST, PORT, DB)
	db, err := gorm.Open(mysql.Open(args))
	if err != nil {
		panic(fmt.Sprintf("failed to connect database with setting: %s", args))
	}

	Connect = &connection{Query: db}

}

func NewDbOnSameConnection(dbName string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASS")
	PORT := os.Getenv("DB_PORT")
	HOST := os.Getenv("DB_HOST")
	args := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		USER, PASS, HOST, PORT, dbName)
	db, err := gorm.Open(mysql.Open(args))
	if err != nil {
		fmt.Println("error ", err)
	}

	Connect = &connection{Query: db}
}


