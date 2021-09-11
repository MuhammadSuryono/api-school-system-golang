package db

import (
	"os"

	"gorm.io/gorm"
)

type databaseHandler struct {
}

type connection struct {
	Query *gorm.DB
}

var (
	Connection connection
)

func CreateConnection() {
	driver := os.Getenv("DB_DRIVER")
	conn := &databaseHandler{}

	if driver == "mysql" {
		conn.InitMysql()
	} else if driver == "pgsql" {
		conn.InitDatabasePostgre()
	}
}
