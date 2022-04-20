package service

import (
	"mediumuz/package/repository"
	"mediumuz/util/logrus"
)

type Authorization interface {
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository, logrus *logrus.Logger) *Service {
	return &Service{}
}
