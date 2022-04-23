package repository

import (
	"fmt"
	"mediumuz/util/logrus"

	"github.com/jmoiron/sqlx"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

const (
	usersTable = "users"
)

func NewPostgresDB(cfg Config, logrus *logrus.Logger) (*sqlx.DB, error) {

	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	logrus.Infof("check db configs %v", cfg)
	if err != nil {
		logrus.Fatalf("failed check db confis.%v", err)
		return nil, err
	}
	logrus.Info("success checked configs.")
	err = db.Ping()
	logrus.Info("send data ping ")
	if err != nil {
		logrus.Fatalf("fail ping to db %v", err)
		return nil, err
	}
	logrus.Info("success ping data. ")
	logrus.Info("successfull connect database")
	return db, nil
}
