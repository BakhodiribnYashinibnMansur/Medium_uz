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
	CheckDataExists(username string, logrus *logrus.Logger) (int, error)
	SaveVerificationCode(username, code string, logrus *logrus.Logger) error
	CheckCode(username, code string, logrus *logrus.Logger) error
	UpdateUserVerified(username string, logrus *logrus.Logger) (effectedRowsNum int64, err error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB, redis *redis.Client) *Repository {
	return &Repository{Authorization: NewAuthPostgres(db, redis)}
}
