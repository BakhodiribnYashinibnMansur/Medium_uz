package service

import (
	"github.com/BakhodiribnYashinibnMansur/Medium_uz/model"
	"github.com/BakhodiribnYashinibnMansur/Medium_uz/package/repository"
	"github.com/BakhodiribnYashinibnMansur/Medium_uz/util/convert"
	"github.com/BakhodiribnYashinibnMansur/Medium_uz/util/logrus"
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

func (service *PostService) UpdatePostImage(id int, filePath string, logrus *logrus.Logger) (int64, error) {
	return service.repo.UpdatePostImage(id, filePath, logrus)
}

func (service *PostService) UpdatePost(id int, input model.Post, logrus *logrus.Logger) (int64, error) {
	var post model.UpdatePost
	post.Title = convert.EmptyStringToNull(input.Title)
	post.Body = convert.EmptyStringToNull(input.Body)
	post.Tags = convert.EmptyArrayStringToNullArray(input.Tags)
	return service.repo.UpdatePost(id, post, logrus)
}

func (service *PostService) DeletePost(userID, postID int, logrus *logrus.Logger) (int64, int64, error) {
	return service.repo.DeletePost(userID, postID, logrus)
}

func (service *PostService) CheckAuthPostId(userID, postID int, logrus *logrus.Logger) (int, error) {
	return service.repo.CheckAuthPostId(userID, postID, logrus)
}
