package service

import (
	"mediumuz/model"
	"mediumuz/package/repository"
	"mediumuz/util/logrus"
)

type Authorization interface {
	CreateUser(user model.User, logrus *logrus.Logger) (int, error)
	GenerateToken(username, password string, logrus *logrus.Logger) (string, error)
	SendMessageEmail(email, username string, logrus *logrus.Logger) error
	VerifyCode(username, code string, logrus *logrus.Logger) (int64, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{Authorization: NewAuthService(repos.Authorization)}
}
