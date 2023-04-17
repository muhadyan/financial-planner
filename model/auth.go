package model

import "github.com/dgrijalva/jwt-go"

type JWTPayload struct {
	ID       int      `json:"id"`
	Username string   `json:"username"`
	Email    string   `json:"email"`
	Fullname string   `json:"fullname"`
	IsActive bool     `json:"is_active"`
	Roles    []string `json:"roles"`
	jwt.StandardClaims
}
