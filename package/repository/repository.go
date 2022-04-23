package repository

import (
	"mediumuz/model"
	"mediumuz/util/logrus"

	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user model.SingUpUserJson, logrus *logrus.Logger) (int, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB, redis *redis.Client) *Repository {
	return &Repository{Authorization: NewAuthPostgres(db, redis)}
}
