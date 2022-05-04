package repository

import (
	"fmt"
	"mediumuz/model"
	"mediumuz/util/logrus"

	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type PostDB struct {
	db    *sqlx.DB
	redis *redis.Client
}

func NewPostDB(db *sqlx.DB, redis *redis.Client) *PostDB {
	return &PostDB{db: db, redis: redis}
}

func (repo *PostDB) CreatePost(post model.Post, logrus *logrus.Logger) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (post_title , post_image_path , post_body , post_tags) VALUES ($1, $2, $3,$4)  RETURNING id", postTable)

	row := repo.db.QueryRow(query, post.Title, post.Image, post.Body, pq.Array(post.Tags))

	if err := row.Scan(&id); err != nil {
		logrus.Infof("ERROR:PSQL Insert error %s", err.Error())
		return 0, err
	}
	logrus.Info("DONE: INSERTED Data PSQL")
	return id, nil
}

func (repo *PostDB) CreatePostUser(userId, postId int, logrus *logrus.Logger) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (post_author_id , post_id ) VALUES ($1, $2)  RETURNING id", postUserTable)
	row := repo.db.QueryRow(query, userId, postId)
	if err := row.Scan(&id); err != nil {
		logrus.Infof("ERROR:PSQL Insert error %s", err.Error())
		return 0, err
	}
	logrus.Info("DONE: INSERTED Data PSQL")
	return id, nil
}

func (repo *PostDB) GetPostById(id int, logrus *logrus.Logger) (post model.PostFull, err error) {
	query := fmt.Sprintf("SELECT post_title ,post_image_path, post_body, post_views_count, post_like_count, post_rated, post_vote, post_tags,  post_date, is_new, is_top_read FROM %s WHERE id = $1", postTable)
	err = repo.db.Get(&post, query, id)
	if err != nil {
		logrus.Errorf("ERROR: don't get users %s", err)
		return post, err
	}
	logrus.Info("DONE:get user data from psql")
	return post, err
}

func (repo *PostDB) CheckPostId(id int, logrus *logrus.Logger) (int, error) {
	var postNumber int
	queryID := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE id=$1", postTable)
	err := repo.db.Get(&postNumber, queryID, id)
	if err != nil {
		logrus.Infof("ERROR:Email query error: %s", err.Error())
		return -1, err
	}
	return postNumber, nil
}
