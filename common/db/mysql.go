package db

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func (d *databaseHandler) InitMysql() {
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

	Connection = connection{Query: db}

}
