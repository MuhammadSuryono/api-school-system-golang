package memberalumni

import "github.com/MuhammadSuryono/siakad-api-golang/model/common"

type MemberAlumni struct {
	ID             int64  `json:"id" gorm:"primaryKey"`
	SchoolYearId   int64  `json:"school_year_id"`
	Name           string `json:"name"`
	GraduationYear string `json:"graduation_year"`
	PhoneNumber    string `json:"phone_number"`
	IsWhatsapp     bool   `json:"is_whatsapp"`
	UrlPhoto       string `json:"url_photo"`
	common.AuditKolom
}

type CreateMemberAlumni struct {
	SchoolYearId   int64  `json:"school_year_id"`
	Name           string `json:"name"`
	GraduationYear string `json:"graduation_year"`
	PhoneNumber    string `json:"phone_number"`
	IsWhatsapp     bool   `json:"is_whatsapp"`
	UrlPhoto       string `json:"url_photo"`
}
