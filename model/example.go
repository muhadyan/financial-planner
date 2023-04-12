package model

import "time"

type Example struct {
	ID        uint   `gorm:"primary_key"`
	Name      string `json:"name"`
	Username  string `json:"username"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
