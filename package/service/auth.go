package service

import (
	"crypto/sha1"
	"fmt"
	"mediumuz/model"
	"mediumuz/package/repository"
	"mediumuz/util/logrus"
	"time"
)

const (
	salt       = "hjqrhjqw124617aj564u564a654u65465aufhajs"
	signingKey = "qrkjk#4#%35FSFJlja#4353K554245987uaS?SFjH"
	tokenTTL   = 12 * time.Hour
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user model.SingUpUserJson, logrus *logrus.Logger) (int, error) {

	user.Password = generatePasswordHash(user.Password)
	logrus.Info("successfully password_hash")

	return s.repo.CreateUser(user, logrus)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
