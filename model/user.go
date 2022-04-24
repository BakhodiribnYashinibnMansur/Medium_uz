package model

type User struct {
	Id         int    `json:"-" `
	Email      string `json:"email" default:"phapp0224mb@gmail.com"`
	Password   string `json:"password" default:"0224mb"`
	FirstName  string `json:"first_name" default:"MRB"`
	SecondName string `json:"second_name" default:"HERO"`
	Image      string `json:"image" `
	City       string `json:"city" default:"Navoi"`
	Phone      string `json:"phone" default:"+9989 93 753 65 71"`
}

type ResponseSign struct {
	Id    int    `json:"id" `
	Token string `json:"token"`
}

type SignInInput struct {
	Username string `json:"username"  default:"MRB"`
	Password string `json:"password" default:"0224mb"`
}
