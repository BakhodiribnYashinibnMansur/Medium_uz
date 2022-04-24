package repository

import (
	"errors"
	"fmt"
	"mediumuz/model"
	"mediumuz/util/logrus"
	"time"

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

func (repo *AuthPostgres) CreateUser(user model.User, logrus *logrus.Logger) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (email, password_hash,firstname,secondname,city,phone) values ($1, $2, $3,$4,$5,$6) RETURNING id", usersTable)

	row := repo.db.QueryRow(query, user.Email, user.Password, user.FirstName, user.SecondName, user.City, user.Phone)

	if err := row.Scan(&id); err != nil {
		logrus.Infof("ERROR:PSQL Insert error %s", err.Error())
		return 0, err
	}
	logrus.Info("DONE: INSERTED Data PSQL")
	return id, nil
}
func (repo *AuthPostgres) CheckDataExists(username string, logrus *logrus.Logger) (int, error) {
	var count int

	query := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE firstname=$1", usersTable)
	err := repo.db.Get(&count, query, username)

	if err != nil {
		logrus.Infof("ERROR:firstname query error: %s", err.Error())
		return 0, err
	}
	return count, nil
}
func (repo *AuthPostgres) GetUserID(username string, logrus *logrus.Logger) (int, error) {
	var id int
	query := fmt.Sprintf("SELECT id FROM %s WHERE firstname=$1 ", usersTable)
	err := repo.db.Get(&id, query, username)
	if err != nil {
		logrus.Errorf("ERROR: don't get users %s", err)
		return 0, errors.New("ERROR: user not found")
	}
	logrus.Info("DONE:get user data from psql")
	return id, nil
}

func (repo *AuthPostgres) UpdateUserVerified(username string, logrus *logrus.Logger) (effectedRowsNum int64, err error) {
	tm := time.Now()
	query := fmt.Sprintf("	UPDATE %s SET isverified = true,verification_date=$1	WHERE firstname = $2  RETURNING id ", usersTable)
	rows, err := repo.db.Exec(query, tm, username)

	if err != nil {
		logrus.Errorf("ERROR: Update verificationCode : %v", err)
		return 0, err
	}
	effectedRowsNum, err = rows.RowsAffected()
	if err != nil {
		logrus.Errorf("ERROR: Update verificationCode effectedRowsNum : %v", err)
		return 0, err
	}
	logrus.Info("DONE:Update verify email")
	return effectedRowsNum, nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// REDIS

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (repo *AuthPostgres) SaveVerificationCode(username, code string, logrus *logrus.Logger) error {
	err := repo.redis.Set(username, code, 180*time.Second).Err()
	if err != nil {
		logrus.Errorf("ERROR:don't save code %s", err)
		return err
	}
	return nil
}

func (repo *AuthPostgres) CheckCode(username, code string, logrus *logrus.Logger) error {
	saveCode, err := repo.redis.Get(username).Result()
	if err != nil {
		logrus.Errorf("ERROR:dont save code %s", err)
		return err
	}
	if saveCode != code {
		return errors.New("code not found ")
	}
	logrus.Info("DONE: verify code")
	return nil
}
