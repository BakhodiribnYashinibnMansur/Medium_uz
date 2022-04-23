package repository

import (
	"mediumuz/model"
	"mediumuz/util/logrus"

	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db    *sqlx.DB
	redis *redis.Client
}

func NewAuthPostgres(db *sqlx.DB, redis *redis.Client) *AuthPostgres {
	return &AuthPostgres{db: db, redis: redis}
}

func (r *AuthPostgres) CreateUser(user model.SingUpUserJson, logrus *logrus.Logger) (int, error) {
	var id int
	// query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id", usersTable)

	// row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	// if err := row.Scan(&id); err != nil {
	// 	return 0, err
	// }
	return id, nil
}
