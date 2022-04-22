package repository

import (
	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB, redis *redis.Client) *Repository {
	return &Repository{}
}
