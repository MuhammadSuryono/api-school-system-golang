package school

import "github.com/MuhammadSuryono/siakad-api-golang/model/common"

type School struct {
	ID          int64  `json:"id" gorm:"primaryKey"`
	NPSN        string `json:"npsn" gorm:"<-:create"`
	SchoolName  string `json:"school_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	common.AuditKolom
}

type DisplayListSchool struct {
	NPSN        string `json:"npsn" gorm:"<-:create"`
	SchoolName  string `json:"school_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
}

type DisplayDetailSchool struct {
	ID          int64  `json:"id"`
	NPSN        string `json:"npsn" gorm:"<-:create"`
	SchoolName  string `json:"school_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	common.AuditKolom
}

type SearchAbleSchool struct {
	NPSN       string `json:"npsn" gorm:"<-:create"`
	SchoolName string `json:"school_name"`
}
