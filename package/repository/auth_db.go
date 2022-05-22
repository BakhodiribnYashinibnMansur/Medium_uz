package repository

import (
	"errors"
	"fmt"

	"github.com/BakhodiribnYashinibnMansur/Medium_uz/model"
	"github.com/BakhodiribnYashinibnMansur/Medium_uz/util/logrus"

	"time"

	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type AuthDB struct {
	db    *sqlx.DB
	redis *redis.Client
}

func NewAuthDB(db *sqlx.DB, redis *redis.Client) *AuthDB {
	return &AuthDB{db: db, redis: redis}
}

func (repo *AuthDB) CreateUser(user model.User, logrus *logrus.Logger) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (email, password_hash,firstname,secondname,nickname,city,phone,interests,bio) values ($1, $2, $3,$4,$5,$6,$7,$8,$9) RETURNING id", usersTable)

	row := repo.db.QueryRow(query, user.Email, user.Password, user.FirstName, user.SecondName, user.NickName, user.City, user.Phone, pq.Array(user.Interests), user.Bio)

	if err := row.Scan(&id); err != nil {
		logrus.Infof("ERROR:PSQL Insert error %s", err.Error())
		return 0, err
	}
	logrus.Info("DONE: INSERTED Data PSQL")
	return id, nil
}

func (repo *AuthDB) CheckDataExistsEmailPassword(email, password string, logrus *logrus.Logger) (int, error) {
	var countUser int

	queryEmail := fmt.Sprintf("SELECT COUNT(id) FROM %s WHERE email=$1 AND password_hash=$2", usersTable)
	err := repo.db.Get(&countUser, queryEmail, email, password)

	if err != nil {
		logrus.Infof("ERROR:Email query error: %s", err.Error())
		return -1, err
	}

	return countUser, nil
}

func (repo *AuthDB) CheckDataExistsEmailNickName(email, nickname string, logrus *logrus.Logger) (int, int, error) {
	var countEmail int
	var countNickName int

	queryEmail := fmt.Sprintf("SELECT COUNT(id) FROM %s WHERE email=$1", usersTable)
	err := repo.db.Get(&countEmail, queryEmail, email)

	if err != nil {
		logrus.Infof("ERROR:Email query error: %s", err.Error())
		return -1, -1, err
	}

	queryNickName := fmt.Sprintf("SELECT COUNT(id) FROM %s WHERE nickname=$1", usersTable)
	err = repo.db.Get(&countNickName, queryNickName, nickname)

	if err != nil {
		logrus.Infof("ERROR:Email query error: %s", err.Error())
		return -1, -1, err
	}
	return countEmail, countNickName, nil
}

func (repo *AuthDB) GetUserID(email string, logrus *logrus.Logger) (int, error) {
	var id int
	query := fmt.Sprintf("SELECT id FROM %s WHERE email=$1 ", usersTable)
	err := repo.db.Get(&id, query, email)
	if err != nil {
		logrus.Errorf("ERROR: don't get users %s", err)
		return 0, errors.New("ERROR: user not found")
	}
	logrus.Info("DONE:get user data from psql")
	return id, nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// REDIS

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (repo *AuthDB) SaveVerificationCode(email, code string, logrus *logrus.Logger) error {
	err := repo.redis.Set(email, code, 180*time.Second).Err()
	if err != nil {
		logrus.Errorf("ERROR:don't save code %s", err)
		return err
	}
	return nil
}
