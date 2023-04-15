package model

import "time"

type User struct {
	ID        uint      `gorm:"primary_key"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	Fullname  string    `json:"fullname"`
	IsActive  bool      `json:"is_active"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SignUpRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
}

type SignUpResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
}

type VerifyRequest struct {
	UserID   int    `query:"user_id"`
	Username string `query:"username"`
}
