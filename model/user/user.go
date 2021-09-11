package user

import "github.com/MuhammadSuryono/siakad-api-golang/model/common"

type User struct {
	ID       int64  `json:"id" gorm:"primaryKey"`
	Username string `json:"username"`
	Password string `json:"password"`
	RoleId   int64  `json:"role_id"`
	common.AuditKolom
}

type CreateUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
	RoleId   int64  `json:"role_id"`
}

type UpdateUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
	RoleId   int64  `json:"role_id"`
}
