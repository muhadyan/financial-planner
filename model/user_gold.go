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

type FindAllUserGoldRequest struct {
	UserID int
	PagiantionReq
}

type UpdateUserGoldRequest struct {
	ID       int
	UserID   int
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
	UserID  int
	UseNPWP bool `json:"use_npwp" form:"use_npwp" query:"use_npwp"`
	PagiantionReq
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

type SellGoldRequest struct {
	ID        int
	UserID    int
	SellPrice float64 `json:"sell_price"`
	SellDate  string  `json:"sell_date"`
}

type DashboardUserGoldRequest struct {
	UserID  int
	UseNPWP bool `json:"use_npwp" form:"use_npwp" query:"use_npwp"`
}

type DashboardUserGoldResponse struct {
	UserID                         int     `json:"user_id"`
	TotalWeight                    float64 `json:"total_weight"`
	AvgBuyPricePerGram             float64 `json:"avg_buy_price_per_gram"`
	AvgSellPricePerGram            float64 `json:"avg_sell_price_per_gram"`
	HighestUnrealizedMargin        float64 `json:"highest_unrealized_margin"`
	HighestUnrealizedMarginPercent float64 `json:"highest_unrealized_margin_percent"`
	LowestUnrealizedMargin         float64 `json:"lowest_unrealized_margin"`
	LowestUnrealizedMarginPercent  float64 `json:"lowest_unrealized_margin_percent"`
	AvgUnrealizedMargin            float64 `json:"avg_unrealized_margin"`
	AvgUnrealizedMarginPercent     float64 `json:"avg_unrealized_margin_percent"`
	HighestRealizedMargin          float64 `json:"highest_realized_margin"`
	HighestRealizedMarginPercent   float64 `json:"highest_realized_margin_percent"`
	LowestRealizedMargin           float64 `json:"lowest_realized_margin"`
	LowestRealizedMarginPercent    float64 `json:"lowest_realized_margin_percent"`
	AvgRealizedMargin              float64 `json:"avg_realized_margin"`
	AvgRealizedMarginPercent       float64 `json:"avg_realized_margin_percent"`
	OldestPossession               string  `json:"oldest_possession"`
	NewestPossession               string  `json:"newest_possession"`
	AvgPossession                  string  `json:"avg_possession"`
}
