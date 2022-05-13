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
	ParseToken(token string) (int, error)
}

type User interface {
	GetUserData(id string, logrus *logrus.Logger) (user model.UserFull, err error)
	VerifyCode(id, email, code string, logrus *logrus.Logger) (int64, error)
	UploadImage(file multipart.File, header *multipart.FileHeader, logrus *logrus.Logger) (filePath string, err error)
	UpdateAccountImage(id int, filePath string, logrus *logrus.Logger) (int64, error)
	UpdateAccount(id int, user model.User, logrus *logrus.Logger) (int64, error)
	CheckUserId(id int, logrus *logrus.Logger) (int, error)
}

type Post interface {
	CreatePost(userId int, post model.Post, logrus *logrus.Logger) (int, error)
	GetPostById(id int, logrus *logrus.Logger) (post model.PostFull, err error)
	CheckPostId(id int, logrus *logrus.Logger) (int, error)
	UpdatePostImage(userID, postID int, filePath string, logrus *logrus.Logger) (int64, error)
	UpdatePost(id int, input model.Post, logrus *logrus.Logger) (int64, error)
	DeletePost(userID, postID int, logrus *logrus.Logger) (int64, int64, error)
	CheckAuthPostId(userID, postID int, logrus *logrus.Logger) (int, error)
}
type Service struct {
	Authorization
	User
	Post
}

func NewService(repos *repository.Repository) *Service {
	return &Service{Authorization: NewAuthService(repos.Authorization), User: NewUserService(repos.User), Post: NewPostService(repos.Post)}
}
