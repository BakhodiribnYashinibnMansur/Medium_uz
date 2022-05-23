package model

type Search struct {
	User []UserFull `json:"user"`
	Post []PostFull `json:"post"`
}
