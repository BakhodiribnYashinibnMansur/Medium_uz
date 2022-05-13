package repository

import (
	"fmt"
	"time"

	"github.com/BakhodiribnYashinibnMansur/Medium_uz/model"
	"github.com/BakhodiribnYashinibnMansur/Medium_uz/util/logrus"

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
	query := fmt.Sprintf("INSERT INTO %s (post_title  , post_body , post_tags) VALUES ($1, $2, $3)  RETURNING id", postTable)

	row := repo.db.QueryRow(query, post.Title, post.Body, pq.Array(post.Tags))

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
	query := fmt.Sprintf("SELECT id , post_title ,post_image_path, post_body, post_views_count, post_like_count, post_rated, post_vote, post_tags,  post_date, is_new, is_top_read FROM %s WHERE id = $1 AND deleted_at IS NULL", postTable)
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
	queryID := fmt.Sprintf("SELECT COUNT(id) FROM %s WHERE id=$1 AND deleted_at IS NULL", postTable)
	err := repo.db.Get(&postNumber, queryID, id)
	if err != nil {
		logrus.Infof("ERROR:check post id query error: %s", err.Error())
		return -1, err
	}
	return postNumber, nil
}

func (repo *PostDB) CheckAuthPostId(userID, postID int, logrus *logrus.Logger) (int, error) {
	var postNumber int
	queryID := fmt.Sprintf("SELECT COUNT(pl.id) FROM %s pl INNER JOIN %s upl  ON pl.id = upl.post_id WHERE  upl.post_author_id = $1 AND upl.post_id = $2 AND pl.deleted_at IS NULL ", postTable, postUserTable)
	err := repo.db.Get(&postNumber, queryID, userID, postID)
	if err != nil {
		logrus.Infof("ERROR:Check auth post id query error: %s", err.Error())
		return -1, err
	}
	return postNumber, nil
}

func (repo *PostDB) UpdatePostImage(userID, postID int, filePath string, logrus *logrus.Logger) (int64, error) {
	tm := time.Now()
	query := fmt.Sprintf("UPDATE %s pl SET post_image_path = $1,updated_at=$2 FROM %s upl   WHERE pl.id = upl.post_id AND upl.post_author_id = $3 AND upl.post_id = $4 RETURNING pl.id", postTable, postUserTable)

	rows, err := repo.db.Exec(query, filePath, tm, userID, postID)

	if err != nil {
		logrus.Errorf("ERROR: Update PostImage : %v", err)
		return 0, err
	}

	effectedRowsNum, err := rows.RowsAffected()

	if err != nil {
		logrus.Errorf("ERROR: Update Post Image effectedRowsNum : %v", err)
		return 0, err
	}
	logrus.Info("DONE:Update Post image")
	return effectedRowsNum, nil

}

func (repo *PostDB) UpdatePost(userID, postID int, post model.UpdatePost, logrus *logrus.Logger) (int64, error) {
	tm := time.Now()
	updateQuery := fmt.Sprintf("UPDATE %s pl SET  post_title=COALESCE($1,post_title) ,post_body=COALESCE($2,post_body), post_tags=COALESCE($3,post_tags) , updated_at=$4 FROM %s upl   WHERE pl.id = upl.post_id AND upl.post_author_id = $5 AND upl.post_id = $6 RETURNING pl.id", postTable, postUserTable)

	rows, err := repo.db.Exec(updateQuery, post.Title, post.Body, pq.Array(post.Tags), tm, userID, postID)

	if err != nil {
		logrus.Errorf("ERROR: Update Post : %v", err)
		return 0, err
	}

	effectedRowsNum, err := rows.RowsAffected()

	if err != nil {
		logrus.Errorf("ERROR: Update Post  effectedRowsNum : %v", err)
		return 0, err
	}
	logrus.Info("DONE:Update Post ")
	return effectedRowsNum, nil
}

func (repo *PostDB) DeletePost(userID, postID int, logrus *logrus.Logger) (int64, int64, error) {
	tm := time.Now()

	deletePostQuery := fmt.Sprintf("UPDATE %s pl SET deleted_at = $1 FROM %s upl   WHERE pl.id = upl.post_id AND upl.post_author_id = $2 AND upl.post_id = $3 RETURNING pl.id", postTable, postUserTable)
	deletePostRow, err := repo.db.Exec(deletePostQuery, tm, userID, postID)

	if err != nil {
		logrus.Errorf("ERROR: Deleted Post : %v", err)
		return -1, -1, err
	}

	deletedPost, err := deletePostRow.RowsAffected()

	if err != nil {
		logrus.Errorf("ERROR: Deleted Post  effectedRowsNum : %v", err)
		return -1, -1, err
	}
	deletePostUserQuery := fmt.Sprintf("UPDATE %s SET deleted_at = $1 WHERE post_id = $2 RETURNING id", postUserTable)
	deletePostUserRow, err := repo.db.Exec(deletePostUserQuery, tm, userID)

	if err != nil {
		logrus.Errorf("ERROR: Deleted Post : %v", err)
		return -1, -1, err
	}

	deletedPostUser, err := deletePostUserRow.RowsAffected()

	if err != nil {
		logrus.Errorf("ERROR: Deleted Post  effectedRowsNum : %v", err)
		return -1, -1, err
	}
	logrus.Info("DONE:Deleted Post ")
	return deletedPost, deletedPostUser, nil
}
