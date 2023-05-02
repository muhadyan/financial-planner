package model

type BasicResp struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Pagination struct {
	BasicResp
	Meta
}

type Meta struct {
	Page         int `json:"page,omitempty"`
	Limit        int `json:"limit,omitempty"`
	TotalRecords int `json:"total_records,omitempty"`
	TotalPages   int `json:"total_pages,omitempty"`
}
