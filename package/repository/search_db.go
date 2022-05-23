package repository

import (
	"fmt"

	"github.com/BakhodiribnYashinibnMansur/Medium_uz/model"
	"github.com/BakhodiribnYashinibnMansur/Medium_uz/util/logrus"
	"github.com/jmoiron/sqlx"
)

type SearchDB struct {
	db *sqlx.DB
}

func NewSearchDB(db *sqlx.DB) *SearchDB {
	return &SearchDB{db: db}
}

func (repo *SearchDB) SearchPost(search string, logrus *logrus.Logger) ([]model.PostFull, error) {
	var searchPost []model.PostFull
	query := fmt.Sprintf("SELECT id , post_title ,post_image_path, post_body, post_views_count, post_like_count, post_rated, post_vote, post_tags,  post_date, is_new, is_top_read FROM %s WHERE post_title  ~*  $1  AND deleted_at IS NULL", postTable)
	err := repo.db.Select(&searchPost, query, search)
	logrus.Info("DONE:get post data from psql")
	logrus.Info(searchPost)
	return searchPost, err
}

func (repo *SearchDB) SearchUser(search string, logrus *logrus.Logger) ([]model.UserFull, error) {
	var searchUser []model.UserFull
	query := fmt.Sprintf("SELECT  	id,	email,	firstname,secondname,nickname,		city,	is_verified,	bio,interests,account_image_path,	phone,	rating,	post_views_count,	follower_count, following_count,like_count,is_super_user	FROM %s WHERE firstname ~* $1 OR secondname ~* $2 AND deleted_at IS NULL ", usersTable)
	err := repo.db.Select(&searchUser, query, search, search)
	return searchUser, err
}
