package repository

import (
	"database/sql"
	"errors"
	"fmt"

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
		if err == sql.ErrNoRows {
			return post, errors.New("ID not found")
		}
		logrus.Errorf("ERROR: don't get users %s", err)
		return post, err
	}
	logrus.Info("DONE:get user data from psql")
	return post, err
}

func (repo *PostDB) UpdatePostImage(userID, postID int, filePath string, logrus *logrus.Logger) (int64, error) {

	query := fmt.Sprintf("UPDATE %s pl SET post_image_path = $1,updated_at=NOW() FROM %s upl   WHERE pl.id = upl.post_id AND upl.post_author_id = $2 AND upl.post_id = $3 RETURNING pl.id", postTable, postUserTable)

	rows, err := repo.db.Exec(query, filePath, userID, postID)

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
	updateQuery := fmt.Sprintf("UPDATE %s pl SET  post_title=COALESCE($1,post_title) ,post_body=COALESCE($2,post_body), post_tags=COALESCE($3,post_tags) , updated_at=NOW() FROM %s upl   WHERE pl.id = upl.post_id AND upl.post_author_id = $4 AND upl.post_id = $5 RETURNING pl.id", postTable, postUserTable)

	rows, err := repo.db.Exec(updateQuery, post.Title, post.Body, pq.Array(post.Tags), userID, postID)

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

	deletePostQuery := fmt.Sprintf("UPDATE %s pl SET deleted_at = NOW() FROM %s upl   WHERE pl.id = upl.post_id AND upl.post_author_id = $1 AND upl.post_id = $2 RETURNING pl.id", postTable, postUserTable)
	deletePostRow, err := repo.db.Exec(deletePostQuery, userID, postID)

	if err != nil {
		logrus.Errorf("ERROR: Deleted Post : %v", err)
		return -1, -1, err
	}

	deletedPost, err := deletePostRow.RowsAffected()

	if err != nil {
		logrus.Errorf("ERROR: Deleted Post  effectedRowsNum : %v", err)
		return -1, -1, err
	}
	deletePostUserQuery := fmt.Sprintf("UPDATE %s SET deleted_at = NOW() WHERE post_id = $1 RETURNING id", postUserTable)
	deletePostUserRow, err := repo.db.Exec(deletePostUserQuery, userID)

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

func (repo *PostDB) LikePost(userID, postID int, logrus *logrus.Logger) (int64, error) {

	likeQuery := fmt.Sprintln("SELECT toggle_comment_like($1,$2)")

	row, err := repo.db.Exec(likeQuery, userID, postID)

	if err != nil {
		logrus.Info("DONE: ERROR  LIKE Data PSQL %s ", err)
		return -1, err
	}
	numRowEffected, err := row.RowsAffected()
	if err != nil {
		logrus.Info("DONE: ERROR  LIKE Data PSQL %s ", err)

		return -1, err
	}
	logrus.Info("DONE: INSERTED  LIKE Data PSQL")
	return numRowEffected, nil
}

func (repo *PostDB) ViewPost(userID, postID int, logrus *logrus.Logger) (int, error) {

	var id int
	query := fmt.Sprintf("INSERT INTO %s (reader_id  , post_id ) VALUES ($1, $2)  RETURNING id", postViewTable)

	row := repo.db.QueryRow(query, userID, postID)

	if err := row.Scan(&id); err != nil {
		logrus.Infof("ERROR:PSQL Insert VIEW error %s", err.Error())
		return 0, err
	}
	logrus.Info("DONE: INSERTED  VIEW Data PSQL")
	return id, nil
}

func (repo *PostDB) RatingPost(userID, postID, userRating int, logrus *logrus.Logger) {}

func (repo *PostDB) CommitPost(input model.CommitPost, logrus *logrus.Logger) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (reader_id,post_id,commits) VALUES ($1, $2, $3)  RETURNING id", postCommitTable)

	row := repo.db.QueryRow(query, input.ReaderID, input.PostID, input.PostCommit)

	if err := row.Scan(&id); err != nil {
		logrus.Infof("ERROR:PSQL Insert error %s", err.Error())
		return 0, err
	}
	logrus.Info("DONE: INSERTED Data PSQL")
	return id, nil
}
