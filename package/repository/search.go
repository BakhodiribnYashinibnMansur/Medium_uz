package repository

import "github.com/jmoiron/sqlx"

type SearchDB struct {
	db *sqlx.DB
}

func NewSearchDB(db *sqlx.DB) *SearchDB {
	return &SearchDB{db: db}
}
