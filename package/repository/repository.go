package repository

import (
	"mediumuz/util/logrus"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB, logrus *logrus.Logger) *Repository {
	return &Repository{}
}
