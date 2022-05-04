package model

import "github.com/lib/pq"

type Post struct {
	ID    int      `json:"-"`
	Title string   `json:"title" default:"Tutorial Golang"`
	Image string   `json:"image" default:""`
	Body  string   `json:"text" default:"Hello World"`
	Tags  []string `json:"tags" default:["Devs"] `
}

type PostFull struct {
	ID             int            `json:"id" db:"id"`
	PostTitle      string         `json:"post_title" db:"post_title"`
	PostImagePath  string         `json:"image" db:"post_image_path"`
	PostBody       string         `json:"post_body" db:"post_body"`
	PostViewsCount int            `json:"post_views_count" db:"post_views_count"`
	PostLikeCount  int            `json:"post_like_count" db:"post_like_count"`
	PostRated      float32        `json:"post_rating" db:"post_rated"`
	PostVote       int            `json:"post_vote" db:"post_vote"`
	PostTags       pq.StringArray `json:"post_tags" db:"post_tags"`
	PostDate       string         `json:"post_date" db:"post_date"`
	IsNew          bool           `json:"is_empty" db:"is_new"`
	IsTopRead      bool           `json:"is_top_read" db:"is_top_read"`
}