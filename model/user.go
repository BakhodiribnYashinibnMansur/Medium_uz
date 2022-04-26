package model

import "database/sql"

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

type UserFull struct {
	Id                int            `json:"-" db:"id"`
	Email             string         `json:"email"  db:"email"`
	FirstName         string         `json:"first_name" db:"firstname"`
	SecondName        string         `json:"second_name" db:"secondname"`
	City              string         `json:"city" db:"city"`
	IsVerified        bool           `json:"is_verified" db:"is_verified"`
	Verification_date sql.NullTime   `json:"verification_date" db:"verification_date"`
	AccountImagePath  sql.NullString `json:"account_image_path" db:"account_image_path"`
	Phone             string         `json:"phone" db:"phone"`
	Rating            string         `json:"rating" db:"rating"`
	PostViews         int            `json:"post_views" db:"post_views"`
	IsSuperUser       bool           `json:"is_super_user" db:"is_super_user"`
	CreatedAt         string         `json:"created_at" db:"created_at"`
}
type ResponseSign struct {
	Id    int    `json:"id" `
	Token string `json:"token"`
}

type SignInInput struct {
	Username string `json:"username"  default:"MRB"`
	Password string `json:"password" default:"0224mb"`
}
