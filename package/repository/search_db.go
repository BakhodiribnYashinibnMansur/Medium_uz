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

func (repo *SearchDB) SearchPost(search string, pagination model.Pagination, logrus *logrus.Logger) ([]model.PostFull, error) {
	var searchPost []model.PostFull
	query := fmt.Sprintf("SELECT  p.id , p.post_title ,p.post_image_path, p.post_views_count, p.post_like_count, p.post_rated, p.post_vote_count, p.post_tags, p.post_date, p.is_new, p.is_top_read,pu.post_author_id,u.firstname,u.secondname,u.account_image_path,u.nickname FROM %s p INNER JOIN %s pu on p.id =pu.post_id INNER JOIN %s u ON u.id= pu.post_author_id WHERE  pu.deleted_at IS NULL AND p.deleted_at IS NULL AND post_title  ~*  $1    OFFSET $2 LIMIT $3", postTable, postUserTable, usersTable)
	err := repo.db.Select(&searchPost, query, search, pagination.Offset, pagination.Limit)
	logrus.Info("DONE:get post data from psql")
	return searchPost, err
}

func (repo *SearchDB) SearchUser(search string, pagination model.Pagination, logrus *logrus.Logger) ([]model.UserFull, error) {
	var searchUser []model.UserFull
	query := fmt.Sprintf("SELECT  	id,	email,	firstname,secondname,nickname,		city,	is_verified,	bio,interests,account_image_path,	phone,	rating,	post_views_count,	saved_post_count,follower_count, following_count,like_count,is_super_user	FROM %s WHERE firstname ~* $1 OR secondname ~* $2 AND deleted_at IS NULL  OFFSET $3 LIMIT $4 ", usersTable)
	err := repo.db.Select(&searchUser, query, search, search, pagination.Offset, pagination.Limit)
	return searchUser, err
}
