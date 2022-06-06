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

func (service *SearchService) SearchPeople(search string, pagination model.Pagination, logrus *logrus.Logger) (resultSearch []model.UserFull, err error) {

	searchArray := strings.Split(search, " ")
	for _, keyword := range searchArray {
		userArray, err := service.repo.SearchUser(keyword, pagination, logrus)
		if err != nil {
			return resultSearch, err
		}
		for _, value := range userArray {
			resultSearch = append(resultSearch, value)
		}
	}
	return resultSearch, nil
}
func (service *SearchService) SearchPost(search string, pagination model.Pagination, logrus *logrus.Logger) ([]model.PostFull, error) {
	return service.repo.SearchPost(search, pagination, logrus)
}
