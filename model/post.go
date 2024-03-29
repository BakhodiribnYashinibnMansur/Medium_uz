package model

import (
	"database/sql"

	"github.com/lib/pq"
)

type Post struct {
	ID    int      `json:"-"`
	Title string   `json:"title" default:"Tutorial Golang"`
	Body  string   `json:"text" default:"Hello World"`
	Tags  []string `json:"tags" default:["Devs"] `
}

type UpdatePost struct {
	ID    int              `json:"-"`
	Title sql.NullString   `json:"title" default:"Tutorial Golang"`
	Body  sql.NullString   `json:"text" default:"Hello World"`
	Tags  []sql.NullString `json:"tags" default:["Devs"] `
}

type PostFull struct {
	ID                   int            `json:"id" db:"id"`
	AuthorID             int            `json:"post_author_id" db:"post_author_id"`
	AuthorFirstName      string         `json:"post_author_first_name" db:"firstname"`
	AuthorSecondName     string         `json:"post_author_second_name" db:"secondname"`
	AuthorImage          string         `json:"post_author_image" db:"account_image_path"`
	AuthorNickname       string         `json:"post_author_nickname" db:"nickname"`
	PostTitle            string         `json:"post_title" db:"post_title"`
	PostImagePath        sql.NullString `json:"image" db:"post_image_path"`
	PostBody             string         `json:"post_body" db:"post_body"`
	PostViewsCount       int            `json:"post_views_count" db:"post_views_count"`
	PostLikeCount        int            `json:"post_like_count" db:"post_like_count"`
	PostRated            float64        `json:"post_rating" db:"post_rated"`
	PostVoteCount        int            `json:"post_vote_count" db:"post_vote_count"`
	PostTags             pq.StringArray `json:"post_tags" db:"post_tags"`
	PostDate             string         `json:"post_date" db:"post_date"`
	IsNew                bool           `json:"is_empty" db:"is_new"`
	IsTopRead            bool           `json:"is_top_read" db:"is_top_read"`
	PostHistoryDate      string         `json:"post_history_date" db:"view_date"`
	PostSavedHistoryDate string         `json:"post_saved_history_date" db:"saved_date"`
}

type CommitPost struct {
	ReaderID   int    `json:"-" db:"reader_id"`
	PostID     int    `json:"post_id" db:"post_id"`
	PostCommit string `json:"post_commit" db:"commits" default:"Wonderful"`
}

type CommitFull struct {
	UserID       int            `json:"reader_id" db:"id"`
	FirstName    string         `json:"first_name" db:"firstname"`
	SecondName   string         `json:"second_name" db:"secondname"`
	UserNickname string         `json:"user_nickname" db:"nickname"`
	UserImage    sql.NullString `json:"user_image" db:"account_image_path"`
	PostCommit   string         `json:"post_commit" db:"commits" `
	CommitData   string         `json:"commit_data" db:"review_data"`
}
