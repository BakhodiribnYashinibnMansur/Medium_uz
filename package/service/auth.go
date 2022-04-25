package service

import (
	"errors"
	"mediumuz/model"
	"mediumuz/package/repository"
	"mediumuz/util/logrus"

	"github.com/dgrijalva/jwt-go"
)

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
	if err != nil {
		return 0, err
	}
	return id, err
}

func (service *AuthService) CheckDataExists(username string, logrus *logrus.Logger) (int, error) {
	count, err := service.repo.CheckDataExists(username, logrus)
	if err != nil {
		logrus.Errorf("ERROR: CheckDataExists  failed: %v", err)
		return 0, err
	}
	return count, nil
}

func (service *AuthService) GenerateToken(username string, logrus *logrus.Logger) (string, error) {
	count, err := service.repo.CheckDataExists(username, logrus)
	if err != nil {
		logrus.Errorf("ERROR: CheckDataExists  failed: %v", err)
		return "", err
	}
	if count == 0 {
		logrus.Errorf("ERROR:  user data not exist : %v", err)
		return "", errors.New("User data not exist")
	}
	id, err := service.repo.GetUserID(username, logrus)
	if err != nil {
		logrus.Errorf("ERROR: Get user ID failed: %v", err)
		return "", err
	}

	token, err := generateJWTToken(id)
	if err != nil {
		logrus.Errorf("ERROR: Get user Token failed: %v", err)
		return "", err
	}
	return token, nil
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}
	return claims.UserId, nil
}

func (service *AuthService) SendMessageEmail(email, username string, logrus *logrus.Logger) error {
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

func (service *AuthService) VerifyCode(username, code string, logrus *logrus.Logger) (int64, error) {
	err := service.repo.CheckCode(username, code, logrus)
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
