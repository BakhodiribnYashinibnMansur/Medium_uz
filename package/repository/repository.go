package repository

import (
	"mediumuz/model"
	"mediumuz/util/logrus"

	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user model.User, logrus *logrus.Logger) (int, error)
	GetUserID(email string, logrus *logrus.Logger) (int, error)
	CheckDataExistsEmailNickName(email, nickname string, logrus *logrus.Logger) (int, int, error)
	SaveVerificationCode(email, code string, logrus *logrus.Logger) error
}

type User interface {
	GetUserData(id string, logrus *logrus.Logger) (model.UserFull, error)
	CheckCode(email, code string, logrus *logrus.Logger) error
	UpdateUserVerified(id string, logrus *logrus.Logger) (effectedRowsNum int64, err error)
	UpdateAccountImage(id int, filePath string, logrus *logrus.Logger) (int64, error)
	UpdateAccount(id int, user model.UpdateUser, logrus *logrus.Logger) (int64, error)
	CheckUserId(id int, logrus *logrus.Logger) (int, error)
}

type Post interface {
	CreatePost(post model.Post, logrus *logrus.Logger) (int, error)
	CreatePostUser(userId, postId int, logrus *logrus.Logger) (int, error)
}

type Repository struct {
	Authorization
	User
	Post
}

func NewRepository(db *sqlx.DB, redis *redis.Client) *Repository {
	return &Repository{Authorization: NewAuthDB(db, redis), User: NewUserDB(db, redis), Post: NewPostDB(db, redis)}
}
