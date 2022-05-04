package service

import (
	"mediumuz/model"
	"mediumuz/package/repository"
	"mediumuz/util/logrus"
)

type PostService struct {
	repo repository.Post
}

func NewPostService(repo repository.Post) *PostService {
	return &PostService{repo: repo}
}

func (service *PostService) CreatePost(userId int, post model.Post, logrus *logrus.Logger) (int, error) {
	postId, err := service.repo.CreatePost(post, logrus)
	if err != nil {
		return 0, err
	}

	postId, err = service.repo.CreatePostUser(userId, postId, logrus)
	if err != nil {
		return 0, err
	}
	return postId, err
}
