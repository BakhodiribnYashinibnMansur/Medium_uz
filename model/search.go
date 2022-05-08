package model

type Search struct {
	Pesson   []PostFull `json:"person"`
	Post     []PostFull `json:"post"`
	Messsage string     `json:"message"`
}
