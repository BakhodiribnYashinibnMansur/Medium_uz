package model

type Post struct {
	ID    int      `json:"-"`
	Title string   `json:"title" default:"Tutorial Golang"`
	Body  string   `json:"text" default:"Hello World"`
	Tags  []string `json:"tags" default:["Devs"] `
}
