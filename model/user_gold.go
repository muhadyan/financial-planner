package model

import "time"

type UserGold struct {
	ID        uint       `gorm:"primary_key"`
	UserID    int        `json:"user_id"`
	BuyPrice  float64    `json:"buy_price"`
	SellPrice *float64   `json:"sell_price"`
	BuyDate   time.Time  `json:"buy_date"`
	SellDate  *time.Time `json:"sell_date"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

type CreateUserGoldRequest struct {
	UserID   int     `json:"user_id"`
	BuyPrice float64 `json:"buy_price"`
	BuyDate  string  `json:"buy_date"`
}

type CreateUserGoldResponse struct {
	ID       int     `json:"id"`
	UserID   int     `json:"user_id"`
	BuyPrice float64 `json:"buy_price"`
	BuyDate  string  `json:"buy_date"`
}
