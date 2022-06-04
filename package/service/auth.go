package service

import (
	"errors"

	"github.com/BakhodiribnYashinibnMansur/Medium_uz/util/logrus"

	"github.com/BakhodiribnYashinibnMansur/Medium_uz/model"
	"github.com/BakhodiribnYashinibnMansur/Medium_uz/package/repository"
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

func (service *AuthService) CheckDataExistsEmailPassword(email, password string, logrus *logrus.Logger) error {
	passwordHash := generatePasswordHash(password)
	logrus.Info("successfully password_hash")
	userCount, err := service.repo.CheckDataExistsEmailPassword(email, passwordHash, logrus)
	if err != nil {
		return err
	}
	if userCount == 0 {
		return errors.New("Email or Password is inCorrect")
	}
	return nil
}

func (service *AuthService) CheckDataExistsEmailNickName(email, nickname string, logrus *logrus.Logger) (model.UserCheck, error) {
	var checkUser model.UserCheck
	countEmail, countNickName, err := service.repo.CheckDataExistsEmailNickName(email, nickname, logrus)
	if err != nil {
		logrus.Errorf("ERROR: CheckDataExists  failed: %v", err)
		return checkUser, err
	}
	checkUser.City = true
	checkUser.FirstName = true
	checkUser.Password = true
	checkUser.Phone = true
	checkUser.SecondName = true

	if countEmail == 0 {
		checkUser.Email = true
	}
	if countNickName == 0 {
		checkUser.NickName = true
	}
	return checkUser, nil
}

func (service *AuthService) GenerateToken(email string, logrus *logrus.Logger) (string, error) {
	id, err := service.repo.GetUserID(email, logrus)
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
	logrus.Infof(email)
	code, err := SendCodeToEmail(email, username, logrus)
	if err != nil {
		logrus.Errorf("ERROR: send email error %v", err)
		return err
	}
	err = service.repo.SaveVerificationCode(email, code, logrus)
	if err != nil {
		logrus.Errorf("ERROR: save verification code error %v", err)
		return err
	}
	return nil
}

func (service *AuthService) RecoveryCheckEmailCode(email string, code string, logrus *logrus.Logger) error {
	err := service.repo.RecoveryCheckEmailCode(email, code, logrus)
	if err != nil {
		logrus.Errorf("ERROR: check error code : %v", err)
		return err
	}
	return nil
}

func (service *AuthService) UpdateAccountPassword(email string, password string, logrus *logrus.Logger) (int64, error) {
	passwordHash := generatePasswordHash(password)
	return service.repo.UpdateAccountPassword(email, passwordHash, logrus)
}
