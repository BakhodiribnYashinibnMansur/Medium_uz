package repository

import (
	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
)

type PostDB struct {
	db    *sqlx.DB
	redis *redis.Client
}

func NewPostDB(db *sqlx.DB, redis *redis.Client) *PostDB {
	return &PostDB{db: db, redis: redis}
}
