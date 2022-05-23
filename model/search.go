package model

type Search struct {
	Person []PostFull `json:"person"`
	Post   []PostFull `json:"post"`
}
