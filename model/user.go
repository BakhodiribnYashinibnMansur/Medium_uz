package model

import (
	"database/sql"

	"github.com/lib/pq"
)

type User struct {
	Id          int      `json:"-" `
	Email       string   `json:"email" default:"phapp0224mb@gmail.com"`
	Password    string   `json:"password" default:"0224mb"`
	NickName    string   `json:"nickname" default:"mrb"`
	FirstName   string   `json:"first_name" default:"MRB"`
	SecondName  string   `json:"second_name" default:"HERO"`
	Interesting []string `json:"interesting" default:""`
	Bio         string   `json:"bio" default:"I am Golang dev"`
	City        string   `json:"city" default:"Navoi"`
	Phone       string   `json:"phone" default:"+9989 93 753 65 71"`
}

type UserCheck struct {
	Email      bool `json:"email"   default:"false"  `
	Password   bool `json:"password" default:"true"   `
	NickName   bool `json:"nickname" default:"false"   `
	FirstName  bool `json:"first_name" default:"true"   `
	SecondName bool `json:"second_name" default:"true"  `
	City       bool `json:"city"   default:"true"  `
	Phone      bool `json:"phone"   default:"true"  `
}

type UserFull struct {
	Id                int            `json:"id" db:"id"`
	Email             string         `json:"email"  db:"email"`
	Nickname          string         `json:"nickname" db:"nickname"`
	Password          string         `json:"password" db:"password_hash"`
	FirstName         string         `json:"first_name" db:"firstname"`
	SecondName        string         `json:"second_name" db:"secondname"`
	Interesting       pq.StringArray `json:"interesting" db:"interesting"`
	Bio               string         `json:"bio" db:"bio"`
	City              string         `json:"city" db:"city"`
	IsVerified        bool           `json:"is_verified" db:"is_verified"`
	Verification_date sql.NullTime   `json:"verification_date" db:"verification_date"`
	AccountImagePath  sql.NullString `json:"account_image_path" db:"account_image_path"`
	Phone             string         `json:"phone" db:"phone"`
	Rating            string         `json:"rating" db:"rating"`
	PostViewsCount    int            `json:"post_views" db:"post_views_count"`
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
	Email    string `json:"email"  default:"phapp0224mb@gmail.com"`
	Password string `json:"password" default:"0224mb"`
}

type UpdateUser struct {
	Id          int              `json:"-" `
	Email       sql.NullString   `json:"email" default:"phapp0224mb@gmail.com"`
	Password    sql.NullString   `json:"password" default:"0224mb"`
	NickName    sql.NullString   `json:"nickname" default:"mrb"`
	FirstName   sql.NullString   `json:"first_name" default:"MRB"`
	SecondName  sql.NullString   `json:"second_name" default:"HERO"`
	Interesting []sql.NullString `json:"interesting" default:""`
	Bio         sql.NullString   `json:"bio" default:"I am Golang dev"`
	City        sql.NullString   `json:"city" default:"Navoi"`
	Phone       sql.NullString   `json:"phone" default:"+9989 93 753 65 71"`
}
