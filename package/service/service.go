package service

import (
	"mime/multipart"

	"github.com/BakhodiribnYashinibnMansur/Medium_uz/model"
	"github.com/BakhodiribnYashinibnMansur/Medium_uz/package/repository"
	"github.com/BakhodiribnYashinibnMansur/Medium_uz/util/logrus"
)

type Authorization interface {
	CreateUser(user model.User, logrus *logrus.Logger) (int, error)
	GenerateToken(email string, logrus *logrus.Logger) (string, error)
	SendMessageEmail(email, username string, logrus *logrus.Logger) error
	CheckDataExistsEmailNickName(email, nickname string, logrus *logrus.Logger) (model.UserCheck, error)
	CheckDataExistsEmailPassword(email, password string, logrus *logrus.Logger) error
	ParseToken(token string) (int, error)
}

type User interface {
	GetUserData(id string, logrus *logrus.Logger) (user model.UserFull, err error)
	VerifyCode(id, email, code string, logrus *logrus.Logger) (int64, error)
	UploadImage(file multipart.File, header *multipart.FileHeader, logrus *logrus.Logger) (filePath string, err error)
	UpdateAccountImage(id int, filePath string, logrus *logrus.Logger) (int64, error)
	UpdateAccount(id int, user model.User, logrus *logrus.Logger) (int64, error)
}

type Post interface {
	CreatePost(userId int, post model.Post, logrus *logrus.Logger) (int, error)
	GetPostById(id int, logrus *logrus.Logger) (post model.PostFull, err error)
	UpdatePostImage(userID, postID int, filePath string, logrus *logrus.Logger) (int64, error)
	UpdatePost(userID, postID int, post model.Post, logrus *logrus.Logger) (int64, error)
	DeletePost(userID, postID int, logrus *logrus.Logger) (int64, int64, error)
	LikePost(userID, postID int, logrus *logrus.Logger) (int, error)
	UnlikePost(userID, postID int, logrus *logrus.Logger) (int64, error)
}

type Search interface {
	UniversalSearch(search string, logrus *logrus.Logger) (model.Search, error)
}

type Service struct {
	Authorization
	User
	Post
	Search
}

func NewService(repos *repository.Repository) *Service {
	return &Service{Authorization: NewAuthService(repos.Authorization), User: NewUserService(repos.User), Post: NewPostService(repos.Post), Search: NewSearchService(repos.Search)}
}
