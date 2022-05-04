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
func (service *PostService) GetPostById(id int, logrus *logrus.Logger) (post model.PostFull, err error) {
	return service.repo.GetPostById(id, logrus)
}
func (service *PostService) CheckPostId(id int, logrus *logrus.Logger) (int, error) {
	return service.repo.CheckPostId(id, logrus)
}