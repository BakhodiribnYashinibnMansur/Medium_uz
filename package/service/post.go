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
func (service *PostService) GetPostBodyById(id int, logrus *logrus.Logger) (post model.PostFull, err error) {
	return service.repo.GetPostBodyById(id, logrus)
}

func (service *PostService) UpdatePostImage(userID, postID int, filePath string, logrus *logrus.Logger) (int64, error) {
	return service.repo.UpdatePostImage(userID, postID, filePath, logrus)
}

func (service *PostService) UpdatePost(userID, postID int, input model.Post, logrus *logrus.Logger) (int64, error) {
	var post model.UpdatePost
	post.Title = convert.EmptyStringToNull(input.Title)
	post.Body = convert.EmptyStringToNull(input.Body)
	post.Tags = convert.EmptyArrayStringToNullArray(input.Tags)
	return service.repo.UpdatePost(userID, postID, post, logrus)
}

func (service *PostService) DeletePost(userID, postID int, logrus *logrus.Logger) (int64, int64, error) {
	return service.repo.DeletePost(userID, postID, logrus)
}

func (service *PostService) LikePost(userID, postID int, logrus *logrus.Logger) (int64, error) {
	return service.repo.LikePost(userID, postID, logrus)
}

func (service *PostService) ViewPost(userID, postID int, logrus *logrus.Logger) (int, error) {
	return service.repo.ViewPost(userID, postID, logrus)
}
func (service *PostService) RatingPost(userID, postID, userRating int, logrus *logrus.Logger) (int64, error) {
	return service.repo.RatingPost(userID, postID, userRating, logrus)
}

func (service *PostService) CommitPost(input model.CommitPost, logrus *logrus.Logger) (int, error) {
	return service.repo.CommitPost(input, logrus)
}

func (service *PostService) GetCommitPost(postID int, pagination model.Pagination, logrus *logrus.Logger) ([]model.CommitFull, error) {
	return service.repo.GetCommitPost(postID, pagination, logrus)
}

func (service *PostService) GetUserPost(userID int, pagination model.Pagination, logrus *logrus.Logger) ([]model.PostFull, error) {
	return service.repo.GetUserPost(userID, pagination, logrus)
}
