package common

import (
	"github.com/MuhammadSuryono/siakad-api-golang/model/school"
	"os/user"
	"time"

	"github.com/MuhammadSuryono/siakad-api-golang/common/db"
)

type Agama struct {
	ID   int64  `json:"id" gorm:"primaryKey"`
	Nama string `json:"nama"`
	AuditKolom
}

type StatusDalamKeluarga struct {
	ID     int64  `json:"id" gorm:"primaryKey"`
	Status string `json:"status"`
	AuditKolom
}

type TinggalBersama struct {
	ID             int64  `json:"id" gorm:"primaryKey"`
	TinggalBersama string `json:"tinggal_bersama"`
	AuditKolom
}

type AuditKolom struct {
	CreatedBy int64      `json:"created_by"`
	UpdatedBy int64      `json:"updated_by"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at" `
}

func MigrationDatabase() {
	db.Connection.Query.AutoMigrate(&user.User{})
	db.Connection.Query.AutoMigrate(&school.School{})
}
