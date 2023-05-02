package model

type CurrentGolds struct {
	Data []CurrentGoldBase `json:"data"`
}

type CurrentGoldBase struct {
	Sell float64 `json:"sell"`
	Buy  float64 `json:"buy"`
	Type string  `json:"type"`
}
