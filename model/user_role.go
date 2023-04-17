package model

type UserRole struct {
	ID     uint `gorm:"primary_key"`
	UserID int  `json:"user_id"`
	RoleID int  `json:"role_id"`
}
