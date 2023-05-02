package model

import "time"

type UserGold struct {
	ID        uint       `gorm:"primary_key"`
	UserID    int        `json:"user_id"`
	Weight    float64    `json:"weight"`
	BuyPrice  float64    `json:"buy_price"`
	SellPrice *float64   `json:"sell_price"`
	BuyDate   time.Time  `json:"buy_date"`
	SellDate  *time.Time `json:"sell_date"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

type CreateUserGoldRequest struct {
	UserID   int     `json:"user_id"`
	Weight   float64 `json:"weight"`
	BuyPrice float64 `json:"buy_price"`
	BuyDate  string  `json:"buy_date"`
}

type CreateUserGoldResponse struct {
	ID       int     `json:"id"`
	UserID   int     `json:"user_id"`
	Weight   float64 `json:"weight"`
	BuyPrice float64 `json:"buy_price"`
	BuyDate  string  `json:"buy_date"`
}

type UserGoldViewParams struct {
	UserID        int
	NullSellPrice bool
	BasedFilter
}

type GetUnrealizedRequest struct {
	UserID int
	Page   int `json:"page" form:"page" query:"page"`
	Limit  int `json:"limit" form:"limit" query:"limit"`
}

type GetUnrealizedResponse struct {
	ID           int     `json:"id"`
	UserID       int     `json:"user_id"`
	Weight       float64 `json:"weight"`
	BuyPrice     float64 `json:"buy_price"`
	Unrealized   float64 `json:"unrealized"`
	Percent      float64 `json:"percent"`
	AvgBuyPrice  float64 `json:"avg_buy_price"`
	CurrentPrice float64 `json:"current_price"`
	BuyDate      string  `json:"buy_date"`
}
