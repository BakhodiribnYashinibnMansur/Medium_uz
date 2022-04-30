package model

import "database/sql"

type User struct {
	Id          int      `json:"-" `
	Email       string   `json:"email" default:"phapp0224mb@gmail.com"`
	Password    string   `json:"password" default:"0224mb"`
	NickName    string   `json:"nickname" default:"mrb"`
	FirstName   string   `json:"first_name" default:"MRB"`
	SecondName  string   `json:"second_name" default:"HERO"`
	Interesting []string `json:"interesting" default:""`
	Bio         string   `json:"bio" default:"I am Golang dev"`
	Image       string   `json:"image" `
	City        string   `json:"city" default:"Navoi"`
	Phone       string   `json:"phone" default:"+9989 93 753 65 71"`
}

type UserCheck struct {
	Email       bool `json:"email"   default:"false"  `
	Password    bool `json:"password" default:"true"   `
	NickName    bool `json:"nickname" default:"false"   `
	FirstName   bool `json:"first_name" default:"true"   `
	SecondName  bool `json:"second_name" default:"true"  `
	Interesting bool `json:"interesting" default:"true"  `
	Bio         bool `json:"bio"  default:"true"  `
	City        bool `json:"city"   default:"true"  `
	Phone       bool `json:"phone"   default:"true"  `
}

type UserFull struct {
	Id                int            `json:"-" db:"id"`
	Email             string         `json:"email"  db:"email"`
	Nickname          string         `json:"nickname" db:"nickname"`
	Password          string         `json:"password" db:"password_hash"`
	FirstName         string         `json:"first_name" db:"firstname"`
	SecondName        string         `json:"second_name" db:"secondname"`
	Interesting       []string       `json:"interesting" db:"interesting"`
	Bio               string         `json:"bio" db:"bio"`
	City              string         `json:"city" db:"city"`
	IsVerified        bool           `json:"is_verified" db:"is_verified"`
	Verification_date sql.NullTime   `json:"verification_date" db:"verification_date"`
	AccountImagePath  sql.NullString `json:"account_image_path" db:"account_image_path"`
	Phone             string         `json:"phone" db:"phone"`
	Rating            string         `json:"rating" db:"rating"`
	PostViewsCount    int            `json:"post_views" db:"post_views"`
	FollowerCount     int            `json:"follower_count" db:"follower_count"`
	FollowingCount    int            `json:"following_count" db:"following_count"`
	LikeCount         int            `json:"like_count" db:"like_count"`
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
