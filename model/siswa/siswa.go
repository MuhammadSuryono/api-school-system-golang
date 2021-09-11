package siswa

import "github.com/MuhammadSuryono/siakad-api-golang/model/common"

type Siswa struct {
	ID                  int64  `json:"id" gorm:"primaryKey"`
	NISN                string `json:"nisn" gorm:"<-:create"`
	NamaLengkap         string `json:"nama_lengkap"`
	TempatLahir         string `json:"tempat_lahir"`
	TanggalLahir        string `json:"tanggal_lahir"`
	JenisKelamin        string `json:"jenis_kelamin"`
	Agama               int64  `json:"agama"`
	AnakKe              int    `json:"anak_ke"`
	StatusDalamKeluarga int64  `json:"status_dalam_keluarga"`
	TinggalBersama      int64  `json:"tinggal_bersama"`
	Alamat              string `json:"alamat"`
	common.AuditKolom
}

type DisplayListSiswa struct {
	NISN                string `json:"nisn"`
	NamaLengkap         string `json:"nama_lengkap"`
	TempatLahir         string `json:"tempat_lahir"`
	TanggalLahir        string `json:"tanggal_lahir"`
	JenisKelamin        string `json:"jenis_kelamin"`
	Agama               string `json:"agama"`
	AnakKe              int    `json:"anak_ke"`
	StatusDalamKeluarga string `json:"status_dalam_keluarga"`
	TinggalBersama      string `json:"tinggal_bersama"`
	Alamat              string `json:"alamat"`
}

type DisplayDetailSiswa struct {
	ID                  int64  `json:"id" gorm:"primaryKey"`
	NISN                string `json:"nisn"`
	NamaLengkap         string `json:"nama_lengkap"`
	TempatLahir         string `json:"tempat_lahir"`
	TanggalLahir        string `json:"tanggal_lahir"`
	JenisKelamin        string `json:"jenis_kelamin"`
	Agama               string `json:"agama"`
	AnakKe              int    `json:"anak_ke"`
	StatusDalamKeluarga string `json:"status_dalam_keluarga"`
	TinggalBersama      string `json:"tinggal_bersama"`
	Alamat              string `json:"alamat"`
	common.AuditKolom
}

type SearchAbleSiswa struct {
	NISN                string `json:"nisn"`
	NamaLengkap         string `json:"nama_lengkap"`
	Agama               int64  `json:"agama"`
	StatusDalamKeluarga int64  `json:"status_dalam_keluarga"`
	TinggalBersama      int64  `json:"tinggal_bersama"`
}