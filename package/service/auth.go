package service

import (
	"crypto/sha1"
	"fmt"
	"mediumuz/model"
	"mediumuz/package/repository"
	"mediumuz/util/logrus"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	salt       = "hjqrhjqw124617aj564u564a654u65465aufhajs"
	signingKey = "qrkjk#4#%35FSFJlja#4353K554245987uaS?SFjH"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}
type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (service *AuthService) CreateUser(user model.User, logrus *logrus.Logger) (int, error) {

	user.Password = generatePasswordHash(user.Password)
	logrus.Info("successfully password_hash")
	id, err := service.repo.CreateUser(user, logrus)

	return id, err

}

func (service *AuthService) GenerateToken(username, password string, logrus *logrus.Logger) (string, error) {
	user, err := service.repo.GetUser(username, logrus)
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})
	return token.SignedString([]byte(signingKey))
}

func (service *AuthService) SendMessageEmail(email, username string, logrus *logrus.Logger) error {
	_, err := service.repo.GetUser(username, logrus)
	if err != nil {
		logrus.Errorf("ERROR: sendEmail error get user error : %v", err)
		return err
	}
	code, err := SendCodeToEmail(email, username, logrus)
	if err != nil {
		logrus.Errorf("ERROR: send email error %v", err)
		return err
	}
	err = service.repo.SaveVerificationCode(username, code, logrus)
	if err != nil {
		logrus.Errorf("ERROR: save verification code error %v", err)
		return err
	}
	return nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (service *AuthService) VerifyCode(username, code string, logrus *logrus.Logger) (int64, error) {
	_, err := service.repo.GetUser(username, logrus)
	if err != nil {
		return 0, err
	}
	err = service.repo.CheckCode(username, code, logrus)
	if err != nil {
		logrus.Errorf("ERROR: check error code : %v", err)
		return 0, err
	}
	effectedRowsNum, err := service.repo.UpdateUserVerified(username, logrus)
	if err != nil {
		logrus.Errorf("ERROR: check error code update : %v", err)
		return 0, err
	}
	return effectedRowsNum, nil
}
