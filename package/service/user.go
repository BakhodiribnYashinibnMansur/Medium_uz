package service

import (
	"mediumuz/model"
	"mediumuz/package/repository"
	"mediumuz/util/logrus"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (service *UserService) GetUserData(id string, logrus *logrus.Logger) (user model.UserFull, err error) {
	user, err = service.repo.GetUserData(id, logrus)
	if err != nil {
		logrus.Error("ERROR: get user Data failed: %v", err)
		return user, err
	}
	return user, nil
}

func (service *UserService) VerifyCode(id, username, code string, logrus *logrus.Logger) (int64, error) {
	err := service.repo.CheckCode(username, code, logrus)
	if err != nil {
		logrus.Errorf("ERROR: check error code : %v", err)
		return 0, err
	}
	effectedRowsNum, err := service.repo.UpdateUserVerified(id, logrus)
	if err != nil {
		logrus.Errorf("ERROR: check error code update : %v", err)
		return 0, err
	}
	return effectedRowsNum, nil
}
