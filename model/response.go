package model

type BasicResp struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
