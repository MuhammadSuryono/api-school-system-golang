package common

import (
	"time"
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
