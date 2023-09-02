package models

type Pagination struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	TotalPages int `json:"totalPages"`
	TotalData  int `json:"totalData"`
}
