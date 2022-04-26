package repository

import (
	"mediumuz/model"
	"mediumuz/util/logrus"

	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user model.User, logrus *logrus.Logger) (int, error)
	GetUserID(username string, logrus *logrus.Logger) (int, error)

	CheckDataExistsUsername(username string, logrus *logrus.Logger) (int, error)
	SaveVerificationCode(username, code string, logrus *logrus.Logger) error
}
type User interface {
	GetUserData(id string, logrus *logrus.Logger) (model.UserFull, error)
	CheckCode(username, code string, logrus *logrus.Logger) error
	UpdateUserVerified(id string, logrus *logrus.Logger) (effectedRowsNum int64, err error)
}

type Repository struct {
	Authorization
	User
}

func NewRepository(db *sqlx.DB, redis *redis.Client) *Repository {
	return &Repository{Authorization: NewAuthDB(db, redis), User: NewUserDB(db, redis)}
}
