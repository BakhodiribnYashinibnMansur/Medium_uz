package service

import (
	"mediumuz/model"
	"mediumuz/package/repository"
	"mediumuz/util/logrus"
	"mime/multipart"
)

type Authorization interface {
	CreateUser(user model.User, logrus *logrus.Logger) (int, error)
	GenerateToken(username string, logrus *logrus.Logger) (string, error)
	SendMessageEmail(email, username string, logrus *logrus.Logger) error
	CheckDataExists(username string, logrus *logrus.Logger) (int, error)
	ParseToken(token string) (int, error)
}

type User interface {
	GetUserData(id string, logrus *logrus.Logger) (user model.UserFull, err error)
	VerifyCode(id string, username, code string, logrus *logrus.Logger) (int64, error)
	UploadAccountImage(file multipart.File, header *multipart.FileHeader, user model.UserFull, logrus *logrus.Logger) (filePath string, err error)
	UpdateAccountImage(id int, filePath string, logrus *logrus.Logger) (int64, error)
}
type Service struct {
	Authorization
	User
}

func NewService(repos *repository.Repository) *Service {
	return &Service{Authorization: NewAuthService(repos.Authorization), User: NewUserService(repos.User)}
}
