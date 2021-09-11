package db

import (
	memberalumni "github.com/MuhammadSuryono/siakad-api-golang/model/alumni/member-alumni"
	"github.com/MuhammadSuryono/siakad-api-golang/model/school"
	"github.com/MuhammadSuryono/siakad-api-golang/model/user"
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

func MigrationDatabase() {
	Connection.Query.AutoMigrate(&user.User{})
	Connection.Query.AutoMigrate(&school.School{})
	Connection.Query.AutoMigrate(&memberalumni.MemberAlumni{})
}
