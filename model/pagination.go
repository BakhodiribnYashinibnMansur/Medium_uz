package model

type Pagination struct {
	Offset int `json:"offset" default:"0"`
	Limit  int `json:"limit" default:"10"`
}
