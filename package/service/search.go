package service

import (
	"strings"

	"github.com/BakhodiribnYashinibnMansur/Medium_uz/model"
	"github.com/BakhodiribnYashinibnMansur/Medium_uz/package/repository"
	"github.com/BakhodiribnYashinibnMansur/Medium_uz/util/logrus"
)

type SearchService struct {
	repo repository.Search
}

func NewSearchService(repo repository.Search) *SearchService {
	return &SearchService{repo: repo}
}

func (service *SearchService) UniversalSearch(search string, logrus *logrus.Logger) (resultSearch model.Search, err error) {

	searchArray := strings.Split(search, " ")
	logrus.Info(searchArray)
	for _, keyword := range searchArray {
		logrus.Info(keyword)
	}
	return resultSearch, err
}
